/*
Copyright 2023 - PRESENT kunalsin9h

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
import Layout from "./components/Layout.tsx";
import Secrets from "./Secrets.tsx";
import Users from "./User/User.tsx";
import Settings from "./Settings.tsx";
import AppsDetail from "./Apps/AppDetail.tsx";
import Logs from "./Logs.tsx";
import LoginPage from "./components/auth/Login.tsx";
import Repos from "./Repos/Repos.tsx";

const router = createBrowserRouter([
  {
    path: "/login",
    element: <LoginPage />,
  },
  {
    path: "/",
    element: <Layout />,
    children: [
      {
        path: "/apps",
        element: <Apps />,
      },
      {
        path: "/apps/:name",
        element: <AppsDetail />,
      },
      { path: "/repos", element: <Repos /> },
      { path: "/secrets", element: <Secrets /> },
      { path: "/users", element: <Users /> },
      { path: "/settings", element: <Settings /> },
      { path: "/logs", element: <Logs /> },
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
