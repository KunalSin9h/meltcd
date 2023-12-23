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

import getTitle from "./lib/getTitle";
import { useEffect, useState } from "react";
import { toast } from "react-hot-toast";

export default function Apps() {
  useEffect(() => {
    document.title = getTitle("Applications");
  }, []);

  const [openWindow, setOpenWindow] = useState(false);

  return (
    <div className="h-screen p-8">
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

      {/* Slider Window
        This is used to create a new application
       */}
      <div
        className={`fixed z-50 bg-white h-full w-[40%] top-0 left-[60%] p-4 text-black ${
          openWindow ? "" : "hidden"
        }`}
      >
        <div className="flex justify-between items-center">
          <button
            onClick={(e) => {
              e.preventDefault();
              setOpenWindow(false);
            }}
            className="text-black py-2 px-4 rounded font-bold bg-red-300 hover:bg-red-400"
          >
            Close
          </button>
          <p className="text-xl">Creating a new Application</p>
        </div>
        <div className="h-full px-8 py-16">
          <NewApplication />
        </div>
      </div>

      {/* Overlay for Slider */}
      <div
        className={`fixed  h-full w-full top-0 left-0 backdrop-blur-sm ${
          openWindow ? "" : "hidden"
        }`}
      ></div>
    </div>
  );
}

type globalResponseData = {
  message: string;
};

// USING api POST /api/apps
/**
  body{
    "name": "string",
    "refresh_timer": "string",
    "source": {
      "path": "string",
      "repoURL": "string",
      "targetRevision": "string"
    },
  } 
 */
function NewApplication() {
  const initialData = {
    name: "",
    refresh_timer: "3m0s",
    source: {
      path: "",
      repoURL: "",
      targetRevision: "HEAD",
    },
  };

  const [bodyData, setBodyData] = useState(initialData);

  return (
    <div className="flex flex-col gap-8">
      <label htmlFor="name" className="flex flex-col">
        <span className="font-semibold my-1">Name</span>
        <input
          id="name"
          required={true}
          className="border p-1 rounded px-2"
          type="text"
          placeholder="auth-backend-server"
          value={bodyData.name}
          onChange={(e) => {
            setBodyData({
              ...bodyData,
              name: e.target.value.trim(),
            });
          }}
        />
      </label>
      <label htmlFor="sync" className="flex flex-col">
        <span className="font-semibold my-1">Sync Timer</span>
        <input
          id="sync"
          className="border p-1 rounded px-2"
          type="text"
          placeholder="3m30s (Default, for 3 minute and 30 seconds)"
          value={bodyData.refresh_timer}
          onChange={(e) => {
            setBodyData({
              ...bodyData,
              refresh_timer: e.target.value.trim(),
            });
          }}
        />
      </label>
      <label htmlFor="repo" className="flex flex-col">
        <span className="font-semibold my-1">Repository URL</span>
        <input
          id="repo"
          required={true}
          className="border p-1 rounded px-2"
          type="text"
          placeholder="https://github.com/username/repo"
          value={bodyData.source.repoURL}
          onChange={(e) => {
            setBodyData({
              ...bodyData,
              source: {
                ...bodyData.source,
                repoURL: e.target.value.trim(),
              },
            });
          }}
        />
      </label>
      <label htmlFor="path" className="flex flex-col">
        <span className="font-semibold my-1">Service File Path</span>
        <input
          id="path"
          required={true}
          className="border p-1 rounded px-2"
          type="text"
          placeholder="deploy/service.yaml"
          value={bodyData.source.path}
          onChange={(e) => {
            setBodyData({
              ...bodyData,
              source: {
                ...bodyData.source,
                path: e.target.value.trim(),
              },
            });
          }}
        />
      </label>
      <label htmlFor="rev" className="flex flex-col">
        <span className="font-semibold my-1">Target Revision</span>
        <input
          id="rev"
          className="border p-1 rounded px-2"
          type="text"
          placeholder="HEAD (Default, can be master, main, my_branch)"
          value={bodyData.source.targetRevision}
          onChange={(e) => {
            setBodyData({
              ...bodyData,
              source: {
                ...bodyData.source,
                targetRevision: e.target.value.trim(),
              },
            });
          }}
        />
      </label>
      <div className="flex items-center gap-4">
        <button
          className="text-black py-2 px-4 rounded font-bold bg-green-400 hover:bg-green-500"
          onClick={async (e) => {
            e.preventDefault();

            if (bodyData.name === "") {
              toast.error("Name is empty");
              return;
            }

            if (bodyData.refresh_timer === "") {
              toast.error("Sync Timer is empty");
              return;
            }

            if (bodyData.source.repoURL === "") {
              toast.error("Repository URL is empty");
              return;
            }

            if (bodyData.source.path === "") {
              toast.error("Service file path is empty");
              return;
            }

            if (bodyData.source.targetRevision === "") {
              toast.error("Target Revision is empty");
              return;
            }

            const request = fetch("/api/apps", {
              method: "POST",
              headers: {
                "Content-Type": "application/json",
              },
              body: JSON.stringify(bodyData),
            });

            toast.promise(request, {
              loading: "Creating new application",
              success: (res) => {
                let good = true;
                if (res.status !== 200) {
                  good = false;
                }

                res
                  .json()
                  .then((data: globalResponseData) => {
                    if (good) {
                      toast.success(data.message);
                    } else {
                      toast.error(data.message);
                    }
                  })
                  .catch((err) => {
                    console.log(err);
                    toast.error("Failed to create new application");
                  });

                return "Executing task";
              },
              error: (err) => {
                console.log(err);
                return "Failed to create new application";
              },
            });
          }}
        >
          Create
        </button>
        <button
          className="text-black py-2 px-4 rounded font-bold border hover:bg-gray-100 border-1 border-black"
          onClick={(e) => {
            e.preventDefault();
            setBodyData(initialData);
            toast.success("Input data reset");
          }}
        >
          Clear
        </button>
      </div>
    </div>
  );
}
