import React from "react";
import {Account} from "./account";
import {BrowserRouter as Router} from "react-router-dom";
import utils from "./utils";
import {HomePage} from "./home";
import {Client as Styletron} from "styletron-engine-atomic";
import {Provider as StyletronProvider} from "styletron-react";
import {LightTheme, BaseProvider, LocaleProvider} from "baseui";
import comps from "./comps"

const Root = new utils.PathSwitch("", function () {
    return <h1>NotFound</h1>;
});

Root.register("/", HomePage);
Root.include(Account);

const engine = new Styletron();

function Inner() {
    const [theme, setTheme] = React.useState(LightTheme);
    utils.glass.setTheme = setTheme;

    return <StyletronProvider value={engine}>
        <BaseProvider theme={theme}>
            <Router>
                {Root.render()}
                {
                    [
                        "/account/login",
                        "/account/register",
                        "/account/profile/zzztttkkk",
                        "/"
                    ].map((item, index) => {
                        let Link = comps.Link;
                        return (
                            <div key={index}>
                                <Link href={item}>{item === "/" ? "Home" : item}</Link>
                            </div>
                        );
                    })
                }
            </Router>
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
