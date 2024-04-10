<script setup>
  const emit = defineEmits(['rail']);
</script>

<template>
  <v-app-bar
    color="teal-darken-4"
    image="https://picsum.photos/1920/1080"
  >
    <template v-slot:image>
      <v-img gradient="to top right, rgba(19,84,122,.8), rgba(128,208,199,.8)"></v-img>
    </template>
    <template v-slot:prepend>
      <v-app-bar-nav-icon @click="emit('rail')"></v-app-bar-nav-icon>
    </template>
    <v-app-bar-title>Title</v-app-bar-title>
    <v-spacer></v-spacer>
    <v-btn icon><v-icon>mdi-magnify</v-icon></v-btn>
    <v-btn icon><v-icon>mdi-dots-vertical</v-icon></v-btn>

  </v-app-bar>
</template>
