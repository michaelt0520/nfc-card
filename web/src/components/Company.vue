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
        <tr>
          <td class="text-left py-3 px-4">1Pac</td>
          <td class="text-left py-3 px-4">86-88 Le Van Duyet, Binh Thanh</td>
          <td class="text-center py-3 px-4"><img width="20" height="20" src="../assets/logo.png" alt="avatar"></td>
          <td class="text-left py-3 px-4 hover:text-blue-500">1Pac.vn</td>
        </tr>
        <tr>
          <td class="text-left py-3 px-4">Tokoinnn</td>
          <td class="text-left py-3 px-4">86-88 Dien Bien Phu, Binh Thanh</td>
          <td class="text-center py-3 px-4"><img width="20" height="20" src="../assets/logo.png" alt="avatar"></td>
          <td class="text-left py-3 px-4 hover:text-blue-500">Tokoin.com.vn.indone</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  name: 'Company'
}
</script>

<style>

</style>
