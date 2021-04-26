import apiCaller from './config/api-caller'

const resources = {
  info: '/api/v1/user/informations/',
}

Object.freeze(resources)

export default {
  createInfo(data) {
    return apiCaller({ method: 'POST', url: resources.info, data: data })
  },

  updateInfo(params, data) {
    return apiCaller({ method: 'PUT', url: `${resources.info}${params}`, data: data })
  },

  deleteInfo(params) {
    return apiCaller({ method: 'DELETE', url: `${resources.info}${params}` })
  }
}
