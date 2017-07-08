<template>
  <div class="book">
    <template v-if="book">
      <md-toolbar>
        <img src="../assets/book-cover.jpg">
      </md-toolbar>
      <md-card>
        <md-card-header>
          <div class="md-title">{{ book.Title }}</div>
          <div class="md-subhead">{{ book.Author }} {{ book.Pub }}</div>
        </md-card-header>
        <md-card-content>
          {{ book.Desc }}
        </md-card-content>
      </md-card>
      <md-bottom-bar>
        <template v-if="borrowedAt">
          <md-button disabled>您在{{ diffInDays }}借阅了这本书</md-button>
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
    </template>
  </div>
</template>

<script>
import moment from 'moment'

moment.locale('zh-cn')
moment.updateLocale('zh-cn', {
  calendar: {
    sameDay: '[今天]LT',
    nextDay: '[明天]LT',
    nextWeek: '[下]dddLT',
    lastDay: '[昨天]LT',
    lastWeek: '[上]dddLT',
    thisWeek: '[这]dddLT',
    sameElse: 'L'
  }
})
moment.calendarFormat = function (myMoment, now) {
  const diff = myMoment.diff(now, 'days', true)
  if (diff >= 0 && diff < 1) {
    return 'sameDay'
  }
  if (diff >= -1 && diff < 0) {
    return 'lastDay'
  }
  if (diff >= 1 && diff < 2) {
    return 'nextDay'
  }
  if (myMoment.week() === now.week() && myMoment.year() === now.year()) {
    return 'thisWeek'
  }
  if (myMoment.week() === (now.week() - 1) && myMoment.year() === now.year()) {
    return 'lastWeek'
  }
  if (myMoment.week() === (now.week() + 1) && myMoment.year() === now.year()) {
    return 'nextWeek'
  }
  return 'sameElse'
}

export default {
  name: 'book',
  data () {
    return {
      book: null,
      borrowedAt: null,
      needConfirmation: false
    }
  },
  computed: {
    diffInDays () {
      if (!this.borrowedAt) {
        return ''
      }
      return moment(this.borrowedAt).calendar()
    }
  },
  created () {
    // fetch the data when the view is created and the data is
    // already being observed
    this.fetchData()
  },
  watch: {
    // call the method again if the route changes
    '$route': 'fetchData'
  },
  methods: {
    fetchData () {
      this.book = null
      this.borrowedAt = null
      this.needConfirmation = false
      this.$http.get(`/api/book/${this.$route.params.id}`).then(resp => {
        this.book = resp.body
      })
      this.$http.get(`/api/borrow?book_id=${this.$route.params.id}&wechat_id=abcdefg`).then(resp => {
        if (resp.body && resp.body.length > 0) {
          this.borrowedAt = resp.body[0].CreatedAt
        }
      })
    },
    borrow () {
      this.needConfirmation = true
    },
    ok () {
      const borrow = {
        BookID: Number(this.$route.params.id),
        WechatID: 'abcdefg'
      }
      this.$http.post('/api/borrow', borrow).then(resp => {
        this.needConfirmation = false
        this.borrowedAt = resp.body.CreatedAt
      })
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
