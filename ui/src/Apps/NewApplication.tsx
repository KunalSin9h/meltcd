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

import { useState } from "react";
import toast from "react-hot-toast";
import { CloseIcon } from "../lib/icon";

type globalResponseData = {
  message: string;
};

export default function NewApplication({
  openWindow,
  setOpenWindow,
  setRefresh,
}: {
  openWindow: boolean;
  setOpenWindow: React.Dispatch<React.SetStateAction<boolean>>;
  setRefresh: React.Dispatch<React.SetStateAction<boolean>>;
}) {
  {
    /* Slider Window
        This is used to create a new application
       */
  }
  return (
    <>
      <div
        className={`fixed z-50 overflow-auto  bg-white h-full w-[40%] top-0 left-[60%] p-4 text-black ${
          openWindow ? "" : "hidden"
        }`}
      >
        <div className="flex justify-between items-center">
          <button
            onClick={(e) => {
              e.preventDefault();
              setOpenWindow(false);
            }}
            className="hover:bg-red-200 p-1 rounded"
          >
            <CloseIcon />
          </button>
          <p className="text-xl">Creating a new Application</p>
        </div>
        <div className="h-full px-8 py-16">
          <CreateApplication
            setRefresh={setRefresh}
            setOpenWindow={setOpenWindow}
          />
        </div>
      </div>
      {/* Overlay for Slider */}
      <div
        className={`fixed  h-full w-full top-0 left-0 bg-black/50 ${
          openWindow ? "" : "hidden"
        }`}
      ></div>
    </>
  );
}

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
function CreateApplication({
  setRefresh,
  setOpenWindow,
}: {
  setRefresh: React.Dispatch<React.SetStateAction<boolean>>;
  setOpenWindow: React.Dispatch<React.SetStateAction<boolean>>;
}) {
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

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

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
              // we are good here
              // and refresh the app list
              setRefresh(true);
              // we can close modal window
              setOpenWindow(false);
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
  };

  return (
    <form className="flex flex-col gap-8" onSubmit={handleSubmit}>
      <InputOption
        name="Name"
        placeholder="auth_backend_server"
        value={bodyData.name}
        onChange={(e: React.ChangeEvent<HTMLInputElement>) => {
          setBodyData({
            ...bodyData,
            name: normalizeInput(e.target.value, ["_"]),
          });
        }}
      />
      <InputOption
        name="Sync Timer"
        placeholder="3m30s (Default, for 3 minute and 30 seconds)"
        value={bodyData.refresh_timer}
        onChange={(e: React.ChangeEvent<HTMLInputElement>) => {
          setBodyData({
            ...bodyData,
            refresh_timer: normalizeInput(e.target.value, []),
          });
        }}
      />
      <InputOption
        name="Repository URL"
        placeholder="https://github.com/username/repo"
        value={bodyData.source.repoURL}
        onChange={(e: React.ChangeEvent<HTMLInputElement>) => {
          setBodyData({
            ...bodyData,
            source: {
              ...bodyData.source,
              repoURL: normalizeInput(e.target.value, [
                ":",
                "/",
                ".",
                "_",
                "-",
              ]),
            },
          });
        }}
      />
      <InputOption
        name="Service File Path"
        placeholder="deploy/service.yaml"
        value={bodyData.source.path}
        onChange={(e: React.ChangeEvent<HTMLInputElement>) => {
          setBodyData({
            ...bodyData,
            source: {
              ...bodyData.source,
              path: normalizeInput(e.target.value, [".", "/", "_"]),
            },
          });
        }}
      />
      <InputOption
        name="Target Revision"
        placeholder="HEAD (Default, can be master, main, my_branch)"
        value={bodyData.source.targetRevision}
        onChange={(e: React.ChangeEvent<HTMLInputElement>) => {
          setBodyData({
            ...bodyData,
            source: {
              ...bodyData.source,
              targetRevision: normalizeInput(e.target.value, []),
            },
          });
        }}
      />
      <div className="flex items-center gap-4">
        <input
          className="text-black py-2 px-4 rounded font-bold bg-green-400 hover:bg-green-500 cursor-pointer"
          value="Create"
          type="submit"
        />
        <input
          className="text-black py-2 px-4 rounded font-bold border hover:bg-gray-100 border-1 border-black cursor-pointer"
          onClick={(e) => {
            e.preventDefault();
            setBodyData(initialData);
            toast.success("Input data reset");
          }}
          type="button"
          value="Clear"
        />
      </div>
    </form>
  );
}

function InputOption({
  name,
  placeholder,
  value,
  onChange,
}: {
  name: string;
  placeholder: string;
  value: string;
  onChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
}) {
  const id = name.replace(" ", "_");

  return (
    <label htmlFor={id} className="flex flex-col">
      <span className="font-semibold my-1">{name}</span>
      <input
        id={id}
        required={true}
        className="border p-1 rounded px-2"
        type="text"
        placeholder={placeholder}
        value={value}
        onChange={onChange}
      />
    </label>
  );
}

export function normalizeInput(
  givenText: string,
  allowed: string[],
  toastBottom?: boolean
): string {
  const len = givenText.length;
  let result = "";

  for (let i = 0; i < len; i++) {
    const code = givenText.charCodeAt(i);

    if (
      (code > 47 && code < 58) || // numeric (0-9)
      (code > 64 && code < 91) || // upper alpha (A-Z)
      (code > 96 && code < 123) || // lower alpha (a-z)
      allowed.includes(givenText[i])
    ) {
      result += givenText[i];
    } else {
      toast.error(`${givenText[i]} is not allowed in input here!`, {
        position: toastBottom ? "bottom-right" : "top-center",
      });
    }
  }

  return result;
}
