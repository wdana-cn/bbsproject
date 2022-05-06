<template>
  <div class="classify-bar" ref="classify">
    <div
      class="item"
      v-for="(item, index) in classifyData"
      :key="item.name"
      :class="index == currentIndex ? 'active' : ''"
      @click="itemClick(index, item.path)"
    >
      {{ item.name }}
    </div>
  </div>
</template>

<script>
export default {
  name: 'ClassifyBar',
  data() {
    return {
      classifyData: [
        {name: '全部', path: '/total'},
        {name: 'IT讨论区', path: '/itchat'},
        {name: '近期热点', path: '/hit'},
        {name: '闲聊分享', path: '/justchat'},
        {name: '好物安利', path: '/share'}
      ],
      currentIndex: 0
    }
  },
  methods: {
    itemClick(index, path) {
      this.currentIndex = index
      this.$router.push(`/home${path}`)
    }
  },
  mounted() {
    let map = {Total: 0, ITCHAT: 1, HIT: 2, JustChat: 3, Share: 4}
    let path = this.$route.params.path
    this.currentIndex = map[path]
  }
}
</script>

<style lang="scss" scoped>
.classify-bar {
  display: flex;
  justify-content: space-around;
  align-items: center;
  flex-direction: row;
  width: 500px;
  height: 45px;
  .item {
    height: 39px;
    width: 80px;
    line-height: 50px;
    text-align: center;
    font-size: 20px;
    color: black;
    cursor: pointer;
    &:hover {
      background-color: rgb(245, 245, 245);
    }
  }
  .active {
    background-color: #cacaca;
    color: rgb(2, 2, 2);
    &:hover {
      background-color: #393949;
      color: rgb(0, 0, 0);
    }
  }
}
</style>
