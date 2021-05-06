<template>
  <div>
    <Header />

    <div v-if="loaded" class="bg-gray-200">
      <div class="flex flex-col sm:flex-row sm:justify-around">
        <admin-side-bar v-model:activeTab="activeTab" :logo="logo" />
        <div class="w-full">
          <user
            v-if="activeTab === 1"
            :users="users"
            v-model:isOpenModal="isOpenModalAddUser"
            @search-user-input="searchUserInput"
            @search-invite-user-input="searchInviteUserInput"
            @select-users-per-page="selectUsersPerPage"
            @on-confirm-edit-user="onConfirmEditUser"
          />
          <card
            v-if="activeTab === 2"
            :cards="cards"
            :users="users"
            :is-show-button-add="true"
            v-model:isOpenModal="isOpenModalAddCard"
            @search-card-input="searchCardInput"
            @select-cards-per-page="selectCardsPerPage"
            @on-click-update-activate-card="onToggleCard"
            @on-change-card-user="onChangeCardUser"
            @on-click-remove-card="onClickRemoveCard"
          />
          <company
            v-if="activeTab === 3"
            :companies="companies"
            :is-show-button-add="true"
            v-model:isOpenModal="isOpenModalAddCompany"
            @select-companies-per-page="selectCompaniesPerPage"
            @select-company-input="searchCompanyInput"
            @on-confirm-edit-company="onConfirmEditCompany"
            @on-click-remove-company="onClickRemoveCompany"
          />

          <modal
            v-model:open="isOpenModalAddUser"
            header="Create User"
            :is-hidden-button-confirm="false"
            @confirm-modal="confirmModalAddUser"
          >
            <template #modal-body>
              <div class="flex flex-row items-center justify-between py-4">
                <label for="name">Full Name</label>
                <input
                  type="text"
                  placeholder="Full name"
                  class="w-4/5 text-sm bg-gray-200 text-gray-800 rounded h-10 p-3 focus:outline-none"
                  v-model="user.name"
                />
              </div>
              <div class="flex flex-row items-center justify-between py-4">
                <label for="email">Email</label>
                <input
                  type="email"
                  placeholder="Email"
                  class="w-4/5 text-sm bg-gray-200 text-gray-800 rounded h-10 p-3 focus:outline-none"
                  v-model="user.email"
                />
              </div>
              <div class="flex flex-row items-center justify-between py-4">
                <label for="username">Username</label>
                <input
                  type="text"
                  placeholder="Username"
                  class="w-4/5 text-sm bg-gray-200 text-gray-800 rounded h-10 p-3 focus:outline-none"
                  v-model="user.username"
                />
              </div>
              <div class="flex flex-row items-center justify-between py-4">
                <label for="password">Password</label>
                <input
                  type="password"
                  placeholder="Password"
                  class="w-4/5 text-sm bg-gray-200 text-gray-800 rounded h-10 p-3 focus:outline-none"
                  v-model="user.password"
                />
              </div>
              <div class="flex flex-row items-center justify-between py-4">
                <label for="address">Address</label>
                <input
                  type="text"
                  placeholder="Address"
                  class="w-4/5 text-sm bg-gray-200 text-gray-800 rounded h-10 p-3 focus:outline-none"
                  v-model="user.address"
                />
              </div>
              <div class="flex flex-row items-center justify-between py-4">
                <label for="phone_number">Phone number</label>
                <input
                  type="text"
                  placeholder="Phone number"
                  class="w-4/5 text-sm bg-gray-200 text-gray-800 rounded h-10 p-3 focus:outline-none"
                  v-model="user.phone_number"
                />
              </div>
              <div class="flex flex-row items-center justify-between py-4">
                <label for="type">Type</label>
                <select
                  class="w-4/5 h-10 rounded border block appearance-none bg-white border-gray-400 text-gray-700 py-2 px-4 pr-8 leading-tight focus:outline-none focus:bg-white focus:border-gray-500"
                  v-model="user.type"
                >
                  <option value="1">Personal</option>
                  <option value="2">Business</option>
                </select>
              </div>
              <div class="flex flex-row items-center justify-between py-4">
                <label for="role">Role</label>
                <select
                  class="w-4/5 h-10 rounded border block appearance-none bg-white border-gray-400 text-gray-700 py-2 px-4 pr-8 leading-tight focus:outline-none focus:bg-white focus:border-gray-500"
                  v-model="user.role"
                >
                  <option value="1">Standard</option>
                  <option value="2">Company Member</option>
                  <option value="3">Company Manager</option>
                  <option value="4">Admin</option>
                </select>
              </div>
            </template>
          </modal>

          <modal
            v-model:open="isOpenModalAddCard"
            header="Create Card"
            :is-hidden-button-confirm="false"
            @confirm-modal="confirmModalAddCard"
          >
            <template #modal-body>
              <div class="flex flex-row items-center justify-between py-4">
                <label for="code">Code</label>
                <input
                  type="text"
                  placeholder="Card code"
                  class="w-4/5 text-sm bg-gray-200 text-gray-800 rounded h-10 p-3 focus:outline-none"
                  v-model="card.code"
                />
              </div>
            </template>
          </modal>

          <modal
            v-model:open="isOpenModalAddCompany"
            header="Create Company"
            :is-hidden-button-confirm="false"
            @confirm-modal="confirmModalAddCompany"
          >
            <template #modal-body>
              <div class="flex flex-row items-center justify-between py-4">
                <label for="name">Name</label>
                <input
                  type="text"
                  placeholder="Search teams or members"
                  class="w-4/5 text-sm bg-gray-200 text-gray-800 rounded h-10 p-3 focus:outline-none"
                  v-model="company.name"
                />
              </div>
              <div class="flex flex-row items-center justify-between py-4">
                <label for="address">Address</label>
                <input
                  type="text"
                  placeholder="Address"
                  class="w-4/5 text-sm bg-gray-200 text-gray-800 rounded h-10 p-3 focus:outline-none"
                  v-model="company.address"
                />
              </div>
              <div class="flex flex-row items-center justify-between py-4">
                <label for="website">website</label>
                <input
                  type="text"
                  placeholder="Website"
                  class="w-4/5 text-sm bg-gray-200 text-gray-800 rounded h-10 p-3 focus:outline-none"
                  v-model="company.website"
                />
              </div>
              <div class="flex flex-row items-center justify-between py-4">
                <label for="hotline">Hotline</label>
                <input
                  type="text"
                  placeholder="Hotline"
                  class="w-4/5 text-sm bg-gray-200 text-gray-800 rounded h-10 p-3 focus:outline-none"
                  v-model="company.hotline"
                />
              </div>
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
import AdminSideBar from "@/components/AdminSideBar";
import Loading from "@/components/Loading";
import User from "@/components/User";
import Card from "@/components/Card";
import Company from "@/components/Company";
import Debounce from "@/mixins/debounce";
import Modal from "@/components/Modal";

export default {
  name: "Admin",

  data() {
    return {
      logo: "http://localhost:8001/public/logos/default_logo.png",
      loaded: false,
      activeTab: 1,
      isOpenModalAddUser: false,
      isOpenModalAddCard: false,
      isOpenModalAddCompany: false,
      user: {
        email: "",
        password: "",
        username: "",
        name: "",
        address: "",
        phone_number: "",
        type: 1,
        role: 1,
      },
      card: {
        code: "",
      },
      company: {
        name: "",
        address: "",
        website: "",
        hotline: "",
      },
    };
  },

  mixins: [Debounce],

  components: {
    Header,
    AdminSideBar,
    Loading,
    Modal,
    User,
    Card,
    Company,
  },

  computed: {
    ...mapState("users", ["users"]),
    ...mapState("cards", ["cards"]),
    ...mapState("companies", ["companies"]),
  },

  methods: {
    ...mapActions("users", ["getUsersList", "createUser", "updateUser"]),
    ...mapActions("cards", [
      "getCardsList",
      "createCard",
      "updateCard",
      "deleteCard",
    ]),
    ...mapActions("companies", [
      "getCompaniesList",
      "createCompany",
      "updateCompany",
      "deleteCompany",
    ]),

    searchUserInput(event) {
      const value = event.target.value.trim();
      if (!value) return;

      this.handleDebounceSearch(this.getUsersList, { q: value }, 1000);
    },

    searchCardInput(event) {
      const value = event.target.value.trim();
      if (!value) return;

      this.handleDebounceSearch(this.getCardsList, { q: value }, 1000);
    },

    searchCompanyInput(event) {
      const value = event.target.value.trim();
      if (!value) return;

      this.handleDebounceSearch(this.getCompaniesList, { q: value }, 1000);
    },

    selectUsersPerPage(event) {
      if (!event.target.value) return;

      this.getUsersList({ per_page: event.target.value });
    },

    selectCardsPerPage(event) {
      if (!event.target.value) return;

      this.getCardsList({ per_page: event.target.value });
    },

    selectCompaniesPerPage(event) {
      if (!event.target.value) return;

      this.getCompaniesList({ per_page: event.target.value });
    },

    confirmModalAddCard() {
      if (this.card.code === "") return;

      this.createCard(this.card).then(() => {
        this.isOpenModalAddCard = false;
        this.card = {};
      });
    },

    confirmModalAddCompany() {
      if (this.company === {}) return;

      this.createCompany(this.company).then(() => {
        this.isOpenModalAddCompany = false;
        this.company = {};
      });
    },

    confirmModalAddUser() {
      if (this.user === {}) return;

      this.createUser(this.user).then(() => {
        this.isOpenModalAddUser = false;
        this.user = {};
      });
    },

    onConfirmEditUser(user) {
      if (user === {}) return;

      const data = {
        params: user.username,
        body: user,
      };

      this.updateUser(data);
    },

    onToggleCard(card, value) {
      const data = {
        params: card.code,
        body: value,
      };
      this.updateCard(data);
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

    onConfirmEditCompany(company) {
      if (company === {}) return;

      const data = {
        params: company.id,
        body: company,
      };

      this.updateCompany(data);
    },

    onClickRemoveCompany(company) {
      if (company === {}) return;

      this.deleteCompany(company.id);
    },

    onClickRemoveCard(card) {
      if (card === {}) return;

      this.deleteCard(card.code);
    },
  },

  async created() {
    await this.getUsersList();
    await this.getCardsList();
    await this.getCompaniesList();

    setTimeout(() => {
      this.loaded = true;
    }, 500);
  },
};
</script>
