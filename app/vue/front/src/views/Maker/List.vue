<template>
  <div class="index">
    <FrontHeader />
    <PageTitle title="メーカー一覧"/>
    <div v-if="nolist.length">
      <h2>データなし</h2>
    </div>
    <div v-else-if="items">
      <b-container>
        <b-list-group v-for="item in items" :key="item.code">
          <b-list-group-item>
            <b-link :href="`/maker/detail/${item.code}`">{{item.name}}</b-link>
          </b-list-group-item>
        </b-list-group>
      </b-container>
    </div>
    <div v-else>
      <h2><b-spinner label="Loading..."></b-spinner></h2>
    </div>
    <FrontFooter />
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, toRefs, onMounted } from '@vue/composition-api'
import FrontHeader from '@/components/Header.vue'
import FrontFooter from '@/components/Footer.vue'
import PageTitle from '@/components/Title.vue'
import axios from 'axios'
interface ReactiveData {
  items: any; // eslint-disable-line @typescript-eslint/no-explicit-any
  nolist: boolean;
}
export default defineComponent({
  components: {
    FrontHeader,
    FrontFooter,
    PageTitle
  },
  setup () {
    const state: ReactiveData = reactive({
      items: {},
      nolist: false
    })

    const render = () => {
      axios.get(`${process.env.VUE_APP_API_ORIGIN}/maker/list`, { timeout: 5000 }).then(
        res => {
          if (!res.data.length) {
            state.nolist = true
          } else {
            state.items = res.data
          }
        }
      ).catch(
        error => {
          alert('データの取得に失敗しました')
          console.log(error)
        }
      )
    }

    onMounted(() => {
      render()
    })

    return {
      ...toRefs(state)
    }
  }
})
</script>
