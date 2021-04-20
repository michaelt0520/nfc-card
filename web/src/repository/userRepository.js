import apiCaller from './config/api-caller'

const resources = {
  users: '/api/v1/users/',
  adminUsers: '/api/v1/admin/users/'
}

Object.freeze(resources)

export default {
  getUsersList() {
    return apiCaller({ method: 'GET', url: resources.adminUsers })
  },

  updateUser(data, params) {
    return apiCaller({ method: 'PUT', url: `${resources.adminUsers}${params}`, data: data })
  }
}
