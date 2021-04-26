// repositoryFactory

import AuthRepository from './authRepository'
import CardRepository from './cardRepository'
import UserRepository from './userRepository'
import InfoRepository from './infoRepository'
import CompanyRepository from './companyRepository'
import ContactRepository from './contactRepository'

const repositories = {
  auth: AuthRepository,
  user: UserRepository,
  info: InfoRepository,
  card: CardRepository,
  company: CompanyRepository,
  contact: ContactRepository
}

export default {
  get: name => repositories[name]
}
