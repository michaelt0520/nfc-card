import {
  INITIAL_STATES_SET_USER,
  FETCH_LIST_USERS,
  SIGN_OUT
} from '../mutation-types'

const mutations = {
  [INITIAL_STATES_SET_USER](state, payload) {
    state.currentUser = payload.user
    state.isAuthenticated = true
    localStorage.setItem('user', JSON.stringify(payload.user))
    localStorage.setItem('is_authenticated', true)
    localStorage.setItem('token', payload.user.jwt)
  },

  [FETCH_LIST_USERS](state, payload) {
    state.users = payload
  },

  [SIGN_OUT](state) {
    state.currentUser = {}
    state.isAuthenticated = false
    localStorage.removeItem('user')
    localStorage.removeItem('token')
    localStorage.removeItem('is_authenticated')
  }
}

export default mutations
