/*
Copyright 2023 - PRESENT Meltred

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

import { Outlet } from "react-router-dom";
import Sidebar from "./Sidebar";
import { Toaster } from "react-hot-toast";
import { useState } from "react";
import { LinkIcon } from "../lib/icon";

export default function Layout() {
  const [openHelpPanel, setOpenHelpPanel] = useState(false);

  return (
    <div className="flex flex-row h-screen w-screen overflow-hidden">
      <Sidebar
        openHelpPanel={openHelpPanel}
        setOpenHelpPanel={setOpenHelpPanel}
      />
      <div className="flex-1 relative">
        <Outlet />
        {/** Help And Support Panel relative to main window */}
        <div
          className={`absolute h-auto rounded bg-sidebar w-48 left-4 bottom-4 p-4 flex flex-col gap-2
            ${openHelpPanel ? "" : "hidden"}
          `}
        >
          <Linker name="Documentation" url="https://cd.meltred.tech/docs" />
          <Linker
            name="File bug or issue"
            url="https://github.com/meltred/meltcd/issues"
          />
          <Linker name="GitHub" url="https://github.com/meltred/meltcd" />
          <Linker name="Discord" url="https://discord.gg/Y2C6mEhhf3" />
          <Linker name="Twitter" url="https://twitter.com/meltredhq" />
        </div>
      </div>
      <Toaster />
    </div>
  );
}

function Linker({ name, url }: { name: string; url: string }) {
  return (
    <a
      href={url}
      target="_blank"
      className="flex items-center gap-2 hover:text-white/70"
    >
      <p>{name}</p>
      <LinkIcon />
    </a>
  );
}
