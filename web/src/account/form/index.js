import React from "react";
import utils from "../../utils"
import {FormControl} from "baseui/form-control";
import {Input} from "baseui/input";

function AccountInput(props) {
    let name = props.name;

    switch (name) {
        case "name": {
            return <FormControl label={props.label}>
                <Input/>
            </FormControl>
        }
        case "password": {
            return <FormControl>
                <Input/>
            </FormControl>
        }
        case "captcha": {
            return <FormControl>
                <Input/>
            </FormControl>
        }
        default : {
            return <h1>A</h1>
        }
    }
}

export function AccountForm(props) {
    return <div
        className={utils.Override.Style(props, "Root", {})}
        {...utils.Override.Props(props, "Root", {})}
    >
        <form {...utils.Override.Props(props, "Form", {})}>
            {
                props.inputs.map(
                    (item, index) => {
                        return <AccountInput {...item} key={index}/>
                    }
                )
            }
        </form>
    </div>
}