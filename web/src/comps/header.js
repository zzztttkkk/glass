import React from "react";

import {
    HeaderNavigation,
    ALIGN,
    StyledNavigationList,
    StyledNavigationItem
} from "baseui/header-navigation";
import {Button} from "baseui/button";
import {Wrapper} from "./wrapper";
import {Avatar} from "baseui/avatar";
import {useStyletron} from "baseui";
import {StatefulPopover, TRIGGER_TYPE, PLACEMENT} from 'baseui/popover';
import {Link, BtnLink} from "./link";

function User(props) {
    let user = props.user;
    const [css, theme] = useStyletron();
    let btnStyle = {width: "6em", display: "block"};
    let linkStyle = {
        color: theme.colors.background,
        textDecoration: "none",
        ":visited": {
            color: theme.colors.background,
            textDecoration: "none",
        }
    };

    return <StatefulPopover
        triggerType={TRIGGER_TYPE.click}
        placement={PLACEMENT.bottom}
        content={
            <div
                className={
                    css({padding: "8px"})
                }
            >
                <BtnLink
                    href={`/account/profile/${user.name}`}
                    btnStyle={btnStyle}
                    linkStyle={linkStyle}
                >Profile</BtnLink>

                <BtnLink
                    href={`/account/settings`}
                    btnStyle={
                        {
                            ...btnStyle,
                            ...{marginTop: "8px",}
                        }
                    }
                    linkStyle={linkStyle}
                >Settings</BtnLink>

                <BtnLink
                    href={"/account/logout"}
                    btnStyle={
                        {
                            ...btnStyle,
                            ...{
                                marginTop: "8px",
                                backgroundColor: theme.colors.negative,
                                ":hover": {
                                    backgroundColor: theme.colors.negative500
                                }
                            }
                        }
                    }
                    linkStyle={linkStyle}
                >Logout</BtnLink>
            </div>
        }
    >
        <div
            className={css({cursor: "pointer"})}
        >
            <Avatar name={user.name} src={user.avatar} size={"scale1200"}/>
        </div>
    </StatefulPopover>
}

export function Header(props) {
    React.useEffect(
        () => {
            let v = props.title || "";
            if (v.length < 1) {
                return;
            }
            if (v.startsWith(" ")) {
                document.title = document.title + v;
            } else {
                document.title = v;
            }
        },
        [props.title]
    )

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
                <StyledNavigationItem>
                    <Link
                        href={"/"}
                        overrides={{Root: {style: {textDecoration: "none"}}}}
                    >
                        <h1>Glass</h1>
                    </Link>
                </StyledNavigationItem>
            </StyledNavigationList>
            <StyledNavigationList $align={ALIGN.center}/>
            <StyledNavigationList $align={ALIGN.right}>
                <StyledNavigationItem><Link href={"/link1"}>Link1</Link></StyledNavigationItem>
                <StyledNavigationItem><Link href={"/editor"}>Editor</Link></StyledNavigationItem>
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