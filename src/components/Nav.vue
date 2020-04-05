<template>
  <div>
    <nav class="nav nav-pills nav-fill">
      <a
        v-on:click="isActive(0)"
        class="nav-item nav-link"
        v-bind:class="isActiveA"
        >Anasayfa</a
      >
      <a
        v-on:click="isActive(1)"
        class="nav-item nav-link"
        v-bind:class="isActiveK"
        >Kaydol</a
      >
      <a
        v-on:click="isActive(2)"
        class="nav-item nav-link"
        v-bind:class="isActiveG"
        >Giriş Yap</a
      >
    </nav>
    <home-page v-bind:ActiveA="isActiveA"></home-page>
    <sign-up v-bind:ActiveK="isActiveK"></sign-up>
    <sign-in v-bind:ActiveG="isActiveG"></sign-in>
    <home-home v-bind:ActiveA="isActiveA"></home-home>
  </div>
</template>

<script>
import Welcome from "./Welcome.vue";
import SignUp from "./SignUp.vue";
import SignIn from "./SignIn.vue";
import Home from "./Home.vue";

export default {
  data() {
    return {
      isActiveA: "active",
      isActiveK: "",
      isActiveG: "",
    };
  },
  methods: {
    isActive: function (which) {
      if (which == 1) {
        (this.isActiveA = ""),
          (this.isActiveK = "active"),
          (this.isActiveG = "");
      } else if (which == 2) {
        (this.isActiveA = ""),
          (this.isActiveK = ""),
          (this.isActiveG = "active");
      } else {
        (this.isActiveA = "active"),
          (this.isActiveK = ""),
          (this.isActiveG = "");
      }
    },
  },

  components: {
    "home-page": Welcome,
    "sign-up": SignUp,
    "sign-in": SignIn,
    "home-home": Home,
  },
};
</script>

<style></style>
