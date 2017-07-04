<template>
  <div class="book">
    <div class="cover">
      <img></img>
    </div>
    <h1>{{ book.title }}</h1>
    <h2>{{ book.desc }}</h2>
    <template v-if="isBorrowed">
      <div class="info">您在今天借阅了这本书</div>
    </template>
    <template v-else-if="needConfirmation">
      <button class="ok" @click="ok">确定</button>
      <button class="cancel" @click="cancel">取消</button>
    </template>
    <template v-else>
      <button class="borrow" @click="borrow">我要借书</button>
    </template>
  </div>
</template>

<script>
export default {
  name: 'book',
  data () {
    return {
      needConfirmation: false,
      isBorrowed: false,
      book: null
    }
  },
  created () {
    this.book = {
      title: '芝麻街小小孩的秘密生活',
      desc: '简介：平装，湖北美术出版社'
    }
  },
  methods: {
    borrow () {
      this.needConfirmation = true
    },
    ok () {
      this.needConfirmation = false
      this.isBorrowed = true
    },
    cancel () {
      this.needConfirmation = false
    }
  }
}
</script>

<style scoped>
.book {
  text-align: center;
}

.cover {
  height: 190px;
  background-color: #63a9e4;
  margin-bottom: 50px;
}

.cover>img {
  margin-top: 120px;
  width: 70px;
  height: 100px;
  border: 1px solid black;
  background-color: white;
}

h1, h2 {
  font-weight: normal;
}

h1 {
  font-size: 1.5em;
}

h2 {
  color: #777;
  font-size: 1em;
}

button.borrow {
  color: white;
  font-size: 1.2em;
  outline: none;
  position: fixed;
  bottom: 0;
  left: 0;
  width: 100%;
  height: 50px;
  line-height: 50px;
  border: 0;
  padding: 0;
  background-color: #63a9e4;
  border-top: 1px solid #4389c4;
}

button.ok {
  color: white;
  font-size: 1.2em;
  outline: none;
  position: fixed;
  bottom: 0;
  left: 0;
  width: 50%;
  height: 50px;
  line-height: 50px;
  border: 0;
  padding: 0;
  background-color: #63a9e4;
  border-top: 1px solid #4389c4;
}

button.cancel {
  color: white;
  font-size: 1.2em;
  outline: none;
  position: fixed;
  bottom: 0;
  right: 0;
  width: 50%;
  height: 50px;
  line-height: 50px;
  border: 0;
  padding: 0;
  background-color: #d9534f;
  border-top: 1px solid #d43f3a;
}

.info {
  background-color: white;
  color: #777;
  font-size: 1em;
  position: fixed;
  bottom: 0;
  left: 0;
  width: 100%;
  height: 50px;
  line-height: 50px;
  border: 0;
  padding: 0;
}
</style>
