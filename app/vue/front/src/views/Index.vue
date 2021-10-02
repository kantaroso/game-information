<template>
  <div class="index">
    <FrontHeader />
    <PageTitle title="トップページ"/>
    <FrontFooter />
  </div>
</template>

<script lang="ts">
import { defineComponent } from '@vue/composition-api'
import FrontHeader from '@/components/Header.vue'
import FrontFooter from '@/components/Footer.vue'
import PageTitle from '@/components/Title.vue'

export default defineComponent({
  components: {
    FrontHeader,
    FrontFooter,
    PageTitle
  }
})
</script>
