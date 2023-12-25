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

import { useEffect } from "react";
import { useParams, useNavigate } from "react-router-dom";
import getTitle from "../lib/getTitle";
import toast from "react-hot-toast";
import { useQuery } from "@tanstack/react-query";
import { MessageWithIcon } from "./AllApplications";
import { ErrorIcon, Spinner } from "../lib/icon";

type respData = {
  created_at: string;
  health: number;
  health_status: string;
  id: number;
  last_synced_at: string;
  name: string;
  refresh_timer: string;
  source: {
    path: string;
    repoURL: string;
    targetRevision: string;
  };
  updated_at: string;
};

export default function AppsDetail() {
  const { name } = useParams();
  const navigate = useNavigate();

  const fetchAppDetail = (): Promise<respData> =>
    fetch(`/api/apps/${name}`).then(async (resp) => await resp.json());

  const { data, isLoading, isError, refetch } = useQuery({
    queryKey: ["GET /api/apps/:name", "GET_APPLICATION_DETAILS"],
    queryFn: fetchAppDetail,
  });

  useEffect(() => {
    let title = name;
    if (title === undefined) {
      title = "Applications";
    }
    document.title = getTitle(title);

    const refetchTimer = setInterval(() => {
      refetch();
    }, 2500);

    return () => {
      clearInterval(refetchTimer);
    };
  }, [name, refetch]);

  if (isError) {
    return (
      <MessageWithIcon
        icon={<ErrorIcon />}
        message="Something wend wrong while fetching application details"
      />
    );
  }

  if (isLoading || data === undefined) {
    return <MessageWithIcon icon={<Spinner />} message="Loading" />;
  }

  return (
    <div className="h-screen p-8">
      <div className="flex justify-between items-center">
        <p className="text-2xl">
          <span
            className="opacity-60 hover:opacity-90 cursor-pointer"
            onClick={(e) => {
              e.preventDefault();
              navigate("/apps");
            }}
          >
            Applications /
          </span>{" "}
          {name}
        </p>
        <button
          onClick={(e) => {
            e.preventDefault();
            const syncAPI = `/api/apps/${name}/refresh`;

            const request = fetch(syncAPI, {
              method: "POST",
            });

            toast.promise(request, {
              loading: "Synchronizing application",
              success: "Application synched successfully",
              error: "Failed to sync application",
            });
          }}
          className="bg-green-500 text-white py-2 px-4 rounded font-bold hover:bg-green-500/90"
        >
          Synchronize
        </button>
      </div>
      <div className="p-8 mt-16">
        <pre>
          <code>{JSON.stringify(data, null, "\t")}</code>
        </pre>
      </div>
    </div>
  );
}
