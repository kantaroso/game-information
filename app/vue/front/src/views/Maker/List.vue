<template>
  <div class="index">
    <FrontHeader />
    <PageTitle title="メーカー一覧"/>
    <h2>メーカー一覧</h2>
    <!--
    <b-list-group v-if="items">
      <b-list-group-item v-for="item in items.videos" :key="item.ID">
      </b-list-group-item>
    </b-list-group>
    <h2 v-else><b-spinner label="Loading..."></b-spinner></h2>
    -->
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
  mounted () {
    axios.get('http://localhost/maker/august').then(
      res => {
        this.items = res.data
      }
    ).catch(
      error => console.log(error)
    )
  }
}
</script>
