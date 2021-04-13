<template>
  <div class="min-h-screen pt-20 pb-6 px-0 body-bg">
    <header class="max-w-lg mx-auto"><p class="text-4xl font-bold text-white text-center">Signin</p></header>

    <main class="bg-white max-w-lg mx-auto p-8 md:p-12 my-10 rounded-lg shadow-2xl">
      <section>
        <h3 class="font-bold text-2xl">Welcome to Card</h3>
        <p class="text-gray-600 pt-2">Sign in to your account.</p>
      </section>
      <section class="mt-10">
        <form class="flex flex-col" method="POST" action="#">
          <div class="mb-6 pt-3 rounded bg-gray-200">
            <label class="block text-gray-700 text-sm font-bold mb-2 ml-3" for="email">Email</label>
            <input type="text" id="email" class="bg-gray-200 rounded w-full text-gray-700 focus:outline-none border-b-4 border-gray-300 focus:border-purple-600 transition duration-500 px-3 pb-3">
          </div>
          <div class="mb-6 pt-3 rounded bg-gray-200">
            <label class="block text-gray-700 text-sm font-bold mb-2 ml-3" for="password">Password</label>
            <input type="password" id="password" class="bg-gray-200 rounded w-full text-gray-700 focus:outline-none border-b-4 border-gray-300 focus:border-purple-600 transition duration-500 px-3 pb-3">
          </div>
          <div class="flex flex-col justify-end">
            <p class="text-right"><a href="#" class="text-sm text-purple-600 hover:text-purple-700 hover:underline mb-6">Forgot your password?</a></p>
            <p class="text-right mb-6"><router-link to="/signup"><a href="#" class="text-sm text-purple-600 hover:text-purple-700 hover:underline mb-6">Do not have account?</a></router-link></p>
          </div>
          <button class="bg-purple-600 hover:bg-purple-700 text-white font-bold py-2 rounded shadow-lg hover:shadow-xl transition duration-200" type="submit">Sign In</button>
        </form>
      </section>
    </main>
  </div>
</template>

<script>
export default {
  name: 'Signin'
}
</script>

<style lang="scss" scoped>
.body-bg {
  background-color: #9921e8;
  background-image: linear-gradient(315deg, #9921e8 0%, #5f72be 74%);
}
</style>
