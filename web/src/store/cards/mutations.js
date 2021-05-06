import {
  FETCH_LIST_CARDS,
  FETCH_CARD,
  CREATE_CARD,
  UPDATE_CARD,
  DELETE_CARD
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
  },

  [UPDATE_CARD](state, payload) {
    const index = state.cards.findIndex(card => card.id === payload.id)
    if (index < 0) return

    state.cards[index].activated = payload.activated
  },

  [DELETE_CARD](state, payload) {
    const index = state.cards.findIndex(card => card.code === payload)
    if (index < 0) return

    state.cards.splice(index, 1)
  }
}

export default mutations
