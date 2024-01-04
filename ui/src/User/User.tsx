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
import getTitle from "../lib/getTitle";
import { NavigateFunction, useNavigate } from "react-router-dom";
import { useQuery } from "@tanstack/react-query";
import {
  DeleteUser,
  EditIcon,
  ErrorIcon,
  Spinner,
  WarningIcon,
} from "../lib/icon";
import { GetSinceTime, MessageWithIcon } from "../Apps/AllApplications";
import { toast } from "react-hot-toast";
import normalizeInput from "../utils/normalizeInput";

type RespData = {
  data: User[];
};

type User = {
  createdAt: string;
  lastLoggedIn: string;
  passwordHash: string;
  role: "admin" | "general";
  updatedAt: string;
  username: string;
};

const fetchApps = (navigate: NavigateFunction): Promise<RespData> =>
  fetch("/api/users").then(async (resp) => {
    if (resp.status === 401) {
      navigate("/login");
    }

    return await resp.json();
  });

export default function Users() {
  useEffect(() => {
    document.title = getTitle("Users");
  }, []);

  return (
    <div className="h-screen p-8 overflow-auto">
      <div className="flex justify-between items-center">
        <p className="text-2xl">Users</p>
        <button
          onClick={(e) => {
            e.preventDefault();
          }}
          // button not in use
          // remove the x from hover
          className="bg-white text-black py-2 px-4 rounded font-bold xhover:bg-white/90 cursor-not-allowed opacity-60"
        >
          New User
        </button>
      </div>
      <div className="m-4 md:m-8 mt-8 md:mt-16 overflow-auto h-[84%]">
        <AllUsers />
      </div>
    </div>
  );
}

function AllUsers() {
  const navigate = useNavigate();

  const { data, isLoading, isError, refetch } = useQuery({
    queryKey: ["GET /api/users", "GET_ALL_USERS"],
    queryFn: () => fetchApps(navigate),
  });

  if (isError) {
    return (
      <MessageWithIcon
        icon={<ErrorIcon />}
        message="Something wend wrong while fetching users"
      />
    );
  }

  if (isLoading || data === undefined) {
    return <MessageWithIcon icon={<Spinner />} message="Loading" />;
  }

  if (data.data.length === 0) {
    return <MessageWithIcon icon={<WarningIcon />} message="No Users" />;
  }

  return (
    <ul className="xl:w-[70%] mx-auto">
      {data.data.map((user, index) => (
        <li
          key={index}
          className="p-2 md:p-4 my-2 md:my-4 rounded bg-[#373d49]/30"
        >
          <div className="flex items-center justify-between">
            <div className="flex  items-center justify-center">
              <span className="md:font-bold md:text-xl mr-1 md:mr-4">
                {user.username}
              </span>
              {user.role === "admin" ? (
                <span className="text-xs text-green-400 font-semibold rounded-lg py-1 px-2  bg-green-400/20 mr-1 md:mr-4">
                  admin
                </span>
              ) : null}
              {localStorage.getItem("username") === user.username ? (
                <span className="text-xs text-yellow-400 font-semibold rounded-lg py-1 px-2 bg-yellow-400/20">
                  you
                </span>
              ) : null}
            </div>
            <EditUser username={user.username} refetch={refetch} />
          </div>
          <div className="md:flex md:items-center md:justify-start mt-4 text-sm md:gap-8">
            <div>
              <span className="opacity-50">Last Logged-In: </span>
              <GetSinceTime time={user.lastLoggedIn} />
            </div>
            <div>
              <span className="opacity-50">Updated: </span>
              <GetSinceTime time={user.updatedAt} />
            </div>{" "}
            <div>
              <span className="opacity-50">Created: </span>
              <GetSinceTime time={user.createdAt} />
            </div>
          </div>
        </li>
      ))}
    </ul>
  );
}

function EditUser({
  username,
  refetch,
}: {
  username: string;
  refetch: () => void;
}) {
  const [openEditModal, setOpenEditModal] = useState(false);
  const [currentPass, setCurrentPass] = useState("");
  const [newPass, setNewPass] = useState("");
  const [newPassConfirm, setNewPassConfirm] = useState("");

  const [newUsername, setNewUsername] = useState("");
  const [mode, setMode] = useState("username");

  return (
    <div className="flex items-center gap-4 relative">
      <div
        className={`relative
                  ${openEditModal ? "bg-sidebarLite rounded" : ""}`}
      >
        <div
          className="cursor-pointer hover:bg-sidebarLite rounded p-2  "
          onClick={() => {
            setOpenEditModal(!openEditModal);
          }}
        >
          <EditIcon />
        </div>
        {/** Popup window for editing user  password */}
        <div
          className={`absolute right-12 -top-0 p-4 rounded bg-sidebarLite z-10
            ${openEditModal ? "" : "hidden"}
            `}
        >
          <div className="flex items-center justify-around mb-4 py-1">
            <p
              className={`px-4 cursor-pointer py-1 ${
                mode === "username" ? "bg-slate-500/60  rounded" : "opacity-80"
              }`}
              onClick={() => setMode("username")}
            >
              Username
            </p>
            <p
              className={`px-4 cursor-pointer py-1 ${
                mode === "password" ? "bg-slate-500/60 rounded" : "opacity-80"
              }`}
              onClick={() => setMode("password")}
            >
              Password
            </p>
          </div>
          <form
            className={`flex flex-col gap-4 ${
              mode !== "username" ? "hidden" : ""
            }`}
          >
            <label>
              <span className="text-sm">New Username:</span>
              <input
                type="text"
                className="rounded bg-gray-100/20 px-2 py-1"
                placeholder="beff_jozos"
                required
                onChange={(e) => {
                  setNewUsername(normalizeInput(e.target.value, ["_"], true));
                }}
                value={newUsername}
              />
            </label>
            <button
              type="submit"
              className={`bg-green-400/80  hover:bg-green-400/60 font-medium rounded-lg text-sm px-5 py-2.5
              `}
              onClick={(e) => {
                e.preventDefault();
                const userApi = `/api/users/${username}/username`;

                const req = fetch(userApi, {
                  method: "PATCH",
                  headers: {
                    "Content-Type": "application/json",
                  },
                  body: JSON.stringify({
                    newUsername: newUsername,
                  }),
                });

                toast.promise(req, {
                  loading: "Changing username",
                  success: (res) => {
                    if (res.status === 200) {
                      localStorage.setItem("username", newUsername);
                      setOpenEditModal(false);
                      refetch();
                      return "Username changed successfully";
                    } else {
                      toast.error("Bad request, try again!");
                    }

                    return "Executing task";
                  },
                  error: "Failed to change username, try again!",
                });
              }}
            >
              Change username
            </button>
          </form>
          <form
            className={`flex flex-col gap-4 ${
              mode !== "password" ? "hidden" : ""
            }`}
          >
            <label>
              <span className="text-sm">Current Password:</span>
              <input
                type="password"
                className="rounded bg-gray-100/20 px-2 py-1"
                placeholder="••••••"
                required
                onChange={(e) => {
                  setCurrentPass(e.target.value);
                }}
                value={currentPass}
              />
            </label>
            <label>
              <span className="text-sm">New Password:</span>
              <input
                type="password"
                className="rounded bg-gray-100/20 px-2 py-1"
                required
                placeholder="••••••"
                onChange={(e) => {
                  setNewPass(e.target.value);
                }}
                value={newPass}
              />
            </label>
            <label>
              <span className="text-sm">Confirm New Password:</span>
              <input
                type="password"
                className="rounded bg-gray-100/20 px-2 py-1"
                required
                placeholder="••••••"
                onChange={(e) => {
                  setNewPassConfirm(e.target.value);
                }}
                value={newPassConfirm}
              />
            </label>
            <p
              className={`px-2 text-sm text-red-400 font-bold text-center
            ${newPass === newPassConfirm ? "hidden" : ""}
            `}
            >
              New Password not matching
            </p>
            <button
              type="submit"
              className={`bg-green-400/80  font-medium rounded-lg text-sm px-5 py-2.5
              ${
                newPass !== newPassConfirm
                  ? "opacity-80 cursor-not-allowed"
                  : "hover:bg-green-400/60"
              }
              `}
              onClick={(e) => {
                e.preventDefault();
                if (newPass !== newPassConfirm) return;

                const passApi = `/api/users/${username}/password`;

                const req = fetch(passApi, {
                  method: "PATCH",
                  headers: {
                    "Content-Type": "application/json",
                  },
                  body: JSON.stringify({
                    currentPassword: currentPass,
                    newPassword: newPass,
                  }),
                });

                toast.promise(req, {
                  loading: "Changing password",
                  success: (res) => {
                    if (res.status === 200) {
                      setOpenEditModal(false);
                      return "Password changed successfully";
                    } else {
                      toast.error("Bad request, try again!");
                    }

                    return "Executing task";
                  },
                  error: "Failed to change password, try again!",
                });
              }}
            >
              Change Password
            </button>
          </form>
        </div>
      </div>
      <div className="rounded p-2 cursor-not-allowed opacity-80  ">
        <DeleteUser />
      </div>
    </div>
  );
}
