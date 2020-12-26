import React from "react";
import {StyledLink} from "baseui/link";
import {useHistory} from "react-router-dom"

export function Link(props) {
    let history = useHistory();
    return <StyledLink
        href={props.href}
        onClick={(evt) => {
            evt.stopPropagation();
            history.push(props.href);
        }}
    >{props.children}</StyledLink>
}