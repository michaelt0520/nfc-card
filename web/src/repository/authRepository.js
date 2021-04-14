import apiCaller from './config/api-caller'

const resources = {
  auth: '/api/v1'
}

Object.freeze(resources)

export default {
  signin (data) {
    return apiCaller({ method: 'POST', url: `${resources.auth}/signin`, data: data })
  },

  signup (data) {
    return apiCaller({ method: 'POST', url: `${resources.auth}/signup`, data: data })
  },

  signout () {
    return apiCaller({ method: 'DELETE', url: `${resources.auth}/signout` })
  }
}
