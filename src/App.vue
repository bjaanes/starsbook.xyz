<script setup lang="ts">
import { ref } from 'vue'
import { RouterView } from 'vue-router'
import { darkTheme, NConfigProvider, NSelect } from 'naive-ui'
import type {SelectOption} from "naive-ui";

const searchTerm = ref("");
const options = ref<SelectOption[]>([])

function handleSearch(query: string) {
  if (!query.length) {
    options.value = []
    return
  }
}
</script>

<template>
  <n-config-provider :theme="darkTheme">
    <div id="app-container">
      <header>
        <img src="@/assets/logo.svg">
        <!--<n-select
            v-model:value="chosenNft"
            placeholder="Search for any stargaze NFT"
            autosize
            :options="options"
            @search="handleSearch"
            style="min-width: 365px;" />-->
      </header>

      <RouterView />
    </div>
  </n-config-provider>

</template>

<style>
@import 'assets/base.css';

header {
  display: flex;
  justify-content: space-between;
  margin: 30px;
}

#app-container {
  max-width: 1200px;
  margin: auto;
}
</style>
