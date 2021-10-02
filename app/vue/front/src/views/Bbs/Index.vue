<template>
  <div class="index">
    <b-overlay :show="overlay" rounded="sm">
      <FrontHeader />
      <Alert ref="alertElement"/>
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
import { defineComponent, reactive, toRefs, onMounted } from '@vue/composition-api'
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
interface ReactiveData {
  overlay: boolean;
  bbsThreads: BbsThread[];
  alertElement: any; // eslint-disable-line @typescript-eslint/no-explicit-any
}
export default defineComponent({
  components: {
    FrontHeader,
    FrontFooter,
    PageTitle,
    BbsContent,
    BbsPost,
    Alert
  },
  setup () {
    const state: ReactiveData = reactive({
      overlay: false,
      bbsThreads: [],
      alertElement: null
    })
    const getBbsQuery = query(collection(firestore, 'bbs'), orderBy('updated_at', 'desc'))

    const renderBbs = async () => {
      state.bbsThreads = []
      const querySnapshot = await getDocs(getBbsQuery)
      querySnapshot.forEach((doc) => {
        const thread = {
          id: doc.id,
          title: doc.get('title'),
          name: doc.get('name'),
          body: doc.get('body'),
          updatedAt: doc.get('updated_at'),
          createdAt: doc.get('created_at')
        }
        state.bbsThreads.push(thread)
      })
    }

    const postStart = () => {
      state.overlay = true
    }

    const postEnd = async (isError: boolean) => {
      console.log('postEnd')
      state.overlay = false
      if (isError) {
        state.alertElement.renderError()
      } else {
        state.alertElement.renderSuccess()
        await renderBbs()
      }
    }

    onMounted(async () => {
      await renderBbs()
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
    })

    return {
      ...toRefs(state),
      postStart,
      postEnd
    }
  }
})
</script>
