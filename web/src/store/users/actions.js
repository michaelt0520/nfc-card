import Repository from '@/repository'
import {
  INITIAL_STATES_SET_USER,
  FETCH_LIST_USERS
} from '../mutation-types'

const AuthRepository = Repository.get('auth')
const UserRepository = Repository.get('user')

const actions = {
  signin ({ commit }, data) {
    return AuthRepository.signin(data)
      .then(res => {
        commit(INITIAL_STATES_SET_USER, res.data.Result)

        return res
      })
  },

  signup (data) {
    return AuthRepository.signup(data)
      .then(res => {

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

  getUsersList ({ commit }) {
    return UserRepository.getUsersList()
      .then(res => {
        commit(FETCH_LIST_USERS, res.data.Result)

        return res
      })
  }
}

export default actions
