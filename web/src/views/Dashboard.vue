<template>
  <div class="flex h-screen bg-b">
    <header-admin />

    <div class="flex flex-row mt-16 w-full">
      <div class="bg-gray-800 shadow-xl w-48 z-10 relative bottom-0">
        <div class="fixed left-0 top-0 text-left justify-between w-48 mt-16">
          <ul class="list-reset text-white flex flex-col text-left px-2">
            <li
              @click="activeTab = 1"
              class="tab cursor-pointer py-3 border-b-2 border-gray-800 text-gray-300 hover:text-white"
              :class="[activeTab === 1 ? 'active' : '']">Card</li>
            <li
              @click="activeTab = 2"
              class="tab cursor-pointer py-3 border-b-2 border-gray-800 text-gray-300 hover:text-white"
              :class="[activeTab === 2 ? 'active' : '']">User</li>
            <li
              @click="activeTab = 3"
              class="tab cursor-pointer py-3 border-b-2 border-gray-800 text-gray-300 hover:text-white"
              :class="[activeTab === 3 ? 'active' : '']">Company</li>
          </ul>
        </div>
      </div>

      <div
        v-if="activeTab === 1"
        class="flex-1 bg-gray-100">
        <div class="bg-gray-800 pt-3">
          <div class="rounded-tl-3xl bg-gradient-to-r from-blue-900 to-gray-800 p-4 shadow text-2xl text-white">
            Card
          </div>
        </div>
        <div class="flex flex-wrap">
          <card
            v-for="card in cards"
            :card="card"
            :key="card.ID"/>
        </div>
      </div>

      <div
        v-if="activeTab === 2"
        class="flex-1 bg-gray-100">
        <div class="bg-gray-800 pt-3">
          <div class="rounded-tl-3xl bg-gradient-to-r from-blue-900 to-gray-800 p-4 shadow text-2xl text-white">
            User
          </div>
        </div>
        <div class="flex flex-wrap">
          <user />
        </div>
      </div>

      <div
        v-if="activeTab === 3"
        class="flex-1 bg-gray-100">
        <div class="bg-gray-800 pt-3">
          <div class="rounded-tl-3xl bg-gradient-to-r from-blue-900 to-gray-800 p-4 shadow text-2xl text-white">
            Company
          </div>
        </div>
        <div class="flex flex-wrap">
          <company />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapState, mapActions } from 'vuex'
import Card from '../components/Card'
import HeaderAdmin from '../components/HeaderAdmin'
import User from '../components/User'
import Company from '../components/Company'

export default {
  name: 'Dashboard',

  components: {
    Card,
    HeaderAdmin,
    User,
    Company
  },

  computed: {
    ...mapState('cards', ['cards']),
  },

  data () {
    return {
      activeTab: 1
    }
  },

  methods: {
    ...mapActions('cards', ['getListCards']),
  },

  created() {
    this.getListCards()
  }
}
</script>

<style lang="scss" scoped>
.tab {
  &.active {
    border-color: rgba(236, 72, 153, var(--tw-border-opacity));
  }
}
</style>
