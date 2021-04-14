import apiCaller from './config/api-caller'

const resources = {
  companies: '/api/v1/companies/',
}

Object.freeze(resources)

export default {
  getCompaniesList() {
    return apiCaller({ method: 'GET', url: resources.companies })
  },

  showCompany(params) {
    return apiCaller({ method: 'GET', url: `${resources.companies}${params}` })
  },

  createConpany(data) {
    return apiCaller({ method: 'POST', url: resources.companies, data: data })
  },

  updateCompany(data, params) {
    return apiCaller({ method: 'PUT', url: `${resources.companies}${params}`, data: data })
  },

  deleteCompany(params) {
    return apiCaller({ method: 'DELETE', url: `${resources.companies}${params}` })
  }
}
