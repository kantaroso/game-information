<template>
  <div class="index">
    <FrontHeader />
    <PageTitle title="オーガスト"/>
    <h2>動画一覧</h2>
    <div v-if="videos">
      <b-list-group>
        <b-list-group-item v-for="video in videos" :key="video.ID">
          <b-embed
            type="iframe"
            aspect="16by9"
            :src="`https://www.youtube.com/embed/${video.ID}`"
            allowfullscreen
          ></b-embed>
        </b-list-group-item>
      </b-list-group>
    </div>
    <div v-else>
      <h3><b-spinner label="Loading..."></b-spinner></h3>
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
  info = null
  videos = null
  mounted () {
    axios.get('http://localhost/maker/detail/august').then(
      res => {
        res.status
        this.info = res.data
      }
    ).catch(
      error => console.log(error)
    )
    axios.get('http://localhost/maker/videos/august').then(
      res => {
        this.videos = res.data
      }
    ).catch(
      error => console.log(error)
    )

  }
}
</script>
