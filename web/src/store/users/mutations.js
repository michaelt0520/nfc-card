import {
  STORE_USER,
} from '../mutation-types'

const mutations = {
  [STORE_USER] (state, payload) {
    state.user = payload.user
  }
}

export default mutations