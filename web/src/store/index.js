import { createStore } from 'vuex'
import users from './users'
import socialInformations from './social-informations'

export default createStore({
  modules: {
    users,
    socialInformations
  }
})
