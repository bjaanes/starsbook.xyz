<script setup lang="ts">
import { NPopover } from "naive-ui";

defineProps<{
  title: string;
  value: string;
  denomination: string;
  description: string;
}>()
</script>

<template>
  <div class="data-text-container">

    <div class="title">
      {{title}}
      <n-popover trigger="hover">
        <template #trigger>
          <img class="qm" src="@/assets/qm.svg">
        </template>
        <span>{{description}}</span>
      </n-popover>
    </div>
    <div class="data">
      <span class="value">{{value}}</span> <span class="denomination">{{denomination}}</span>
    </div>
  </div>
</template>

<style scoped>
.data-text-container {
  display: inline-block;
  height: fit-content;
}

.qm {
  position: absolute;
  top: -5px;
  right: -25px;
}

.title {
  display: inline-block;
  position: relative;
  font-size: 36px;
}

.data {
  margin-top: 8px;
  font-family: 'Space Grotesk', sans-serif;
}

.value {
  font-size: 30px;
}

.denomination {
  font-size: 22px;
  color: #5375EE;
}
</style>
