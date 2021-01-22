import React from "react";
import {Page} from "../comps/page";
import utils from "../utils"
import {useStyletron} from "baseui";
import {useHistory} from "react-router-dom"

export function ErrorPage(props) {
	let localization = utils.glass.localization;
	let [css] = useStyletron();
	let history = useHistory();

	React.useEffect(
		function () {
			window.setTimeout(
				() => {
					history.push("/")
				},
				5000
			)
		},
		// eslint-disable-next-line
		[]
	);


	return <Page>
		<h1
			className={css({textAlign: "center"})}
		>{localization.glass.status[props.status || 404]}</h1>
	</Page>
}

export function ErrNotFoundPage(props) {
	return <ErrorPage status={404}/>
}