import Repository from '@/repository'
import {
  FETCH_COMPANY,
  FETCH_COMPANY_USERS,
  FETCH_COMPANY_CARDS,
  FETCH_LIST_COMPANIES,
  UPDATE_CURRENT_COMPANY,
  CREATE_COMPANY,
  UPDATE_COMPANY,
  DELETE_COMPANY,
  UPDATE_CARD,
  CREATE_USER,
  UPDATE_USER
} from '../mutation-types'

const CompanyRepository = Repository.get('company')
const UserRepository = Repository.get('user')
const CardRepository = Repository.get('card')
const UploadRepository = Repository.get('upload')

const actions = {
  getCurrentCompany({ commit }) {
    return CompanyRepository.getCurrentCompany()
      .then(res => {
        commit(FETCH_COMPANY, res.data.result)

        return res
      })
  },

  getCurrentCompanyUsers({ commit }, params) {
    return UserRepository.getCurrentCompanyUsers(params)
      .then(res => {
        commit(FETCH_COMPANY_USERS, res.data.result)

        return res
      })
  },

  getCurrentCompanyCards({ commit }, params) {
    return CardRepository.getCurrentCompanyCards(params)
      .then(res => {
        commit(FETCH_COMPANY_CARDS, res.data.result)

        return res
      })
  },

  updateCurrentCompany({ commit }, data) {
    return CompanyRepository.updateCurrentCompany(data)
      .then(res => {
        commit(UPDATE_CURRENT_COMPANY, res.data.result)

        return res
      })
  },

  // eslint-disable-next-line no-unused-vars
  createLogo({ _commit }, data) {
    return UploadRepository.createLogo(data)
      .then(res => {

        return res
      })
  },

  getCompaniesList({ commit }, params) {
    return CompanyRepository.getCompaniesList(params)
      .then(res => {
        commit(FETCH_LIST_COMPANIES, res.data.result)

        return res
      })
  },

  createCompany({ commit }, data) {
    return CompanyRepository.createConpany(data)
      .then(res => {
        commit(CREATE_COMPANY, res.data.result)

        return res
      })
  },

  updateCard({ commit }, data) {
    return CardRepository.updateCompanyCard(data.body, data.params)
      .then(res => {
        commit(UPDATE_CARD, res.data.result)

        return res
      })
  },

  createCompanyUser({ commit }, data) {
    return UserRepository.createCompanyUser(data)
      .then(res => {
        commit(CREATE_USER, res.data.result)

        return res
      })
  },

  updateCompanyUser({ commit }, data) {
    return UserRepository.updateCompanyUser(null, data.params)
      .then(res => {
        commit(UPDATE_USER, data.params)

        return res
      })
  },

  updateCompany({ commit }, data) {
    return CompanyRepository.updateCompany(data.body, data.params)
      .then(res => {
        commit(UPDATE_COMPANY, res.data.result)

        return res
      })
  },

  deleteCompany({ commit }, data) {
    return CompanyRepository.deleteCompany(data)
      .then(res => {
        commit(DELETE_COMPANY, data)

        return res
      })
  }
}

export default actions
