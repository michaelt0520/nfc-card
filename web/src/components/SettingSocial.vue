<template>
  <section class="py-8 pb-40 bg-gray-100 bg-opacity-50">
    <div class="mx-auto container max-w-2xl md:w-3/4 shadow-md">
      <div
        class="bg-gray-100 p-4 border-t-2 bg-opacity-5 border-indigo-400 rounded-t"
      >
        <div class="max-w-sm mx-auto md:w-full md:mx-0">
          <div class="inline-flex items-center space-x-4">
            <h1 class="text-gray-600">Social informations</h1>
          </div>
        </div>
      </div>

      <div class="bg-white space-y-6">
        <div class="flex pt-4 md:w-11/12 mx-auto">
          <div class="relative inline-flex">
            <ICOArrow
              class="shadow w-2 h-2 absolute top-0 right-0 m-4 pointer-events-none"
            />
            <select
              class="shadow border border-gray-300 rounded-full text-gray-600 h-10 pl-5 pr-10 bg-white hover:border-gray-400 focus:outline-none appearance-none"
              v-model="addSocial.name"
            >
              <option v-for="social in socialsJson" :key="social.index">
                {{ social.name }}
              </option>
            </select>
          </div>
          <input
            class="shadow appearance-none border border-gray-300 rounded-full w-full py-2 px-3 mx-4 h-10 text-grey-darker"
            placeholder="Add Information"
            v-model="addSocial.data"
          />
          <button
            class="flex-no-shrink p-2 border-2 rounded text-teal border-teal hover:text-white hover:bg-teal"
            @click="onClickAddInfo"
          >
            Add
          </button>
        </div>

        <div class="md:w-11/12 mx-auto py-5">
          <div
            class="flex mb-4 items-center"
            v-for="info in socialsList"
            :key="info.id"
          >
            <div class="w-full text-grey-darkest">
              <p class="flex flex-col">
                <label class="text-sm text-gray-400">{{ info.name }}</label>
                {{ info.data }}
              </p>
            </div>
            <label class="flex items-center">
              <input
                class="relative w-10 h-5 transition-all duration-200 ease-in-out bg-gray-400 rounded-full shadow-inner outline-none appearance-none"
                type="checkbox"
                :checked="info.visibled"
                @click="onClickUpdateVisibleInfo(info)"
              />
            </label>
            <button
              class="flex-no-shrink p-2 ml-2 border-2 rounded text-red border-red hover:text-red-500 hover:bg-red"
              @click="onClickDeleteInfo(info.id)"
            >
              <ICODelete />
            </button>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
import { mapActions } from "vuex";
import SocialsJson from "@/assets/socials.json";
import ICODelete from "@/assets/icons/ICODelete";
import ICOArrow from "@/assets/icons/ICOArrow";

export default {
  name: "SettingSocial",

  props: {
    socialsList: {
      type: Array,
      require: true,
    },
  },

  data() {
    return {
      addSocial: {
        name: "",
        data: "",
      },

      socialsJson: SocialsJson.socials,
    };
  },

  components: {
    ICODelete,
    ICOArrow,
  },

  methods: {
    ...mapActions("users", ["createInfo", "updateInfo", "deleteInfo"]),

    onClickAddInfo() {
      this.createInfo(this.addSocial);
    },

    onClickUpdateVisibleInfo(info) {
      const data = {
        id: info.id,
        value: {
          visibled: !info.visibled,
        },
      };
      this.updateInfo(data);
    },

    onClickDeleteInfo(infoId) {
      this.deleteInfo(infoId);
    },
  },
};
</script>

<style>
input:before {
  content: "";
  position: absolute;
  width: 1.25rem;
  height: 1.25rem;
  border-radius: 50%;
  top: 0;
  left: 0;
  transform: scale(1.1);
  box-shadow: 0 0.125rem 0.5rem rgba(0, 0, 0, 0.2);
  background-color: white;
  transition: 0.2s ease-in-out;
}

input:checked {
  background-color: #7f9cf5;
}

input:checked:before {
  left: 1.25rem;
}
</style>
