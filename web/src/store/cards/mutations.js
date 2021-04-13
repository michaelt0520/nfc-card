import {
  FETCH_LIST_CARDS,
} from '../mutation-types'

const mutations = {
  [FETCH_LIST_CARDS] (state, payload) {
    state.cards = payload
  }
}

export default mutations
