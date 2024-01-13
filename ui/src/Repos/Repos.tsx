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
import normalizeInput from "../utils/normalizeInput";
import { CloseIcon } from "../lib/icon";
import toast from "react-hot-toast";

export default function Repos() {
  const [newRepoOpen, setNewRepoOpen] = useState(false);
  const [repoURL, setRepoURL] = useState("");
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

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
            className="bg-white text-black py-2 px-4 rounded font-bold hover:bg-white/90"
          >
            {newRepoOpen ? <CloseIcon /> : "New Repository"}
          </button>
          <div
            className={`absolute top-16 bg-white right-0 p-4 flex flex-col gap-4 text-black w-96 rounded
            ${!newRepoOpen ? "hidden" : ""}
          `}
          >
            <label className="flex flex-col">
              <span className="font-bold">Repository Url</span>
              <input
                placeholder="https://github.com/k9exp/infra-test"
                className="bg-white border p-2 rounded mt-1"
                onChange={(e) => {
                  setRepoURL(
                    normalizeInput(
                      e.target.value,
                      [":", "/", ".", "_", "-"],
                      true
                    )
                  );
                }}
                value={repoURL}
              />
            </label>
            <div className="text-center">
              <span className="rounded-full px-2 py-1 border border-green-400 bg-green-200 ">
                &darr; Basic Auth
              </span>
            </div>
            <label className="flex flex-col">
              <span className="font-bold">Username</span>
              <input
                placeholder="beff_jezos"
                type="text"
                className="bg-white border p-2 rounded mt-1"
                onChange={(e) => setUsername(e.target.value)}
                value={username}
              />
            </label>
            <label className="flex flex-col">
              <span className="font-bold">Password</span>
              <input
                placeholder="••••••••"
                type="password"
                className="bg-white border p-2 rounded mt-1"
                onChange={(e) => setPassword(e.target.value)}
                value={password}
              />
            </label>
            <button
              className="text-black py-2 px-4 rounded font-bold bg-green-400 hover:bg-green-500 cursor-pointer"
              onClick={() => {
                const api = "/api/repo";

                const req = fetch(api, {
                  method: "POST",
                  headers: {
                    "Content-Type": "application/json",
                  },
                  body: JSON.stringify({
                    username,
                    password,
                    url: repoURL,
                  }),
                });

                toast.promise(req, {
                  loading: "Adding new Repository",
                  success: (res) => {
                    let msg = "Successfully created new repository";

                    if (res.status === 401) {
                      msg = "Unauthorized";
                      // we can navigate to /login, but coming here is theoretically
                      // impossible
                    } else if (res.status == 400) {
                      msg = "Bad request";
                    } else if (res.status === 500) {
                      msg = "Internal Server Error";
                    }

                    toast.success(msg);
                    return "Executing task";
                  },
                  error: "Failed to add new Repository",
                });
              }}
            >
              Add
            </button>
          </div>
        </div>
      </div>
    </div>
  );
}
