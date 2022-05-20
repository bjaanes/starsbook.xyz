<script setup lang="ts">
import {onMounted, ref, watch, computed} from "vue";
import {useRoute, useRouter} from "vue-router";
import { NPagination, NSpin, NSelect, NFormItem } from "naive-ui";
import NftInfo from "@/components/NftInfo.vue";
import {SearchClient as TypesenseSearchClient} from "typesense";

const route = useRoute()
const router = useRouter()

const loading = ref(false);
const comingSoon = ref(false);
const projectName = ref("");
const numberOfNfts = ref(0);
const page = ref(1);
const nfts = ref<Array<{title: string, img: string, id: string, rarityRank: number}>>([])
const sortBy = ref("nftId");
const sortOptions = computed(() => {
  const o = [
    {
      label: "ID",
      value: "nftId",
    }
  ]

  if (!comingSoon.value || import.meta.env.VITE_SEE_ALL === "true") {
    o.push(
        {
          label: "Rank/Rarity",
          value: "rarityRank"
        }
    )
  }

  return o
})

let projectShortName = "";
let client = new TypesenseSearchClient({
  nodes: [{
    host: 'starsbook-typesense.gjermund.tech',
    port: 443,
    protocol: 'https'
  }],
  apiKey: 'UzmYK08FCqsPOkMjz35eQEqsFPDahIOz',
  connectionTimeoutSeconds: 2
});

const updateCollectionView = async () => {
  projectShortName = String(route.name).split('-')[0]
  const projectInfo = await fetch(`https://starsbook-assets.storage.googleapis.com/${projectShortName}/min_project.json`).then(res => res.json());
  projectName.value = projectInfo.name;
  numberOfNfts.value = projectInfo.numberOfNfts;

  comingSoon.value = projectInfo.comingSoon;

  await loadNfts(1);
}

const changePage = async (page: number) => {
  await loadNfts(page, sortBy.value);
}

const updateSortBy = async(sortBy: string) => {
  await loadNfts(1, sortBy)
}

const loadNfts = async (page: number, _sortBy?: string) => {
  loading.value = true;
  nfts.value = []

  const sortBy = _sortBy ? _sortBy : "nftId";

  const res = await client.collections("nfts").documents().search({
    q: projectShortName,
    query_by: "collectionShortName",
    sort_by: `${sortBy}:asc`,
    page: page,
    per_page: 9,
  }, {})

  if (!res.hits) {
    loading.value = false;
    return;
  }

  for (let hit of res.hits) {
    const document = hit.document as {nftId: string, name: string, imageFileName: string, rarityRank: number}
    nfts.value.push({
      title: document.name,
      img: `https://starsbook-assets.storage.googleapis.com/${projectShortName}/imgs_min/${document.imageFileName}`,
      id: document.nftId,
      rarityRank: document.rarityRank,
    })
  }

  loading.value = false;
}

watch(() => route.params, updateCollectionView);
onMounted(updateCollectionView)

function selectNft(id: string) {
  router.push(`/${projectShortName}/${id}`);
}

</script>

<template>
  <n-spin :show="loading">
    <h1>{{projectName}}</h1>

    <n-form-item label="Sort by">
      <n-select @update:value="updateSortBy" v-model:value="sortBy" :options="sortOptions" />
    </n-form-item>


    <div class="nft-container">
      <NftInfo class="nft" v-for="(nft, index) of nfts" :key="index" :title="nft.title" :img="nft.img" :rank="nft.rarityRank.toString()" @click="selectNft(nft.id)" :coming-soon="comingSoon"></NftInfo>
    </div>

    <div class="paginator-container">
      <n-pagination v-on:update:page="changePage" v-model:page="page" :page-count="Math.ceil(numberOfNfts / 9)" />
    </div>

  </n-spin>
</template>

<style scoped>
.nft-container {
  display: flex;
  justify-content: space-between;
  gap: 50px;
  flex-wrap: wrap;
}

.nft {
  cursor: pointer;
}

.paginator-container {
  width: fit-content;
  margin-top: 25px;
  margin-left: auto;
  margin-right: auto;
}
</style>