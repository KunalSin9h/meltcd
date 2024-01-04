import { useEffect, useState } from "react";
import getTitle from "../../lib/getTitle";
import MeltcdBranding from "../Branding";
import { Spinner } from "../../lib/icon";
import { useNavigate } from "react-router-dom";
import { Toaster, toast } from "react-hot-toast";

export default function LoginPage() {
  const [showSpinner, setShowSpinner] = useState(false);
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  const navigate = useNavigate();

  useEffect(() => {
    document.title = getTitle("Login");
  });

  return (
    <div className="h-screen bg-rootBg flex items-center justify-center">
      <Toaster />
      <div className="flex flex-col items-center justify-center px-6 py-8 mx-auto lg:py-0">
        <MeltcdBranding />
        <div className="rounded-lg shadow border border-gray-700 bg-sidebar">
          <div className="p-6 space-y-4 md:space-y-6 sm:p-8">
            <h1 className="text-xl font-bold leading-tight tracking-tight text-gray-900 md:text-2xl text-white md:px-8">
              Sign in to your account
            </h1>
            <form className="space-y-4 md:space-y-6" action="#">
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
                  onChange={(e) => setUsername(e.target.value)}
                  value={username}
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
                  onChange={(e) => setPassword(e.target.value)}
                  value={password}
                />
              </div>
              <button
                className="w-full bg-sidebarLite hover:bg-sidebarLite/70 font-medium rounded-lg text-sm px-5 py-2.5 text-center"
                onClick={(e) => {
                  e.preventDefault();
                  setShowSpinner(true);

                  if (username === "" || password === "") {
                    toast.error("Some input is still empty!");
                    setShowSpinner(false);
                    return;
                  }

                  const loginAPI = "/api/login";
                  fetch(loginAPI, {
                    method: "POST",
                    headers: {
                      Authorization: `Basic ${getBasicAuthToken(
                        username,
                        password
                      )}`,
                    },
                  })
                    .then((resp) => {
                      setShowSpinner(false);

                      if (resp.status === 400) {
                        toast.error("Bad request!");
                        return;
                      } else if (resp.status === 401) {
                        toast.error("Unauthorized!");
                        return;
                      } else if (resp.status === 500) {
                        toast.error("Internal server error, try again!");
                        return;
                      }

                      navigate("/apps");
                    })
                    .catch(() => {
                      toast.error("Something went  wrong, try again!");
                      setShowSpinner(false);
                    });
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

function getBasicAuthToken(username: string, password: string): string {
  const userPass = `${username}:${password}`;

  return btoa(userPass);
}
