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
        if (this.parent !== null) {
            return this.parent.full() + this.prefix;
        }
        return this.prefix;
    }

    render() {
        let l = this.children.length;
        return <Switch key={this.index}>
            {
                this.children.map(
                    (item, index) => {
                        return <Route key={index} path={item.prefix}>
                            {item.render()}
                        </Route>
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
                this.notfound != null ? <Route component={this.notfound}/> : null
            }
        </Switch>
    }
}