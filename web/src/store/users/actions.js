import Repository from '@/repository'
import {
  INITIAL_STATES_SET_USER,
  CREATE_USER,
  FETCH_LIST_USERS
} from '../mutation-types'

const AuthRepository = Repository.get('auth')
const UserRepository = Repository.get('user')

const actions = {
  login ({ commit }, data) {
    return AuthRepository.login(data)
      .then(res => {
        commit(INITIAL_STATES_SET_USER, res.data.Result)

        return res
      })
  },

  signup ({ commit }, data) {
    return AuthRepository.signup(data)
      .then(res => {
        commit(CREATE_USER, res.data.Result)

        return res
      })
  },

  signout () {
    return AuthRepository.signout()
      .then(res => {
        localStorage.removeItem('user')
        localStorage.removeItem('token')
        localStorage.removeItem('is_authenticated')

        return res
      })
  },

  getListUsers ({ commit }) {
    return UserRepository.index()
      .then(res => {
        commit(FETCH_LIST_USERS, res.data.Result)

        return res
      })
  }
}

export default actions
