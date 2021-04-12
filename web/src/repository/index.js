// repositoryFactory

import AuthRepository from './authRepository'

const repositories = {
  auth: AuthRepository
}

export default {
  get: name => repositories[name]
}
