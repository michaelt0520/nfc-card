import Repository from '@/repository'
import {
  FETCH_CARD,
  FETCH_LIST_CARDS
} from '../mutation-types'

const CardRepository = Repository.get('card')

const actions = {
  getCardsList ({ commit }) {
    return CardRepository.getCardsList()
      .then(res => {
        commit(FETCH_LIST_CARDS, res.data.Result)

        return res
      })
  },

  getCard ({ commit }, data) {
    return CardRepository.getCard(data)
      .then(res => {
        commit(FETCH_CARD, res.data.Result)

        return res
      })
  }
}

export default actions
