import { useEffect, useState } from "react";
import getTitle from "../../lib/getTitle";
import { normalizeInput } from "../../Apps/NewApplication";
import { Toaster } from "react-hot-toast";

export default function SignUp() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");

  useEffect(() => {
    document.title = getTitle("Sign Up");
  });

  return (
    <div className="h-screen bg-rootBg flex items-center justify-center">
      <Toaster />
      <div className="flex flex-col items-center justify-center px-6 py-8 mx-auto lg:py-0">
        <div className="flex items-center mb-6 text-2xl font-semibold text-white">
          <img
            className="w-8 h-8 mr-2 select-none"
            src="/logo.png"
            alt="logo"
          />
          Meltcd
        </div>
        <div className="rounded-lg shadow border border-gray-700 bg-sidebar">
          <div className="p-6 space-y-4 md:space-y-6 sm:p-8">
            <h1 className="text-xl font-bold leading-tight tracking-tight text-gray-900 md:text-2xl text-white md:px-16">
              Create an account
            </h1>
            <form className="space-y-4 md:space-y-6" action="/api/signup">
              <div>
                <label
                  htmlFor="username"
                  className="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                >
                  Your Username
                </label>
                <input
                  type="text"
                  name="username"
                  id="username"
                  className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                  placeholder="kunalsin9h"
                  required={true}
                  onChange={(e) => {
                    setUsername(normalizeInput(e.target.value, ["_"], true));
                  }}
                  value={username}
                />
              </div>
              <div>
                <label
                  htmlFor="password"
                  className="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                >
                  Password
                </label>
                <input
                  type="password"
                  name="password"
                  id="password"
                  placeholder="••••••••"
                  className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                  required={true}
                  onChange={(e) => {
                    setPassword(e.target.value);
                  }}
                  value={password}
                />
              </div>
              <div>
                <label
                  htmlFor="confirm-password"
                  className="block mb-2 text-sm font-medium text-white flex items-center gap-4"
                >
                  <p>Confirm Password</p>
                  <p
                    className={`text-sm font-bold text-red-300 text-center ${
                      password === confirmPassword ? "hidden" : ""
                    }`}
                  >
                    Password must match
                  </p>
                </label>
                <input
                  type="password"
                  name="confirm-password"
                  id="confirm-password"
                  placeholder="••••••••"
                  className="border sm:text-sm rounded-lg  block w-full p-2.5 bg-gray-700 placeholder-gray-400 text-white focus:ring-blue-500 focus:border-blue-500"
                  required={true}
                  onChange={(e) => {
                    setConfirmPassword(e.target.value);
                  }}
                  value={confirmPassword}
                />
              </div>
              <button
                type="submit"
                className="w-full bg-sidebarLite hover:bg-sidebarLite/70 text-white bg-primary-600 hover:bg-primary-700 focus:ring-4 focus:outline-none focus:ring-primary-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-primary-600 dark:hover:bg-primary-700 dark:focus:ring-primary-800"
              >
                Sign Up
              </button>
            </form>
          </div>
        </div>
      </div>
    </div>
  );
}
