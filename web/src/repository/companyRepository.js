import apiCaller from './config/api-caller'

const resources = {
  company: '/api/v1/company/',
  adminCompanies: '/api/v1/admin/companies/'
}

Object.freeze(resources)

export default {
  getCurrentCompany() {
    return apiCaller({ method: 'GET', url: resources.company })
  },

  updateCurrentCompany(data) {
    return apiCaller({ method: 'PUT', url: resources.company, data: data })
  },

  getCompaniesList(params) {
    return apiCaller({ method: 'GET', url: resources.adminCompanies, params: params })
  },

  createConpany(data) {
    return apiCaller({ method: 'POST', url: resources.adminCompanies, data: data })
  },

  updateCompany(data, params) {
    return apiCaller({ method: 'PUT', url: `${resources.adminCompanies}${params}`, data: data })
  },

  deleteCompany(params) {
    return apiCaller({ method: 'DELETE', url: `${resources.adminCompanies}${params}` })
  }
}
