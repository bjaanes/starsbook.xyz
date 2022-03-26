<script setup lang="ts">
import CollectionInfo from "@/components/CollectionInfo.vue";
import projects from "@/generated/projects"
import {useRouter} from "vue-router";

const router = useRouter()

function selectProject(projectShortName: string) {
  router.push(`/${projectShortName}`);
}
</script>

<template>
  <div>
    <h1>Welcome to Starsbook!</h1>

    <div class="collection-container">
      <CollectionInfo class="collection" v-for="(project, index) in projects" :key="index" :collection-name="project.name" :img="project.img" @click="selectProject(project.shortName)"></CollectionInfo>
    </div>

  </div>
</template>

<style scoped>
.collection-container {
  display: flex;
  justify-content: space-between;
  gap: 50px;
  flex-wrap: wrap;
}

.collection {
  cursor: pointer;
}
</style>