import Logout from '@/views/Logout.vue';
import Home from '@/views/Home.vue';
import Auth from '@/views/Auth.vue';
import Error from '@/views/Error.vue';

export const ROUTE_NAME = {
  HOME: 'home',
  LOGOUT: 'logout',
  AUTH: 'auth',
  ERROR: 'error',
};

export const ROUTES = [
  {
    path: '/',
    name: ROUTE_NAME.HOME,
    component: Home,
  },
  {
    path: '/logout',
    name: ROUTE_NAME.LOGOUT,
    component: Logout,
    meta: {
      skipAuth: true,
    },
  },
  {
    path: '/callback',
    name: ROUTE_NAME.AUTH,
    component: Auth,
    meta: {
      skipAuth: true,
    },
  },
  {
    path: '/error',
    name: ROUTE_NAME.ERROR,
    component: Error,
    props: true,
  },
];
