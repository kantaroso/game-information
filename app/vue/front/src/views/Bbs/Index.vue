<template>
  <div class="index">
    <b-overlay :show="overlay" rounded="sm">
      <FrontHeader />
      <Alert ref="alert"/>
      <PageTitle title="掲示板"/>
      <p>
        <b-container>
          <b-row>
            <b-col cols="auto">
              <BbsPost @startProcessing="postStart" @endProcessing="postEnd" />
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
import Alert from '@/components/Alert.vue'
import BbsContent from '@/components/Bbs/Content.vue'
import BbsPost from '@/components/Bbs/Post.vue'
import firestore from '@/lib/firebase/firestore'
import { BbsThread } from '@/lib/interface/bbs'
import { collection, query, orderBy, getDocs } from 'firebase/firestore'
import { getAuth, signInAnonymously, onAuthStateChanged } from 'firebase/auth'
@Component({
  components: {
    FrontHeader,
    FrontFooter,
    PageTitle,
    BbsContent,
    BbsPost,
    Alert
  }
})
export default class Index extends Vue {
  overlay = false
  bbsThreads: BbsThread[] = []
  query = query(collection(firestore, 'bbs'), orderBy('updated_at', 'desc'))

  postStart () {
    this.overlay = true
  }

  async postEnd (isError: boolean) {
    this.overlay = false
    const myAlert = this.$refs.alert as Alert
    if (isError) {
      myAlert.renderError()
    } else {
      myAlert.renderSuccess()
      await this.renderBbs()
    }
  }

  renderError () {
    console.log('エラー')
  }

  async renderBbs () {
    this.bbsThreads = []
    const querySnapshot = await getDocs(this.query)
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

  async mounted () {
    await this.renderBbs()
    const auth = getAuth()
    signInAnonymously(auth)
      .then(() => {
        // Signed in..
      })
      .catch((error) => {
        // const errorCode = error.code
        const errorMessage = error.message
        alert(errorMessage)
      })
    onAuthStateChanged(auth, (user) => {
      if (user) {
        // User is signed in, see docs for a list of available properties
        // https://firebase.google.com/docs/reference/js/firebase.User
        // const uid = user.uid
        // console.log(uid)
      }
    })
  }
}
</script>
