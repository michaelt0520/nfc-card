import apiCaller from './config/api-caller'

const resources = {
  company: '/api/v1/company/',
  users: '/api/v1/company/users',
  cards: '/api/v1/company/cards',
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

  getCurrentCompanyUsers(params) {
    return apiCaller({ method: 'GET', url: resources.users, params: params })
  },

  getCurrentCompanyCards(params) {
    return apiCaller({ method: 'GET', url: resources.cards, params: params })
  },

  getCompaniesList() {
    return apiCaller({ method: 'GET', url: resources.adminCompanies })
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
