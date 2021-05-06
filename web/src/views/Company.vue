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
            v-model:isOpenModal="isOpenModalInviteUser"
            @search-user-input="searchUserInput"
            @search-invite-user-input="searchInviteUserInput"
            @select-users-per-page="selectUsersPerPage"
            @on-click-remove-user-from-company="onClickRemoveCompanyUser"
          />

          <card
            v-if="activeTab === 2"
            :cards="cards"
            :users="users"
            @search-card-input="searchCardInput"
            @select-cards-per-page="selectCardsPerPage"
            @on-click-update-activate-card="onUpdateCard"
            @on-change-card-user="onChangeCardUser"
          />

          <setting-company v-if="activeTab === 3" :company="company" />

          <modal
            v-model:open="isOpenModalInviteUser"
            header="Invite"
            :is-hidden-button-confirm="false"
            @confirm-modal="onConfirmInviteUser"
          >
            <template #modal-body>
              <input
                type="text"
                placeholder="Search teams or members"
                class="my-2 w-full text-sm bg-gray-200 text-gray-800 rounded h-10 p-3 focus:outline-none"
                v-model="inviteUserEmail"
              />
            </template>
          </modal>
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
import Modal from "@/components/Modal";

export default {
  name: "Company",

  data() {
    return {
      loaded: false,
      activeTab: 1,
      isOpenModalInviteUser: false,
      inviteUserEmail: "",
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
    Modal,
  },

  computed: {
    ...mapState("companies", ["company", "users", "cards"]),
  },

  methods: {
    ...mapActions("companies", [
      "getCurrentCompany",
      "getCurrentCompanyUsers",
      "getCurrentCompanyCards",
      "updateCard",
      "createCompanyUser",
      "updateCompanyUser",
    ]),

    searchUserInput(event) {
      const value = event.target.value.trim();
      if (!value) return;

      this.handleDebounceSearch(
        this.getCurrentCompanyUsers,
        { q: value },
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

    onConfirmInviteUser() {
      if (this.inviteUserEmail === "") return;

      this.createCompanyUser({ email: this.inviteUserEmail }).then(() => {
        (this.isOpenModalInviteUser = false), (this.inviteUserEmail = "");
      });
    },

    selectUsersPerPage(event) {
      if (!event.target.value) return;

      this.getCurrentCompanyUsers({ per_page: event.target.value });
    },

    selectCardsPerPage(event) {
      if (!event.target.value) return;

      this.getCurrentCompanyCards({ per_page: event.target.value });
    },

    onUpdateCard(card, value) {
      const data = {
        params: card.code,
        body: value,
      };
      this.updateCard(data);
    },

    onClickRemoveCompanyUser(user) {
      this.updateCompanyUser({ params: user.username });
    },

    onChangeCardUser(card, event) {
      const data = {
        params: card.code,
        body: {
          user_id: event.target.value,
        },
      };
      this.updateCard(data);
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
