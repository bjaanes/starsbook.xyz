<script setup lang="ts">
import NftInfo from "@/components/NftInfo.vue";
import DataText from "@/components/DataText.vue";
import { useRoute } from 'vue-router'
import {onMounted, ref, watch} from "vue";

const route = useRoute()

const nftTitle = ref("");
const nftImg = ref("")
const attributeTags = ref([]);
const collectionName = ref("");

const rarityScore = ref("");
const rarityRank = ref("");
const KRV = ref("");
const KARV = ref("");

const updateNftView = async () => {
  const projectShortName = String(route.name)
  const nftId = route.params.id;
  const projectInfo = await fetch(`/${projectShortName}/min_project.json`).then(res => res.json());
  const nft = await fetch(`/${projectShortName}/nfts/${nftId}.json`).then(res => res.json());

  nftTitle.value = nft.name;
  nftImg.value = `/${projectShortName}/imgs/${nft.img}`;
  attributeTags.value = nft.attributes
      .filter((attr: { ignoredForDisplay: boolean }) => !attr.ignoredForDisplay);
  collectionName.value = projectInfo.name;
  rarityScore.value = parseInt(nft.rarityScore) + '';
  rarityRank.value = nft.rarityRank;
  KRV.value = parseInt(nft.prices.KRV) + '';
  KARV.value = `${parseInt(nft.prices.KARV_NOW)} - ${parseInt(nft.prices.KARV_FUTURE)}`
}

watch(() => route.params, updateNftView);
onMounted(updateNftView)


</script>

<template>
  <div class="nft-view-container">
    <nft-info :title="nftTitle" :img="nftImg" :attributeTags="attributeTags" :collection-name="collectionName"></nft-info>
    <div class="data-container">
      <data-text title="Starvalue" value="Coming soon" denomination="">
        Coming soon
      </data-text>
      <data-text title="Rarity" :value="rarityScore" denomination="">
        Rarity score. More coming soon...
      </data-text>
      <data-text title="Rarity rank" :value="rarityRank" denomination="">
        The rarity rank in the collection.
      </data-text>
      <data-text title="KRV" :value="KRV" denomination="STARS">
        The Kryoten Rarity Valuation.<br>
        This is an algorithm designed by Kryoten to come up with a PRE-Market Valuation based on solely rarity score. This is the simplest metric and as valued by Kryoten, should be the lowest expected price to trade a Fren for given there are no official market values or floor prices yet.
      </data-text>
      <data-text title="KARV" :value="KARV" denomination="STARS">
        The Kryoten Accelerated Rarity Valuation.<br>
        This is the companion algorithm to give a more speculative price based on other NFT projects, also accounting for value and growth in floor prices, with similar supply. This number is a range speculating todays KARV, and a future price based on a time span of 2 years.
      </data-text>
    </div>

  </div>
</template>

<style scoped>
.nft-view-container {
  display: flex;
  padding: 10px 50px;
}

@media only screen and (max-width: 800px) {
  .nft-view-container {
    flex-wrap: wrap;
    gap: 50px;
  }
}

.data-container {
  display: flex;
  flex-wrap: wrap;
  gap: 50px;
}

.data-container * {
  margin-left: 25px;
}

</style>
