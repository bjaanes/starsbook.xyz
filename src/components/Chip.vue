<script setup lang="ts">
import { computed } from 'vue'

const props = defineProps<{
  text: string;
  small: boolean;
  color: string;
}>()

const chipClasses = computed(() => {
  const cssClasses = {} as any;

  if (props.small) {
    cssClasses['small-chip'] = true
  } else {
    cssClasses['big-chip'] = true
  }

  cssClasses[props.color] = true;

  console.log(cssClasses);

  return cssClasses;
})
</script>

<template>
  <div class="chip" :class="chipClasses">{{ text }}</div>
</template>

<style scoped>
.chip {
  border-radius: 8px;
  display: inline-block;
}

.big-chip {
  height: 29px;
  line-height: 29px;
  padding-left: 6px;
  padding-right: 6px;
  font-size: 14px;
}

.small-chip {
  height: 23px;
  line-height: 23px;
  font-size: 12px;
  padding-left: 6px;
  padding-right: 6px;
}

.blue {
  background: rgba(83, 117, 238, 0.2);
  color: #5375EE;
}

.pink {
  background: rgba(254, 157, 250, 0.2);
  color: #FE9DFA;
}

.yellow {
  background: rgba(254, 215, 157, 0.2);
  color: #FED79D;
}

.green {
  background: rgba(194, 254, 157, 0.2);
  color: #C2FE9D;
}

.teal {
  background: rgba(157, 231, 254, 0.2);
  color: #9DE7FE;
}

.orange {
  background: rgba(254, 180, 157, 0.2);
  color: #FEB49D;
}

.red {
  background: rgba(254, 157, 157, 0.2);
  color: #FE9D9D;
}
</style>
