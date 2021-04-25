import Repository from '@/repository'
import {
  FETCH_LIST_CONTACTS,
} from '../mutation-types'

const ContactRepository = Repository.get('contact')

const actions = {
  // eslint-disable-next-line no-unused-vars
  createContact({ commit }, data) {
    return ContactRepository.createContact(data)
      .then(res => {

        return res
      })
  },

  getContactsList({ commit }) {
    return ContactRepository.getContactsList()
      .then(res => {
        commit(FETCH_LIST_CONTACTS, res.data.result)

        return res
      })
  }
}

export default actions
