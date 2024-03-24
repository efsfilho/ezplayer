package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/yaml.v3"
)

const CONFIG_FILE = "config.yaml"

type Config struct {
	Server struct {
		Port uint16 `yaml:"port"`
	}
	Eznvr struct {
		Storage string `yaml:"storage"`
	}
}

var (
	SERVER_PORT   uint16
	EZNVR_STORAGE string
)

func loadConfig() {
	var config Config

	// Cheks config file
	createFile := false
	if _, err := os.Stat(CONFIG_FILE); err != nil {
		if createFile = os.IsNotExist(err); !createFile {
			log.Fatal(err)
		}
	}

	// // Creates a default config
	// if createFile {
	// 	conf = Config{
	// 		struct {
	// 			Port uint16 "yaml:\"port\""
	// 		}{4000},
	// 		struct {
	// 			Storage string "yaml:\"storage\""
	// 		}{"/path/to/eznvr/storage"},
	// 	}
	// 	var out []byte
	// 	var newFile *os.File
	// 	var err error

	// 	if out, err = yaml.Marshal(conf); err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	if newFile, err = os.Create(CONFIG_FILE); err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	defer newFile.Close()
	// 	if _, err = newFile.Write(out); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	// Reads config
	var file []byte
	var err error
	if file, err = os.ReadFile(CONFIG_FILE); err != nil {
		log.Fatal(err)
	}
	if err = yaml.Unmarshal(file, &config); err != nil {
		log.Fatal(err)
	}

	SERVER_PORT = config.Server.Port
	EZNVR_STORAGE = config.Eznvr.Storage

	if _, err := os.Stat(EZNVR_STORAGE); err != nil {
		if os.IsNotExist(err) {
			log.Fatal(err)
		}
	}
	log.Println("EZ-NVR storage:", EZNVR_STORAGE)
}

type Video struct {
	Name             string    `json:"name"`
	Location         string    `json:"location"`
	LastModification time.Time `json:"-"`
}

func main() {
	loadConfig()

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339_nano}] - ${remote_ip} \"${method} ${uri} ${protocol}\" ${status} ${bytes_out} - ${user_agent} \n",
	}))
	e.Use(middleware.Recover())
	// e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", "http://10.0.0.11:3000"},
	}))

	// Routes
	e.Static("/", "web")
	e.Static("/recordings", EZNVR_STORAGE) //TODO
	e.GET("/list-cams", listCams)
	e.GET("/list-videos/:cam", listVideos)

	// Start server
	e.HideBanner = true
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", SERVER_PORT)))
}

func listCams(c echo.Context) error {
	log.Println()
	var cams []string
	files, err := os.ReadDir(EZNVR_STORAGE)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		// TODO remove filter
		if strings.HasPrefix(file.Name(), "portao") {
			cams = append(cams, file.Name())
			fmt.Println(file.Name(), file.IsDir())
		}
	}

	return c.JSON(http.StatusOK, cams)
}

func listFiles(vid Video) []Video {

	var videos []Video
	dirEntry, err := os.ReadDir(vid.Location)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range dirEntry {
		filePath := path.Join(vid.Location, file.Name())
		if file.IsDir() {
			video := Video{
				path.Base(filePath),
				filePath,
				time.Now(),
			}
			videos = append(videos, listFiles(video)...)
		} else {
			// TODO make a different restriction to match other files than mkv files
			if strings.HasSuffix(filePath, ".mkv") {
				lastModification := time.Now()
				info, e := file.Info()
				if e == nil {
					lastModification = info.ModTime()
				}
				location := strings.Replace(filePath, "..", "", 1)
				location = strings.Replace(location, EZNVR_STORAGE, "", 1)
				video := Video{
					path.Base(filePath),
					location,
					lastModification,
				}
				videos = append(videos, video)
			}
		}
	}
	return videos
}

func listVideos(c echo.Context) error {

	cam := c.Param("cam")
	videosPath := path.Join(EZNVR_STORAGE, cam)
	_, err := os.Stat(videosPath)

	if err != nil {
		if os.IsNotExist(err) {
			return echo.NewHTTPError(http.StatusNotFound, "Camera not found")
		} else {
			log.Fatal(err)
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
	}

	var files1, files2 []Video
	files1 = listFiles(Video{"", videosPath, time.Now()})

	for i := len(files1) - 1; i >= 0; i-- {
		files2 = append(files2, files1[i])
	}

	return c.JSON(200, files2)
}
