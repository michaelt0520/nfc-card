import Repository from '@/repository'
import {
  FETCH_LIST_COMPANIES
} from '../mutation-types'

const CompanyRepository = Repository.get('company')

const actions = {
  getListCompanies ({ commit }) {
    return CompanyRepository.index()
      .then(res => {
        commit(FETCH_LIST_COMPANIES, res.data.Result)

        return res
      })
  }
}

export default actions
