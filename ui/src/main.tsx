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

import React from "react";
import ReactDOM from "react-dom/client";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import "./index.css";
import Apps from "./Apps/Apps.tsx";
import Login from "./components/signup/Login.tsx";
import Layout from "./components/Layout.tsx";
import Repos from "./Repos.tsx";
import Secrets from "./Secrets.tsx";
import User from "./User.tsx";
import Settings from "./Settings.tsx";

const router = createBrowserRouter([
  {
    path: "/",
    element: <Login />,
  },
  {
    path: "/dash",
    element: <Layout />,
    children: [
      {
        index: true,
        element: <Apps />,
      },
      { path: "/dash/repos", element: <Repos /> },
      { path: "/dash/secrets", element: <Secrets /> },
      { path: "/dash/user", element: <User /> },
      { path: "/dash/settings", element: <Settings /> },
    ],
  },
]);

const queryClient = new QueryClient();

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <QueryClientProvider client={queryClient}>
      <RouterProvider router={router} />
    </QueryClientProvider>
  </React.StrictMode>
);
