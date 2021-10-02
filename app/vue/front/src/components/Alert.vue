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
import { defineComponent, reactive, toRefs } from '@vue/composition-api'
interface ReactiveData {
  dismissCountDownSuccess: number;
  dismissCountDownError: number;
}
export default defineComponent({
  setup () {
    const dismissSecs = 3
    const state: ReactiveData = reactive({
      dismissCountDownSuccess: 0,
      dismissCountDownError: 0
    })

    const countDownChangedSuccess = (dismissCountDown: number) => {
      state.dismissCountDownSuccess = dismissCountDown
    }

    const countDownChangedError = (dismissCountDown: number) => {
      state.dismissCountDownError = dismissCountDown
    }

    const renderSuccess = () => {
      state.dismissCountDownSuccess = dismissSecs
    }

    const renderError = () => {
      state.dismissCountDownError = dismissSecs
    }

    return {
      ...toRefs(state),
      countDownChangedSuccess,
      countDownChangedError,
      renderSuccess,
      renderError
    }
  }
})
</script>
