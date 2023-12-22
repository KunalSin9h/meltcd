import { useState } from "react";
import { PanelIcon } from "../lib/icon";

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
    </div>
  );
}
