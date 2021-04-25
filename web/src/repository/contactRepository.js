import apiCaller from './config/api-caller'

const resources = {
  contact: '/api/v1/app/contact',
  adminContacts: '/api/v1/admin/contacts/'
}

Object.freeze(resources)

export default {
  getContactsList() {
    return apiCaller({ method: 'GET', url: resources.adminContacts })
  },

  createContact(data) {
    return apiCaller({ method: 'POST', url: resources.contact, data: data })
  },

  deleteContact(params) {
    return apiCaller({ method: 'DELETE', url: `${resources.adminContacts}${params}` })
  }
}
