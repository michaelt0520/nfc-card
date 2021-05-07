<template>
  <div>
    <Header />

    <div v-if="loaded" class="bg-gray-100 pt-14">
      <div class="flex justify-self-center pt-40 pb-8 mx-auto h-screen">
        <div class="w-1/2 flex flex-row mx-auto">
          <div class="w-1/4">
            <div class="mx-auto container w-60 shadow-md">
              <div
                class="bg-gray-100 p-4 border-t-2 bg-opacity-5 border-indigo-400 rounded-t"
              >
                <div class="max-w-sm mx-auto md:w-full md:mx-0">
                  <div class="inline-flex items-center space-x-4">
                    <h1 class="text-gray-600">Menu</h1>
                  </div>
                </div>
              </div>

              <div class="bg-white space-y-6">
                <div class="mx-auto">
                  <div class="w-full flex flex-col">
                    <button
                      class="border-t-2 text-center border-gray-100 font-medium text-gray-600 py-4 px-4 w-full block hover:bg-gray-100 transition duration-150"
                      @click="activeTab = 1"
                    >
                      Profile
                    </button>
                    <button
                      class="border-t-2 text-center border-gray-100 font-medium text-gray-600 py-4 px-4 w-full block hover:bg-gray-100 transition duration-150"
                      @click="activeTab = 2"
                    >
                      Password
                    </button>
                    <button
                      class="border-t-2 text-center border-gray-100 font-medium text-gray-600 py-4 px-4 w-full block hover:bg-gray-100 transition duration-150"
                      @click="activeTab = 3"
                    >
                      Cards
                    </button>
                    <button
                      class="border-t-2 text-center border-gray-100 font-medium text-gray-600 py-4 px-4 w-full block hover:bg-gray-100 transition duration-150"
                      @click="activeTab = 4"
                    >
                      Informations
                    </button>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div class="w-3/4">
            <setting-profile :user="user" v-if="activeTab === 1" />
            <setting-password v-if="activeTab === 2" />
            <setting-card :cards-list="user.cards" v-if="activeTab === 3" />
            <setting-social
              :socials-list="user.informations"
              v-if="activeTab === 4"
            />
          </div>
        </div>
      </div>
    </div>
    <div v-else>
      <loading />
    </div>
  </div>
</template>

<script>
import { mapState, mapActions } from "vuex";
import Header from "@/components/Header";
import SettingSocial from "@/components/SettingSocial";
import SettingProfile from "@/components/SettingProfile";
import SettingCard from "@/components/SettingCard";
import SettingPassword from "@/components/SettingPassword";
import Loading from "@/components/Loading";

export default {
  name: "Setting",

  components: {
    Header,
    SettingProfile,
    SettingSocial,
    SettingCard,
    SettingPassword,
    Loading,
  },

  data() {
    return {
      activeTab: 1,
      loaded: false,
    };
  },

  computed: {
    ...mapState("users", ["user"]),
  },

  methods: {
    ...mapActions("users", ["getCurrentUser"]),
  },

  created() {
    this.getCurrentUser().then(() => {
      setTimeout(() => {
        this.loaded = true;
      }, 500);
    });
  },
};
</script>
