
export const api = {
    get: (path, options) => request('GET', path, options),
    post: (path, options) => request('POST', path, options),
    put: (path, options) => request('PUT', path, options),
    delete: (path, options) => request('DELETE', path, options),
    getJSON: (path, options) => requestJSON('GET', path, options),
    postJSON: (path, options) => requestJSON('POST', path, options),
    putJSON: (path, options) => requestJSON('PUT', path, options),
    deleteJSON: (path, options) => requestJSON('DELETE', path, options),
}

async function request(method, path, options = {}) {
    const init = {
        method,
        credentials: 'same-origin',
        ...options,
    }
    try {
        const response = await fetch('/api' + path, init)
        const data = response.status === 204 ? null : await response.text()
        return {
            status: response.status,
            ok: response.ok,
            data,
            response,
        }
    } catch (err) {
        if (err.name === 'AbortError') {
            return {
                status: 0,
                ok: false,
                data: null,
                aborted: true,
                err,
            }
        }
        throw err
    }
}

async function requestJSON(method, path, options = {}) {
    const { body, ...init } = options
    if (body) {
        init.body = JSON.stringify(body)
    }
    const response = await request(method, path, init)
    if (response.ok) {
        response.data = JSON.parse(response.data)
    }
    return response
}
