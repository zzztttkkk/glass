import React from "react";
import {StyledLink} from "baseui/link";
import {useHistory} from "react-router-dom"
import utils from "../utils"

export function Link(props) {
    let history = useHistory();
    return <StyledLink
        className={utils.Override.Style(props, "Root", {})}
        href={props.href}
        onClick={(evt) => {
            evt.stopPropagation();
            history.push(props.href);
        }}
        {...utils.Override.Props(props, "Root", {})}
    >{props.children}</StyledLink>
}