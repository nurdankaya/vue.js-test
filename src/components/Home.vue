<template id="main">
  <div id="center">
    <button v-on:click="clicked = true" type="button" class="btn btn-light">
      Post Ekle
    </button>
    <div v-if="clicked">
      <form action="">
        <div class="form-group">
          <label>Post Konusu:</label>
          <input v-model="post.subject" type="text" class="form-control" />
        </div>
        <div class="form-group">
          <textarea
            v-model="post.maintext"
            class="form-control"
            placeholder="Bir şeyler yaz..."
            rows="3"
          ></textarea>
        </div>
      </form>
      <p>Konu: {{ post.subject }}</p>
      <p>Post: {{ post.maintext }}</p>
    </div>
  </div>
</template>

<script>
export default {
  props: ["ActiveA"],
  data() {
    return {
      clicked: false,
      post: {
        subject: "",
        maintext: "",
      },
    };
  },
  methods: {},
};
</script>

<style lang="css" scoped>
#center {
  margin: 25%;
}
</style>
