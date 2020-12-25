// @flow

import * as React from "react";

import {Route, Switch} from "react-router-dom";


export class PathSwitch {
    prefix : string;
    parent : ?PathSwitch;
    m : {string: PathSwitch};
    children : Array<PathSwitch>;
    notfound: (any)=>React.Node;

    constructor(prefix: string, notfound: (any)=>React.Node) {
        this.prefix = prefix;
        this.parent = null;
        this.m = {};
        this.children = [];

        this.notfound = notfound;
    }

    register(path: string, val: PathSwitch) {
        if (typeof this.m[path] != "undefined") {
            throw path;
        }
        this.m[path] = val;
    }

    include(ps: PathSwitch) {
        ps.parent = this;
        this.children.push(ps);
    }

    full(): string {
        if (this.parent) {
            return this.parent.full() + this.prefix;
        }
        return this.prefix;
    }

    render(): React.Node {
        let l = this.children.length;
        return <Switch>
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