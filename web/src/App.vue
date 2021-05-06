<template>
  <div>
    <router-view />
    <Alert />
  </div>
</template>

<script>
import Alert from "@/components/Alert";

export default {
  name: "App",

  components: {
    Alert,
  },
};
</script>
