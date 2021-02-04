import React from "react";

import {
    HeaderNavigation,
    ALIGN,
    StyledNavigationList,
    StyledNavigationItem
} from "baseui/header-navigation";
import {Wrapper} from "./wrapper";
import {Avatar} from "baseui/avatar";
import {useStyletron} from "baseui";
import {StatefulPopover, TRIGGER_TYPE, PLACEMENT} from 'baseui/popover';
import {Link, BtnLink} from "./link";
import utils from "../utils";

function User(props) {
    let user = props.user;
    const [css, theme] = useStyletron();
    let btnStyle = {
        display: "block",
        width: "6em",
    };

    return <StatefulPopover
        triggerType={TRIGGER_TYPE.click}
        placement={PLACEMENT.bottom}
        content={
            <div className={css({padding: "8px"})}>
                <BtnLink
                    href={`/account/profile/${user.name}`}
                    BTN={{BaseButton: {style: {...btnStyle}}}}>
                    Profile
                </BtnLink>
                <BtnLink
                    href={`/account/settings`}
                    BTN={{BaseButton: {style: {...btnStyle, ...{marginTop: "8px", marginBottom: "8px"}}}}}>
                    Settings
                </BtnLink>
                <BtnLink
                    href={"/account/logout"}
                    BTN={{
                        BaseButton: {
                            style: {
                                ...btnStyle,
                                ...{
                                    backgroundColor: theme.colors.negative,
                                    ":hover": {
                                        backgroundColor: theme.colors.negative500
                                    }
                                }
                            }
                        }
                    }}>
                    Logout
                </BtnLink>
            </div>
        }
    >
        <div className={css({cursor: "pointer"})}>
            <Avatar name={user.name} src={user.avatar} size={"scale1200"}/>
        </div>
    </StatefulPopover>
}

export function Header(props) {
    const [css] = useStyletron();
    const [title, setTitle] = React.useState(props.title || "");
    utils.glass.setTitle = setTitle;
    const user = utils.glass.useUser();

    React.useEffect(
        () => {
            if (title.startsWith(" ")) {
                document.title = document.title + title;
            } else {
                document.title = title;
            }
        },
        [title]
    )

    return <Wrapper
        overrides={{
            Root: {
                style: function (theme) {
                    return {
                        borderBottomStyle: "solid",
                        borderBottomWidth: "1px",
                        borderBottomColor: theme.colors.border,
                        marginBottom: "16px"
                    }
                }
            },
            Content: {
                style: {
                    boxSizing: "border-box",
                    padding: "0 1em"
                }
            }
        }}
    >
        <HeaderNavigation overrides={{Root: {style: {borderBottomWidth: 0,}}}}>
            <StyledNavigationList $align={ALIGN.left}>
                <StyledNavigationItem className={css({paddingLeft: "0!important"})}>
                    <Link href={"/"} overrides={{Root: {style: {textDecoration: "none"}}}}>
                        <h1>Glass</h1>
                    </Link>
                </StyledNavigationItem>
                {/*<StyledNavigationItem><Link href={"/link1"}>Link1</Link></StyledNavigationItem>*/}
                <StyledNavigationItem><Link href={"/editor"}>Editor</Link></StyledNavigationItem>
            </StyledNavigationList>

            <StyledNavigationList $align={ALIGN.center}/>

            <StyledNavigationList $align={ALIGN.right}>
                <StyledNavigationItem>
                    {
                        user
                            ?
                            <User user={user}/>
                            :
                            (
                                window.location.pathname === "/account/login"
                                    ?
                                    <BtnLink href={"/account/register"}>
                                        {utils.glass.localization.glass.account.register}
                                    </BtnLink>
                                    :
                                    (
                                        <BtnLink href={`/account/login?ref=${window.location.pathname}`}>
                                            {utils.glass.localization.glass.account.login}
                                        </BtnLink>
                                    )
                            )
                    }
                </StyledNavigationItem>
            </StyledNavigationList>
        </HeaderNavigation>
    </Wrapper>
}