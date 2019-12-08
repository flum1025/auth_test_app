import createPersistedState from 'vuex-persistedstate';
import { MUTATION_TYPES } from '@/store/modules/login';
import { LOCAL_STORAGE_KEY } from '@/config/setting';

const targetMutationTypes = [`login/${MUTATION_TYPES.SET_LOGIN}`];

export default createPersistedState({
  key: LOCAL_STORAGE_KEY.LOGIN_STORE_KEY,
  paths: ['login.idToken', 'login.name', 'login.tokenExpiry'],
  filter({ type }) {
    return targetMutationTypes.includes(type);
  },
  storage: localStorage,
});
