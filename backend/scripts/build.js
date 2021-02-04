const exec = require("child_process").execSync;

exec(`go build ${(require("./common").args)} ./cmd/main`, {stdio: "inherit"});
