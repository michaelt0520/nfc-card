import {
  FETCH_COMPANY,
  FETCH_COMPANY_USERS,
  FETCH_COMPANY_CARDS,
  FETCH_LIST_COMPANIES,
  UPDATE_COMPANY
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
  }
}

export default mutations
