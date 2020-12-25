import {Account} from "./account"
import {BrowserRouter as Router, Link} from "react-router-dom";
import utils from "./utils"
import home from "./home"

const Root = new utils.PathSwitch("", function () {
    return <h1>Not found!</h1>
});

Root.register("/", home.Page);
Root.include(Account);

function App() {
    return (
        <Router>
            {Root.render()}
            {
                [
                    "/account/login",
                    "/account/register",
                    "/account/profile/zzztttkkk",
                    "/"
                ].map(
                    (item, index) => {
                        return <div key={index}>
                            <Link to={item}>{item}</Link>
                        </div>
                    }
                )
            }
        </Router>
    );
}

export default App;
