import apiCaller from './config/api-caller'

const resources = {
  companies: '/api/v1/companies/',
  adminCompanies: '/api/v1/admin/companies/'
}

Object.freeze(resources)

export default {
  getCompaniesList() {
    return apiCaller({ method: 'GET', url: resources.adminCompanies })
  },

  showCompany(params) {
    return apiCaller({ method: 'GET', url: `${resources.adminCompanies}${params}` })
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
