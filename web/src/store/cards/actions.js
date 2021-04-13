import Repository from '@/repository'
import {
  FETCH_CARD,
  FETCH_LIST_CARDS
} from '../mutation-types'

const CardRepository = Repository.get('card')

const actions = {
  getListCards ({ commit }) {
    return CardRepository.index()
      .then(res => {
        commit(FETCH_LIST_CARDS, res.data.Result)

        return res
      })
  },

  show ({ commit }, data) {
    return CardRepository.show(data)
      .then(res => {
        commit(FETCH_CARD, res.data.Result)

        return res
      })
  }
}

export default actions
