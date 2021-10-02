<template>
  <div>
    <b-alert
      :show="dismissCountDownSuccess"
      variant="success"
      dismissible
      @dismissed="dismissCountDownSuccess=0"
      @dismiss-count-down="countDownChangedSuccess"
      >
      投稿しました
    </b-alert>
    <b-alert
      :show="dismissCountDownError"
      variant="danger"
      dismissible
      @dismissed="dismissCountDownError=0"
      @dismiss-count-down="countDownChangedError"
      >
      投稿に失敗しました
    </b-alert>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'

@Component
export default class Alert extends Vue {
  dismissSecs = 3
  dismissCountDownSuccess = 0
  dismissCountDownError = 0

  countDownChangedSuccess (dismissCountDown: number) {
    this.dismissCountDownSuccess = dismissCountDown
  }

  countDownChangedError (dismissCountDown: number) {
    this.dismissCountDownError = dismissCountDown
  }

  renderSuccess () {
    this.dismissCountDownSuccess = this.dismissSecs
  }

  renderError () {
    this.dismissCountDownError = this.dismissSecs
  }
}
</script>
