let flags = {
    "glass/internal.BuiltTime": Date.now().toString(),
};

function makeFlags() {
    return Object.entries(flags).reduce((rv, [k, v]) => `${rv} ${k}=${v}`.trim(), "")
}

module.exports = {
    args: `-ldflags "-X '${makeFlags()}'"`
}
