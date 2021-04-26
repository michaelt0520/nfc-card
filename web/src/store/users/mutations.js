import {
  FETCH_LIST_USERS,
  FETCH_USER,
  UPDATE_USER,
  CREATE_INFORMATION,
  UPDATE_INFORMATION,
  DELETE_INFORMATION
} from '../mutation-types'

const mutations = {
  [FETCH_LIST_USERS](state, payload) {
    state.users = payload
  },

  [FETCH_USER](state, payload) {
    state.user = payload
  },

  [UPDATE_USER](state, payload) {
    state.user = payload
  },

  [CREATE_INFORMATION](state, payload) {
    state.user.informations.push(payload)
  },

  [UPDATE_INFORMATION](state, payload) {
    const index = state.user.informations.findIndex(socialInfo => socialInfo.id === payload.id)
    if (index < 0) return

    state.user.informations[index].visibled = payload.visibled
  },

  [DELETE_INFORMATION](state, payload) {
    const index = state.user.informations.findIndex(socialInfo => socialInfo.id === payload)
    if (index < 0) return

    state.user.informations.splice(index, 1)
  }
}

export default mutations
