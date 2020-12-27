import React from "react";
import utils from "../utils"

export function Wrapper(props) {
    return <div
        className={utils.Override.Style(
            props, "Root", {width: "100%", position: "relative"}
        )}
        {...utils.Override.Props(props, "Root", null)}
    >
        <div
            className={
                utils.Override.Style(
                    props, "Content", {maxWidth: "1600px", margin: "0 auto"}
                )
            }
            {...utils.Override.Props(props, "Content", null)}
        >
            {props.children}
        </div>
    </div>
}