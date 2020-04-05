<template>
  <div v-if="ActiveG" id="ground">
    <div class="container">
      <div class="row">
        <div class="col-sm"></div>
        <div class="col-sm">
          <form>
            <div class="form-group">
              <label>Email adresi</label>
              <input type="email" class="form-control" />
            </div>
            <div class="form-group">
              <label>Parola</label>
              <input type="password" class="form-control" />
            </div>

            <button type="submit" class="btn btn-danger">
              Giriş
            </button>
          </form>
        </div>
        <div class="col-sm"></div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: ["ActiveG"],
  data() {
    return {
      imgPath:
        "https://images.pexels.com/photos/842339/pexels-photo-842339.jpeg?auto=compress&cs=tinysrgb&dpr=3&h=750&w=1260",
    };
  },
};
</script>

<style lang="css" scoped>
#ground {
  background-color: #a9a9a9;
  margin-top: 200px;
  margin-bottom: 200px;
  padding-top: 50px;
  padding-right: 50px;
  padding-bottom: 50px;
  padding-left: 50px;
}
</style>
