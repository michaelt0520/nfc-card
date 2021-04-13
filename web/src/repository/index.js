// repositoryFactory

import AuthRepository from './authRepository'
import CardRepository from './cardRepository'
import UserRepository from './userRepository'
import CompanyRepository from './companyRepository'

const repositories = {
  auth: AuthRepository,
  user: UserRepository,
  card: CardRepository,
  company: CompanyRepository
}

export default {
  get: name => repositories[name]
}
