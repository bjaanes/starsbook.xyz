<script setup lang="ts">
import {ref} from "vue";
import { SearchClient as TypesenseSearchClient } from "typesense";
import {NSelect} from "naive-ui";
import type {SelectOption} from "naive-ui";
import {DatabaseSearch24Regular, Search24Regular} from '@vicons/fluent'
import { useRouter } from 'vue-router'

let client = new TypesenseSearchClient({
  nodes: [{
    host: 'starsbook-typesense.gjermund.tech',
    port: 443,
    protocol: 'https'
  }],
  apiKey: 'lwf86ywndIydr1P4qGvfy9VaSd6ORGoN',
  connectionTimeoutSeconds: 2
});

const router = useRouter()

const loading = ref(false);
const options = ref<SelectOption[]>([]);
const value = ref(null);
const show = ref(false);

async function handleSearch(query: string) {
  if (!query.length) {
    options.value = []
    return
  }
  loading.value = true
  const res = await client.collections("nfts").documents().search({
    q: query,
    query_by: "name,collectionName",
    sort_by: "_text_match:desc,nftId:asc",
    filter_by: "hidden:=false",
  }, {})
  loading.value = false;

  if (res.hits) {
    options.value = res.hits.map((hit) => {
      const document = hit.document as { id: string; name: string }
      return {
        value: document.id,
        label: document.name,
      }
    })
  }

}

function handleSelect(selectedValue: string) {
  const split = selectedValue.split('_')
  router.push(`/${split[0]}/${split[1]}`);
}

</script>

<template>
  <n-select
      v-model:value="value"
      v-model:show="show"
      filterable
      placeholder="Search NFTs"
      :options="options"
      :loading="loading"
      clearable
      remote
      @search="handleSearch"
      @update:value="handleSelect"
      style="width: 365px;"
  >
    <template #arrow>
      <transition name="slide-left">
          <DatabaseSearch24Regular v-if="!show" />
          <Search24Regular v-else />
      </transition>
    </template>
  </n-select>
</template>

<style scoped>
:deep(.slide-left-enter-active),
:deep(.slide-left-leave-active) {
  transition: transform 0.3s ease, opacity 0.3s ease;
}

:deep(.slide-left-enter-from),
:deep(.slide-left-leave-to) {
  position: absolute;
  opacity: 0;
}

:deep(.slide-left-enter-from) {
  transform: translateX(-10px);
}

:deep(.slide-left-leave-to) {
  transform: translateX(10px);
}
</style>
