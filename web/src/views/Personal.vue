<template>
  <div class="text-gray-900 leading-normal tracking-wider bg-cover" style="background-image: url('https://source.unsplash.com/1L71sPT5XKc')">
    <div class="max-w-4xl flex items-center h-screen flex-wrap mx-auto">
      <div class="profile w-full lg:w-3/5 rounded-lg lg:rounded-l-lg lg:rounded-r-none shadow-2xl bg-white opacity-75 mx-6 lg:mx-0">
        <div class="p-4 md:p-12 text-left">
          <div class="block lg:hidden rounded-full shadow-xl mx-auto -mt-16 h-48 w-48 bg-cover bg-center" style="background-image: url('https://source.unsplash.com/MP0IUfwrn0A')"></div>
          <h1 class="text-3xl font-bold pt-8 lg:pt-0">{{ card.user?.name }}</h1>
          <div class="mx-auto lg:mx-0 w-4/5 pt-3 border-b-2 border-green-500 opacity-25"></div>
          <p class="pt-4">Company name: {{ card.user?.company.name }}</p>
          <p class="pt-2">Company website: {{ card.user?.company.website }}</p>
          <p class="pt-2">Company address: {{ card.user?.company.address }}</p>
        </div>
      </div>

      <div class="w-full lg:w-2/5">
        <img :src="card.User?.avatar" class="rounded-none lg:rounded-lg shadow-2xl hidden lg:block">
      </div>
    </div>
  </div>
</template>

<script>
import { mapState, mapActions } from 'vuex'

export default {
  name: 'Personal',

  data () {
    return {
      test: {}
    }
  },

  computed: {
    ...mapState('cards', ['card'])
  },

  methods: {
    ...mapActions('cards', ['getCard']),
  },

  mounted() {
    if (process.env.NODE_ENV === 'production') window.history.replaceState({}, '', `/personal/${this.card.User?.username}`)
  },

  created() {
    this.getCard(this.$route.params.code)
  }
}
</script>
