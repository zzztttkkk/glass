import {Cookie} from "./cookie";

let user = undefined;

export const glass = {
    setTheme: null,
    setLocate: null,

    TryGetUser: async function () {
        if (user) {
            return user;
        }

        if (!Cookie.get(process.env.REACT_APP_AUTH_COOKIE_NAME)) {
            return null
        }
        // todo fetch api
        return null;
    },
};