import apiCaller from './config/api-caller'

const resources = {
  signup: '/api/v1/signup',
  login: '/api/v1/signin',
  logout: '/api/v1/signout'
}

Object.freeze(resources)

export default {
  login (data) {
    return apiCaller({ method: 'POST', url: resources.login, data: data })
  },

  signup (data) {
    return apiCaller({ method: 'POST', url: resources.signup, data: data })
  },

  signout () {
    return apiCaller({ method: 'DELETE', url: resources.logout })
  }
}
