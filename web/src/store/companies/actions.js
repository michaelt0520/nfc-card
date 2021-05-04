import Repository from '@/repository'
import {
  FETCH_COMPANY,
  FETCH_COMPANY_USERS,
  FETCH_COMPANY_CARDS,
  FETCH_LIST_COMPANIES,
  UPDATE_COMPANY
} from '../mutation-types'

const CompanyRepository = Repository.get('company')
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
    return CompanyRepository.getCurrentCompanyUsers(params)
      .then(res => {
        commit(FETCH_COMPANY_USERS, res.data.result)

        return res
      })
  },

  getCurrentCompanyCards({ commit }, params) {
    return CompanyRepository.getCurrentCompanyCards(params)
      .then(res => {
        commit(FETCH_COMPANY_CARDS, res.data.result)

        return res
      })
  },

  updateCurrentCompany({ commit }, data) {
    return CompanyRepository.updateCurrentCompany(data)
      .then(res => {
        commit(UPDATE_COMPANY, res.data.result)

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

  getCompaniesList({ commit }) {
    return CompanyRepository.getCompaniesList()
      .then(res => {
        commit(FETCH_LIST_COMPANIES, res.data.result)

        return res
      })
  }
}

export default actions
