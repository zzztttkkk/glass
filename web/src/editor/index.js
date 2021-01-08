import React from "react";
import {Page} from "../comps/page";
import {Wrapper} from "../comps/wrapper";
import {glass} from "../utils/glass";
import "./editor.css"
import {Input} from "baseui/input";
import {useStyletron} from "baseui";
import {Button} from "baseui/button";
import utils from "../utils"

function newQuill(Q, setContentSize) {
	glass.Editor.quill = new Q(
		"#editor",
		{
			theme: "snow",
			modules: {
				toolbar: [
					[{'font': []}, {'size': []}],
					['bold', 'italic', 'underline', 'strike'],
					[{'color': []}, {'background': []}],
					[{'script': 'super'}, {'script': 'sub'}],
					['link', 'blockquote', 'code-block', 'formula'],
					[{'list': 'ordered'}, {'list': 'bullet'}, {'indent': '-1'}, {'indent': '+1'}],
					[{'align': []}],
					['clean'],
				],
				history: {
					delay: 2000,
					maxStack: 500,
					userOnly: false
				},
			}
		},
	);


	glass.Editor.load();
	setContentSize(glass.Editor.getText().length);
	glass.Editor.on(
		"text-change", function () {
			setContentSize(glass.Editor.getText().length);
		}
	);

	let interval = window.setInterval(
		() => {
			let toolbarContainer = document.querySelector(".ql-toolbar");
			if (!toolbarContainer) return;
			window.clearInterval(interval);

			new utils.Fragment().appendElement(
				// undo
				function () {
					return document.createElement("button")
				}(),
				// redo
				function () {
					return document.createElement("button")
				}(),
				// save
				function () {
					return document.createElement("button")
				}(),
				// close input method
				function () {
					return document.createElement("button")
				}(),
			);
		},
		20
	);
}


function Editor() {
	const [css] = useStyletron();
	const [title, _setTitle] = React.useState("");
	const [contentSize, setContentSize] = React.useState(0);

	function setTitle(v) {
		_setTitle(v);
		glass.Editor.__title = v;
	}

	glass.Editor.setTitle = setTitle;

	React.useEffect(
		() => {
			window.onload = function () {
				new utils.Fragment().appendCSS(
					"https://cdn.jsdelivr.net/npm/quill@1.3.7/dist/quill.snow.css"
				).appendCSS(
					"https://cdn.jsdelivr.net/npm/katex@0.12.0/dist/katex.min.css"
				).appendJavaScript(
					"https://cdn.jsdelivr.net/npm/katex@0.12.0/dist/katex.min.js"
				).appendJavaScript(
					"https://cdn.jsdelivr.net/npm/quill@1.3.7/dist/quill.min.js"
				).insertInto(document.head);

				let interval = window.setInterval(
					() => {
						let Q = window["Quill"];
						if (!Q) return;
						window.clearInterval(interval);
						newQuill(Q, setContentSize);
					},
					20,
				);

				window.addEventListener(
					"keydown",
					function (evt) {
						if (evt.key === "s" && evt.ctrlKey) {
							evt.preventDefault();
							glass.Editor.save();
						}
					}
				)
			};

		},
		[]
	);

	const [showSubmitDialog, setShowSubmitDialog] = React.useState(false);

	function submit() {
		setShowSubmitDialog(true);
	}

	return <>
		<Wrapper overrides={{Root: {style: {marginBottom: "32px"}}}}>
			<div className={css({marginBottom: "16px", display: "flex"})}>
				<Input overrides={{Input: {style: {fontSize: "2em"}}}}
					   placeholder={"Title"}
					   value={title}
					   onChange={(event => setTitle(event.currentTarget.value))}
				/>
				<Button disabled={contentSize < 200 || title.length < 1}
						overrides={{BaseButton: {style: {marginLeft: "16px"}}}}
						onClick={submit}
				>Submit</Button>
			</div>
			<div id={"editor"}/>
		</Wrapper>
		{
			showSubmitDialog &&
			<div>Dialog</div>
		}
	</>
}

export function EditorPage() {
	return <Page title={" Editor"}>
		<Editor/>
	</Page>
}