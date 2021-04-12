import Repository from '@/repository'
import {
  FETCH_CARD
} from '../mutation-types'

const CardRepository = Repository.get('card')

const actions = {
  fetch ({ commit }, data) {
    return CardRepository.fetch(data)
      .then(res => {
        commit(FETCH_CARD, res.data.Result)

        return res
      })
  }
}

export default actions
