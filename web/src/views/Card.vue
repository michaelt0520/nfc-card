<template></template>

<script>
import { mapState, mapActions } from "vuex";

export default {
  data() {
    return {
      card: {},
    };
  },

  computed: {
    ...mapState("cards", ["user"]),
  },

  methods: {
    ...mapActions("cards", ["fetch"]),
  },

  created() {
    this.fetch(this.$route.params.code);
  },
};
</script>
