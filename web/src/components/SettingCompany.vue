<template>
  <section class="py-8 bg-gray-100 bg-opacity-50">
    <div class="mx-auto container max-w-2xl md:w-3/4 shadow-md">
      <div
        class="bg-gray-100 p-4 border-t-2 bg-opacity-5 border-indigo-400 rounded-t"
      >
        <div class="max-w-sm mx-auto md:w-full md:mx-0">
          <div class="inline-flex items-center space-x-4">
            <h1>Setting</h1>
          </div>
        </div>
      </div>
      <div class="bg-white space-y-6">
        <div
          class="md:inline-flex space-y-4 md:space-y-0 w-full p-4 text-gray-500 items-center"
        >
          <h2 class="md:w-1/3 max-w-sm mx-auto">Logo</h2>
          <div class="md:w-2/3 max-w-sm mx-auto">
            <div>
              <div class="py-3 center mx-auto">
                <div
                  class="bg-white px-4 py-5 rounded-lg shadow-lg text-center w-48"
                >
                  <div class="mb-4">
                    <img
                      class="w-auto mx-auto rounded-full object-cover object-center"
                      :src="updateCompany.logo"
                      alt="Avatar Upload"
                    />
                  </div>
                  <label class="cursor-pointer mt-6">
                    <span
                      class="mt-2 leading-normal px-4 py-2 bg-blue-500 text-white text-sm rounded-full"
                      >Select Logo</span
                    >
                    <input
                      type="file"
                      ref="file"
                      class="hidden"
                      :accept="accept"
                      @change="handleLogoUpload"
                    />
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
          <h2 class="md:w-1/3 mx-auto max-w-sm">Information</h2>
          <div class="md:w-2/3 mx-auto max-w-sm space-y-5">
            <div>
              <label class="text-sm text-gray-400">Name</label>
              <div class="w-full inline-flex border">
                <div class="w-1/12 pt-2 bg-gray-100">
                  <ICOPerson />
                </div>
                <input
                  type="text"
                  class="w-11/12 focus:outline-none focus:text-gray-600 p-2"
                  v-model="updateCompany.name"
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
                  v-model="updateCompany.address"
                />
              </div>
            </div>
            <div>
              <label class="text-sm text-gray-400">Website</label>
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
                  v-model="updateCompany.website"
                />
              </div>
            </div>
            <div>
              <label class="text-sm text-gray-400">Hotline</label>
              <div class="w-full inline-flex border">
                <div class="pt-2 w-1/12 bg-gray-100">
                  <ICOPhoneNumber />
                </div>
                <input
                  type="text"
                  class="w-11/12 focus:outline-none focus:text-gray-600 p-2"
                  v-model="updateCompany.hotline"
                />
              </div>
            </div>
          </div>
        </div>

        <hr />
        <div class="md:inline-flex w-full pb-8 text-gray-500">
          <button
            class="text-white max-w-xs mx-auto rounded-md bg-indigo-400 py-2 px-4 inline-flex items-center focus:outline-none md:float-right"
            @click="onClickUpdateCompany"
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
import ICOPerson from "@/assets/icons/ICOPerson";
import ICOPhoneNumber from "@/assets/icons/ICOPhoneNumber";
import ICOUpdate from "@/assets/icons/ICOUpdate";

export default {
  name: "SettingProfile",

  props: {
    company: {
      type: Object,
      require: true,
    },
  },

  data() {
    return {
      updateCompany: {},
    };
  },

  components: {
    ICOPerson,
    ICOPhoneNumber,
    ICOUpdate,
  },

  methods: {
    ...mapActions("companies", ["updateCurrentCompany", "createLogo"]),

    onClickUpdateCompany() {
      const data = {
        logo: this.updateCompany.logo,
        name: this.updateCompany.name,
        address: this.updateCompany.address,
        website: this.updateCompany.website,
        hotline: this.updateCompany.hotline,
      };
      this.updateCurrentCompany(data);
      this.$emit("update:isDisplay", true);
    },

    handleLogoUpload() {
      let formData = new FormData();
      formData.append("file", this.$refs.file.files[0]);

      this.createLogo(formData).then((res) => {
        this.updateCompany.logo = res.data.result;
      });
    },
  },

  created() {
    this.updateCompany = { ...this.company };
  },
};
</script>
