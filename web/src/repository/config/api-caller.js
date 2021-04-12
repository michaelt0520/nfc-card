import axios from 'axios'

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
