import apiCaller from './config/api-caller'

const resources = {
  cards: '/api/v1/cards/',
}

Object.freeze(resources)

export default {
  index() {
    return apiCaller({ method: 'GET', url: resources.cards })
  },

  show(params) {
    return apiCaller({ method: 'GET', url: `${resources.cards}${params}` })
  },

  create(data) {
    return apiCaller({ method: 'POST', url: resources.cards, data: data })
  },

  update(data, params) {
    return apiCaller({ method: 'PUT', url: `${resources.cards}${params}`, data: data })
  },

  delete(params) {
    return apiCaller({ method: 'DELETE', url: `${resources.cards}${params}` })
  }
}
