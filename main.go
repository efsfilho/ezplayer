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
)

var storage = "/"

type Video struct {
	Name             string    `json:"name"`
	Location         string    `json:"location"`
	LastModification time.Time `json:"-"`
}

func main() {

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
	e.Static("/recordings", storage)
	e.GET("/list-cams", listCams)
	e.GET("/list-videos/:cam", listVideos)

	// Start server
	e.Logger.Fatal(e.Start(":4000"))
}

func listCams(c echo.Context) error {

	var cams []string
	files, err := os.ReadDir(storage)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if strings.HasPrefix(file.Name(), "portao") {
			cams = append(cams, file.Name())
			fmt.Println(file.Name(), file.IsDir())
		}
	}

	return c.JSON(http.StatusOK, cams)
}

func listFiles(vid Video) []Video {

	var videos []Video
	files, err := os.ReadDir(vid.Location)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
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
				location = strings.Replace(location, storage, "", 1)
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
	videosPath := path.Join(storage, cam)
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
