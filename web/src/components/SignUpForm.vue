<template>
  <div
    id="join"
    class="flex justify-center w-full antialiased mx-auto my-12 px-8 mb-14 p-9 bg-gray-100"
  >
    <div class="flex mb-4 max-w-6xl row-signup">
      <div class="w-3/5 row-right">
        <div class="text-dig">
          <p
            class="text-4xl sm:text-5xl text-gray-800 dark:text-white font-extrabold tracking-tight"
          >
            Join with us! <br />
            <br />
          </p>
          <div class="text-center mb-4 w-3/5">
            <img src="@/assets/workflow_4x.png" />
          </div>
        </div>
      </div>
      <div class="w-2/5 row-left">
        <div class="text-singup mb-8">
          <h2 class="">Join us for free</h2>
        </div>
        <div class="form-signup-dig">
          <div class="w-full max-w-xs">
            <form class="max-w-md mb-4 form-input" @submit.prevent>
              <div class="mb-4">
                <label
                  class="block text-grey-darker text-sm font-bold mb-2"
                  for="fullname"
                >
                  Full name
                </label>
                <input
                  class="shadow appearance-none border rounded w-full h-12 py-2 px-3 text-grey-darker leading-tight focus:outline-none focus:shadow-outline"
                  id="fullname"
                  type="text"
                  placeholder="Full name"
                  v-model="fullname"
                />
              </div>
              <div class="mb-4">
                <label
                  class="block text-grey-darker text-sm font-bold mb-2"
                  for="username"
                >
                  Username
                </label>
                <input
                  class="shadow appearance-none border rounded w-full h-12 py-2 px-3 text-grey-darker leading-tight focus:outline-none focus:shadow-outline"
                  id="username"
                  type="text"
                  placeholder="Username"
                  v-model="username"
                />
              </div>
              <div class="mb-4">
                <label
                  class="block text-grey-darker text-sm font-bold mb-2"
                  for="email"
                >
                  Email
                </label>
                <input
                  class="shadow appearance-none border rounded h-12 w-full py-2 px-3 text-grey-darker mb-3 leading-tight focus:outline-none focus:shadow-outline"
                  id="email"
                  type="email"
                  placeholder="email"
                  v-model="email"
                />
              </div>
              <div class="mb-6">
                <label
                  class="block text-grey-darker text-sm font-bold mb-2"
                  for="password"
                >
                  Password
                </label>
                <input
                  class="shadow appearance-none border rounded w-full h-12 py-2 px-3 text-grey-darker mb-3 leading-tight focus:outline-none focus:shadow-outline"
                  id="password"
                  type="password"
                  placeholder="******************"
                  v-model="password"
                />
              </div>
              <div class="flex items-center justify-between">
                <button
                  class="bg-blue hover:bg-blue-dark text-black font-bold w-full h-12 py-2 px-4 rounded focus:text-gray-900 hover:bg-gray-200"
                  type="button"
                  @click="onClickSignUp"
                >
                  Sign Up
                </button>
              </div>
            </form>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapActions } from "vuex";

export default {
  name: "SignUpForm",

  data() {
    return {
      email: "",
      fullname: "",
      username: "",
      password: "",
    };
  },

  methods: {
    ...mapActions("users", ["signup"]),

    onClickSignUp() {
      const data = {
        email: this.email,
        name: this.fullname,
        username: this.username,
        password: this.password,
      };
      this.email = this.fullname = this.username = this.password = "";
      this.signup(data);
    },
  },
};
</script>
