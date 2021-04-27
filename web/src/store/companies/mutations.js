import {
  FETCH_COMPANY,
  FETCH_LIST_COMPANIES,
} from '../mutation-types'

const mutations = {
  [FETCH_COMPANY](state, payload) {
    state.company = payload
  },

  [FETCH_LIST_COMPANIES](state, payload) {
    state.companies = payload
  }
}

export default mutations
