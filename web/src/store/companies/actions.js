import Repository from '@/repository'
import {
  FETCH_COMPANY,
  FETCH_LIST_COMPANIES
} from '../mutation-types'

const CompanyRepository = Repository.get('company')

const actions = {
  getCurrentCompany({ commit }) {
    return CompanyRepository.getCurrentCompany()
      .then(res => {
        commit(FETCH_COMPANY, res.data.result)

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
