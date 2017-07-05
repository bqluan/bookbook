<template>
  <div class="book">
    <md-toolbar>
      <img src="../assets/book-cover.jpg">
    </md-toolbar>
    <md-card>
      <md-card-header>
        <div class="md-title">{{ book.title }}</div>
        <div class="md-subhead">{{ book.author }} {{ book.publisher }}</div>
      </md-card-header>
      <md-card-content>
        {{ book.desc }}
      </md-card-content>
    </md-card>
    <md-bottom-bar>
      <template v-if="borrowedAt">
        <md-button disabled>您在今天借阅了这本书</md-button>
      </template>
      <template v-else-if="book.qty <= 0">
        <md-button disabled>这本书已经被借走了</md-button>
      </template>
      <template v-else-if="needConfirmation">
        <md-button class="md-raised md-primary" @click="ok">确定</md-button>
        <md-button class="md-raised md-accent" @click="cancel">取消</md-button>
      </template>
      <template v-else>
        <md-button class="md-raised md-primary" @click="borrow">我要借书</md-button>
      </template>
    </md-bottom-bar>
  </div>
</template>

<script>
export default {
  name: 'book',
  data () {
    return {
      book: null,
      borrowedAt: null,
      needConfirmation: false
    }
  },
  created () {
    this.book = {
      title: '为何是他：怀疑主义时代的信仰',
      author: '[美]提摩太·凯勒',
      publisher: '上海三联书店',
      desc: '在一个怀疑的时代，我们何以知道基督教是惟一可信的宗教？上帝为何允许苦难发生？一位自称是爱的上帝为何会把人送到地狱去？教会是否要为诸多的不公义负责？科学是否已经否定了基督教？提摩太·凯勒借着文学、哲学、日常生活谈话以及严谨的逻辑思考，回答了怀疑基督教信仰的人甚或热心信徒常有的一些疑问，解释了为什么相信基督教的上帝是十分明智与合理的选择。',
      qty: 1
    }
  },
  methods: {
    borrow () {
      this.needConfirmation = true
    },
    ok () {
      this.needConfirmation = false
      this.borrowedAt = new Date()
    },
    cancel () {
      this.needConfirmation = false
    }
  }
}
</script>

<style scoped>
.book {
  padding-bottom: 56px;
}
.md-toolbar {
  height: 180px;
  margin-bottom: 25px;
}
.md-toolbar img {
  margin: 130px auto 0 auto;
  max-width: 70px;
  max-height: 100px;
  width: auto;
  height: auto;
  box-shadow: 0 1px 5px rgba(0, 0, 0, .2), 0 2px 2px rgba(0, 0, 0, .14), 0 3px 1px -2px rgba(0, 0, 0, .12);
  background-color: white;
}
.md-card {
  box-shadow: none;
}
.md-title,
.md-subhead {
  text-align: center;
}
.md-bottom-bar {
  position: fixed;
  bottom: 0;
  z-index: 999;
}
.md-bottom-bar button {
  width: 100%;
  margin: 0;
  border-radius: 0;
  font-size: 18px;
}
</style>
