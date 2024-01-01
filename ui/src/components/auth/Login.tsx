import { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import getTitle from "../../lib/getTitle";

export default function LoginPage() {
  const navigate = useNavigate();

  useEffect(() => {
    document.title = getTitle("Login");
  });

  return (
    <div className="h-screen bg-rootBg flex items-center justify-center">
      <div className="flex flex-col items-center justify-center px-6 py-8 mx-auto lg:py-0">
        <div className="flex items-center mb-6 text-2xl font-semibold text-gray-900 dark:text-white">
          <img
            className="w-8 h-8 mr-2 select-none"
            src="/logo.png"
            alt="logo"
          />
          Meltcd
        </div>
        <div className="rounded-lg shadow border border-gray-700 bg-sidebar">
          <div className="p-6 space-y-4 md:space-y-6 sm:p-8">
            <h1 className="text-xl font-bold leading-tight tracking-tight text-gray-900 md:text-2xl text-white md:px-8">
              Sign in to your account
            </h1>
            <form
              className="space-y-4 md:space-y-6"
              method="POST"
              action="/api/login"
            >
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
                />
              </div>
              <div
                className="text-sm font-medium hover:underline text-primary-500 cursor-pointer"
                onClick={() => {
                  navigate("/password");
                }}
              >
                Forgot password?
              </div>
              <button
                type="submit"
                className="w-full bg-sidebarLite hover:bg-sidebarLite/70 text-white bg-primary-600 hover:bg-primary-700 focus:ring-4 focus:outline-none focus:ring-primary-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-primary-600 dark:hover:bg-primary-700 dark:focus:ring-primary-800"
              >
                Sign in
              </button>
              <p className="text-sm font-light text-gray-500 dark:text-gray-400">
                Don’t have an account yet?{" "}
                <div
                  className="font-medium hover:underline inline"
                  onClick={() => {
                    navigate("/signup");
                  }}
                >
                  Sign up
                </div>
              </p>
            </form>
          </div>
        </div>
      </div>
    </div>
  );
}
