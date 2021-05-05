import {
  FETCH_COMPANY,
  FETCH_COMPANY_USERS,
  FETCH_COMPANY_CARDS,
  FETCH_LIST_COMPANIES,
  UPDATE_COMPANY,
  CREATE_COMPANY,
  UPDATE_CARD,
  UPDATE_USER
} from '../mutation-types'

const mutations = {
  [FETCH_COMPANY](state, payload) {
    state.company = payload
  },

  [FETCH_COMPANY_USERS](state, payload) {
    state.users = payload
  },

  [FETCH_COMPANY_CARDS](state, payload) {
    state.cards = payload
  },

  [FETCH_LIST_COMPANIES](state, payload) {
    state.companies = payload
  },

  [UPDATE_COMPANY](state, payload) {
    state.company = payload
  },

  [CREATE_COMPANY](state, payload) {
    state.companies.push(payload)
  },

  [UPDATE_CARD](state, payload) {
    const index = state.cards.findIndex(card => card.id === payload.id)
    if (index < 0) return

    state.cards[index].activated = payload.activated
  },

  [UPDATE_USER](state, payload) {
    const index = state.users.findIndex(user => user.username === payload)
    if (index < 0) return

    state.users.slice(index, 1)
  }
}

export default mutations
