<template>
  <div class="index">
    <FrontHeader/>
    <h1>new page</h1>
    <FrontFooter/>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import FrontHeader from '@/components/Header.vue' // @ is an alias to /src
import FrontFooter from '@/components/Footer.vue' // @ is an alias to /src

@Component({
  components: {
    FrontHeader,
    FrontFooter
  }
})
export default class Index extends Vue {}
</script>
