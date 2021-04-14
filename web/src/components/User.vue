<template>
  <div class="w-full bg-white overflow-auto p-6 bg-gray-100">
    <table class="min-w-full bg-white">
      <thead class="bg-gray-800 text-white">
        <tr>
          <th class="text-left py-3 px-4 uppercase font-semibold text-sm">Avatar</th>
          <th class="text-left py-3 px-4 uppercase font-semibold text-sm">Name</th>
          <th class="text-left py-3 px-4 uppercase font-semibold text-sm">Username</th>
          <th class="text-left py-3 px-4 uppercase font-semibold text-sm">Email</th>
          <th class="text-left py-3 px-4 uppercase font-semibold text-sm">Type</th>
          <th class="text-left py-3 px-4 uppercase font-semibold text-sm">Company</th>
        </tr>
      </thead>

      <tbody>
        <tr v-for="user in users" :key="user.id">
          <td class="text-center py-3 px-4"><img width="20" height="20" :src="user.avatar" alt="avatar"></td>
          <td class="text-left py-3 px-4">{{ user.name }}</td>
          <td class="text-left py-3 px-4">{{ user.username }}</td>
          <td class="text-left py-3 px-4 hover:text-blue-500">{{ user.email }}</td>
          <td class="text-left py-3 px-4">{{ user.type }}</td>
          <td class="text-left py-3 px-4">{{ user.Company.name }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
import { mapState, mapActions } from 'vuex'

export default {
  name: 'User',

  computed: {
    ...mapState('users', ['users']),
  },

  methods: {
    ...mapActions('users', ['getUsersList']),
  },

  created() {
    if (!(this.users.length > 1)) this.getUsersList()
  }
}
</script>
