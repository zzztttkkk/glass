const exec = require("child_process").execSync;

exec(`go run ${require("./common").args} ./cmd/main -c ./conf`, {stdio: "inherit"});