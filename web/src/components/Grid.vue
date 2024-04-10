<script setup>
  import { ref } from 'vue'
  const page = ref(1);
  const emit = defineEmits(['get-video-scr']);
  const videos = defineProps(['videos']);
</script>

<template>
  <v-container>
    <!-- <v-btn @click="teste"> teste</v-btn> -->
    <v-row>
      <v-col xs="12" sm="12" md="10" xxl="10" >
        <v-data-iterator
          :items="videos.videos"
          items-per-page=12
          :page="page"
        >
          <template v-slot:default="{ items }">
            <v-container fluid>
              <v-row dense>
                <v-col
                  v-for="(item, i) in items"
                  :key="i"
                  xs="12"
                  sm="12"
                  md="10"
                  lg="6"
                  xl="4"
                  xxl="4"
                >
                  <v-card height="400">
                    <v-card-title>{{item.raw.name}}</v-card-title>
                    <video v-if="videos.videos.length > 0" ref="videoRef" :src=item.raw.location type="video/mp4" id="video-container" width="100%" controls></video>
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

      </v-col>
    </v-row>
  </v-container>
</template>
