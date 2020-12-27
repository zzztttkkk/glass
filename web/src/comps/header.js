import React from "react";

import {
    HeaderNavigation,
    ALIGN,
    StyledNavigationList,
    StyledNavigationItem
} from "baseui/header-navigation";
import {Button} from "baseui/button";
import {Wrapper} from "./wrapper";

import {Link} from "./link";

function User(props) {
    return <div>U</div>
}

export function Header(props) {
    return <Wrapper
        overrides={{
            Root: {
                style: function (theme) {
                    return {
                        borderBottomStyle: "solid",
                        borderBottomWidth: "1px",
                        borderBottomColor: theme.colors.border,
                        marginBottom: "18px",
                    }
                }
            }
        }}
    >
        <HeaderNavigation overrides={{Root: {style: {borderBottomWidth: 0,}}}}>
            <StyledNavigationList $align={ALIGN.left}>
                <StyledNavigationItem><Link href={"/"}>Glass</Link></StyledNavigationItem>
            </StyledNavigationList>
            <StyledNavigationList $align={ALIGN.center}/>
            <StyledNavigationList $align={ALIGN.right}>
                <StyledNavigationItem><Link href={"/link1"}>Link1</Link></StyledNavigationItem>
                <StyledNavigationItem><Link href={"/link2"}>Link2</Link></StyledNavigationItem>
            </StyledNavigationList>
            <StyledNavigationList $align={ALIGN.right}>
                <StyledNavigationItem>
                    {
                        props.user
                            ?
                            <User user={props.user}/>
                            :
                            <Link href={`/account/login?ref=${window.location.pathname}`}><Button>Login</Button></Link>
                    }
                </StyledNavigationItem>
            </StyledNavigationList>
        </HeaderNavigation>
    </Wrapper>
}