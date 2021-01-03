import React from "react";
import {BrowserRouter as Router} from "react-router-dom";
import {Client as Styletron} from "styletron-engine-atomic";
import {Provider as StyletronProvider} from "styletron-react";
import {LightTheme, BaseProvider, LocaleProvider, useStyletron} from "baseui";
import utils from "./utils";


// pages
import {HomePage} from "./home";
import {Account as AccountPages} from "./account";
import {EditorPage} from "./editor";


const Root = new utils.PathSwitch(
    "",
    function () {
        return <h1>NotFound</h1>;
    }
);

Root.register("/", HomePage);
Root.register("/editor", EditorPage);
Root.include(AccountPages);

const styletron = new Styletron();

function makeResponsiveTheme(theme) {
    const breakpoints = {
        small: 769,
        medium: 1024,
        large: 1216,
    };
    const ResponsiveTheme = Object.keys(breakpoints).reduce(
        (acc, key) => {
            acc.mediaQuery[
                key
                ] = `@media screen and (min-width: ${breakpoints[key]}px)`;
            return acc;
        },
        {
            breakpoints,
            mediaQuery: {},
        },
    );
    return {...theme, ...ResponsiveTheme};
}

function GlassOverrideSetup(props) {
    const [css, theme] = useStyletron();
    utils.glass.css = css;
    utils.glass.theme = theme;
    return <>{props.children}</>
}

function Inner() {
    const [theme, setTheme] = React.useState(makeResponsiveTheme(LightTheme));
    utils.glass.setTheme = function (v) {
        setTheme(makeResponsiveTheme(v));
    };
    return <StyletronProvider value={styletron}>
        <BaseProvider theme={theme}>
            <GlassOverrideSetup>
                <Router>{Root.render()}</Router>
            </GlassOverrideSetup>
        </BaseProvider>
    </StyletronProvider>
}

function App() {
    const [locate, setLocate] = React.useState(null);
    utils.glass.setLocate = setLocate;

    if (!locate) {
        return <Inner/>
    }
    return (
        <LocaleProvider locale={locate}>
            <Inner/>
        </LocaleProvider>
    );
}

export default App;
