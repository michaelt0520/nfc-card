import Repository from '@/repository'
import {
  INITIAL_STATES_SET_USER,
  SIGN_OUT
} from '../mutation-types'

const AuthRepository = Repository.get('auth')

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
}

export default actions
