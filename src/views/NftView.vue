<script setup lang="ts">
import NftInfo from "@/components/NftInfo.vue";
import DataText from "@/components/DataText.vue";
import { useRoute } from 'vue-router'
import {onMounted, ref} from "vue";

const route = useRoute()

const nftTitle = ref("");
const nftImg = ref("")
const attributeTags = ref([]);
const collectionName = ref("");

const rarityScore = ref("");
const rarityRank = ref("");
const KRV = ref("");
const KARV = ref("");

onMounted(async () => {
  const projectShortName = String(route.name)
  const nftId = route.params.id;
  const projectInfo = await fetch(`/${projectShortName}/min_project.json`).then(res => res.json());
  const nft = await fetch(`/${projectShortName}/nfts/${nftId}.json`).then(res => res.json());

  nftTitle.value = nft.name;
  nftImg.value = `/${projectShortName}/imgs/${nft.img}`;
  attributeTags.value = nft.attributes
      .filter((attr: { rarityScore: number }) => attr.rarityScore != 0);
  collectionName.value = projectInfo.name;
  rarityScore.value = parseInt(nft.rarityScore) + '';
  rarityRank.value = nft.rarityRank;
  KRV.value = parseInt(nft.prices.KRV) + '';
  KARV.value = `${parseInt(nft.prices.KARV_NOW)} - ${parseInt(nft.prices.KARV_FUTURE)}`
})
</script>

<template>
  <div class="nft-view-container">
    <nft-info :title="nftTitle" :img="nftImg" :attributeTags="attributeTags" :collection-name="collectionName"></nft-info>
    <div class="data-container">
      <data-text title="Starvalue" value="Coming soon" denomination="" description="TODO"></data-text>
      <data-text title="Rarity" :value="rarityScore" denomination="" description="hoppsi"></data-text>
      <data-text title="Rarity rank" :value="rarityRank" denomination="" description="hoppsi"></data-text>
      <data-text title="KRV" :value="KRV" denomination="STARS" description="hoppsi"></data-text>
      <data-text title="KARV" :value="KARV" denomination="STARS" description="hoppsi"></data-text>
    </div>

  </div>
</template>

<style scoped>
.nft-view-container {
  display: flex;
  padding: 10px 50px;
}

.data-container {
  display: flex;
  flex-wrap: wrap;
}

.data-container * {
  margin-left: 25px;
  margin-right: 25px;
}

</style>
