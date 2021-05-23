<template>
  <div class="index">
    <FrontHeader />
    <div v-if="info">
      <PageTitle :title="`${info.name}`"/>
      <div v-if="info.ohp" class="m-3">
        <b-link target="_blank" :href="`${info.ohp}`">公式ホームページ</b-link>
      </div>
    </div>
    <div v-else>
      <h3><b-spinner label="Loading..."></b-spinner></h3>
    </div>
    <b-container>
      <b-row>
        <b-col cols="8">
          <div v-if="videos">
            <b-list-group v-if="videos.length > 0">
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
            <b-spinner label="Loading..."></b-spinner>
          </div>
        </b-col>
        <b-col cols="4">
          <a v-if="info && info.twitter_name" class="twitter-timeline" data-lang="ja" data-theme="dark" :href="`https://twitter.com/${info.twitter_name}`">Tweets by {{info.twitter_name}}</a>
          <div v-else>
            <h3><b-spinner label="Loading..."></b-spinner></h3>
          </div>
        </b-col>
      </b-row>
    </b-container>

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
  code = this.$route.params.path
  mounted () {
    axios.get(`${process.env.VUE_APP_API_ORIGIN}/maker/detail/${this.code}`, { timeout: 5000 }).then(
      res => {
        this.info = res.data
      }
    ).catch(
      error => {
        alert('データの取得に失敗しました')
        console.log(error)
      }
    )
    axios.get(`${process.env.VUE_APP_API_ORIGIN}/maker/videos/${this.code}`, { timeout: 5000 }).then(
      res => {
        this.videos = res.data
      }
    ).catch(
      error => {
        console.log(error)
      }
    )
  }
}
</script>
