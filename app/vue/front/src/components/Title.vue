<template>
  <h1 class="m-4">{{title}}</h1>
</template>

<script lang="ts">
import { defineComponent } from '@vue/composition-api'
export default defineComponent({
  props: {
    title: {
      type: String,
      required: true
    }
  }
})
</script>
