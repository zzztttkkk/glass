import {PathSwitch} from "./router"
import {Cookie} from "./cookie";
import {glass} from "./glass";
import {Override} from "./overrides";
import {Fragment} from "./fragment";
import {Fetch} from "./fetch";

global.glass = glass;

// eslint-disable-next-line
export default {
	PathSwitch,
	params: function (props, key) {
		let match = props.match;
		if (!match) {
			return null;
		}
		return match.params ? match.params[key] : null
	},
	glass,
	Cookie,
	Override,
	Fragment,
	Fetch
}