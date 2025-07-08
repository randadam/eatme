import * as api from "./client"

export class HttpApiError extends Error {
    public readonly status: number
    public readonly apiError: api.ModelsApiError

    constructor(
        status: number,
        apiError: api.ModelsApiError,
    ) {
        super(apiError.message)
        this.status = status
        this.apiError = apiError
    }
}

api.defaults.baseUrl = "/api"

api.defaults.fetch = async (url, opts = {}) => {
    const token = localStorage.getItem("token")
    if (token) {
        opts.headers = {
            "Authorization": token ? `Bearer ${token}` : "",
            ...opts.headers,
        }
    }
    const res = await fetch(url, opts)

    if (!res.ok) {
        let parsed: api.ModelsApiError = {
            code: "UNKNOWN_ERROR",
            message: `HTTP ${res.status}`
        }
        try {
            const body = await res.json();
            if (body?.error) {
                parsed = body.error
            }
            console.error(parsed)
        } catch (e) {
            console.error(`Error extracting API error details: ${e}`)
            console.error(`Response body: ${await res.text()}`)
        }

        throw new HttpApiError(res.status, parsed)
    }

    return res
}

export default api