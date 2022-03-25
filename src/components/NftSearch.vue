<script setup lang="ts">
import {ref} from "vue";
import { NSelect } from 'naive-ui'
import type {SelectOption} from "naive-ui";

const searchTerm = ref("");
const loading = ref(false)
const options = ref<SelectOption[]>([])

function handleSearch(query: string) {
  if (!query.length) {
    options.value = []
    return
  }

  loading.value = true

  window.setTimeout(() => {
    options.value = [
      {
        label: 'Drive My Car',
        value: 'song1'
      }
    ];
    loading.value = false
  }, 1000)

}
</script>

<template>
  <n-select
      v-model:value="chosenNft"
      placeholder="Search for any stargaze NFT"
      filterable
      clearable
      remote
      autosize
      :loading="loading"
      :options="options"
      @search="handleSearch"
      style="width: 365px;" />
</template>

<style scoped>
</style>
