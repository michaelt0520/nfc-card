import axios from 'axios'

axios.interceptors.response.use(
  function (response) {
    alertTimeOut('#alert-success')

    return response
  },

  function (error) {
    if (error.response && error.response.status === 401) {
      localStorage.removeItem('user')
      localStorage.removeItem('token')
      localStorage.removeItem('is_authenticated')
    }

    alertTimeOut('#alert-error')

    return Promise.reject(error)
  }
)

const alertTimeOut = id => {
  const alertEl = document.querySelector(id)
  alertEl.classList.remove('hidden')

  setTimeout(() => {
    alertEl.classList.add('hidden')
  }, 1500)
}

export default function apiCaller(
  payload = {
    params: null,
    data: null
  }
) {
  const { method, url, params, data, headers } = payload
  headers && Object.assign(axios.defaults.headers.common, headers)

  axios.defaults.headers.common['Accept'] = 'application/json'
  axios.defaults.headers.common['Authorization'] = localStorage.getItem('token')

  return axios({
    method,
    url: `${process.env.VUE_APP_DOMAIN}${url}`,
    params,
    data
  })
}
