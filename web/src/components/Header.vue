<template>
  <div class="min-w-full fixed top-0 z-50">
    <div class="antialiased bg-gray-100 dark-mode:bg-gray-900">
      <div
        class="w-full text-gray-700 bg-white dark-mode:text-gray-200 dark-mode:bg-gray-800"
      >
        <div
          x-data="{ open: true }"
          class="flex flex-col max-w-screen-xl px-4 mx-auto md:items-center md:justify-between md:flex-row md:px-6 lg:px-8"
        >
          <div class="flex flex-row items-center justify-between p-4">
            <router-link
              to="/"
              class="text-lg font-semibold tracking-widest text-gray-900 uppercase rounded-lg dark-mode:text-white focus:outline-none focus:shadow-outline"
            >
              ONETAP
            </router-link>
            <button
              class="rounded-lg md:hidden focus:outline-none focus:shadow-outline"
              @click="open = !open"
            >
              <svg fill="currentColor" viewBox="0 0 20 20" class="w-6 h-6">
                <path
                  x-show="!open"
                  fill-rule="evenodd"
                  d="M3 5a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM3 10a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM9 15a1 1 0 011-1h6a1 1 0 110 2h-6a1 1 0 01-1-1z"
                  clip-rule="evenodd"
                ></path>
                <path
                  x-show="open"
                  fill-rule="evenodd"
                  d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z"
                  clip-rule="evenodd"
                ></path>
              </svg>
            </button>
          </div>
          <nav
            :class="{ flex: open, hidden: !open }"
            class="flex-col flex-grow hidden pb-4 md:pb-0 md:flex md:justify-end md:flex-row"
          >
            <router-link
              to="/"
              class="px-4 py-2 mt-2 text-sm font-semibold bg-transparent rounded-lg dark-mode:bg-transparent dark-mode:hover:bg-gray-600 dark-mode:focus:bg-gray-600 dark-mode:focus:text-white dark-mode:hover:text-white dark-mode:text-gray-200 md:mt-0 md:ml-4 hover:text-gray-900 focus:text-gray-900 hover:bg-gray-200 focus:bg-gray-200 focus:outline-none focus:shadow-outline"
            >
              Features
            </router-link>
            <router-link
              to="#join"
              class="px-4 py-2 mt-2 text-sm font-semibold bg-transparent rounded-lg dark-mode:bg-transparent dark-mode:hover:bg-gray-600 dark-mode:focus:bg-gray-600 dark-mode:focus:text-white dark-mode:hover:text-white dark-mode:text-gray-200 md:mt-0 md:ml-4 hover:text-gray-900 focus:text-gray-900 hover:bg-gray-200 focus:bg-gray-200 focus:outline-none focus:shadow-outline"
            >
              Join
            </router-link>
            <router-link
              to="#price"
              class="px-4 py-2 mt-2 text-sm font-semibold bg-transparent rounded-lg dark-mode:bg-transparent dark-mode:hover:bg-gray-600 dark-mode:focus:bg-gray-600 dark-mode:focus:text-white dark-mode:hover:text-white dark-mode:text-gray-200 md:mt-0 md:ml-4 hover:text-gray-900 focus:text-gray-900 hover:bg-gray-200 focus:bg-gray-200 focus:outline-none focus:shadow-outline"
            >
              Price
            </router-link>
            <router-link
              to="#contact"
              class="px-4 py-2 mt-2 text-sm font-semibold bg-transparent rounded-lg dark-mode:bg-transparent dark-mode:hover:bg-gray-600 dark-mode:focus:bg-gray-600 dark-mode:focus:text-white dark-mode:hover:text-white dark-mode:text-gray-200 md:mt-0 md:ml-4 hover:text-gray-900 focus:text-gray-900 hover:bg-gray-200 focus:bg-gray-200 focus:outline-none focus:shadow-outline"
            >
              Contact
            </router-link>
            <div
              class="relative"
              x-data="{ open: true }"
              v-if="!isAuthenticated"
            >
              <button
                @click="open = !open"
                :class="{ 'bg-gray-200': open }"
                class="flex flex-row text-gray-900 items-center w-full px-4 py-2 mt-2 text-sm font-semibold text-left bg-transparent rounded-lg dark-mode:bg-transparent dark-mode:focus:text-white dark-mode:hover:text-white dark-mode:focus:bg-gray-600 dark-mode:hover:bg-gray-600 md:w-auto md:inline md:mt-0 md:ml-4 hover:text-gray-900 focus:text-gray-900 hover:bg-gray-200 focus:bg-gray-200 focus:outline-none focus:shadow-outline"
              >
                <span>Login</span>
                <svg
                  fill="currentColor"
                  viewBox="0 0 20 20"
                  :class="{ 'rotate-180': open, 'rotate-0': !open }"
                  class="inline w-4 h-4 mt-1 ml-1 transition-transform duration-200 transform md:-mt-1"
                >
                  <path
                    fill-rule="evenodd"
                    d="M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z"
                    clip-rule="evenodd"
                  ></path>
                </svg>
              </button>
              <div
                v-if="open"
                x-transition:enter="transition ease-out duration-100"
                x-transition:enter-start="transform opacity-0 scale-95"
                x-transition:enter-end="transform opacity-100 scale-100"
                x-transition:leave="transition ease-in duration-75"
                x-transition:leave-start="transform opacity-100 scale-100"
                x-transition:leave-end="transform opacity-0 scale-95"
                class="absolute right-0 w-full md:max-w-screen-sm md:w-screen mt-2 origin-top-right"
              >
                <login-form v-model:open="open" />
              </div>
            </div>
            <div class="relative" x-data="{ open: true }" v-else>
              <button
                @click="open = !open"
                :class="{ 'bg-gray-200': open }"
                class="flex flex-row text-gray-900 items-center w-full px-4 py-2 mt-2 text-sm font-semibold text-left bg-transparent rounded-lg dark-mode:bg-transparent dark-mode:focus:text-white dark-mode:hover:text-white dark-mode:focus:bg-gray-600 dark-mode:hover:bg-gray-600 md:w-auto md:inline md:mt-0 md:ml-4 hover:text-gray-900 focus:text-gray-900 hover:bg-gray-200 focus:bg-gray-200 focus:outline-none focus:shadow-outline"
              >
                <span>Hello {{ currentUser.name }}</span>
                <svg
                  fill="currentColor"
                  viewBox="0 0 20 20"
                  :class="{ 'rotate-180': open, 'rotate-0': !open }"
                  class="inline w-4 h-4 mt-1 ml-1 transition-transform duration-200 transform md:-mt-1"
                >
                  <path
                    fill-rule="evenodd"
                    d="M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z"
                    clip-rule="evenodd"
                  ></path>
                </svg>
              </button>
              <div
                v-if="open"
                x-transition:enter="transition ease-out duration-100"
                x-transition:enter-start="transform opacity-0 scale-95"
                x-transition:enter-end="transform opacity-100 scale-100"
                x-transition:leave="transition ease-in duration-75"
                x-transition:leave-start="transform opacity-100 scale-100"
                x-transition:leave-end="transform opacity-0 scale-95"
                class="absolute right-0 w-full md:max-w-screen-sm mt-2 origin-top-right"
              >
                <div class="bg-white shadow-md rounded flex flex-col">
                  <div class="flex flex-col justify-items-start">
                    <button
                      v-if="currentUser.role === 4"
                      class="bg-blue hover:bg-blue-dark text-gray-900 font-bold py-2 px-4 rounded focus:text-gray-900 hover:bg-gray-200"
                      type="button"
                    >
                      <router-link to="/admin"> Admin </router-link>
                    </button>
                    <button
                      v-if="currentUser.role === 3"
                      class="bg-blue hover:bg-blue-dark text-gray-900 font-bold py-2 px-4 rounded focus:text-gray-900 hover:bg-gray-200"
                      type="button"
                    >
                      <router-link to="/company"> Company </router-link>
                    </button>
                    <button
                      class="bg-blue hover:bg-blue-dark text-gray-900 font-bold py-2 px-4 rounded focus:text-gray-900 hover:bg-gray-200"
                      type="button"
                    >
                      <router-link to="/setting"> Setting </router-link>
                    </button>
                    <button
                      class="bg-blue hover:bg-blue-dark text-gray-900 font-bold py-2 px-4 rounded focus:text-gray-900 hover:bg-gray-200"
                      @click="onClickSignout"
                    >
                      Logout
                    </button>
                  </div>
                </div>
              </div>
            </div>
          </nav>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapState, mapActions } from "vuex";
import LoginForm from "./LoginForm";

export default {
  name: "Header",

  components: {
    LoginForm,
  },

  data() {
    return {
      open: false,
    };
  },

  computed: {
    ...mapState("users", ["currentUser", "isAuthenticated"]),
  },

  methods: {
    ...mapActions("users", ["signout"]),

    onClickSignout() {
      this.signout();
      this.open = false;
      this.$router.push("/");
    },
  },
};
</script>
