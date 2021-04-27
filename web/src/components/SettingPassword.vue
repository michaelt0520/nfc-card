<template>
  <section class="py-8 bg-gray-100 bg-opacity-50">
    <div class="mx-auto container max-w-2xl md:w-3/4 shadow-md">
      <div
        class="bg-gray-100 p-4 border-t-2 bg-opacity-5 border-indigo-400 rounded-t"
      >
        <div class="max-w-sm mx-auto md:w-full md:mx-0">
          <div class="inline-flex items-center space-x-4">
            <h1 class="text-gray-600">Update Password</h1>
          </div>
        </div>
      </div>
      <div class="bg-white space-y-6">
        <div
          class="md:inline-flex w-full space-y-4 md:space-y-0 p-8 text-gray-500 items-center"
        >
          <h2 class="md:w-1/3 max-w-sm mx-auto">Change password</h2>
          <div class="md:w-2/3 max-w-sm mx-auto">
            <div>
              <label class="text-sm text-gray-400">Old Password</label>
              <div class="w-full inline-flex border-b">
                <div class="w-1/12 pt-2">
                  <ICOLock />
                </div>
                <input
                  type="password"
                  class="w-11/12 focus:outline-none focus:text-gray-600 p-2 ml-4"
                  placeholder="Old password"
                  v-model="password.old_password"
                />
              </div>
            </div>
            <div>
              <label class="text-sm text-gray-400">New Password</label>
              <div class="w-full inline-flex border-b">
                <div class="w-1/12 pt-2">
                  <ICOLock />
                </div>
                <input
                  type="password"
                  class="w-11/12 focus:outline-none focus:text-gray-600 p-2 ml-4"
                  placeholder="New password"
                  v-model="password.new_password"
                />
              </div>
            </div>
            <div>
              <label class="text-sm text-gray-400">Password Confirmation</label>
              <div class="w-full inline-flex border-b">
                <div class="w-1/12 pt-2">
                  <ICOLock />
                </div>
                <input
                  type="password"
                  class="w-11/12 focus:outline-none focus:text-gray-600 p-2 ml-4"
                  placeholder="Password confirmation"
                  v-model="password.password_confirmation"
                />
              </div>
            </div>
          </div>
        </div>

        <hr />
        <div class="md:inline-flex w-full pb-8 text-gray-500">
          <button
            class="text-white max-w-xs mx-auto rounded-md bg-indigo-400 py-2 px-4 inline-flex items-center focus:outline-none md:float-right"
            @click="onClickUpdatePassword"
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
import ICOLock from "@/assets/icons/ICOLock";
import ICOUpdate from "@/assets/icons/ICOUpdate";

export default {
  name: "SettingPassword",

  data() {
    return {
      password: {
        old_password: "",
        new_password: "",
        password_confirmation: "",
      },
    };
  },

  components: {
    ICOLock,
    ICOUpdate,
  },

  methods: {
    ...mapActions("users", ["updateCurrentUserPassword"]),

    onClickUpdatePassword() {
      this.updateCurrentUserPassword(this.password);
    },
  },
};
</script>
