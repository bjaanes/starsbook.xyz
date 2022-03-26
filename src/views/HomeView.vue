<script setup lang="ts">
import CollectionInfo from "@/components/CollectionInfo.vue";
import projects from "@/generated/projects"

</script>

<template>
  <div>
    <h1>Welcome to Starsbook!</h1>

    <div class="collection-container">
      <CollectionInfo v-for="(project, index) in projects" :key="index" :collection-name="project.name" :img="project.img"></CollectionInfo>
    </div>

  </div>
</template>

<style scoped>
.collection-container {
  display: flex;
  gap: 50px;
  flex-wrap: wrap;
}
</style>