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
import { Component, Vue, Prop } from 'vue-property-decorator'
import firestore from '@/lib/firebase/Firestore'
interface Input {
  name: string | null;
  body: string | null;
}
@Component
export default class Post extends Vue {
  @Prop({ type: String, required: false })
  mode!: string

  input: Input = {
    name: '',
    body: ''
  }

  onSubmit (event: Event) {
    event.preventDefault()
    console.log(this.input)
    this.$root.$emit('bv::hide::modal', 'modal-post')
    console.log(firestore)
    // this.$emit('startProcessing')
    // this.$emit('endProcessing')
  }
}
</script>
