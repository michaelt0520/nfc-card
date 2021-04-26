import Repository from '@/repository'
import {
  FETCH_LIST_USERS,
  FETCH_USER,
  UPDATE_USER,
  CREATE_INFORMATION,
  UPDATE_INFORMATION,
  DELETE_INFORMATION
} from '../mutation-types'

const UserRepository = Repository.get('user')
const InfoRepository = Repository.get('info')

const actions = {
  getUsersList({ commit }) {
    return UserRepository.getUsersList()
      .then(res => {
        commit(FETCH_LIST_USERS, res.data.result)

        return res
      })
  },

  getCurrentUser({ commit }) {
    return UserRepository.getCurrentUser()
      .then(res => {
        commit(FETCH_USER, res.data.result)

        return res
      })
  },

  updateCurrentUser({ commit }, data) {
    return UserRepository.updateCurrentUser(data)
      .then(res => {
        commit(UPDATE_USER, res.data.result)

        return res
      })
  },

  createInfo({ commit }, data) {
    return InfoRepository.createInfo(data)
      .then(res => {
        commit(CREATE_INFORMATION, res.data.result)

        return
      })
  },

  updateInfo({ commit }, data) {
    return InfoRepository.updateInfo(data.id, data.value)
      .then(res => {
        commit(UPDATE_INFORMATION, res.data.result)

        return
      })
  },

  deleteInfo({ commit }, params) {
    return InfoRepository.deleteInfo(params)
      .then(() => {
        commit(DELETE_INFORMATION, params)

        return
      })
  }
}

export default actions
