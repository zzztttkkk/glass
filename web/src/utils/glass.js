import {Cookie} from "./cookie";
import {User} from "./user";

let user = new User("ztk");

export const glass = {
    setTheme: null,
    setLocate: null,
    css: (v) => {
        return ""
    },
    theme: {},
    Editor: {
        quill: null,
        getContents: function () {
            return this.quill.getContents();
        },
        setContents: function (c) {
            this.quill.setContents(c);
        },
        getText: function () {
            return this.quill.getText()
        },
        getTitle: function () {
        },
        setTitle: function () {
        },
        save: function () {
            window.localStorage.setItem("editor::content", JSON.stringify(this.getContents()));
            window.localStorage.setItem("editor::title", this.getTitle());
        },
        load: function () {
            let v = window.localStorage.getItem("editor::content");
            if (v) {
                try {
                    this.setContents(JSON.parse(v));
                } catch (e) {
                    window.localStorage.removeItem("editor");
                }
            }
            v = window.localStorage.getItem("editor::title");
            if (v) {
                this.setTitle(v);
            }
        },
        on: function (v, f) {
            return this.quill.on(v, f)
        }
    },

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