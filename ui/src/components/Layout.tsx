/*
Copyright 2023 - PRESENT kunalsin9h

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

import { Outlet, useNavigate, useLocation } from "react-router-dom";
import Sidebar from "./Sidebar";
import toast, { Toaster } from "react-hot-toast";
import { createContext, useEffect, useState } from "react";
import { LinkIcon } from "../lib/icon";
import version from "../version";
import MeltcdBranding from "./Branding";

type ReactSetState<T> = React.Dispatch<React.SetStateAction<T>>;

type authContext = {
  username: string;
  setUsername: ReactSetState<string | null>;
};

export const Ctx = createContext<authContext | null>(null);

export default function Layout() {
  const [openHelpPanel, setOpenHelpPanel] = useState(false);
  const [username, setUsername] = useState<string | null>("admin");
  const navigate = useNavigate();
  const location = useLocation();

  // check login here
  // and if not authorized then redirect to /login
  useEffect(() => {
    const getUser = async () => {
      try {
        const res = await fetch("/api/users/current");

        if (res.status === 401) {
          navigate("/login");
        } else if (res.status === 200) {
          const username = await res.text();

          // username will be the 404 page of react-router-don in local vite dev environment
          // so check if the username is valid and return from the function
          if (username.split(" ").length != 1) {
            return;
          }

          setUsername(username);

          if (location.pathname === "/") {
            navigate("/apps");
          }
        } else {
          toast.error("Something wend wrong, server does not respond with 200");
        }
      } catch (err) {
        toast.error("Something wend wrong, try again!");
      }
    };

    getUser();
  });

  if (!username) {
    return (
      <div className="h-screen w-screen fixed top-0 left-0 bg-inherit flex justify-center items-center">
        <MeltcdBranding />
      </div>
    );
  }

  return (
    <div className="flex flex-row h-screen w-screen overflow-hidden">
      <Ctx.Provider
        value={{
          username,
          setUsername,
        }}
      >
        <Sidebar
          openHelpPanel={openHelpPanel}
          setOpenHelpPanel={setOpenHelpPanel}
        />
      </Ctx.Provider>
      <div className="flex-1 relative">
        <Ctx.Provider
          value={{
            username,
            setUsername,
          }}
        >
          <Outlet />
        </Ctx.Provider>

        {/** Help And Support Panel relative to main window */}
        <div
          className={`absolute h-auto rounded bg-sidebar w-48 left-4 bottom-4 p-4 flex flex-col gap-2
            ${openHelpPanel ? "" : "hidden"}
          `}
        >
          <Linker name="Documentation" url="https://cd.kunalsin9h.com/docs" />
          <Linker
            name="File bug or issue"
            url="https://github.com/kunalsin9h/meltcd/issues"
          />
          <Linker name="GitHub" url="https://github.com/kunalsin9h/meltcd" />
          <Linker name="Discord" url="https://discord.gg/Y2C6mEhhf3" />
          <Linker name="Twitter" url="https://twitter.com/kunalsin9hhq" />
          <div className="text-center border-t border-t-white/20 pt-2">
            v<span className="font-bold">{version}</span>
          </div>
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
