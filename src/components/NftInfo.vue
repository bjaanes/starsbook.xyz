<script setup lang="ts">
import Chip from "@/components/Chip.vue";
import {computed} from "vue";

const COLORS = [
    "blue",
    "pink",
    "yellow",
    "green",
    "teal",
    "orange",
    "red",
];

const props = defineProps<{
  img: string;
  title: string;
  rank: string;
  collectionName?: string,
  attributeTags?: Array<{type: string, value: string}>;
  comingSoon?: boolean,
}>()

const attributeChips = computed(() => {
  let currentIndex = 0;

  if (props.attributeTags) {
    return props.attributeTags.map(at => {
      if (currentIndex >= COLORS.length) {
        currentIndex = 0;
      }

      const color = COLORS[currentIndex];
      currentIndex++;

      return {
        title: at.value,
        type: at.type,
        color
      }
    })
  } else {
    return [];
  }
})
</script>

<template>
  <div class="nft-info-container">
    <img v-if="img" class="nft-image" :src="img">
    <div class="summary-chips">
      <chip v-if="collectionName" :text="collectionName" :small="false" color="blue"></chip>
      <chip v-if="!comingSoon" :text="'Rank: ' + rank" :small="false" color="pink"></chip>
    </div>
    <div class="title">{{ title }}</div>
    <div v-if="attributeChips.length > 0" class="traits-title">Traits</div>
    <div v-if="attributeChips.length > 0" class="traits-chips">
      <chip v-for="(ac, index) in attributeChips" :key="index" :text="ac.title" :small="true" :color="ac.color" :title="ac.type"></chip>
    </div>
  </div>
</template>

<style scoped>
.nft-info-container {
  background-color: #141416;
  border-radius: 20px;
  width: 250px;
  padding: 10px;
}

.nft-image {
  height: 250px;
  width: 250px;
  border-radius: 16px;
  margin-left: auto;
  margin-right: auto;
  display: block;
}

.summary-chips {
  margin-top: 10px;
}

.summary-chips * {
  margin-left: 5px;
  margin-right: 5px;
}

.summary-chips :first-child {
  margin-left: 0;
}

.summary-chips :last-child {
  margin-right: 0;
}

.title {
  margin-top: 10px;
  font-size: 24px;
}

.traits-title {
  margin-top: 10px;
  font-size: 18px;
}

.traits-chips {

}

.traits-chips * {
  margin-top: 10px;
  margin-right: 10px;
}

</style>
