import React from "react";
import {Page} from "../comps/page";
import {Wrapper} from "../comps/wrapper";
import {glass} from "../utils/glass";
import "./editor.css"
import {Input} from "baseui/input";
import {useStyletron} from "baseui";
import {Button} from "baseui/button";


function Editor() {
    const [css] = useStyletron();
    const [title, setTitle] = React.useState("");
    const [contentSize, setContentSize] = React.useState(0);

    React.useEffect(
        () => {
            glass.Editor.getTitle = function () {
                return title;
            }
        },
        [title]
    )

    React.useEffect(
        () => {
            window.onload = function () {
                let link = document.createElement("link");
                link.setAttribute("href", "https://cdn.jsdelivr.net/npm/quill@1.3.7/dist/quill.snow.css");
                link.setAttribute("rel", "stylesheet");

                let quill = document.createElement("script");
                quill.setAttribute("src", "https://cdn.jsdelivr.net/npm/quill@1.3.7/dist/quill.min.js");

                document.head.appendChild(link);
                document.head.appendChild(quill);

                let interval = window.setInterval(
                    () => {
                        let Q = window["Quill"];
                        if (!Q) {
                            return;
                        }
                        window.clearInterval(interval);
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
                                        ['link', 'blockquote', 'code-block'],
                                        [{'list': 'ordered'}, {'list': 'bullet'}, {'indent': '-1'}, {'indent': '+1'}],
                                        [{'align': []}],
                                        ['clean']
                                    ]
                                }
                            },
                        );
                        glass.Editor.setTitle = function (t) {
                            setTitle(t);
                        }

                        glass.Editor.load();
                        setContentSize(glass.Editor.getText().length);
                        glass.Editor.on(
                            "text-change", function () {
                                setContentSize(glass.Editor.getText().length);
                            }
                        );
                    },
                    100,
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
        <Wrapper>
            <div className={css({marginBottom: "16px", display: "flex"})}>
                <Input overrides={{Input: {style: {fontSize: "2em"}}}}
                       placeholder={"Title"}
                       value={title}
                       onChange={event => {
                           event.stopPropagation();
                           setTitle(event.currentTarget.value);
                       }}
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