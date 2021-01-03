import React from "react";
import {StyledLink} from "baseui/link";
import {Button} from "baseui/button";
import {useHistory} from "react-router-dom"
import utils from "../utils"
import {useStyletron} from "baseui";

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

export function BtnLink(props) {
    const [css] = useStyletron();
    return <Button className={css(props.btnStyle || {})}>
        <Link {...props} overrides={{Root: {style: props.linkStyle || {}}}}/>
    </Button>
}