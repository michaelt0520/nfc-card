<template>
  <div class="w-full overflow-auto p-6 bg-gray-100">
    <table class="min-w-full bg-white">
      <thead class="bg-gray-800 text-white">
        <tr>
          <th class="text-left py-3 px-4 uppercase font-semibold text-sm">
            Code
          </th>
          <th class="text-left py-3 px-4 uppercase font-semibold text-sm">
            User's name
          </th>
          <th class="text-left py-3 px-4 uppercase font-semibold text-sm">
            Activated
          </th>
        </tr>
      </thead>

      <tbody>
        <tr v-for="card in cards" :key="card">
          <td class="text-left py-3 px-4">{{ card.code }}</td>
          <td class="text-left py-3 px-4">{{ card.user.name }}</td>
          <td class="text-left py-3 px-4 hover:text-blue-500">
            {{ card.activated }}
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
import { mapState, mapActions } from "vuex";

export default {
  name: "Card",

  computed: {
    ...mapState("cards", ["cards"]),
  },

  methods: {
    ...mapActions("cards", ["getCardsList"]),
  },

  created() {
    this.getCardsList();
  },
};
</script>
