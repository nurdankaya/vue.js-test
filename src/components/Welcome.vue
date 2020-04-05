<template>
  <div v-if="ActiveA">
    <img :src="imgPath" class="img-fluid img-responsive" />
    <div class="carousel-caption">
      <h1 class="display-1">Hoşgeldin!</h1>
    </div>
  </div>
</template>

<script>
export default {
  props: ["ActiveA"],
  data() {
    return {
      imgPath:
        "https://images.pexels.com/photos/842339/pexels-photo-842339.jpeg?auto=compress&cs=tinysrgb&dpr=3&h=750&w=1260",
    };
  },
};
</script>

<style></style>
