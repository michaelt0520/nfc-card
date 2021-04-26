import {
  INITIAL_STATES_SET_USER,
  SIGN_OUT
} from '../mutation-types'

const mutations = {
  [INITIAL_STATES_SET_USER](state, payload) {
    state.currentUser = payload.user
    state.isAuthenticated = true
    localStorage.setItem('user', JSON.stringify(payload.user))
    localStorage.setItem('is_authenticated', true)
    localStorage.setItem('token', `Bearer ${payload.user.jwt}`)
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
