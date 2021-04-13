import { createStore, createLogger } from 'vuex'
import users from './users'
import informations from './informations'
import cards from './cards'
import companies from './companies'

const debug = process.env.NODE_ENV === 'development'

export default createStore({
  modules: {
    users,
    informations,
    cards,
    companies
  },
  strict: debug,
  plugins: debug ? [createLogger()] : []
})
