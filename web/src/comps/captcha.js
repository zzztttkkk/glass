import React from "react";
import utils from "../utils"


export function Captcha(props) {
	const [src, setSrc] = React.useState("");
	const [state, setState] = React.useState(0);
	const [voidClickCount, setVoidClickCount] = React.useState(0);

	const imgRef = React.createRef();
	const toaster = utils.glass.toaster;
	const locate = utils.glass.localization;
	const key = `${props.where}.captcha.last`

	const incr = function () {
		let v = state + 1;
		if (v > 10) {
			v = 1;
		}
		setState(v);
	}
	props.obj.incr = incr;

	React.useEffect(
		function () {
			if (voidClickCount === 10) {
				setVoidClickCount(0);
				toaster.autoClose(locate.glass.common.captchaWarning, 1500, toaster.warning);
			}
		},
		// eslint-disable-next-line
		[voidClickCount]
	)

	React.useEffect(
		function () {
			async function reCaptchaImage() {
				let lastTime = Number(window.sessionStorage.getItem(key));
				if (lastTime !== 0 && !isNaN(lastTime) && Date.now() - lastTime < 65 * 1000) {
					if (src) {
						setVoidClickCount(state + 1);
						return;
					}
				}

				let res = await utils.Fetch.GET("/api/captcha.png");
				if (res.status === 200) {
					window.sessionStorage.setItem(key, Date.now());
					setSrc(URL.createObjectURL(await res.blob()));
				} else {
					toaster.autoClose(locate.glass.status[res.status], 2000, toaster.negative);
					window.sessionStorage.removeItem(key);
				}
			}

			if (state > 0) {
				reCaptchaImage();
			}
		},
		// eslint-disable-next-line
		[state]
	)


	return <div
		className={utils.Override.Style(
			props,
			"Root",
			{display: src ? "block" : "none", margin: "16px 0"}
		)}
		{...utils.Override.Props(props, "Root", {})}
	>
		<img
			ref={imgRef}
			alt={"captcha"} className={utils.Override.Style(props, "Image",
			{display: "block", margin: "0 auto", cursor: "pointer"}
		)}
			src={src}
			onClick={(event => incr())}
			{...utils.Override.Props(props, "Image", {})}
		/>
	</div>
}