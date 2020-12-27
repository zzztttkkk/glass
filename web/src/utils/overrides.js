import {glass} from "./glass";

function getOverrideData(props, position, type) {
    let od = props["overrides"];
    if (!od) return null;
    let pod = od[position];
    if (!pod) return null;
    return pod[type];
}

export class Override {
    static Style(props, position, dist) {
        let style = getOverrideData(props, position, "style");
        if (!style) {
            return glass.css(dist);
        }
        if (typeof style === "function") {
            return glass.css({...dist, ...style(glass.theme)});
        }
        return glass.css({...dist, ...style});
    }

    static Props(props, position, dist) {
        let pprops = getOverrideData(props, position, "props");
        if (!pprops) {
            return dist;
        }
        return {...dist, ...pprops};
    }
}