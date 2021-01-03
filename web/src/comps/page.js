import React from "react";
import {Header} from "./header";
import utils from "../utils"

export function Page(props) {
    const [user, setUser] = React.useState(null);
    let tryGetUser = async () => {
        let u = await utils.glass.TryGetUser();
        if (u) {
            setUser(u);
        }
    }

    React.useEffect(
        () => {
            tryGetUser();
        },
        []
    )

    return <div id="page">
        <Header user={user} title={props.title}/>
        {props.children}
    </div>
}