<template>
  <div class="w-full bg-white overflow-auto p-6 bg-gray-100">
    <table class="min-w-full bg-white">
      <thead class="bg-gray-800 text-white">
        <tr>
          <th class="text-left py-3 px-4 uppercase font-semibold text-sm">Name</th>
          <th class="text-left py-3 px-4 uppercase font-semibold text-sm">Address</th>
          <th class="text-left py-3 px-4 uppercase font-semibold text-sm">Logo</th>
          <th class="text-left py-3 px-4 uppercase font-semibold text-sm">Website</th>
        </tr>
      </thead>

      <tbody>
        <tr v-for="company in companies" :key="company.id">
          <td class="text-center py-3 px-4"><img width="20" height="20" :src="company.logo" alt="avatar"></td>
          <td class="text-left py-3 px-4">{{ company.name }}</td>
          <td class="text-left py-3 px-4">{{ company.address }}</td>
          <td class="text-left py-3 px-4 hover:text-blue-500">{{ company.website }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
import { mapState, mapActions } from 'vuex'

export default {
  name: 'Company',

  computed: {
    ...mapState('companies', ['companies']),
  },

  methods: {
    ...mapActions('companies', ['getListCompanies']),
  },

  created() {
    if (!(this.companies.length > 1)) this.getListCompanies()
  }
}
</script>
