import Repository from '@/repository'
import {
  INITIAL_STATES_SET_USER,
  SIGN_OUT,
  FETCH_LIST_USERS,
  FETCH_USER,
  UPDATE_USER,
  SEARCH_USERS,
  CREATE_INFORMATION,
  UPDATE_INFORMATION,
  DELETE_INFORMATION
} from '../mutation-types'

const AuthRepository = Repository.get('auth')
const UserRepository = Repository.get('user')
const InfoRepository = Repository.get('info')
const UploadRepository = Repository.get('upload')

const actions = {
  signin({ commit }, data) {
    return AuthRepository.signin(data)
      .then(res => {
        commit(INITIAL_STATES_SET_USER, res.data.result)

        return res
      })
  },

  // eslint-disable-next-line no-unused-vars
  signup({ _commit }, data) {
    return AuthRepository.signup(data)
      .then(res => {

        return res
      })
  },

  signout({ commit }) {
    return AuthRepository.signout()
      .then(res => {
        commit(SIGN_OUT)

        return res
      })
  },

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

  updateCurrentUserPassword({ commit }, data) {
    return UserRepository.updateCurrentUserPassword(data)
      .then(res => {
        commit(SIGN_OUT)

        return res
      })
  },

  // eslint-disable-next-line no-unused-vars
  createAvatar({ _commit }, data) {
    return UploadRepository.createAvatar(data)
      .then(res => {

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
  },

  searchUser({ commit }, params) {
    return UserRepository.searchUser(params)
      .then(res => {
        commit(SEARCH_USERS, res.data.result)

        return res
      })
  }
}

export default actions
