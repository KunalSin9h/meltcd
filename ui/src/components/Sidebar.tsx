import { useEffect } from "react";
import getTitle from "../lib/getTitle";

export default function Sidebar() {
  useEffect(() => {
    document.title = getTitle("Dashboard");
  }, []);

  return (
    <div className="bg-sidebar flex flex-col w-64 p-4">
      <div className="flex items-center gap-2 px-1 py-3">
        <img src="/logo.png" className="h-8 w-8" alt="Meltcd Logo" />
        <p className="font-bold text-lg">Meltcd</p>
      </div>
    </div>
  );
}
