import { createStore, createLogger } from 'vuex'
import users from './users'
import socialInformations from './social-informations'

const debug = process.env.NODE_ENV === 'development'

export default createStore({
  modules: {
    users,
    socialInformations
  },
  strict: debug,
  plugins: debug ? [createLogger()] : []
})
