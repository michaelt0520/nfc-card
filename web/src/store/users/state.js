const state = {
  currentUser: localStorage.getItem('user') && JSON.parse(localStorage.getItem('user')) || {},
  isAuthenticated: localStorage.getItem('is_authenticated') || false,
  user: {},
  users: []
}

export default state
