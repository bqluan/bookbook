import Vue from 'vue'
import Router from 'vue-router'

// lazy-loading components
const Book = resolve => require(['@/components/Book'], resolve)

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/book/:id',
      name: 'Book',
      component: Book
    }
  ]
})
