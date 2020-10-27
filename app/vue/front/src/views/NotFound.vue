<template>
  <div class="index">
    <FrontHeader />
    <PageTitle title="ページが存在しません"/>
    <FrontFooter />
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import FrontHeader from '@/components/Header.vue'
import FrontFooter from '@/components/Footer.vue'
import PageTitle from '@/components/Title.vue'
@Component({
  components: {
    FrontHeader,
    FrontFooter,
    PageTitle
  }
})
export default class Index extends Vue {}
</script>
