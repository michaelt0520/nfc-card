import Repository from '@/repository'
import {
  FETCH_LIST_USERS,
  FETCH_USER,
  UPDATE_USER
} from '../mutation-types'

const UserRepository = Repository.get('user')

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
  }
}

export default actions
