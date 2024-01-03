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

import getTitle from "../lib/getTitle";
import { useEffect, useState } from "react";
import NewApplication from "./NewApplication";
import AllApplications from "./AllApplications";

export default function Apps() {
  const [openWindow, setOpenWindow] = useState(false);
  // this is used to refetch the app list when new app is
  // created
  const [refresh, setRefresh] = useState(false);

  useEffect(() => {
    document.title = getTitle("Applications");
  }, []);

  return (
    <div className="h-screen p-8 overflow-auto">
      <div className="flex justify-between items-center">
        <p className="text-2xl">Applications</p>
        <button
          onClick={(e) => {
            e.preventDefault();
            setOpenWindow(true);
          }}
          className="bg-white text-black py-2 px-4 rounded font-bold hover:bg-white/90"
        >
          New Application
        </button>
      </div>

      {/* New Application window modal*/}
      <NewApplication
        openWindow={openWindow}
        setOpenWindow={setOpenWindow}
        setRefresh={setRefresh}
      />

      {/* Showing All Applications */}
      <div className="m-4 md:m-8 mt-8 md:mt-16 overflow-auto h-[84%]">
        <AllApplications refresh={refresh} />
      </div>
    </div>
  );
}
