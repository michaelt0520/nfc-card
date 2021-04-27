<template>
  <div>
    <Header />

    <div v-if="loaded" class="bg-gray-200">
      <div class="flex flex-col sm:flex-row sm:justify-around">
        <company-side-bar v-model:activeTab="activeTab" :logo="company.logo" />
        <div class="w-full">
          <user
            :users="company.users"
            v-model:isClickAddUser="isClickAddUser"
          />
        </div>
      </div>

      <personal-users-list v-if="isClickAddUser" />
    </div>
    <div v-else>
      <loading />
    </div>
  </div>
</template>

<script>
import { mapState, mapActions } from "vuex";
import Header from "@/components/Header";
import CompanySideBar from "../components/CompanySideBar";
import Loading from "@/components/Loading";
import User from "@/components/User";
import PersonalUsersList from "@/components/PersonalUsersList";

export default {
  name: "Company",

  components: {
    Header,
    CompanySideBar,
    Loading,
    User,
    PersonalUsersList,
  },

  data() {
    return {
      loaded: false,
      activeTab: 1,
      isClickAddUser: false,
    };
  },

  computed: {
    ...mapState("companies", ["company"]),
  },

  methods: {
    ...mapActions("companies", ["getCurrentCompany"]),
  },

  created() {
    this.getCurrentCompany().then(() => {
      setTimeout(() => {
        this.loaded = true;
      }, 500);
    });
  },
};
</script>

<style>
</style>
