<template>
  <div class="w-full bg-white overflow-auto p-6 bg-gray-100">
    <table class="min-w-full bg-white">
      <thead class="bg-gray-800 text-white">
        <tr>
          <th class="text-left py-3 px-4 uppercase font-semibold text-sm">Name</th>
          <th class="text-left py-3 px-4 uppercase font-semibold text-sm">Avatar</th>
          <th class="text-left py-3 px-4 uppercase font-semibold text-sm">Email</th>
          <th class="text-left py-3 px-4 uppercase font-semibold text-sm">Type</th>
          <th class="text-left py-3 px-4 uppercase font-semibold text-sm">Company</th>
        </tr>
      </thead>

      <tbody>
        <tr>
          <td class="text-left py-3 px-4">Nghia</td>
          <td class="text-center py-3 px-4"><img width="20" height="20" src="../assets/logo.png" alt="avatar"></td>
          <td class="text-left py-3 px-4 hover:text-blue-500">tran.nghia@1pac.vn</td>
          <td class="text-left py-3 px-4 hover:text-blue-500">Something</td>
          <td class="text-left py-3 px-4 hover:text-blue-500">1Pac</td>
        </tr>
        <tr>
          <td class="text-left py-3 px-4">Michael</td>
          <td class="text-center py-3 px-4"><img width="20" height="20" src="../assets/logo.png" alt="avatar"></td>
          <td class="text-left py-3 px-4 hover:text-blue-500">michakeak.nghia@1pac.vn</td>
          <td class="text-left py-3 px-4 hover:text-blue-500">Somethingsas</td>
          <td class="text-left py-3 px-4 hover:text-blue-500">Tokoinnnn</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  name: 'User'
}
</script>

<style>

</style>
