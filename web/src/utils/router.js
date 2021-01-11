import React from "react";

import {Route, Switch} from "react-router-dom";


function isLazyObject(obj) {
	return obj.$$typeof && typeof obj.$$typeof === "symbol" && obj.$$typeof.toString() === "Symbol(react.lazy)"
}


export class PathSwitch extends React.Component {
    constructor(prefix, notfound) {
        super();
        this.prefix = prefix;
        this.parent = null;
        this.m = {};
        this.children = [];

        this.notfound = notfound;
    }

	register(path, val) {
		if (isLazyObject(val)) {
			// todo support lazy
			throw new Error("todo lazy");
		}

		if (typeof this.m[path] != "undefined") {
			throw path;
		}
		this.m[path] = val;
	}

	include(ps) {
		ps.parent = this;
		this.children.push(ps);
	}

	full() {
		if (this.parent) {
			return this.parent.full() + this.prefix;
		}
		return this.prefix;
	}

	render() {
		let l = this.children.length;
		return <Switch>
			{
				this.children.map(
					(item, index) => {
						return <Route key={index} path={item.prefix}>{item.render()}</Route>
					}
				)
			}
			{
				Object.entries(this.m).map(
					([k, v], index) => {
						return <Route exact={this.notfound != null} key={l + index} path={this.full() + k}>{v}</Route>
					}
				)
			}
			{
				this.notfound != null && <Route component={this.notfound}/>
			}
		</Switch>
	}
}
