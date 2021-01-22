import utils from "../utils";
import {RegisterPage} from "./register";
import {LoginPage} from "./login"
import {ProfilePage} from "./profile"
import {ErrNotFoundPage} from "../error";

export const Account = new utils.PathSwitch("/account", ErrNotFoundPage);

Account.register("/register", RegisterPage);

Account.register("/login", LoginPage);

Account.register("/profile/:name", ProfilePage)

