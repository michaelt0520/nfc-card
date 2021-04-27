<template>
  <div class="py-64 bg-gray-100" v-if="loaded">
    <div
      class="bg-white relative shadow-xl w-5/6 md:w-4/6 lg:w-3/6 xl:w-2/6 mx-auto"
    >
      <div class="flex justify-center">
        <img
          :src="card.user.avatar"
          alt=""
          class="rounded-full mx-auto absolute -top-20 w-32 h-32 shadow-2xl border-4 border-white"
        />
      </div>

      <div class="mt-16">
        <h1 class="font-bold text-center text-3xl text-gray-900">
          {{ card.user.name }}
        </h1>
        <p
          v-if="card.user.company"
          class="text-center text-sm text-gray-400 font-medium"
        >
          <span class="font-bold">@Company: </span>
          <a :href="card.user.company.website">
            {{ card.user.company.name }}
          </a>
        </p>
        <p
          v-if="card.user.address"
          class="text-center text-sm text-gray-400 font-medium"
        >
          <span class="font-bold">@Address: </span>
          {{ card.user.address }}
        </p>
        <p
          v-if="card.user.phone_number"
          class="text-center text-sm text-gray-400 font-medium"
        >
          <span class="font-bold">@Phone: </span>
          {{ card.user.phone_number }}
        </p>
        <div class="my-5">
          <a
            :href="`mailto:${card.user.email}`"
            class="text-indigo-200 block text-center font-medium leading-6 px-6 py-3 bg-indigo-600"
            >Connect <span class="font-bold">{{ card.user.email }}</span></a
          >
        </div>

        <div class="w-full">
          <div class="mt-5 w-full flex flex-col">
            <a
              v-for="info in card.user.informations"
              :key="info"
              class="border-t-2 text-center border-gray-100 font-medium text-gray-600 py-4 px-4 w-full block hover:bg-gray-100 transition duration-150"
              :href="info.data"
            >
              {{ info.name }}
            </a>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapState, mapActions } from "vuex";

export default {
  name: "Card",

  data() {
    return {
      onConfirmEdit: false,
      loaded: false,
    };
  },

  computed: {
    ...mapState("cards", ["card"]),
  },

  methods: {
    ...mapActions("cards", ["getCard"]),
  },

  mounted() {
    if (process.env.NODE_ENV === "production")
      window.history.replaceState(
        {},
        "",
        `/personal/${this.card.User?.username}`
      );
  },

  created() {
    this.getCard(this.$route.params.code).finally(() => (this.loaded = true));
  },
};
</script>
