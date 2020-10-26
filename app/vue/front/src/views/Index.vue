<template>
  <div class="index">
    <FrontHeader />
    <h1>動画一覧</h1>
    <b-list-group v-if="items">
      <b-list-group-item v-for="item in items.movies" :key="item.ID">
        <b-embed
          type="iframe"
          aspect="16by9"
          :src="`https://www.youtube.com/embed/${item.ID}`"
          allowfullscreen
        ></b-embed>
      </b-list-group-item>
    </b-list-group>
    <h2 v-else><b-spinner label="Loading..."></b-spinner></h2>
    <FrontFooter />
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import FrontHeader from '@/components/Header.vue' // @ is an alias to /src
import FrontFooter from '@/components/Footer.vue' // @ is an alias to /src
import axios from 'axios'
@Component({
  components: {
    FrontHeader,
    FrontFooter
  }
})
export default class Index extends Vue {
  items = null
  mounted () {
    axios.get('http://localhost/maker/list').then(
      res => {
        this.items = res.data
      }
    ).catch(
      error => console.log(error)
    )
  }
}
</script>
