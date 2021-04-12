import { createStore, createLogger } from 'vuex'
import users from './users'
import informations from './informations'
import cards from './cards'

const debug = process.env.NODE_ENV === 'development'

export default createStore({
  modules: {
    users,
    informations,
    cards
  },
  strict: debug,
  plugins: debug ? [createLogger()] : []
})
