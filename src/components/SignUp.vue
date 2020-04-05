<template id="main">
  <div v-if="ActiveK">
    <div>
      <img
        id="imgSignUp"
        src="https://images.pexels.com/photos/1089438/pexels-photo-1089438.jpeg?auto=compress&cs=tinysrgb&dpr=2&h=650&w=940"
        class="img-fluid img-responsive"
      />
      <div class="carousel-caption">
        <h1 class="display-1">Hesabını oluştur.</h1>
      </div>
      <div class="container">
        <div class="row">
          <div class="col-sm"></div>
          <div class="col-sm">
            <form id="form">
              <div class="form-group">
                <label>Kullanıcı Adı</label>
                <input v-model="userName" type="text" class="form-control" />
              </div>
              <div class="form-group">
                <label>Email adresi</label>
                <input v-model="userMail" type="email" class="form-control" />
                <small id="emailHelp" class="form-text text-muted"
                  >We'll never share your email with anyone else.</small
                >
              </div>
              <div class="form-group">
                <label>Parola</label>
                <input
                  v-model="userPassword"
                  type="password"
                  class="form-control"
                />
              </div>

              <button
                v-on:click.prevent="post"
                type="submit"
                class="btn btn-danger"
              >
                Kayıt Ol
              </button>
            </form>
          </div>
          <div class="col-sm"></div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: ["ActiveK"],
  data() {
    return {
      userName: "",
      userMail: "",
      userPassword: "",
      imgPath:
        "https://images.pexels.com/photos/842339/pexels-photo-842339.jpeg?auto=compress&cs=tinysrgb&dpr=3&h=750&w=1260",
    };
  },
  methods: {
    post: function () {
      this.$http.post("database/test.db", {
        username: this.userName,
        email: this.userMail,
        password: this.userPassword,
      });
    },
  },
};
</script>

<style lang="css" scoped>
#form {
  margin-top: 200px;
  margin-bottom: 200px;
}
#main {
  background-color: #a9a9a9;
}
#imgSignUp {
  width: 100%;
}
</style>
