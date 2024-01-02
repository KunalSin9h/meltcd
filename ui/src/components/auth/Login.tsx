import { useEffect, useState } from "react";
import getTitle from "../../lib/getTitle";
import MeltcdBranding from "../Branding";
import { Spinner } from "../../lib/icon";
// import { Spinner } from "../../lib/icon";

export default function LoginPage() {
  const [showSpinner, setShowSpinner] = useState(false);

  useEffect(() => {
    document.title = getTitle("Login");
  });

  return (
    <div className="h-screen bg-rootBg flex items-center justify-center">
      <div className="flex flex-col items-center justify-center px-6 py-8 mx-auto lg:py-0">
        <MeltcdBranding />
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
                  className="block mb-2 text-sm font-medium text-white"
                >
                  Your Username
                </label>
                <input
                  type="text"
                  name="username"
                  id="username"
                  className="border sm:text-sm rounded-lg  block w-full p-2.5 bg-gray-700 border-gray-600 placeholder-gray-400 text-white focus:ring-blue-500 focus:border-blue-500"
                  placeholder="admin"
                  required={true}
                />
              </div>
              <div>
                <label
                  htmlFor="password"
                  className="block mb-2 text-sm font-medium text-white"
                >
                  Password
                </label>
                <input
                  type="password"
                  name="password"
                  id="password"
                  placeholder="••••••••"
                  className="border sm:text-sm rounded-lg  block w-full p-2.5 bg-gray-700 border-gray-600 placeholder-gray-400 text-white focus:ring-blue-500 focus:border-blue-500"
                  required={true}
                />
              </div>
              <button
                type="submit"
                className="w-full bg-sidebarLite hover:bg-sidebarLite/70 font-medium rounded-lg text-sm px-5 py-2.5 text-center"
                onClick={() => {
                  setShowSpinner(true);
                }}
              >
                {showSpinner ? (
                  <div className="flex justify-center items-center">
                    <Spinner />
                  </div>
                ) : (
                  "Sing in"
                )}
              </button>
            </form>
          </div>
        </div>
      </div>
    </div>
  );
}
