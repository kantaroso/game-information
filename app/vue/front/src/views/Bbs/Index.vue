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
          <b-container v-for="item in bbsThreads" :key="item.id">
            <p><BbsContent :bbsThread="item" /></p>
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
import firestore from '@/lib/firebase/Firestore'
import { BbsThread } from '@/lib/firestore/Interface'
import { collection, query, orderBy, getDocs } from 'firebase/firestore'
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
  bbsThreads: BbsThread[] = []

  showOverlay () {
    this.overlay = true
  }

  hideOverlay () {
    this.overlay = false
  }

  renderError () {
    console.log('エラー')
  }

  async mounted () {
    const q = query(collection(firestore, 'bbs'), orderBy('updated_at', 'desc'))
    const querySnapshot = await getDocs(q)
    querySnapshot.forEach((doc) => {
      const thread = {
        id: doc.id,
        title: doc.get('title'),
        name: doc.get('name'),
        body: doc.get('body'),
        updatedAt: doc.get('updated_at'),
        createdAt: doc.get('created_at')
      }
      this.bbsThreads.push(thread)
    })
  }
}
</script>
