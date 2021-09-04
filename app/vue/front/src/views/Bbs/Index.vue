<template>
  <div class="index">
    <b-overlay :show="overlay" rounded="sm">
      <FrontHeader />
      <PageTitle title="掲示板"/>
      <p>
          <b-container>
              <b-row>
                  <b-col cols="auto">
                    <BbsPost @startProcessing="showOverlay" @endProcessing="hideOverlay" />
                  </b-col>
              </b-row>
          </b-container>
      </p>
      <p>
          <b-container>
              <BbsContent body="投稿テスト" :publishAt="1630741459558" />
          </b-container>
      </p>
      <FrontFooter />
    </b-overlay>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import FrontHeader from '@/components/Header.vue'
import FrontFooter from '@/components/Footer.vue'
import PageTitle from '@/components/Title.vue'
import BbsContent from '@/components/Bbs/Content.vue'
import BbsPost from '@/components/Bbs/Post.vue'
@Component({
  components: {
    FrontHeader,
    FrontFooter,
    PageTitle,
    BbsContent,
    BbsPost
  }
})
export default class Index extends Vue {
  overlay=false

  showOverlay () {
    this.overlay = true
  }

  hideOverlay () {
    this.overlay = false
  }
}
</script>
