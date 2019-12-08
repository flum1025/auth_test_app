import store from '@/store';
import { authorize } from '@/services/auth';
import { GETTER_TYPES } from '@/store/modules/login';

export default (router) => {
  router.beforeEach(({ meta: { skipAuth }, name: routeName }, from, next) => {
    if (skipAuth) {
      next();
      return;
    }

    const isLoggedIn = store.getters[`login/${GETTER_TYPES.isLoggedIn}`];

    if (!isLoggedIn) {
      authorize();
      return;
    }

    next();
  });
};
