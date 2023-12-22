import { useEffect } from "react";
import { Outlet } from "react-router-dom";
import Sidebar from "./Sidebar";
import getTitle from "../lib/getTitle";

export default function Layout() {
  useEffect(() => {
    document.title = getTitle("Dashboard");
  }, []);

  return (
    <div className="flex flex-row h-screen w-screen overflow-hidden">
      <Sidebar />
      <div className="flex-1 relative">
        <Outlet />
      </div>
    </div>
  );
}
