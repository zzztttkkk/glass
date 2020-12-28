import React from "react";
import {Page} from "../../comps/page";
import {Wrapper} from "../../comps/wrapper";
import {FormControl} from "baseui/form-control";
import {Input} from "baseui/input";

function FormGroup(props) {
    const [username, setUsername] = React.useState("");
    const [usernameInErr, setUsernameInErr] = React.useState(false);

    return <div>
        <FormControl>
            <Input
                clearable clearOnEscape required startEnhancer={"Username"} type={"text"} value={username}
                error={usernameInErr}
                onChange={(evt) => {
                    evt.stopPropagation();
                    setUsernameInErr(false);
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
            <Input clearable startEnhancer={"Password"} type={"password"}/>
        </FormControl>
    </div>
}

export function Login() {
    return <Page>
        <Wrapper>
            <FormGroup/>
        </Wrapper>
    </Page>
}