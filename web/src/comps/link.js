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
	let history = useHistory();
	let [, theme] = useStyletron();
	let fn = (evt) => {
		evt.stopPropagation();
		history.push(props.href);
	};

	return <Button overrides={props.BTN || {}} tabIndex={-1} onClick={fn}>
		<StyledLink
			className={
				utils.Override.Style(
					props,
					"Link",
					{
						color: `${theme.colors.background}`,
						":visited": {color: `${theme.colors.background}`}
					}
				)
			}
			href={props.href} onClick={fn}
			{...utils.Override.Props(props, "Link", {})}
		>{props.children}</StyledLink>
	</Button>
}