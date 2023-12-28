import Tippy from "@tippyjs/react";

export default function Tooltip(props: {
  content: string;
  children: JSX.Element;
}) {
  return (
    <Tippy
      content={props.content}
      className="p-2 rounded mx-2 bg-sidebarLite"
      interactive={true}
      delay={0}
      duration={0}
    >
      {props.children}
    </Tippy>
  );
}
