import apiCaller from './config/api-caller'

const resources = {
  fetch: '/api/v1/cards',
  create: '/api/v1/cards',
  update: '/api/v1/cards'
}

Object.freeze(resources)

export default {
  fetch (data) {
    return apiCaller({ method: 'GET', url: `${resources.fetch}/${data}` })
  },

  create (data) {
    return apiCaller({ method: 'POST', url: resources.create, data: data })
  },

  update (data, params) {
    return apiCaller({ method: 'PUT', url: `${resources.update}/${params}`, data: data })
  }
}
