import Vue from 'vue';
import Vuex from 'vuex';
import login from './modules/login';
import persistLoginStorePlugin from './plugins/persistLoginStorePlugin';

Vue.use(Vuex);

export default new Vuex.Store({
  strict: process.env.NODE_ENV !== 'production',
  modules: {
    login,
  },
  plugins: [persistLoginStorePlugin],
});
