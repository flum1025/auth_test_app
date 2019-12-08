const initialState = {
  idToken: null,
  name: null,
  tokenExpiry: null,
};

const state = {
  ...initialState,
};

export const MUTATION_TYPES = {
  SET_LOGIN: 'SET/LOGIN/LOGIN',
};

const mutations = {
  [MUTATION_TYPES.SET_LOGIN](localState, { idToken, name, tokenExpiry }) {
    localState.idToken = idToken;
    localState.name = name;
    localState.tokenExpiry = tokenExpiry;
  },
};

export const GETTER_TYPES = {
  isLoggedIn: 'isLoggedIn',
};

const getters = {
  [GETTER_TYPES.isLoggedIn](localState) {
    return localState.idToken && Date.now() < new Date(localState.tokenExpiry * 1000);
  },
};

export const ACTION_TYPES = {
  logIn: 'logIn',
  logOut: 'logOut',
};

const actions = {
  [ACTION_TYPES.logIn]({ commit }, { idToken, name, tokenExpiry }) {
    commit(MUTATION_TYPES.SET_LOGIN, {
      idToken,
      name,
      tokenExpiry,
    });
  },
  [ACTION_TYPES.logOut]({ commit }) {
    commit(MUTATION_TYPES.SET_LOGIN, { ...initialState });
  },
};

export default {
  namespaced: true,
  state,
  mutations,
  getters,
  actions,
};
