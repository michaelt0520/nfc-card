import apiCaller from './config/api-caller'

const resources = {
  users: '/api/v1/users/',
}

Object.freeze(resources)

export default {
  getUsersList() {
    return apiCaller({ method: 'GET', url: resources.users })
  },

  updateUser(data, params) {
    return apiCaller({ method: 'PUT', url: `${resources.users}${params}`, data: data })
  }
}
