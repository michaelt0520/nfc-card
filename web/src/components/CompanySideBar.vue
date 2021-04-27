<template>
  <div class="w-64 h-screen bg-white">
    <div class="flex items-center justify-center mt-10">
      <img class="h-16" :src="logo" />
    </div>

    <nav class="mt-10">
      <a
        class="flex items-center mt-5 py-2 px-8 border-r-4"
        :class="activeClass(1)"
        @click="$emit('update:activeTab', 1)"
        href="#"
      >
        <ICOAccount class="h-5 w-5" />

        <span class="mx-4 font-medium">Accounts</span>
      </a>

      <a
        class="flex items-center mt-5 py-2 px-8 border-r-4"
        :class="activeClass(2)"
        @click="$emit('update:activeTab', 2)"
        href="#"
      >
        <ICOCard class="h-5 w-5" />
        <span class="mx-4 font-medium">Cards</span>
      </a>

      <a
        class="flex items-center mt-5 py-2 px-8 border-r-4"
        :class="activeClass(3)"
        @click="$emit('update:activeTab', 3)"
        href="#"
      >
        <ICOSetting class="h-5 w-5" />

        <span class="mx-4 font-medium">Settings</span>
      </a>
    </nav>

    <div class="absolute bottom-0 my-10">
      <a
        class="flex items-center py-2 px-8 text-gray-500 hover:text-gray-600"
        href="#"
      >
        <ICOQuestionMark class="h-5 w-5" />

        <span class="mx-4 font-medium">Support</span>
      </a>
    </div>
  </div>
</template>

<script>
import ICOCard from "@/assets/icons/ICOCard";
import ICOAccount from "@/assets/icons/ICOAccount";
import ICOSetting from "@/assets/icons/ICOSetting";
import ICOQuestionMark from "@/assets/icons/ICOQuestionMark";

export default {
  name: "CompanySideBar",

  props: {
    logo: {
      type: String,
      requre: true,
    },

    activeTab: {
      type: Number,
      require: true,
      default: 1,
    },
  },

  components: {
    ICOCard,
    ICOAccount,
    ICOSetting,
    ICOQuestionMark,
  },

  methods: {
    activeClass(activeTab) {
      return this.activeTab === activeTab
        ? "bg-gray-200 text-gray-700 border-gray-700"
        : "text-gray-600 border-white hover:bg-gray-200 hover:text-gray-700 hover:border-gray-700";
    },
  },
};
</script>
