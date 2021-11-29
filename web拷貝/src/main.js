import Vue from 'vue'
import App from './App.vue'
import router from './router'

//---- ElementUI
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css'
Vue.use(ElementUI);

//---- Bootstrap
import Bootstrap from 'bootstrap'
import BootstrapVue from 'bootstrap-vue'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
Vue.use(BootstrapVue)

//---- CSS
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import '../src/less/style_admin.less'

Vue.config.productionTip = false

import Http from './util/http'

Http.install = function (Vue) {
  Vue.prototype.$http = Http;
};
Vue.use(Http);

import Public from './util/public'

Public.install = function (Vue) {
  Vue.prototype.$public = Public;
};
Vue.use(Public);

new Vue({
  Bootstrap,
  router,
  render: h => h(App),
}).$mount('#app')
