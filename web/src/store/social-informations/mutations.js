import {
  ADD_SOCIAL_INFORMATION,
} from '../mutation-types'

const mutations = {
  [ADD_SOCIAL_INFORMATION] (state, payload) {
    state.socialInfos = [...state.socialInfos, ...payload.socialInfos]
    state.meta = payload.meta
  }
}

export default mutations