import Repository from '@/repository'
import {
  FETCH_CARD,
  FETCH_LIST_CARDS,
  CREATE_CARD,
  UPDATE_CARD,
  DELETE_CARD
} from '../mutation-types'

const CardRepository = Repository.get('card')

const actions = {
  getCardsList({ commit }, params) {
    return CardRepository.getCardsList(params)
      .then(res => {
        commit(FETCH_LIST_CARDS, res.data.result)

        return res
      })
  },

  getCard({ commit }, data) {
    return CardRepository.getCard(data)
      .then(res => {
        commit(FETCH_CARD, res.data.result)

        return res
      })
  },

  createCard({ commit }, data) {
    return CardRepository.createCard(data)
      .then(res => {
        commit(CREATE_CARD, res.data.result)

        return res
      })
  },

  updateCard({ commit }, data) {
    return CardRepository.updateAdminCard(data.body, data.params)
      .then(res => {
        commit(UPDATE_CARD, res.data.result)

        return res
      })
  },

  deleteCard({ commit }, data) {
    return CardRepository.deleteAdminCard(data)
      .then(res => {
        commit(DELETE_CARD, data)

        return res
      })
  }
}

export default actions
