<template>
  <div>
    <Header />

    <div v-if="loaded">
      <setting-profile :user="user" />
      <setting-password />
      <setting-social :socials-list="user.informations" />
    </div>
    <div v-else>
      <loading />
    </div>
  </div>
</template>

<script>
import { mapState, mapActions } from "vuex";
import Header from "@/components/Header";
import SettingSocial from "@/components/SettingSocial";
import SettingProfile from "@/components/SettingProfile";
import SettingPassword from "@/components/SettingPassword";
import Loading from "@/components/Loading";

export default {
  name: "Setting",

  components: {
    Header,
    SettingProfile,
    SettingSocial,
    SettingPassword,
    Loading,
  },

  data() {
    return {
      loaded: false,
    };
  },

  computed: {
    ...mapState("users", ["user"]),
  },

  methods: {
    ...mapActions("users", ["getCurrentUser"]),
  },

  created() {
    this.getCurrentUser().then(() => {
      setTimeout(() => {
        this.loaded = true;
      }, 500);
    });
  },
};
</script>
