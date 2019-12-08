import Vue from 'vue';
import Router from 'vue-router';
import { ROUTES } from './config';
import navigationGuard from './auth';

Vue.use(Router);

const router = new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: ROUTES,
});

navigationGuard(router);

export default router;
