import { useState } from "react";
import { AppsIcon, PanelIcon, ReposIcon } from "../lib/icon";
import { NavLink } from "react-router-dom";

export default function Sidebar() {
  const [panelOpen, setPanelOpen] = useState(true);

  return (
    <div
      className={`bg-sidebar animate-in flex flex-col ${
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
      <div className="mt-8 flex flex-col gap-4">
        <Item name="Apps" to="/" icon={<AppsIcon />} panelOpen={panelOpen} />
        <Item
          name="Repos"
          to="/repos"
          icon={<ReposIcon />}
          panelOpen={panelOpen}
        />
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
      className={`hover:bg-sidebarLite rounded px-2 flex gap-2 items-center ${
        panelOpen ? "py-1" : "justify-center py-2"
      }`}
    >
      {icon}
      <span className={`${panelOpen ? "" : "hidden"}`}>{name}</span>
    </NavLink>
  );
}
