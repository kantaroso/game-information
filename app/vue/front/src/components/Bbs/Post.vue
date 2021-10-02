<template>
  <div>
    <b-button variant="success" v-b-modal.modal-post>投稿する</b-button>
      <b-modal
        id="modal-post"
        title="掲示板に投稿する"
        hide-footer
        v-if="input"
        >
        <b-form @submit="onSubmit">
          <b-form-group
            label="名前"
            label-for="input-name"
          >
            <b-form-input
              id="input-name"
              v-model="input.name"
              required
            ></b-form-input>
          </b-form-group>
          <b-form-group
            label="タイトル"
            label-for="input-title"
          >
            <b-form-input
              id="input-title"
              v-model="input.title"
              required
            ></b-form-input>
          </b-form-group>
          <b-form-group
            label="本文"
            label-for="input-body"
          >
            <b-form-textarea
              id="input-body"
              v-model="input.body"
              required
            ></b-form-textarea>
          </b-form-group>
          <b-button type="submit" variant="primary">投稿する</b-button>
        </b-form>
      </b-modal>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, toRefs } from '@vue/composition-api'
import firestore from '@/lib/firebase/firestore'
import { doc, setDoc } from 'firebase/firestore'
import { nanoid } from 'nanoid'
import { BbsThread } from '@/lib/interface/bbs'
import { toUnderscoreCaseObject } from '@/lib/interface/common'

interface Input {
  name: string;
  title: string;
  body: string;
}
interface ReactiveData {
  input: Input;
}
export default defineComponent({
  props: {
    mode: {
      type: String,
      required: false
    }
  },
  setup (_, context) {
    const state: ReactiveData = reactive({
      input: {
        name: '',
        title: '',
        body: ''
      }
    })

    const onSubmit = (event: Event) => {
      event.preventDefault()
      context.emit('startProcessing')
      context.root.$root.$emit('bv::hide::modal', 'modal-post')
      const uuid = nanoid()
      const now = Date.now()
      const docData: BbsThread = {
        id: uuid,
        title: state.input.title,
        name: state.input.name,
        body: state.input.body,
        createdAt: now,
        updatedAt: now
      }
      context.emit('endProcessing')
      const promise = setDoc(doc(firestore, 'bbs', uuid), toUnderscoreCaseObject(docData))
      promise.then(() => {
        state.input = {
          name: '',
          title: '',
          body: ''
        }
        context.emit('endProcessing')
      }).catch(() => {
        state.input = {
          name: '',
          title: '',
          body: ''
        }
        context.emit('endProcessing', true)
      })
    }
    return {
      ...toRefs(state),
      onSubmit
    }
  }
})
</script>
