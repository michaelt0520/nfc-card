import {
  INITIAL_STATES_SET_USER,
  CREATE_USER
} from '../mutation-types'

const mutations = {
  [INITIAL_STATES_SET_USER](state, payload) {
    state.currentUser = payload.User
    localStorage.setItem('user', payload.User)
    localStorage.setItem('token', payload.Token)
  },

  [CREATE_USER](state, payload) {
    state.currentUser = payload
  }
}

export default mutations
