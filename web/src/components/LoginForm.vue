<template>
  <div class="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4 flex flex-col">
    <div class="mb-4">
      <label
        class="block text-grey-darker text-sm font-bold mb-2"
        for="username"
      >
        Email
      </label>
      <input
        class="shadow appearance-none border rounded w-full py-2 px-3 text-grey-darker"
        id="email"
        type="text"
        placeholder="Email"
        v-model="loginEmail"
        @keyup.enter="onClickSignin"
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
        class="shadow appearance-none border border-red rounded w-full py-2 px-3 text-grey-darker mb-3"
        id="password"
        type="password"
        placeholder="******************"
        @keyup.enter="onClickSignin"
        v-model="loginPassword"
      />
    </div>
    <div class="flex items-center justify-between">
      <button
        class="bg-blue hover:bg-blue-dark text-gray-900 font-bold py-2 px-4 rounded focus:text-gray-900 hover:bg-gray-200"
        type="button"
        @click="onClickSignin"
      >
        Sign In
      </button>
      <a
        class="bg-blue hover:bg-blue-dark text-gray-900 font-bold py-2 px-4 rounded focus:text-gray-900 hover:bg-gray-200"
        href="#"
      >
        Forgot Password?
      </a>
    </div>
  </div>
</template>

<script>
import { mapActions } from "vuex";

export default {
  name: "LoginForm",

  props: {
    open: {
      type: Boolean,
      default: false,
    },
  },

  data() {
    return {
      loginEmail: "",
      loginPassword: "",
    };
  },

  methods: {
    ...mapActions("users", ["signin"]),

    onClickSignin() {
      const data = {
        email: this.loginEmail,
        password: this.loginPassword,
      };
      this.signin(data);
      this.$emit("update:open", false);
    },
  },
};
</script>
