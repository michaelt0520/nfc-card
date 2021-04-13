<template>
  <div class="w-1/3 p-6">
    <div class="bg-gradient-to-b from-green-200 to-green-100 border-b-4 border-green-600 rounded-lg shadow-xl p-5">
      <p>ID: {{ card.id }}</p>
      <p>Code: {{ card.code }}</p>
      <p>Activate: {{ card.activated }}</p>
      <p>User's ID: {{ card.user_id }}</p>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Card',

  props: {
    card: {
      id: Number,
      code: String,
      user_id: Number,
      activated: Boolean
    }
  }
}
</script>
