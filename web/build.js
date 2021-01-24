const exec = require("child_process").execSync;
const fs = require("fs");

function onErr(err) {
    if (err !== null) {
        throw err;
    }
}


exec("react-scripts build", onErr);

exec("rm -rf ../backend/dist/webbuild/*", onErr);

exec("cp -r ./build/* ../backend/dist/webbuild/");

fs.writeFileSync("../backend/dist/webbuild/time.txt", Date.now().toString(), onErr);

console.log("build ok");