import React from "react";

import {
    HeaderNavigation,
    ALIGN,
    StyledNavigationList,
    StyledNavigationItem
} from "baseui/header-navigation";
import {Button} from "baseui/button";

import {Link} from "./link";

export function Header(props) {
    return <HeaderNavigation>
        <StyledNavigationList $align={ALIGN.left}>
            <StyledNavigationItem><Link href={"/"}>Glass</Link></StyledNavigationItem>
        </StyledNavigationList>
        <StyledNavigationList $align={ALIGN.center}/>
        <StyledNavigationList $align={ALIGN.right}>
            <StyledNavigationItem><Link href={"/link1"}>Link1</Link></StyledNavigationItem>
            <StyledNavigationItem><Link href={"/link2"}>Link2</Link></StyledNavigationItem>
        </StyledNavigationList>
        <StyledNavigationList $align={ALIGN.right}>
            <StyledNavigationItem><Button>Login</Button></StyledNavigationItem>
        </StyledNavigationList>
    </HeaderNavigation>
}