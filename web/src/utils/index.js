// @flow

import {PathSwitch} from "./router"

// eslint-disable-next-line
export default {
    PathSwitch,
    params: function (props: any, key: string): ?string {
        let match = props.match;
        if (!match) {
            return null;
        }
        return match.params ? match.params[key] : null
    }
}