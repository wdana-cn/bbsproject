<template>
  <div id="app">
    <home-header />
    <div class="bg"></div>
    <div class="img">
      <img :src="url" alt="" srcset="" />
    </div>
    <div class="content">
      <keep-alive>
        <router-view />
      </keep-alive>
    </div>
  </div>
</template>

<script>
import HomeHeader from 'components/HomeHeader'
export default {
  name: 'App',
  data() {
    return {
      url: require('./assets/bgc.png')
    }
  },
  components: {
    HomeHeader
  }
}
</script>

<style lang="scss">
.content {
  width: calc(100vw - 17px);
  xbackground-color: rgb(223, 223, 223);
}
.bg {
  position: fixed;
  height: 100vh;
  width: 100vw;
  top: 0;
  left: 0;
  background-image: repeating-linear-gradient(to bottom,#FFFFFF,#EE82EE);
  z-index: -1;
}
.img {
  height: 784px;
  width: 354px;
  position: fixed;
  top: 120px;
  left: 70%;
  img {
    height: 580px;
    width: 400px;
  }
}
.content {
  width: 1154px;
  margin: auto;
}
@media (max-width: 1126) {
  .img {
    display: none;
  }
}
</style>
