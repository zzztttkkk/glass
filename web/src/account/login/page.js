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
	const [password, setPassword] = React.useState("");
	const [captcha, setCaptcha] = React.useState("");
	const [obj] = React.useState({incr: (v => void 0)});

	async function submit(evt) {
		evt.preventDefault();
		let res = await utils.Fetch.GET("api/account/login", evt.currentTarget);
		if (res.status !== 200) {
			utils.glass.toaster.onFetchError(res.status);
			return
		}
		console.log(res)
	}

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
		<form onSubmit={submit}>
			<FormControl>
				<Input required name={"name"}
					   overrides={{StartEnhancer: {style: {width: "4em"}}}}
					   startEnhancer={utils.glass.localization.glass.account.username}
					   type={"text"} value={username}
					   onChange={(evt) => setUsername(evt.currentTarget.value)}
				/>
			</FormControl>
			<FormControl>
				<Input clearable required name={"password"}
					   value={password} onChange={(event => setPassword(event.currentTarget.value))}
					   overrides={{StartEnhancer: {style: {width: "4em"}}}}
					   startEnhancer={utils.glass.localization.glass.account.password}
					   type={"password"}/>
			</FormControl>

			<FormControl>
				<Input clearable
					   onFocus={(event => obj.incr())}
					   value={captcha} onChange={(event => setCaptcha(event.currentTarget.value))}
					   overrides={{StartEnhancer: {style: {width: "4em"}}}}
					   startEnhancer={utils.glass.localization.glass.common.captcha}
					   type={"text"}/>
			</FormControl>
			<Captcha obj={obj} where={"account.login"}/>
			<div>
				<BtnLink href={"/account/register"} BTN={{BaseButton: {style: {marginRight: "16px"}}}}>
					{utils.glass.localization.glass.account.register}
				</BtnLink>
				<BtnLink href={"/account/repwd"}>
					{utils.glass.localization.glass.account.fgpwd}
				</BtnLink>
				<Button type={"submit"}
						overrides={{BaseButton: {style: {float: "right"}}}}>{utils.glass.localization.glass.common.submit}
				</Button>
			</div>
		</form>
	</div>
}

export function Login() {
	return <Page>
		<Wrapper>
			<FormGroup/>
		</Wrapper>
	</Page>
}