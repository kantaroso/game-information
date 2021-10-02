<template>
    <header>
        <div>
            <b-navbar toggleable type="dark" variant="dark">
                <b-navbar-brand href="#">
                    <strong>PV:{{ pv }}</strong>
                </b-navbar-brand>

                <b-navbar-toggle target="navbar-toggle-collapse">
                <template v-slot:default="{ expanded }">
                    <b-icon v-if="expanded" icon="chevron-bar-up"></b-icon>
                    <b-icon v-else icon="chevron-bar-down"></b-icon>
                </template>
                </b-navbar-toggle>

                <b-collapse id="navbar-toggle-collapse" is-nav>
                <b-navbar-nav class="ml-auto">
                    <b-nav-item href="/maker/list">メーカー一覧</b-nav-item>
                    <b-nav-item href="/bbs">掲示板</b-nav-item>
                    <b-nav-item href="#">メニュー3</b-nav-item>
                </b-navbar-nav>
                </b-collapse>
            </b-navbar>
        </div>
    </header>
</template>

<script lang="ts">
import { defineComponent, reactive, toRefs, onMounted } from '@vue/composition-api'
import axios from 'axios'
interface ReactiveData {
  pv: string;
}
export default defineComponent({
  setup () {
    const state: ReactiveData = reactive({
      pv: '-'
    })

    onMounted(() => {
      axios.get(`${process.env.VUE_APP_API_ORIGIN}/common/pv`).then(
        res => {
          state.pv = String(res.data.pv)
        }
      ).catch(
        error => console.log(error)
      )
    })

    return {
      ...toRefs(state)
    }
  }
})
</script>
