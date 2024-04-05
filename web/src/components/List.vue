<template>
  <v-app id="inspire">
    <v-app-bar flat>
      <v-container class="mx-auto d-flex align-center justify-center">
        <v-avatar
          class="me-4 "
          color="grey-darken-1"
          size="32"
        ></v-avatar>

        <v-btn
          v-for="link in links"
          :key="link"
          :text="link"
          variant="text"
        ></v-btn>

        <v-spacer></v-spacer>

        <v-responsive max-width="160">
          <v-text-field
            density="compact"
            flat
            hide-details
            label="Search"
            rounded="lg"
            single-line
            variant="solo-filled"
          ></v-text-field>
        </v-responsive>
      </v-container>
    </v-app-bar>

    <v-main class="bg-grey-lighten-3">
      <v-container>

        <!-- <v-row>
          <v-col
            v-for="n in 24"
            :key="n"
            cols="4"
          >
            <v-card height="200"></v-card>
          </v-col>
        </v-row> -->

        <!-- <v-row>
          <v-col cols="6" xs="12">
            <v-card height="100" color="success"></v-card>
          </v-col>

          <v-col cols="6" xs="2">
            <v-card height="100" color="warning"></v-card>
          </v-col>
        </v-row> -->

        <v-row>
          <!-- LEFT MENU -->
          <v-col xs="12" sm="12" md="2" xxl="1" class="mt-4">
            <v-sheet rounded="lg">
              <v-list rounded="lg">
                <v-list-item
                  v-for="(c, i) in cams"
                  :title="c"
                  link
                  @click="loadVideos(c)"
                  :key="i"
                >
                </v-list-item>

                <v-divider class="my-2"></v-divider>

                <v-list-item
                  color="grey-lighten-4"
                  link
                  title="Carregar cams"
                  @click="loadCams"
                ></v-list-item>
              </v-list>
            </v-sheet>
          </v-col>

          <v-col xs="12" sm="12" md="10" xxl="10" >
            <!-- <v-container>
              <v-row> -->
                <v-data-iterator
                  :items="videos"
                  :items-per-page="listWidth"
                  :page="page"
                >
                  <template v-slot:default="{ items }">
                    <v-container fluid>
                      <v-row dense>
                        <v-col
                          v-for="item in items"
                          :key="item"
                          xs="12"
                          sm="12"
                          md="10"
                          lg="6"
                          xl="4"
                          xxl="4"
                        >
                          <v-card height="400">
                            <v-card-title>{{item.raw.name}}</v-card-title>
                            <video v-if="videos.length > 0" ref="videoRef" :src=getVideoSrc(item.raw.location) type="video/mp4" id="video-container" width="100%" controls></video>
                          </v-card>
                        </v-col>
                      </v-row>
                    </v-container>
                  </template>




                  <template v-slot:footer="{ page, pageCount, prevPage, nextPage }">
                    <div class="d-flex align-center justify-center pa-4">
                      <v-btn
                        :disabled="page === 1"
                        icon="mdi-arrow-left"
                        density="comfortable"
                        variant="tonal"
                        rounded
                        @click="prevPage"
                      ></v-btn>

                      <div class="mx-2 text-caption">
                        Page {{ page }} of {{ pageCount }}
                      </div>

                      <v-btn
                        :disabled="page >= pageCount"
                        icon="mdi-arrow-right"
                        density="comfortable"
                        variant="tonal"
                        rounded
                        @click="nextPage"
                      ></v-btn>
                    </div>
                  </template>
                </v-data-iterator>
                <!-- <v-col
                  v-for="v in videos"
                  :key="v"
                  xs="12"
                  sm="4"
                > -->
                  <!-- <v-card height="200"></v-card> -->
                  <!-- <video v-if="videos.length > 0" ref="videoRef" :src=videoServer+v id="video-container" width="100%" controls></video> -->
                <!-- </v-col> -->
              <!-- </v-row>
            </v-container> -->
          </v-col>

        </v-row>
      </v-container>
    </v-main>
  </v-app>
</template>

<script>

  import videojs from 'video.js';
  window.videojs = videojs;
  import VideoPlayer from '@/components/Player.vue';
  import { useDisplay } from 'vuetify/lib/framework.mjs';

  export default {
    components: {
      VideoPlayer
    },
    setup(){
      const { name } = useDisplay()
      return {
        name
      }
    },
    async mounted() {
      if (__API_HOST__ && __API_PORT__) {
        this.serverAdress = `http://${__API_HOST__}:${__API_PORT__}`;
      }

      this.loadCams();
      this.loadVideos()
    },
    data: () => ({
      cams: [],
      videos: [],
      serverPort: __API_PORT__,
      serverHost: __API_HOST__,
      serverAdress: '',
      // videoServer: 'http://10.0.0.15:4000',
      page: 1,
      items: Array.from({ length: 40 }, (k, v) => ({
        title: 'Item ' + v + 1,
        text: 'Lorem ipsum dolor sit amet consectetur adipisicing elit. Commodi, ratione debitis quis est labore voluptatibus! Eaque cupiditate minima, at placeat totam, magni doloremque veniam neque porro libero rerum unde voluptatem!',
      })),
      links: [
        'Dashboard',
        'Messages',
        'Profile',
        'Updates',
      ],
      // videoOptions: {
      //     autoplay: false,
      //     controls: false,
      //     // height: 300,
      //     // width: '200',
      //     // fluid: true,
      //     // liveui: true,
      //     responsive: true,
      //     sources: [
      //       {
      //         src:
      //           // 'https://vjs.zencdn.net/v/oceans.mp4',
      //           'http://10.0.0.11:8080/teste.mkv',
      //           type: 'video/mp4'
      //       }
      //     ]
      //   }
    }),
    computed: {
      listWidth() {
        switch (this.name) {
          case "xs":
            return 1;
          case "sm":
            return 10;
          case "md":
            return 6;
          case "lg":
            return 6;
          case "xl":
            return 9;
          case "xxl":
            return 9;
          default:
            return 12
        }
      },
    },
    methods: {
      getVideoSrc(videoPath) {
        return `${this.serverAdress}/recordings/${videoPath}`
      },
      async loadCams() {
        fetch(`${this.serverAdress}/list-cams`).then(res => {
          if (!res.ok) {
            throw new Error(`HTTP error! Status: ${res.status}`);
          }

          return res.json();
        }).then(res => {
          this.cams = res;
        })
      },
      async loadVideos(cam) {

        fetch(`${this.serverAdress}/list-videos/${cam}`).then(res => {

          if (!res.ok) {
            throw new Error(`HTTP error! Status: ${res.status}`);
          }

          return res.json();
        }).then(res => {
          console.log('videos: ', res);
          this.videos = res;
        })
      }
    }
  }

  // export default {
  //   name: 'VideoExample',
  //   components: {
  //     VideoPlayer
  //   },
  //   data() {
  //     return {
  //       videoOptions: {
  //         autoplay: true,
  //         controls: true,
  //         sources: [
  //           {
  //             src:
  //               'http://localhost/recordings/teste.mp4',
  //               type: 'video/mp4'
  //           }
  //         ]
  //       }
  //     };
  //   }
  // };
</script>
