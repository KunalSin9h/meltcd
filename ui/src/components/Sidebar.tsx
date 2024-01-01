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

import { useEffect, useState } from "react";
import {
  AppsIcon,
  HelpIcon,
  LogIcon,
  PanelIcon,
  ReposIcon,
  SecretIcon,
  SettingIcon,
  UserIcon,
} from "../lib/icon";
import { NavLink } from "react-router-dom";

export default function Sidebar({
  openHelpPanel,
  setOpenHelpPanel,
}: {
  openHelpPanel: boolean;
  setOpenHelpPanel: React.Dispatch<React.SetStateAction<boolean>>;
}) {
  const [panelOpen, setPanelOpen] = useState(true);
  const [user, setUser] = useState("Admin");

  useEffect(() => {
    setUser(localStorage.getItem("username") as string);
  }, [user]);

  window.onresize = () => {
    const width = window.outerWidth;
    if (width <= 1000) {
      setPanelOpen(false);
    }
  };

  return (
    <div
      className={`bg-sidebar relative flex flex-col ${
        panelOpen ? "w-64" : "w-18"
      } p-4`}
    >
      <div
        className={`flex ${
          panelOpen ? "" : "flex-col"
        } justify-between items-center`}
      >
        <div className="flex items-center gap-2 px-1 py-3">
          <img src="/logo.png" className="h-8 w-8" alt="Meltcd Logo" />
          <p className={`font-bold text-lg ${panelOpen ? "" : "hidden"} `}>
            Meltcd
          </p>
        </div>
        <span
          onClick={(e) => {
            e.preventDefault();
            setPanelOpen(!panelOpen);
          }}
          className={`${
            panelOpen ? "" : "rotate-180 mt-4"
          } hover:bg-sidebarLite rounded p-2 cursor-pointer`}
        >
          <PanelIcon />
        </span>
      </div>
      <div className="mt-8 flex flex-col gap-4 flex-1">
        <Item
          name="Apps"
          to="/apps"
          icon={<AppsIcon />}
          panelOpen={panelOpen}
        />
        <Item
          name="Repos"
          to="/repos"
          icon={<ReposIcon />}
          panelOpen={panelOpen}
        />
        <Item
          name="Secrets"
          to="/secrets"
          icon={<SecretIcon />}
          panelOpen={panelOpen}
        />
        <Item name="Logs" to="/logs" icon={<LogIcon />} panelOpen={panelOpen} />
      </div>
      <div className="mb-2 flex flex-col gap-4">
        <Item
          name={user}
          to={`/user`}
          icon={<UserIcon />}
          panelOpen={panelOpen}
        />
        <Item
          name="Settings"
          to="/settings"
          icon={<SettingIcon />}
          panelOpen={panelOpen}
        />
        <div
          className={`hover:bg-sidebarLite hover:border-l hover:border-l-[5px] hover:border-white/40 rounded-r px-2 flex gap-2 items-center cursor-pointer ${
            panelOpen ? "py-1" : "justify-center py-2"
          }
          ${openHelpPanel ? "bg-sidebarLite rounded" : ""}
          `}
          onClick={(e) => {
            e.preventDefault();
            setOpenHelpPanel(!openHelpPanel);
          }}
        >
          <HelpIcon />
          <span className={`text-lg ${panelOpen ? "" : "hidden"}`}>
            Help & Support
          </span>
        </div>
      </div>
    </div>
  );
}

function Item({
  name,
  to,
  icon,
  panelOpen,
}: {
  name: string;
  to: string;
  icon: JSX.Element;
  panelOpen: boolean;
}) {
  return (
    <NavLink
      to={to}
      className={`hover:bg-sidebarLite hover:border-l hover:border-l-[5px] hover:border-white/40 rounded-r px-2 flex gap-2 items-center ${
        panelOpen ? "py-1" : "justify-center py-2"
      }`}
    >
      {icon}
      <span className={`text-lg ${panelOpen ? "" : "hidden"}`}>{name}</span>
    </NavLink>
  );
}
