<template>
  <div class="about">
    <h1>This is an about page</h1>
    <el-form ref="form" :model="form" label-width="120px">
      <el-form-item label="username">
        <el-input v-model="form.username"></el-input>
      </el-form-item>
      <el-form-item label="Password">
        <el-input
          type="password"
          v-model="form.password"
          autocomplete="off"
        ></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">Create</el-button>
        <el-button>Cancel</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { mapActions } from "vuex";

export default {
  name: "About",

  data() {
    return {
      form: {
        username: "",
        password: "",
      },
    };
  },

  methods: {
    ...mapActions("users", ["login", "signup"]),

    onSubmit() {
      const data = {
        username: this.form.username,
        password: this.form.password,
      };
      this.login(data);
    },
  },
};
</script>
