const state = {
  currentUser: localStorage.getItem('user') && JSON.parse(localStorage.getItem('user')) || {},
  isAuthenticated: localStorage.getItem('is_authenticated') || false
}

export default state
