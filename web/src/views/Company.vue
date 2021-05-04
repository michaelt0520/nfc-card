<template>
  <div>
    <Header />

    <div v-if="loaded" class="bg-gray-200">
      <div class="flex flex-col sm:flex-row sm:justify-around">
        <company-side-bar v-model:activeTab="activeTab" :logo="company.logo" />
        <div class="w-full">
          <user
            v-if="activeTab === 1"
            :users="users"
            :invite-users="searchUsers"
            @search-user-input="searchUserInput"
            @search-invite-user-input="searchInviteUserInput"
            @select-users-per-page="selectUsersPerPage"
          />
          <card
            v-if="activeTab === 2"
            :cards="cards"
            @search-card-input="searchCardInput"
            @select-cards-per-page="selectCardsPerPage"
          />
          <setting-company v-if="activeTab === 3" :company="company" />
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
import CompanySideBar from "../components/CompanySideBar";
import Loading from "@/components/Loading";
import User from "@/components/User";
import Debounce from "@/mixins/debounce";
import Card from "@/components/Card";
import SettingCompany from "@/components/SettingCompany";

export default {
  name: "Company",

  data() {
    return {
      loaded: false,
      activeTab: 1,
    };
  },

  mixins: [Debounce],

  components: {
    Header,
    CompanySideBar,
    Loading,
    User,
    Card,
    SettingCompany,
  },

  computed: {
    ...mapState("companies", ["company", "users", "cards"]),
    ...mapState("users", ["searchUsers"]),
  },

  methods: {
    ...mapActions("companies", [
      "getCurrentCompany",
      "getCurrentCompanyUsers",
      "getCurrentCompanyCards",
    ]),

    ...mapActions("users", ["searchUser"]),

    searchUserInput(fillterBy, event) {
      const value = event.target.value.trim();
      if (!value) return;

      this.handleDebounceSearch(
        this.getCurrentCompanyUsers,
        { [fillterBy]: value },
        1000
      );
    },

    searchCardInput(event) {
      const value = event.target.value.trim();
      if (!value) return;

      this.handleDebounceSearch(
        this.getCurrentCompanyCards,
        { code: value },
        1000
      );
    },

    searchInviteUserInput(event) {
      const value = event.target.value.trim();
      if (!value) return;

      this.handleDebounceSearch(this.searchUser, { email: value }, 500);
    },

    selectUsersPerPage(event) {
      if (!event.target.value) return;

      this.getCurrentCompanyUsers({ per_page: event.target.value });
    },

    selectCardsPerPage(event) {
      if (!event.target.value) return;

      this.getCurrentCompanyCards({ per_page: event.target.value });
    },
  },

  async created() {
    await this.getCurrentCompany();
    await this.getCurrentCompanyUsers();
    await this.getCurrentCompanyCards();

    setTimeout(() => {
      this.loaded = true;
    }, 500);
  },
};
</script>

<style>
</style>
