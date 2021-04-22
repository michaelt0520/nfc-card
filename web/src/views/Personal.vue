<template>
  <div class="text-gray-900 leading-normal tracking-wider bg-cover" style="background-image: url('https://source.unsplash.com/1L71sPT5XKc')">
    <div class="max-w-4xl flex items-center h-screen flex-wrap mx-auto">
      <div class="profile w-full lg:w-3/5 rounded-lg lg:rounded-l-lg lg:rounded-r-none shadow-2xl bg-white opacity-75 mx-6 lg:mx-0">
        <div class="p-4 md:p-12 text-left">
          <img width="120" height="120" class="mx-auto -mt-16 block lg:hidden" :src="card.user?.avatar" alt="avatar">
          <h1 class="text-3xl font-bold pt-8 lg:pt-0">{{ card.user?.name }}</h1>
          <div class="mx-auto lg:mx-0 w-4/5 pt-3 border-b-2 border-green-500 opacity-25"></div>
          <p class="pt-4"><span class="font-bold">Company name:</span> {{ card.user?.company.name }}</p>
          <p class="pt-2"><span class="font-bold">Company website:</span> {{ card.user?.company.website }}</p>
          <p class="pt-2"><span class="font-bold">Company address:</span> {{ card.user?.company.address }}</p>

          <p
            v-for="item in card.user?.informations"
            :key="item.index"
            class="pt-2">
            <span class="font-bold">{{ item.name }}: </span>
            <a target="_blank" :href="item.data">{{ item.data }}</a>
          </p>
        </div>
        <div class="text-right mr-6 mb-6">
          <router-link :to="`${this.$route.fullPath}/edit`">
            <button @click="onConfirmEdit = true" class="focus:outline-none focus:ring-2 focus:ring-purple-600 focus:ring-opacity-50 p-2 rounded-md bg-gray-300 hover:bg-gray-400">Edit</button>
          </router-link>
        </div>
      </div>

      <div class="w-full lg:w-2/5">
        <img :src="card.user?.avatar" class="rounded-none lg:rounded-lg shadow-2xl hidden lg:block">
      </div>
    </div>

    <!-- <modal
      v-if="onConfirmEdit"
      v-model:open="onConfirmEdit"
      header="Edit information"
      @confirm-dialog="onConfirmEditInfo"
      isHiddenButtonConfirm="false">
      <template v-slot:modal-body>
        <p>body</p>
        <p>body</p>
        <p>body</p>
        <p>body</p>
        <p>body</p>
      </template>
    </modal> -->
  </div>
</template>

<script>
import { mapState, mapActions } from 'vuex'

export default {
  name: 'Personal',

  data () {
    return {
      onConfirmEdit: false
    }
  },

  computed: {
    ...mapState('cards', ['card'])
  },

  methods: {
    ...mapActions('cards', ['getCard'])
  },

  mounted() {
    if (process.env.NODE_ENV === 'production') window.history.replaceState({}, '', `/personal/${this.card.User?.username}`)
  },

  created() {
    this.getCard(this.$route.params.code)
  }
}
</script>
