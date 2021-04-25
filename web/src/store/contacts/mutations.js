import {
  FETCH_LIST_CONTACTS
} from '../mutation-types'

const mutations = {
  [FETCH_LIST_CONTACTS](state, payload) {
    state.contacts = payload
  },
}

export default mutations
