import utils from "../utils";
import register from "./register";
import login from "./login"
import profile from "./profile"

export const Account = new utils.PathSwitch("/account", function () {
    return <h1>account not found</h1>
});

Account.register("/register", register.Page);

Account.register("/login", login.Page);

Account.register("/profile/:name", profile.Page)

