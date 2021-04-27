import apiCaller from './config/api-caller'

const resources = {
  avatar: '/api/v1/upload/avatar',
  logo: '/api/v1/upload/logo'
}

Object.freeze(resources)

export default {
  createAvatar(data) {
    return apiCaller({ method: 'POST', url: resources.avatar, data: data, headers: { 'Content-Type': 'multipart/form-data' } })
  },

  createLogo(data) {
    return apiCaller({ method: 'POST', url: resources.logo, data: data, headers: { 'Content-Type': 'multipart/form-data' } })
  },
}
