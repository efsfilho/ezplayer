<template>
  <v-app id="inspire">

    <NavigationDrawer
      :rail="rail"
      :cameras="cameras"
      @load-cameras="loadCams"
      @select-camera="loadVideos"
    />

    <AppBar
      @rail="rail = !rail"
    />

    <v-main>
      <Grid
        v-if="videos.length > 0"
        @get-video-scr="getVideoSrc"
        :videos="videos"
      />
    </v-main>

  </v-app>
</template>

<script >
  export default {
    data: () => ({
      rail: false,
      cameras: [],
      videos: [],
      serverAdress: 'http://localhost:4000'
    }),
    // computed: {
    //   listWidth() {
    //     switch (this.name) {
    //       case "xs":
    //         return 1;
    //       case "sm":
    //         return 10;
    //       case "md":
    //         return 6;
    //       case "lg":
    //         return 6;
    //       case "xl":
    //         return 9;
    //       case "xxl":
    //         return 9;
    //       default:
    //         return 12
    //     }
    //   },
    // },
    methods: {
      async loadCams() {
        fetch(`${this.serverAdress}/list-cams`).then(res => {
          if (!res.ok) {
            throw new Error(`HTTP error! Status: ${res.status}`);
          }
          return res.json();
        }).then(res => {
          this.cameras = res;
        })
      },

      async loadVideos(cam) {
        fetch(`${this.serverAdress}/list-videos/${cam}`).then(res => {
          if (!res.ok) {
            throw new Error(`HTTP error! Status: ${res.status}`);
          }
          return res.json();
        }).then(res => {
          this.videos = res.map(v => ({name: v.name, location: this.getVideoSrc(v.location)}));
        });
      },
      getVideoSrc(videoPath) {
         return `${this.serverAdress}/recordings/${videoPath}`
      },
    }
  }
</script>
