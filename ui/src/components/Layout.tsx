import { Outlet } from "react-router-dom";
import Sidebar from "./Sidebar";

export default function Layout() {
  return (
    <div className="flex flex-row h-screen w-screen overflow-hidden">
      <Sidebar />
      <div className="flex-1 relative">
        <div className="flex flex-col justify-between h-full">
          {/* <Header /> */}
          {/* <ScrollArea className="flex-1 bg-white rounded p-2 xl:p-4 mt-1 xl:mt-2 ml-2 xl:ml-4 mb-2 xl:mb-4 mr-2 xl:mr-4"> */}
          <Outlet />
          {/* </ScrollArea> */}
          {/* <Footer /> */}
        </div>
      </div>
      {/* <Toaster /> */}
    </div>
  );
}
