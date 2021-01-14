import {Cookie} from "./cookie";
import {toaster} from "baseui/toast";
import {glass as lg} from "../languages/glass";
import {Fetch} from "./fetch";

let user = null;

export const glass = {
	isMobile: false,
	setTheme: null,
	locate: {glass: lg},
	setLocate: null,
	css: (v) => {
		return ""
	},
	theme: {},
	toaster: {
		toaster,
		autoClose: function (msg, timeout, level = "info") {
			let key = this.toaster[level](msg);
			window.setTimeout((() => this.toaster.clear(key)), timeout)
		}
	},
	setTitle: function (t) {
	},
	Editor: {
		quill: null,
		__title: "",
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
			return this.__title;
		},
		setTitle: function (title) {
		},
		save: function () {
			window.localStorage.setItem("editor::content", JSON.stringify(this.getContents()));
			window.localStorage.setItem("editor::title", this.getTitle());
			glass.toaster.autoClose(glass.locate.glass.common.saved, 500, "positive");
		},
		load: function () {
			let v = window.localStorage.getItem("editor::content");
			if (v) {
				try {
					this.setContents(JSON.parse(v));
				} catch (e) {
					window.localStorage.removeItem("editor::content");
					window.localStorage.removeItem("editor::title");
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
		let data = await Fetch.GET("api/account/info");
		console.log(data);
		return data;
	},
};