export class Fragment {
	constructor() {
		this.ele = document.createDocumentFragment();
	}

	appendElement(ele) {
		this.ele.appendChild(ele);
		return this;
	}

	appendElements(eles) {
		eles.forEach((item => this.appendElement(item)));
		return this;
	}

	appendCSS(href) {
		let link = document.createElement("link");
		link.setAttribute("href", href);
		link.setAttribute("rel", "stylesheet");
		this.ele.appendChild(link);
		return this;
	}

	appendJavaScript(src) {
		let quill = document.createElement("script");
		quill.setAttribute("src", src);
		this.ele.appendChild(quill);
		return this;
	}

	insertInto(ele) {
		ele.appendChild(this.ele);
	}
}