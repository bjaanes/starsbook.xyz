<script setup lang="ts">
import { RouterView } from 'vue-router'
import { darkTheme, NConfigProvider } from 'naive-ui'
import NftSearch from "@/components/NftSearch.vue";
</script>

<template>
  <n-config-provider :theme="darkTheme">
    <div id="app-container">
      <header>
        <img src="@/assets/logo.svg">
        <nft-search></nft-search>
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
