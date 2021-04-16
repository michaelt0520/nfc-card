<template>
  <div class="w-1/3 p-6">
    <div class="bg-gradient-to-b from-green-200 to-green-100 border-b-4 border-green-600 rounded-lg shadow-xl p-5">
      <p>Code: {{ card.code }}</p>
      <p>Activate: {{ card.activated }}</p>
      <p>User's name: {{ card.user.name }}</p>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Card',

  props: {
    card: {
      type: Object,
      required: true,
    }
  }
}
</script>
