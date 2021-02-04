const exec = require("child_process").execSync;
const fs = require("fs");


if (!fs.existsSync("../backend/dist/webbuild")) {
    fs.mkdirSync("../backend/dist/webbuild")
}

exec("react-scripts build", {stdio: "inherit"});

exec("rm -rf ../backend/dist/webbuild/*", {stdio: "inherit"});

exec("cp -r ./build/* ../backend/dist/webbuild/", {stdio: "inherit"});

fs.writeFileSync("../backend/dist/webbuild/web-built-time.txt", Date.now().toString());

fs.writeFileSync("../backend/dist/webbuild/version.txt", exec("git rev-parse HEAD").toString().trim())

console.log("build ok");