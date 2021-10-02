<template>
  <b-card
    no-body
    bg-variant="light"
    align="left"
    border-variant="dark"
  >
    <b-card-header>
      <b-container class="m-0 p-0">
        <b-row align-v="center">
          <b-col cols="7">
            {{bbsThread.title}}
          </b-col>
          <b-col cols="5">
            <p class="m-0 p-0"><small class="text-muted">投稿者:{{bbsThread.name}}</small></p>
            <p class="m-0 p-0"><small class="text-muted">投稿日時:{{getDateString(bbsThread.createdAt)}}</small></p>
            <p class="m-0 p-0" v-if="isRenderUpdatedAt(bbsThread)"><small class="text-muted">更新日時:{{getDateString(bbsThread.updatedAt)}}</small></p>
          </b-col>
        </b-row>
      </b-container>
    </b-card-header>
    <b-card-body>
      <b-card-text>{{bbsThread.body}}</b-card-text>
    </b-card-body>
  </b-card>
</template>

<script lang="ts">
import { defineComponent } from '@vue/composition-api'
import { BbsThread } from '@/lib/interface/bbs'

export default defineComponent({
  props: {
    bbsThread: {
      type: Object,
      required: true
    }
  },
  setup () {
    const getDateString = (timestamp: number): string => {
      const dt = new Date(timestamp)
      return dt.getFullYear() + '年' + (dt.getMonth() + 1) + '月' + dt.getDate() + '日 ' + dt.getHours() + '時' + dt.getMinutes() + '分'
    }

    const isRenderUpdatedAt = (bbsThread: BbsThread): boolean => {
      return bbsThread.createdAt !== bbsThread.updatedAt
    }

    return {
      getDateString,
      isRenderUpdatedAt
    }
  }
})
</script>
