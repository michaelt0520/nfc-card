import {
  FETCH_LIST_CARDS,
  FETCH_CARD,
  CREATE_CARD
} from '../mutation-types'

const mutations = {
  [FETCH_LIST_CARDS](state, payload) {
    state.cards = payload
  },

  [FETCH_CARD](state, payload) {
    state.card = payload
  },

  [CREATE_CARD](state, payload) {
    state.cards.push(payload)
  }
}

export default mutations
