<script setup>
  const props =defineProps(['rail','cameras','location']);
  const emit = defineEmits(['load-cameras','select-camera']);
</script>
<template>
  <v-navigation-drawer
    expand-on-hover
    permanent
    :rail="rail"
  >
    <v-list>
      <v-list-item class="pl-2">
        <v-icon icon="mdi-cctv" size="x-large" ></v-icon>
      </v-list-item>
    </v-list>

    <v-divider></v-divider>

    <v-list rounded="lg" nav>
      <v-list-item
        v-for="(c, i) in cameras"
        :title="c"
        prepend-icon="mdi-folder"
        link
        @click="emit('select-camera', c)"
        :key="i"
        >
      </v-list-item>

      <v-divider class="my-2"></v-divider>

      <v-list-item
        color="grey-lighten-4"
        link
        title="UPDATE"
        @click="emit('load-cameras')"
      ></v-list-item>
    </v-list>

  </v-navigation-drawer>
</template>

