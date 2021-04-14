import apiCaller from './config/api-caller'

const resources = {
  cards: '/api/v1/cards/'
}

Object.freeze(resources)

export default {
  getCardsList() {
    return apiCaller({ method: 'GET', url: resources.cards })
  },

  getCard(params) {
    return apiCaller({ method: 'GET', url: `${resources.cards}${params}` })
  },

  createCard(data) {
    return apiCaller({ method: 'POST', url: resources.cards, data: data })
  },

  updateCard(data, params) {
    return apiCaller({ method: 'PUT', url: `${resources.cards}${params}`, data: data })
  },

  deleteCard(params) {
    return apiCaller({ method: 'DELETE', url: `${resources.cards}${params}` })
  }
}
