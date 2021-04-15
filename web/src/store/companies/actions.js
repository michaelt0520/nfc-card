import Repository from '@/repository'
import {
  FETCH_LIST_COMPANIES
} from '../mutation-types'

const CompanyRepository = Repository.get('company')

const actions = {
  getCompaniesList({ commit }) {
    return CompanyRepository.getCompaniesList()
      .then(res => {
        commit(FETCH_LIST_COMPANIES, res.data.result)

        return res
      })
  }
}

export default actions
