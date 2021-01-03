import utils from "../utils";
import {RegisterPage} from "./register";
import {LoginPage} from "./login"
import {ProfilePage} from "./profile"

export const Account = new utils.PathSwitch(
    "/account",
    function () {
        return <h1>account not found</h1>
    }
);

Account.register("/register", RegisterPage);

Account.register("/login", LoginPage);

Account.register("/profile/:name", ProfilePage)

