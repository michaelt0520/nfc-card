import {
  FETCH_LIST_USERS,
  FETCH_USER,
  UPDATE_USER
} from '../mutation-types'

const mutations = {
  [FETCH_LIST_USERS](state, payload) {
    state.users = payload
  },

  [FETCH_USER](state, payload) {
    state.user = payload
  },

  [UPDATE_USER](state, payload) {
    state.user = payload
  }
}

export default mutations
