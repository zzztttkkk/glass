export class Cookie {
    static get(key) {
        return document.cookie.split(" ").reduce(
            (prev, item) => {
                let [k, ...v] = item.split("=");
                v = v.join("=")
                prev[k] = v;
                return prev;
            },
            {}
        )[key];
    }

    static set(key, val) {
        document.cookie = `${key}=${encodeURIComponent(val)}`
    }
}