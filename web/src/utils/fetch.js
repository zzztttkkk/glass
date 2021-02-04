export class Fetch {
	static async request(method, path, query, headers, body, sendJson) {
		let isCORS = false;
		if (!path.startsWith("http://") && !path.startsWith("https://") && !path.startsWith("://")) {
			path = `${window.location.origin}/${path}`
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
			url.search = url.search.length > 0 ? `${url.search}&${qs.join("&")}` : qs.join("&");
		}

		let options = {headers: {}, method: method};
		if (isCORS) {
			options["mode"] = "cors"
		}
		if (headers) {
			options["headers"] = headers;
		}

		if (body) {
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
		}
		return await fetch(url.toString(), options);
	}

	static async GET(path, query, headers) {
		if (query instanceof Element) {
			query = new FormData(query);
		}

		if (query instanceof FormData) {
			query = Array.from(query.entries()).reduce((prev, [k, v]) => (prev[k] = v || true) && prev, {})
		}
		return await Fetch.request("get", path, query, headers);
	}

	static async DELETE(path, query, headers) {
		return await Fetch.request("delete", path, query, headers);
	}

	static async POST(path, query, headers, body) {
		return await Fetch.request("post", path, query, headers, body);
	}

	static async PUT(path, query, headers, body) {
		return await Fetch.request("put", path, query, headers, body);
	}

	static async PostJSON(path, query, headers, body) {
		return await Fetch.request("post", path, query, headers, body, true);
	}
}