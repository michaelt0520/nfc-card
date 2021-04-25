import apiCaller from './config/api-caller'

const resources = {
  card: '/api/v1/app/card/',
  adminCards: '/api/v1/admin/cards/'
}

Object.freeze(resources)

export default {
  getCardsList() {
    return apiCaller({ method: 'GET', url: resources.adminCards })
  },

  getCard(params) {
    return apiCaller({ method: 'GET', url: `${resources.card}${params}` })
  },

  createCard(data) {
    return apiCaller({ method: 'POST', url: resources.adminCards, data: data })
  },

  updateCard(data, params) {
    return apiCaller({ method: 'PUT', url: `${resources.adminCards}${params}`, data: data })
  },

  deleteCard(params) {
    return apiCaller({ method: 'DELETE', url: `${resources.adminCards}${params}` })
  }
}
