<script setup lang="ts">
import {RouterView, useRouter} from 'vue-router'
import {darkTheme, NConfigProvider} from 'naive-ui'
import type {GlobalThemeOverrides} from 'naive-ui';
import NftSearch from "@/components/NftSearch.vue";

const router = useRouter()

const themeOverrides: GlobalThemeOverrides = {
  common: {

  }
}

function goHome() {
  router.push("/");
}
</script>

<template>
  <n-config-provider class="config-container" :theme-overrides="themeOverrides" :theme="darkTheme">
    <div id="app-container">
      <header>
        <img @click="goHome" src="@/assets/logo.svg">
        <nft-search></nft-search>
      </header>

      <RouterView />

      <footer>
        ❤️ Built with love ❤️<br />
        Starsbook donation address:
        juno1vrgrym3nnazzvkgwzxp7tvq9es0awv00zsz9fg7e9lp837kul62q6qds2s
      </footer>
    </div>
  </n-config-provider>

</template>

<style>
@import 'assets/base.css';

.config-container,#app-container {
  height: 100%;
}

header {
  display: flex;
  justify-content: space-between;
  margin: 30px;
}

header img {
  cursor: pointer;
}

#app-container {
  max-width: 1200px;
  margin: auto;
  display: flex;
  flex-direction: column;
}

footer {
  margin-top: auto;
  padding-top: 50px;
  text-align: center;
  overflow-wrap: anywhere;
}
</style>
