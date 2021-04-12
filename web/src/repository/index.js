// repositoryFactory

import AuthRepository from './authRepository'
import CardRepository from './cardRepository'

const repositories = {
  auth: AuthRepository,
  card: CardRepository,
}

export default {
  get: name => repositories[name]
}
