import { useEffect } from "react";
import getTitle from "./lib/getTitle";

export default function Repos() {
  useEffect(() => {
    document.title = getTitle("Repositories");
  }, []);

  return (
    <div className="h-screen p-8">
      <div className="flex justify-between items-center">
        <p className="text-2xl">Repositories</p>
        <button
          onClick={(e) => {
            e.preventDefault();
          }}
          className="bg-white text-black py-2 px-4 rounded font-bold border-dashed hover:bg-inherit hover:text-white border-2 border-white transition ease-in-out delay-50 hover:-translate-y-1 duration-100"
        >
          New Repository
        </button>
      </div>
    </div>
  );
}
