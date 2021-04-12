<template>
  <div class="home">
    <img src="../assets/logo.png" />
    <h2>{{ isSignedUp }}</h2>
    <div>
      <el-form ref="form" :model="form" label-width="120px">
        <el-form-item label="Email">
          <el-input v-model="form.email"></el-input>
        </el-form-item>
        <el-form-item label="Password">
          <el-input
            type="password"
            v-model="form.password"
            autocomplete="off"
          ></el-input>
        </el-form-item>
        <el-form-item label="Username">
          <el-input v-model="form.username"></el-input>
        </el-form-item>
        <el-form-item label="Name">
          <el-input v-model="form.name"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="onSubmit">Create</el-button>
          <el-button>Cancel</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
import { mapState, mapActions } from "vuex";

export default {
  name: "Home",

  data() {
    return {
      form: {
        email: "",
        password: "",
        username: "",
        name: "",
      },
    };
  },

  computed: {
    ...mapState("users", ["user"]),

    isSignedUp() {
      return this.user ? this.user.email : "-1";
    },
  },

  methods: {
    ...mapActions("users", ["login", "signup"]),

    onSubmit() {
      const data = {
        email: this.form.email,
        password: this.form.password,
        username: this.form.username,
        name: this.form.name,
      };
      this.signup(data);
      console.log('1')
    },
  }
};
</script>
