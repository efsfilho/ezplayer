# EZPLAYER

This is a very simple media server built with: Go framework echo, Vuetifyjs and Video.js(video player), that you can use to browser and play video files. 

It was primarily build to be used as a web interface for [EZ-NVR](https://github.com/cyb3rdoc/ez-nvr), a really interesting NVR project. But it can be used to play any mapped mkv/mp4 files.


## Usage

```bash
docker run -d --name ezplayer -v /your/video/storage:/storage -p 4000:4000 efsfilho/ezplayer:latest
```
The ezplayer will use the port ```4000```

Replace ```/your/video/storage``` with the location of your video files. In the web interface, videos will be grouped by "cameras" as it uses the EZ-NVR recording folder structure, where folders at the first level are created for each camera connection.


The web interface will be available at ```http://localhost:4000```
