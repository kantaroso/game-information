<template>
  <div class="index">
    <FrontHeader />
    <PageTitle title="メーカー一覧"/>
    <div v-if="nolist.length">
      <h2>データなし</h2>
    </div>
    <div v-else-if="items">
      <b-list-group v-for="item in items" :key="item.code">
        <b-list-group-item>
          <b-link :href="`/maker/detail/${item.code}`">{{item.name}}</b-link>
        </b-list-group-item>
      </b-list-group>
    </div>
    <div v-else>
      <h2><b-spinner label="Loading..."></b-spinner></h2>
    </div>
    <FrontFooter />
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import FrontHeader from '@/components/Header.vue'
import FrontFooter from '@/components/Footer.vue'
import PageTitle from '@/components/Title.vue'
import axios from 'axios'
@Component({
  components: {
    FrontHeader,
    FrontFooter,
    PageTitle
  }
})
export default class Index extends Vue {
  items = null
  nolist = false
  mounted () {
    axios.get('http://localhost/maker/list').then(
      res => {
        if (!res.data.length) {
          this.nolist = true
        } else {
          this.items = res.data
        }
      }
    ).catch(
      error => console.log(error)
    )
  }
}
</script>
