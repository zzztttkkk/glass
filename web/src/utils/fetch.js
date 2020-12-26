export async function request(method, path, query, headers, body, sendJson) {
    let isCORS = false;
    if (!path.startsWith("http://") && !path.startsWith("https://") && !path.startsWith("://")) {
        if (path.startsWith("/")) {
            path = `${process.env.REACT_APP_HTTP_SCHEME}://${process.env.REACT_APP_DOMAIN}${path}`
        } else {
            path = `${process.env.REACT_APP_HTTP_SCHEME}://${process.env.REACT_APP_DOMAIN}${process.env.REACT_APP_API_PATH_PREFIX}${path}`
        }
    } else {
        isCORS = true;
    }

    let url = new URL(path);
    if (query) {
        let qs = Object.entries(query).reduce(
            (prev, [key, val]) => {
                if (val instanceof Array) {
                    val.forEach(
                        (item) => {
                            prev.push(`${encodeURIComponent(key)}=${encodeURIComponent(item)}`)
                        }
                    )
                } else {
                    prev.push(`${encodeURIComponent(key)}=${encodeURIComponent(val.toString())}`)
                }
                return prev;
            },
            []
        );
        url.search = `${url.search}&${qs.join("&")}`
    }

    let options = {headers: {}};
    if (isCORS) {
        options["mode"] = "cors"
    }
    if (headers) {
        options["headers"] = headers;
    }
    if (sendJson) {
        options.headers["Content-Type"] = "application/json"
        options["body"] = JSON.stringify(body);
    } else {
        if (body instanceof FormData) {
            options["body"] = body;
        } else {
            if (typeof body === "string") {
                options["body"] = body;
                options.headers["Content-Type"] = "text/plain"
            } else if (typeof body === "object") {
                let fa = new FormData();
                Object.entries(body).forEach(
                    ([key, val]) => {
                        if (val instanceof Array) {
                            val.forEach(
                                (item) => {
                                    fa.append(key, item.toString())
                                }
                            )
                        } else {
                            fa.append(key, val.toString());
                        }
                    }
                );
                options["body"] = fa;
            } else {
                throw new Error(`unexpected request body: "${body}"`);
            }
        }
    }
    return await fetch(url.toString(), options);
}