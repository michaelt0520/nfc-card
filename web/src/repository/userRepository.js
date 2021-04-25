import apiCaller from './config/api-caller'

const resources = {
  user: '/api/v1/user/',
  adminUsers: '/api/v1/admin/users/'
}

Object.freeze(resources)

export default {
  getCurrentUser() {
    return apiCaller({ method: 'GET', url: resources.user })
  },

  getUsersList() {
    return apiCaller({ method: 'GET', url: resources.adminUsers })
  },

  updateCurrentUser(data) {
    return apiCaller({ method: 'PUT', url: resources.user, data: data })
  }
}
