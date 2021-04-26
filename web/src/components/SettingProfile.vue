<template>
  <section class="pt-40 pb-8 bg-gray-100 bg-opacity-50">
    <div class="mx-auto container max-w-2xl md:w-3/4 shadow-md">
      <div
        class="bg-gray-100 p-4 border-t-2 bg-opacity-5 border-indigo-400 rounded-t"
      >
        <div class="max-w-sm mx-auto md:w-full md:mx-0">
          <div class="inline-flex items-center space-x-4">
            <img
              class="w-10 h-10 object-cover rounded-full"
              alt="User avatar"
              :src="updateUser.avatar"
            />

            <h1 class="text-gray-600">{{ updateUser.name }}</h1>
          </div>
        </div>
      </div>
      <div class="bg-white space-y-6">
        <div
          class="md:inline-flex space-y-4 md:space-y-0 w-full p-4 text-gray-500 items-center"
        >
          <h2 class="md:w-1/3 max-w-sm mx-auto">Account</h2>
          <div class="md:w-2/3 max-w-sm mx-auto">
            <div>
              <label class="text-sm text-gray-400">Username</label>
              <div class="w-full inline-flex border">
                <div class="pt-2 w-1/12 bg-gray-100 bg-opacity-50">
                  <svg
                    fill="none"
                    class="w-6 text-gray-400 mx-auto"
                    viewBox="0 0 24 24"
                    stroke="currentColor"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z"
                    />
                  </svg>
                </div>
                <input
                  type="username"
                  class="w-11/12 focus:outline-none focus:text-gray-600 p-2"
                  :placeholder="updateUser.username"
                  disabled
                />
              </div>
            </div>
            <div>
              <label class="text-sm text-gray-400">Email</label>
              <div class="w-full inline-flex border">
                <div class="pt-2 w-1/12 bg-gray-100 bg-opacity-50">
                  <ICOEmail />
                </div>
                <input
                  type="email"
                  class="w-11/12 focus:outline-none focus:text-gray-600 p-2"
                  :placeholder="updateUser.email"
                  disabled
                />
              </div>
            </div>
          </div>
        </div>

        <hr />
        <div
          class="md:inline-flex space-y-4 md:space-y-0 w-full p-4 text-gray-500 items-center"
        >
          <h2 class="md:w-1/3 max-w-sm mx-auto">Avatar</h2>
          <div class="md:w-2/3 max-w-sm mx-auto">
            <div>
              <div class="py-3 center mx-auto">
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
                    <input type="file" class="hidden" :accept="accept" />
                  </label>
                </div>
              </div>
            </div>
          </div>
        </div>

        <hr />
        <div
          class="md:inline-flex space-y-4 md:space-y-0 w-full p-4 text-gray-500 items-center"
        >
          <h2 class="md:w-1/3 mx-auto max-w-sm">Personal info</h2>
          <div class="md:w-2/3 mx-auto max-w-sm space-y-5">
            <div>
              <label class="text-sm text-gray-400">Full name</label>
              <div class="w-full inline-flex border">
                <div class="w-1/12 pt-2 bg-gray-100">
                  <ICOPerson />
                </div>
                <input
                  type="text"
                  class="w-11/12 focus:outline-none focus:text-gray-600 p-2"
                  v-model="updateUser.name"
                />
              </div>
            </div>
            <div>
              <label class="text-sm text-gray-400">Address</label>
              <div class="w-full inline-flex border">
                <div class="pt-2 w-1/12 bg-gray-100">
                  <svg
                    fill="none"
                    class="w-6 text-gray-400 mx-auto"
                    viewBox="0 0 24 24"
                    stroke="currentColor"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M12 18h.01M8 21h8a2 2 0 002-2V5a2 2 0 00-2-2H8a2 2 0 00-2 2v14a2 2 0 002 2z"
                    />
                  </svg>
                </div>
                <input
                  type="text"
                  class="w-11/12 focus:outline-none focus:text-gray-600 p-2"
                  v-model="updateUser.address"
                />
              </div>
            </div>
            <div>
              <label class="text-sm text-gray-400">Phone number</label>
              <div class="w-full inline-flex border">
                <div class="pt-2 w-1/12 bg-gray-100">
                  <ICOPhoneNumber />
                </div>
                <input
                  type="text"
                  class="w-11/12 focus:outline-none focus:text-gray-600 p-2"
                  v-model="updateUser.phone_number"
                />
              </div>
            </div>
          </div>
        </div>

        <hr />
        <div
          class="md:inline-flex w-full space-y-4 md:space-y-0 p-8 text-gray-500 items-center"
        >
          <h2 class="md:w-4/12 max-w-sm mx-auto">Change password</h2>

          <div class="md:w-2/3 mx-auto max-w-sm space-y-5">
            <div class="w-full inline-flex border-b">
              <div class="w-1/12 pt-2">
                <ICOLock />
              </div>
              <input
                type="password"
                class="w-11/12 focus:outline-none focus:text-gray-600 p-2 ml-4"
                placeholder="New"
              />
            </div>
          </div>
        </div>

        <hr />
        <div class="md:inline-flex w-full pb-8 text-gray-500">
          <button
            class="text-white max-w-xs mx-auto rounded-md bg-indigo-400 py-2 px-4 inline-flex items-center focus:outline-none md:float-right"
          >
            <ICOUpdate />
            Update
          </button>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
import { mapActions } from "vuex";
import ICOEmail from "@/assets/icons/ICOEmail";
import ICOPerson from "@/assets/icons/ICOPerson";
import ICOPhoneNumber from "@/assets/icons/ICOPhoneNumber";
import ICOLock from "@/assets/icons/ICOLock";
import ICOUpdate from "@/assets/icons/ICOUpdate";

export default {
  name: "SettingProfile",

  props: {
    user: {
      type: Object,
      require: true,
    },
  },

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
    ICOEmail,
    ICOPerson,
    ICOPhoneNumber,
    ICOLock,
    ICOUpdate,
  },

  methods: {
    ...mapActions("users", ["updateCurrentUser"]),

    onClickUpdate() {
      this.updateCurrentUser(this.updateUser);
    },
  },

  created() {
    this.updateUser = this.user;
  },
};
</script>
