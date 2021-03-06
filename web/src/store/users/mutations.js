import {
  INITIAL_STATES_SET_USER,
  SIGN_OUT,
  FETCH_LIST_USERS,
  FETCH_USER,
  CREATE_USER,
  UPDATE_CURRENT_USER,
  UPDATE_USER,
  SEARCH_USERS,
  CREATE_INFORMATION,
  UPDATE_INFORMATION,
  DELETE_INFORMATION,
  ADD_CARD,
  UPDATE_CARD,
  REMOVE_CARD
} from '../mutation-types'

const mutations = {
  [INITIAL_STATES_SET_USER](state, payload) {
    state.currentUser = payload.user
    state.isAuthenticated = true
    localStorage.setItem('user', JSON.stringify(payload.user))
    localStorage.setItem('token', `Bearer ${payload.token}`)
    localStorage.setItem('is_authenticated', true)
  },

  [SIGN_OUT](state) {
    state.currentUser = {}
    state.isAuthenticated = false
    localStorage.removeItem('user')
    localStorage.removeItem('token')
    localStorage.removeItem('is_authenticated')
  },

  [FETCH_LIST_USERS](state, payload) {
    state.users = payload
  },

  [FETCH_USER](state, payload) {
    state.user = payload
  },

  [UPDATE_CURRENT_USER](state, payload) {
    state.user = payload
    state.currentUser = payload
    localStorage.setItem('user', JSON.stringify(payload))
  },

  [CREATE_INFORMATION](state, payload) {
    state.user.informations.push(payload)
  },

  [UPDATE_INFORMATION](state, payload) {
    const index = state.user.informations.findIndex(socialInfo => socialInfo.id === payload.id)
    if (index < 0) return

    state.user.informations[index].visibled = payload.visibled
  },

  [DELETE_INFORMATION](state, payload) {
    const index = state.user.informations.findIndex(socialInfo => socialInfo.id === payload)
    if (index < 0) return

    state.user.informations.splice(index, 1)
  },

  [SEARCH_USERS](state, payload) {
    state.searchUsers = payload
  },

  [CREATE_USER](state, payload) {
    state.users.push(payload)
  },

  [UPDATE_USER](state, payload) {
    const index = state.users.findIndex(user => user.id === payload.id)
    if (index < 0) return

    state.users[index] = payload
  },

  [ADD_CARD](state, payload) {
    state.user.cards.push(payload)
  },

  [UPDATE_CARD](state, payload) {
    const index = state.user.cards.findIndex(card => card.code === payload.code)
    if (index < 0) return

    state.user.cards[index] = payload
  },

  [REMOVE_CARD](state, payload) {
    const index = state.user.cards.findIndex(cards => cards.code === payload)
    if (index < 0) return

    state.user.cards.splice(index, 1)
  }
}

export default mutations
