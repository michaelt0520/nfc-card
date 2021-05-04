import apiCaller from './config/api-caller'

const resources = {
  user: '/api/v1/user/',
  users: '/api/v1/app/users',
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

  searchUser(params) {
    return apiCaller({ method: 'GET', url: resources.users, params: params })
  },

  updateCurrentUser(data) {
    return apiCaller({ method: 'PUT', url: resources.user, data: data })
  },

  updateCurrentUserPassword(data) {
    return apiCaller({ method: 'PUT', url: `${resources.user}password`, data: data })
  },
}
