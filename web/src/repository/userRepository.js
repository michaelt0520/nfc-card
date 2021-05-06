import apiCaller from './config/api-caller'

const resources = {
  user: '/api/v1/user/',
  users: '/api/v1/app/users',
  companyUsers: '/api/v1/company/users',
  adminUsers: '/api/v1/admin/users/'
}

Object.freeze(resources)

export default {
  getCurrentUser() {
    return apiCaller({ method: 'GET', url: resources.user })
  },

  searchUser(params) {
    return apiCaller({ method: 'GET', url: resources.users, params: params })
  },

  updateCurrentUser(data) {
    return apiCaller({ method: 'PUT', url: resources.user, data: data })
  },

  updateCurrentUserPassword(data) {
    return apiCaller({ method: 'PUT', url: `${resources.user}password`, data: data })
  },

  getCurrentCompanyUsers(params) {
    return apiCaller({ method: 'GET', url: resources.companyUsers, params: params })
  },

  createCompanyUser(data) {
    return apiCaller({ method: 'POST', url: resources.companyUsers, data: data })
  },

  updateCompanyUser(data, params) {
    return apiCaller({ method: 'PUT', url: `${resources.companyUsers}/${params}`, data: data })
  },

  getUsersList(params) {
    return apiCaller({ method: 'GET', url: resources.adminUsers, params: params })
  },

  createUser(data) {
    return apiCaller({ method: 'POST', url: resources.adminUsers, data: data })
  },

  updateUser(data, params) {
    return apiCaller({ method: 'PUT', url: `${resources.adminUsers}${params}`, data: data })
  }
}
