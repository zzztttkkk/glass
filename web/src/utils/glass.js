import {Cookie} from "./cookie";

let user = undefined;

export const glass = {
    setTheme: null,
    setLocate: null,
    css: (v) => {
        return ""
    },
    theme: {},

    TryGetUser: async function () {
        if (user) {
            return user;
        }

        if (!Cookie.get(process.env.REACT_APP_AUTH_COOKIE_NAME)) {
            return null
        }
        return null;
    },
};