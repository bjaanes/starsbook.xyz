<script setup lang="ts">
defineProps<{
  img: string;
  collectionName: string,
}>()
</script>

<template>
  <div class="collection-info-container">
    <img class="collection-image" :src="img">
    <div class="title">{{ collectionName }}</div>
  </div>
</template>

<style scoped>
.collection-info-container {
  background-color: #141416;
  border-radius: 20px;
  width: 250px;
  padding: 10px;
}

.collection-image {
  height: 250px;
  width: 250px;
  border-radius: 16px;
  margin-left: auto;
  margin-right: auto;
  display: block;
}

.title {
  margin-top: 10px;
  font-size: 24px;
}

</style>
