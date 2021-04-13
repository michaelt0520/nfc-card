import apiCaller from './config/api-caller'

const resources = {
  users: '/api/v1/users/',
}

Object.freeze(resources)

export default {
  index() {
    return apiCaller({ method: 'GET', url: resources.users })
  },

  update(data, params) {
    return apiCaller({ method: 'PUT', url: `${resources.users}${params}`, data: data })
  }
}
