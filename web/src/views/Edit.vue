<template>
  <form @submit.prevent>
    <div class="bg-gray-200">
      <div class="min-h-screen flex items-center justify-center px-4">
        <div class="max-w-4xl  bg-white w-full rounded-lg shadow-xl">
          <div class="p-4 border-b">
            <h2 class="text-2xl ">
              Information App
            </h2>
            <p class="text-sm text-gray-500">
              Personal details and application.
            </p>
          </div>

          <div>
            <div class="md:grid md:grid-cols-2 hover:bg-gray-50 md:space-y-0 space-y-1 p-4 border-b">
              <p class="text-gray-600">
                Full name
              </p>
              <input v-model="username" class="bg-gray-200 rounded w-full text-gray-700 focus:outline-none border-b-4 border-gray-300 focus:border-purple-600 transition duration-500 px-3 pb-3" type="text" :placeholder="card.user?.name">
            </div>

            <div
              v-if="!isPersonalCard"
              class="md:grid md:grid-cols-2 hover:bg-gray-50 md:space-y-0 space-y-1 p-4 border-b">
              <p class="text-gray-600">
                Company name
              </p>
              <input v-model="companyName" class="bg-gray-200 rounded w-full text-gray-700 focus:outline-none border-b-4 border-gray-300 focus:border-purple-600 transition duration-500 px-3 pb-3" type="text" :placeholder="card.user?.company.name">
            </div>

            <div
              v-if="!isPersonalCard"
              class="md:grid md:grid-cols-2 hover:bg-gray-50 md:space-y-0 space-y-1 p-4 border-b">
              <p class="text-gray-600">
                Company website
              </p>
              <input v-model="companyWebsite" class="bg-gray-200 rounded w-full text-gray-700 focus:outline-none border-b-4 border-gray-300 focus:border-purple-600 transition duration-500 px-3 pb-3" type="text" :placeholder="card.user?.company.website">
            </div>

            <div
              v-if="!isPersonalCard"
              class="md:grid md:grid-cols-2 hover:bg-gray-50 md:space-y-0 space-y-1 p-4 border-b">
              <p class="text-gray-600">
                Company address
              </p>
              <input v-model="companyAddress" class="bg-gray-200 rounded w-full text-gray-700 focus:outline-none border-b-4 border-gray-300 focus:border-purple-600 transition duration-500 px-3 pb-3" type="text" :placeholder="card.user?.company.address">
            </div>

            <div
              v-for="item in card.user?.informations"
              :key="item.index"
              class="md:grid md:grid-cols-2 hover:bg-gray-50 md:space-y-0 space-y-1 p-4 border-b">
              <p class="text-gray-600">
                {{ item.name }}
              </p>
              <input :v-model="`mxh${item.name}`" class="bg-gray-200 rounded w-full text-gray-700 focus:outline-none border-b-4 border-gray-300 focus:border-purple-600 transition duration-500 px-3 pb-3" type="text" :placeholder="item.data">
            </div>
          </div>

          <div class="text-right m-2">
            <button @click="onSubmit" class="bg-purple-600 hover:bg-purple-700 text-white font-bold p-2 rounded shadow-lg hover:shadow-xl transition duration-200" type="submit">Edit</button>
          </div>
        </div>
      </div>
    </div>
  </form>
</template>

<script>
import { mapState, mapActions } from 'vuex'

export default {
  name: 'Edit',

  data () {
    return {
      username: '',
      companyName: '',
      companyWebsite: '',
      companyAddress: '',
      Facebook: '',
      Instagram: '',
      Youtube: '',
      WhatsApp: ''
    }
  },

  computed: {
    ...mapState('cards', ['card']),

    isPersonalCard () {
      return this.card.user?.type === 'Personal'
    }
  },

  methods: {
    ...mapActions('cards', ['getCard']),

    onSubmit() {

      console.log(this.mxhFacebook)
    }
  },

  created() {
    this.getCard(this.$route.params.code)
  }
}
</script>

<style>

</style>
