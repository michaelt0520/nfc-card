<template>
  <div>
    <Header />

    <div>
      <setting-profile :user="user" />
      <setting-social :socials-list="user.informations" />
    </div>
  </div>
</template>

<script>
import { mapState, mapActions } from "vuex";
import Header from "@/components/Header";
import SettingSocial from "../components/SettingSocial";
import SettingProfile from "../components/SettingProfile";

export default {
  name: "Setting",

  components: {
    Header,
    SettingProfile,
    SettingSocial,
  },

  computed: {
    ...mapState("users", ["user"]),
  },

  methods: {
    ...mapActions("users", ["getCurrentUser"]),
  },

  created() {
    this.getCurrentUser();
  },
};
</script>
