<template>
  <b-card bg-variant="dark" text-variant="white" class="text-center">
    <b-card-text>{{body}}</b-card-text>
    <template #footer>
      <small class="text-muted">{{getPublishAtString(publishAt)}}</small>
    </template>
  </b-card>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'

@Component
export default class Content extends Vue {
  @Prop({ type: String, required: true })
  body!: string

  @Prop({ type: Number, required: true })
  publishAt!: number

  getPublishAtString (publishAt: number): string {
    const dt = new Date(publishAt)
    return dt.getFullYear() + '年' + dt.getMonth() + '月' + dt.getDay() + '日 ' + dt.getHours() + '時' + dt.getMinutes() + '分'
  }
}
</script>
