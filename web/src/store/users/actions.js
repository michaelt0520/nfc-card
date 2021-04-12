import Repository from '@/repository'
import {
  INITIAL_STATES_SET_USER,
  CREATE_USER
} from '../mutation-types'

const AuthRepository = Repository.get('auth')

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

        return res
      })
  }
}

export default actions
