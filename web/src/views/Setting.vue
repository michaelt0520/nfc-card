<template>
  <div>
    <Header />

    <div>
      <form @submit.prevent>
        <div class="bg-gray-200">
          <div class="min-h-screen flex items-center justify-center px-4">
            <div class="max-w-4xl bg-white w-full rounded-lg shadow-xl">
              <div class="p-4 border-b">
                <h2 class="text-2xl">Information App</h2>
                <p class="text-sm text-gray-500">
                  Personal details and application.
                </p>
              </div>

              <div>
                <div
                  class="md:grid md:grid-cols-2 hover:bg-gray-50 md:space-y-0 space-y-1 p-4 border-b"
                >
                  <div
                    class="bg-white px-4 py-5 rounded-lg shadow-lg text-center w-48"
                  >
                    <div class="mb-4">
                      <img
                        class="w-auto mx-auto rounded-full object-cover object-center"
                        :src="updateUser.avatar"
                        alt="Avatar Upload"
                      />
                    </div>
                    <label class="cursor-pointer mt-6">
                      <span
                        class="mt-2 leading-normal px-4 py-2 bg-blue-500 text-white text-sm rounded-full"
                        >Select Avatar</span
                      >
                      <input
                        type="file"
                        class="hidden"
                        :multiple="multiple"
                        :accept="accept"
                      />
                    </label>
                  </div>
                </div>

                <div
                  class="md:grid md:grid-cols-2 hover:bg-gray-50 md:space-y-0 space-y-1 p-4 border-b"
                >
                  <p class="text-gray-600">Username</p>
                  <input
                    class="bg-gray-200 rounded w-full text-gray-700 focus:outline-none border-b-4 border-gray-300 focus:border-purple-600 transition duration-500 px-3 pb-3"
                    type="text"
                    disabled
                    :value="updateUser.username"
                  />
                </div>

                <div
                  class="md:grid md:grid-cols-2 hover:bg-gray-50 md:space-y-0 space-y-1 p-4 border-b"
                >
                  <p class="text-gray-600">Full name</p>
                  <input
                    id="name"
                    class="bg-gray-200 rounded w-full text-gray-700 focus:outline-none border-b-4 border-gray-300 focus:border-purple-600 transition duration-500 px-3 pb-3"
                    type="text"
                    v-model="updateUser.name"
                  />
                </div>

                <div
                  class="md:grid md:grid-cols-2 hover:bg-gray-50 md:space-y-0 space-y-1 p-4 border-b"
                >
                  <p class="text-gray-600">Email</p>
                  <input
                    id="email"
                    class="bg-gray-200 rounded w-full text-gray-700 focus:outline-none border-b-4 border-gray-300 focus:border-purple-600 transition duration-500 px-3 pb-3"
                    type="text"
                    v-model="updateUser.email"
                  />
                </div>

                <div
                  class="md:grid md:grid-cols-2 hover:bg-gray-50 md:space-y-0 space-y-1 p-4 border-b"
                >
                  <p class="text-gray-600">Address</p>
                  <input
                    id="address"
                    class="bg-gray-200 rounded w-full text-gray-700 focus:outline-none border-b-4 border-gray-300 focus:border-purple-600 transition duration-500 px-3 pb-3"
                    type="text"
                    v-model="updateUser.address"
                  />
                </div>

                <div
                  class="md:grid md:grid-cols-2 hover:bg-gray-50 md:space-y-0 space-y-1 p-4 border-b"
                >
                  <p class="text-gray-600">Phone number</p>
                  <input
                    id="phone_number"
                    class="bg-gray-200 rounded w-full text-gray-700 focus:outline-none border-b-4 border-gray-300 focus:border-purple-600 transition duration-500 px-3 pb-3"
                    type="tel"
                    v-model="updateUser.phone_number"
                  />
                </div>

                <div
                  v-for="item in user.informations"
                  :key="item.index"
                  class="md:grid md:grid-cols-2 hover:bg-gray-50 md:space-y-0 space-y-1 p-4 border-b"
                >
                  <p class="text-gray-600">
                    {{ item.name }}
                  </p>
                  <input
                    :v-model="`mxh${item.name}`"
                    class="bg-gray-200 rounded w-full text-gray-700 focus:outline-none border-b-4 border-gray-300 focus:border-purple-600 transition duration-500 px-3 pb-3"
                    type="text"
                    :placeholder="item.data"
                  />
                </div>
              </div>

              <div class="text-right m-2">
                <button
                  @click="onClickUpdate"
                  class="bg-purple-600 hover:bg-purple-700 text-white font-bold p-2 rounded shadow-lg hover:shadow-xl transition duration-200"
                  type="submit"
                >
                  Update
                </button>
              </div>
            </div>
          </div>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
import { mapState, mapActions } from "vuex";
import Header from "../components/Header";

export default {
  name: "Edit",

  data() {
    return {
      updateUser: {
        avatar: "",
        name: "",
        email: "",
        address: "",
        phone_number: "",
      },
    };
  },

  components: {
    Header,
  },

  computed: {
    ...mapState("users", ["user"]),
  },

  methods: {
    ...mapActions("users", ["getCurrentUser", "updateCurrentUser"]),

    onClickUpdate() {
      this.updateCurrentUser(this.updateUser);
    },
  },

  async created() {
    await this.getCurrentUser();
    this.updateUser = JSON.parse(JSON.stringify(this.user));
  },
};
</script>

<style>
</style>
