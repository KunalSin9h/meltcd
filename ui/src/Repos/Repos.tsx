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
import setTitle from "../lib/setTitle";
import { CloseIcon } from "../lib/icon";
import NewRepository from "./NewRepo";
import AllRepos from "./AllRepos";

export default function Repos() {
  const [newRepoOpen, setNewRepoOpen] = useState(false);
  const [refreshSignal, setRefreshSignal] = useState(false);

  useEffect(() => {
    document.title = setTitle("Repositories");
  }, []);

  return (
    <div className="h-screen p-8">
      <div className="flex justify-between items-center">
        <p className="text-2xl">Repositories</p>
        <div className="relative">
          <button
            onClick={() => {
              setNewRepoOpen(!newRepoOpen);
            }}
            className={`bg-white text-black py-2 rounded font-bold hover:bg-white/90
              ${newRepoOpen ? "px-2" : "px-4"}
            `}
          >
            {newRepoOpen ? <CloseIcon /> : "New Repository"}
          </button>

          <NewRepository
            newRepoOpen={newRepoOpen}
            closeNewRepoOpen={setNewRepoOpen}
            setRefreshSignal={setRefreshSignal}
          />
        </div>
      </div>

      {/** Show All repositories */}
      <AllRepos refresh={refreshSignal} />
    </div>
  );
}
