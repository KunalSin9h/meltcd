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

  const [deleteModal, setDeleteModal] = useState(false);

  useEffect(() => {
    let title = name;
    if (title === undefined) {
      title = "Applications";
    }
    document.title = getTitle(title);
  }, [name]);

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
        <div className="flex items-center gap-4">
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
          <button
            className="text-red-400 py-2 px-4 rounded border border-red-400 font-bold hover:bg-red-400/20"
            onClick={(e) => {
              e.preventDefault();
              setDeleteModal(true);
            }}
          >
            Delete
          </button>

          {/**Modal window for deleting application */}
          <div
            className={`overflow-y-auto overflow-x-hidden fixed top-0 right-0 left-0 z-50 flex justify-center items-center h-full w-full bg-black/40
            ${deleteModal ? "" : "hidden"}
            `}
          >
            <div className="bg-sidebar p-4 rounded flex flex-col gap-2">
              <p className="text-xl font-bold">Delete Application?</p>
              <p>Are you sure you want to delete application!</p>
              <div className="flex justify-end gap-4 items-center">
                <button
                  className={`py-1 px-2 border border-white/30 rounded font-bold hover:bg-white/10`}
                  onClick={(e) => {
                    e.preventDefault();
                    setDeleteModal(false);
                  }}
                >
                  Cancel
                </button>
                <button
                  className={`py-1 px-2 rounded bg-red-500 text-white hover:bg-red-400 font-bold`}
                  onClick={(e) => {
                    e.preventDefault();
                    const deleteAPI = `/api/apps/${name}`;

                    const request = fetch(deleteAPI, {
                      method: "DELETE",
                    });

                    toast.promise(request, {
                      loading: "Deleting application",
                      success: "Application deleted successfully",
                      error: "Failed to delete application",
                    });

                    setDeleteModal(false);
                  }}
                >
                  Delete
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div className="p-8 mt-16">
        <ShowAppDetails name={name} />
      </div>
    </div>
  );
}

function ShowAppDetails({ name }: { name: string | undefined }) {
  const fetchAppDetail = (): Promise<respData> =>
    fetch(`/api/apps/${name}`).then(async (resp) => await resp.json());

  const { data, isLoading, isError, refetch } = useQuery({
    queryKey: ["GET /api/apps/:name", "GET_APPLICATION_DETAILS"],
    queryFn: fetchAppDetail,
  });

  useEffect(() => {
    const refetchTimer = setInterval(() => {
      refetch();
    }, 5000);

    return () => {
      clearInterval(refetchTimer);
    };
  }, [refetch]);

  if (isError) {
    return (
      <MessageWithIcon
        icon={<ErrorIcon />}
        message="Something wend wrong while fetching application details"
      />
    );
  }

  if (isLoading || data === undefined || name == undefined) {
    return <MessageWithIcon icon={<Spinner />} message="Loading" />;
  }

  return (
    <div>
      <pre>
        <code>{JSON.stringify(data, null, "\t")}</code>
      </pre>
    </div>
  );
}
