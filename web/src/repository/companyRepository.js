import apiCaller from './config/api-caller'

const resources = {
  companies: '/api/v1/companies/',
}

Object.freeze(resources)

export default {
  index() {
    return apiCaller({ method: 'GET', url: resources.companies })
  },

  show(params) {
    return apiCaller({ method: 'GET', url: `${resources.companies}${params}` })
  },

  create(data) {
    return apiCaller({ method: 'POST', url: resources.companies, data: data })
  },

  update(data, params) {
    return apiCaller({ method: 'PUT', url: `${resources.companies}${params}`, data: data })
  },

  delete(params) {
    return apiCaller({ method: 'DELETE', url: `${resources.companies}${params}` })
  }
}
