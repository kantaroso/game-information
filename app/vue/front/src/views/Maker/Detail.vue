<template>
  <div class="index">
    <FrontHeader />
    <div v-if="info">
      <PageTitle :title="`${info.name}`" />
      <div v-if="info.ohp" class="m-3">
        <b-link target="_blank" :href="`${info.ohp}`">公式ホームページ</b-link>
      </div>
    </div>
    <div v-else>
      <h3><b-spinner label="Loading..."></b-spinner></h3>
    </div>
    <p v-if="info && info.twitter_name">
      <b-container>
        <b-row align-h="center">
          <b-col cols="10">
            <a
              class="twitter-timeline"
              data-height="500"
              data-lang="ja"
              data-theme="dark"
              :href="`https://twitter.com/${info.twitter_name}`"
              >Tweets by {{ info.twitter_name }}</a
            >
          </b-col>
        </b-row>
      </b-container>
    </p>
    <p v-else>
      <b-spinner label="Loading..."></b-spinner>
    </p>
    <p v-if="splitVideos.length > 0">
      <b-container>
        <b-row
          v-for="splitVideo in splitVideos"
          :key="splitVideo.index"
          cols="2"
          fluid
        >
          <b-col
            v-for="video in splitVideo.videos"
            :key="video.id"
            align-h="center"
            fluid
          >
            <b-card>
              <div class="movie-thumbnail-list" @click="showModal(video.id)">
                <b-card-img-lazy
                  thumbnail
                  :src="`https://i.ytimg.com/vi/${video.id}/hqdefault.jpg`"
                >
                </b-card-img-lazy>
                <div class="_mask">
                  <div class="_text">
                    {{ video.title }}
                  </div>
                  <div class="_img">
                    <b-img-lazy fluid src="/img/movie_play.png"></b-img-lazy>
                  </div>
                </div>
              </div>
            </b-card>
          </b-col>
        </b-row>
      </b-container>
    </p>
    <p v-else>
      <b-spinner label="Loading..."></b-spinner>
    </p>
    <b-modal
      ref="moviePlayModal"
      ok-only
      ok-title="閉じる"
      size="xl"
      hide-header
      hide-footer
      body-bg-variant="dark"
    >
      <b-embed
        type="iframe"
        aspect="16by9"
        :src="`https://www.youtube.com/embed/${modelVideoID}`"
        allowfullscreen
      ></b-embed>
    </b-modal>
    <FrontFooter />
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, toRefs, onMounted } from '@vue/composition-api'
import FrontHeader from '@/components/Header.vue'
import FrontFooter from '@/components/Footer.vue'
import PageTitle from '@/components/Title.vue'
import axios from 'axios'
interface ResponseVideo {
  id: string;
  title: string;
}
interface SpritVideo {
  index: number;
  videos: Array<ResponseVideo>;
}
interface ReactiveData {
  info: any; // eslint-disable-line @typescript-eslint/no-explicit-any
  splitVideos: Array<SpritVideo>;
  moviePlayModal: any; // eslint-disable-line @typescript-eslint/no-explicit-any
  modelVideoID: string;
}
export default defineComponent({
  components: {
    FrontHeader,
    FrontFooter,
    PageTitle
  },
  setup (_, context) {
    const videoColumnNumber = 2
    const code = context.root.$route.params.path
    const state: ReactiveData = reactive({
      info: null,
      splitVideos: [],
      moviePlayModal: null,
      modelVideoID: ''
    })

    const showModal = (id: string) => {
      state.modelVideoID = id
      state.moviePlayModal.show()
    }

    const hideModal = () => {
      state.moviePlayModal.hide()
    }

    const render = () => {
      axios
        .get(`${process.env.VUE_APP_API_ORIGIN}/maker/detail/${code}`, {
          timeout: 5000
        })
        .then((res) => {
          state.info = res.data
          const twitterScript = document.createElement('script')
          twitterScript.setAttribute(
            'src',
            'https://platform.twitter.com/widgets.js'
          )
          twitterScript.setAttribute('async', 'async')
          document.head.appendChild(twitterScript)
        })
        .catch(() => {
          alert('メーカーデータの取得に失敗しました')
        })
      axios
        .get(`${process.env.VUE_APP_API_ORIGIN}/maker/videos/${code}`, {
          timeout: 5000
        })
        .then((res) => {
          const tmp: Array<SpritVideo> = []
          res.data.forEach((element: ResponseVideo, index: number) => {
            const key = Math.floor(index / videoColumnNumber)
            if (!tmp[key]) {
              tmp[key] = {
                index: key,
                videos: []
              }
            }
            tmp[key].videos.push(element)
          })
          state.splitVideos = tmp
        })
        .catch(() => {
          alert('メーカー動画データの取得に失敗しました')
        })
    }

    onMounted(() => {
      render()
    })

    return {
      ...toRefs(state),
      showModal,
      hideModal
    }
  }
})
</script>

<style>
.movie-thumbnail-list {
  position: relative;
}
.movie-thumbnail-list ._mask {
  position: absolute;
  display: grid;
  align-items: center;
  justify-items: center;
  top: 0;
  right: 0;
  bottom: 0;
  left: 0;
  background-color: black;
  opacity: 0.5;
}
.movie-thumbnail-list ._img {
  position: absolute;
  width: 30%;
}
.movie-thumbnail-list ._text {
  position: absolute;
  top: 5%;
  left: 5%;
  right: 5%;
  color: white;
  font-weight: bold;
  display: -webkit-box;
  -webkit-box-orient: vertical;
  -webkit-line-clamp: 2;
  overflow: hidden;
}
@media (orientation: portrait) {
  .movie-thumbnail-list ._text {
    font-size: 70%;
  }
}
</style>
