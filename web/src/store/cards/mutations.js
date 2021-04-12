import {
  FETCH_CARD,
} from '../mutation-types'

const mutations = {
  [FETCH_CARD] (state, payload) {
    state.card = payload.card
  }
}

export default mutations
