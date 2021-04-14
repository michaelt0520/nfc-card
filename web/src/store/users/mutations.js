import {
  INITIAL_STATES_SET_USER,
  FETCH_LIST_USERS
} from '../mutation-types'

const mutations = {
  [INITIAL_STATES_SET_USER](state, payload) {
    state.currentUser = payload.User
    state.isAuthenticated = true
    localStorage.setItem('user', JSON.stringify(payload.User))
    localStorage.setItem('is_authenticated', true)
    localStorage.setItem('token', payload.Token)
  },

  [FETCH_LIST_USERS](state, payload) {
    state.users = payload
  }
}

export default mutations
