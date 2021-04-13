import {
  FETCH_LIST_COMPANIES,
} from '../mutation-types'

const mutations = {
  [FETCH_LIST_COMPANIES] (state, payload) {
    state.companies = payload
  }
}

export default mutations
