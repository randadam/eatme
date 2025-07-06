import * as api from './client'

api.defaults.baseUrl = '/api'

api.defaults.fetch = (url, opts = {}) => {
    const token = localStorage.getItem('token')
    if (token) {
        opts.headers = {
            'Authorization': token ? `Bearer ${token}` : '',
            ...opts.headers,
        }
    }
    return fetch(url, opts)
}

export default api