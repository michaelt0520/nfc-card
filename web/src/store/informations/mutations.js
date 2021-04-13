import {
  CREATE_INFORMATION,
} from '../mutation-types'

const mutations = {
  [CREATE_INFORMATION] (state, payload) {
    state.socialInfos = [...state.socialInfos, ...payload.socialInfos]
    state.meta = payload.meta
  }
}

export default mutations
