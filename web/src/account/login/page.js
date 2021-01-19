import React from "react";
import {Page} from "../../comps/page";
import {Wrapper} from "../../comps/wrapper";
import {FormControl} from "baseui/form-control";
import {Input} from "baseui/input";
import {useStyletron} from "baseui";
import {BtnLink} from "../../comps/link";
import {Button} from "baseui/button";
import utils from "../../utils"
import {Captcha} from "../../comps/captcha";


function FormGroup(props) {
	const [username, setUsername] = React.useState("");
	const [usernameErr, setUsernameErr] = React.useState("");
	const [password, setPassword] = React.useState("");
	const [passwordErr, setPasswordErr] = React.useState("");
	const [captcha, setCaptcha] = React.useState("");
	const [captchaErr, setCaptchaErr] = React.useState("");
	const [obj] = React.useState({incr: (v => void 0)});

	const [css, theme] = useStyletron();
	return <div
		className={css(
			{
				width: "80%",
				margin: "0 auto",
				[theme.mediaQuery.large]: {maxWidth: "64em", width: "50%", margin: "0 auto"},
			}
		)}
	>
		<FormControl>
			<Input
				clearable clearOnEscape required
				overrides={{StartEnhancer: {style: {width: "4em"}}}}
				startEnhancer={utils.glass.locate.glass.account.username}
				type={"text"}
				value={username}
				error={usernameErr}
				onChange={(evt) => {
					evt.stopPropagation();
					setUsernameErr(false);
					setUsername(evt.currentTarget.value);
				}}
				onBlur={
					(evt) => {
						evt.stopPropagation();

					}
				}
			/>
		</FormControl>
		<FormControl>
			<Input clearable value={password} onChange={(event => setPassword(event.currentTarget.value))}
				   overrides={{StartEnhancer: {style: {width: "4em"}}}}
				   startEnhancer={utils.glass.locate.glass.account.password}
				   type={"password"}/>
		</FormControl>

		<FormControl>
			<Input clearable
				   onFocus={(event => obj.incr())}
				   overrides={{StartEnhancer: {style: {width: "4em"}}}}
				   startEnhancer={utils.glass.locate.glass.common.captcha}
				   type={"text"}/>
		</FormControl>
		<Captcha obj={obj} where={"account.login"}/>
		<div>
			<BtnLink href={"/account/register"} BTN={{BaseButton: {style: {marginRight: "16px"}}}}>
				{utils.glass.locate.glass.account.register}
			</BtnLink>
			<BtnLink href={"/account/repwd"}>
				{utils.glass.locate.glass.account.fgpwd}
			</BtnLink>
			<Button
				overrides={{BaseButton: {style: {float: "right"}}}}>{utils.glass.locate.glass.common.submit}</Button>
		</div>
	</div>
}

export function Login() {
	return <Page>
		<Wrapper>
			<FormGroup/>
		</Wrapper>
	</Page>
}