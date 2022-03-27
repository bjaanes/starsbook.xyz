<script setup lang="ts">
import {onMounted, ref, watch} from "vue";
import {useRoute, useRouter} from "vue-router";
import { NPagination, NSpin } from "naive-ui";
import NftInfo from "@/components/NftInfo.vue";

const route = useRoute()
const router = useRouter()

const loading = ref(false);
const projectName = ref("");
const numberOfNfts = ref(0);
const page = ref(1);
const nfts = ref<Array<{title: string, img: string, id: number, rarityRank: string}>>([])

let projectShortName = "";

const updateCollectionView = async () => {
  projectShortName = String(route.name).split('-')[0]
  const projectInfo = await fetch(`/${projectShortName}/min_project.json`).then(res => res.json());
  projectName.value = projectInfo.name;
  numberOfNfts.value = projectInfo.numberOfNfts;

  await loadNfts(1, 9);
}

const loadNfts = async (from: number, to: number) => {

  loading.value = true;
  nfts.value = []
  for (let i = from; i <= to; ++i) {
    const nft = await fetch(`/${projectShortName}/nfts/${i}.json`).then(res => res.json());
    nfts.value.push({
      title: nft.name,
      img: `/${projectShortName}/imgs/min_${nft.img}`,
      id: i,
      rarityRank: nft.rarityRank,
    })
  }
  loading.value = false;
}

watch(() => route.params, updateCollectionView);
onMounted(updateCollectionView)

function selectNft(id: number) {
  router.push(`/${projectShortName}/${id}`);
}

async function selectPage(page: number) {
  const from = (page-1) * 9 + 1;
  let to = page * 9;

  if (to > numberOfNfts.value) {
    to = numberOfNfts.value
  }

  await loadNfts(from, to);
}

</script>

<template>
  <n-spin :show="loading">
    <h1>{{projectName}}</h1>

    <div class="nft-container">
      <NftInfo class="nft" v-for="(nft, index) of nfts" :key="index" :title="nft.title" :img="nft.img" :rank="nft.rarityRank" @click="selectNft(nft.id)"></NftInfo>
    </div>

    <div class="paginator-container">
      <n-pagination v-on:update:page="selectPage" v-model:page="page" :page-count="Math.ceil(numberOfNfts / 9)" />
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