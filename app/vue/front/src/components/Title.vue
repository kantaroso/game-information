<template>
  <h1 class="m-4">{{title}}</h1>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'
@Component
export default class Title extends Vue {
  @Prop({ type: String, required: true })
  title!: string
}
</script>
