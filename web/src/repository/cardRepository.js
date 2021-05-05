import apiCaller from './config/api-caller'

const resources = {
  card: '/api/v1/app/cards/',
  companyCards: '/api/v1/company/cards',
  adminCards: '/api/v1/admin/cards/'
}

Object.freeze(resources)

export default {
  getCard(params) {
    return apiCaller({ method: 'GET', url: `${resources.card}${params}` })
  },

  getCurrentCompanyCards(params) {
    return apiCaller({ method: 'GET', url: resources.companyCards, params: params })
  },

  updateCompanyCard(data, params) {
    return apiCaller({ method: 'PUT', url: `${resources.companyCards}/${params}`, data: data })
  },

  getCardsList(params) {
    return apiCaller({ method: 'GET', url: resources.adminCards, params: params })
  },

  createCard(data) {
    return apiCaller({ method: 'POST', url: resources.adminCards, data: data })
  },

  updateAdminCard(data, params) {
    return apiCaller({ method: 'PUT', url: `${resources.adminCards}${params}`, data: data })
  },

  deleteCard(params) {
    return apiCaller({ method: 'DELETE', url: `${resources.adminCards}${params}` })
  }
}
