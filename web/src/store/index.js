import { createStore, createLogger } from 'vuex'
import users from './users'
import cards from './cards'
import companies from './companies'
import contacts from './contacts'

const debug = process.env.NODE_ENV === 'development'

export default createStore({
  modules: {
    users,
    cards,
    companies,
    contacts
  },
  strict: debug,
  plugins: debug ? [createLogger()] : []
})
