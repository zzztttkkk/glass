import React from "react";
import {Route, Switch} from "react-router-dom";

export class PathSwitch {
	constructor(prefix, notfound) {
		this.prefix = prefix;
		this.parent = null;
		this.m = {};
		this.children = [];

		this.notfound = notfound;
	}

	register(path, val) {
		if (typeof val !== "function" && val instanceof React.Component) {
			throw new Error(`${val} is not a React Component`);
		}

		if (typeof this.m[path] != "undefined") {
			throw path;
		}
		this.m[path] = val;
	}

	include(ps) {
		if (ps instanceof PathSwitch) {
			ps.parent = this;
			this.children.push(ps);
		} else {
			throw new Error(`${ps} is not a PathSwitch`);
		}
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
						let Comp = v;
						return <Route exact={this.notfound != null} key={l + index} path={this.full() + k}>
							<Comp/>
						</Route>
					}
				)
			}
			{
				this.notfound != null && <Route component={this.notfound}/>
			}
		</Switch>
	}
}
