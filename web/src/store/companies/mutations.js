import {
  FETCH_COMPANY,
  FETCH_COMPANY_USERS,
  FETCH_COMPANY_CARDS,
  FETCH_LIST_COMPANIES,
  UPDATE_CURRENT_COMPANY,
  CREATE_COMPANY,
  UPDATE_COMPANY,
  DELETE_COMPANY,
  UPDATE_CARD,
  CREATE_USER,
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

  [UPDATE_CURRENT_COMPANY](state, payload) {
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

  [CREATE_USER](state, payload) {
    state.users.push(payload)
  },

  [UPDATE_USER](state, payload) {
    const index = state.users.findIndex(user => user.username === payload)
    if (index < 0) return

    state.users.splice(index, 1)
  },

  [UPDATE_COMPANY](state, payload) {
    console.log("Before", payload)
    const index = state.companies.findIndex(company => company.id === payload.id)
    if (index < 0) return

    console.log("Before", state.companies)
    console.log("Before", state.companies[index])

    state.companies[index] = payload
  },

  [DELETE_COMPANY](state, payload) {
    const index = state.companies.findIndex(company => company.id === payload)
    if (index < 0) return

    state.companies.splice(index, 1)
  }
}

export default mutations
