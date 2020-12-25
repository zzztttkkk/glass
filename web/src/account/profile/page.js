import React from "react";

export function Profile(props) {
    return <div>
        <h1>Profile of "{props.match.params.name}"</h1>
    </div>
}