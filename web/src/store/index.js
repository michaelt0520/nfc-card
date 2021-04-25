import { createStore, createLogger } from 'vuex'
import auths from './auths'
import users from './users'
import informations from './informations'
import cards from './cards'
import companies from './companies'
import contacts from './contacts'

const debug = process.env.NODE_ENV === 'development'

export default createStore({
  modules: {
    auths,
    users,
    informations,
    cards,
    companies,
    contacts
  },
  strict: debug,
  plugins: debug ? [createLogger()] : []
})
