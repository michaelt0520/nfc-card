<template>
  <div class="min-h-screen pt-20 pb-6 px-0 body-bg">
    <header class="max-w-lg mx-auto"><p class="text-4xl font-bold text-white text-center">Signup</p></header>

    <main class="bg-white max-w-lg mx-auto p-8 md:p-12 my-10 rounded-lg shadow-2xl">
      <section>
        <h3 class="font-bold text-2xl">Welcome to Card</h3>
        <p class="text-gray-600 pt-2">Signup to your account.</p>
      </section>
      <section class="mt-10" v-if="isLogedIn">
        <form @submit.prevent class="flex flex-col">
          <div class="mb-6 pt-3 rounded bg-gray-200">
            <label class="block text-gray-700 text-sm font-bold mb-2 ml-3" for="email">Email</label>
            <input v-model="email" type="email" id="email" class="bg-gray-200 rounded w-full text-gray-700 focus:outline-none border-b-4 border-gray-300 focus:border-purple-600 transition duration-500 px-3 pb-3">
          </div>
          <div class="mb-6 pt-3 rounded bg-gray-200">
            <label class="block text-gray-700 text-sm font-bold mb-2 ml-3" for="fullname">Full name</label>
            <input v-model="fullname" type="text" id="fullname" class="bg-gray-200 rounded w-full text-gray-700 focus:outline-none border-b-4 border-gray-300 focus:border-purple-600 transition duration-500 px-3 pb-3">
          </div>
          <div class="mb-6 pt-3 rounded bg-gray-200">
            <label class="block text-gray-700 text-sm font-bold mb-2 ml-3" for="username">Username</label>
            <input v-model="username" type="text" id="username" class="bg-gray-200 rounded w-full text-gray-700 focus:outline-none border-b-4 border-gray-300 focus:border-purple-600 transition duration-500 px-3 pb-3">
          </div>
          <div class="mb-6 pt-3 rounded bg-gray-200">
            <label class="block text-gray-700 text-sm font-bold mb-2 ml-3" for="password">Password</label>
            <input v-model="password" type="password" id="password" class="bg-gray-200 rounded w-full text-gray-700 focus:outline-none border-b-4 border-gray-300 focus:border-purple-600 transition duration-500 px-3 pb-3">
          </div>
          <div class="flex flex-col justify-end">
            <p class="text-right"><a href="#" class="text-sm text-purple-600 hover:text-purple-700 hover:underline mb-6">Forgot your password?</a></p>
            <p class="text-right mb-6"><router-link to="/signup"><a href="#" class="text-sm text-purple-600 hover:text-purple-700 hover:underline mb-6">Do not have account?</a></router-link></p>
          </div>
          <button @click="onSubmit" class="bg-purple-600 hover:bg-purple-700 text-white font-bold py-2 rounded shadow-lg hover:shadow-xl transition duration-200" type="submit">Sign Up</button>
        </form>
      </section>
    </main>
  </div>
</template>

<script>
import { mapState, mapActions } from 'vuex'

export default {
  name: 'Signup',

  data() {
    return {
      email: '',
      fullname: '',
      username: '',
      password: ''
    }
  },

  computed: {
    ...mapState('users', ['currentUser']),

    isLogedIn() {
      return this.currentUser !== null
    }
  },

  methods: {
    ...mapActions('users', ['signup']),

    onSubmit() {
      const data = {
        email: this.email,
        name: this.fullname,
        username: this.username,
        password: this.password
      }
      this.email = this.fullname = this.username = this.password = ''
      this.signup(data)
    }
  }
}
</script>

<style lang="scss" scoped>
.body-bg {
  background-color: #9921e8;
  background-image: linear-gradient(315deg, #9921e8 0%, #5f72be 74%);
}
</style>
