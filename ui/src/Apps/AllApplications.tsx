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

import { useQuery } from "@tanstack/react-query";
import { ErrorIcon, Spinner, WarningIcon } from "../lib/icon";
import { useEffect } from "react";

type respData = {
  data: appData[];
};

type appData = {
  created_at: string;
  health: string;
  id: number;
  last_synced_at: string;
  name: string;
  updated_at: string;
};

const fetchApps = (): Promise<respData> =>
  fetch("/api/apps").then(async (resp) => await resp.json());

export default function AllApplications({ refresh }: { refresh: boolean }) {
  const { data, isLoading, isError, refetch } = useQuery({
    queryKey: ["GET /api/apps", "GET_ALL_APPLICATIONS"],
    queryFn: fetchApps,
  });

  // fetching the current status of application on regular interval
  useEffect(() => {
    const refreshing = setInterval(() => {
      console.log("Refetching");
      refetch();
    }, 2500);

    return () => {
      clearInterval(refreshing);
    };
  });

  // when adding a new application
  // this refresh will be updated by other component NewApplication.tsx
  // so this will also fe updated
  if (refresh === true) {
    refetch();
  }

  if (isError || data === undefined) {
    return (
      <MessageWithIcon
        icon={<ErrorIcon />}
        message="Something wend wrong while fetching applications"
      />
    );
  }

  if (isLoading) {
    return <MessageWithIcon icon={<Spinner />} message="Loading" />;
  }

  if (data.data === null || data.data.length === 0) {
    return <MessageWithIcon icon={<WarningIcon />} message="No Application" />;
  }

  return (
    <table className="w-full">
      <thead>
        <tr>
          <th>S.NO</th>
          <th>Name</th>
          <th>Last Synched At</th>
          <th>Created At</th>
          <th>Updated At</th>
        </tr>
      </thead>
      <tbody>
        {data.data.map((app, index) => (
          <tr
            className="group/app hover:bg-[#414b58] cursor-pointer"
            key={index}
          >
            <td className={`font-bold ${getBgColorForHealth(app.health)}`}>
              {app.id}
            </td>
            <td>{app.name}</td>
            <td>{getTimeDate(app.last_synced_at)}</td>
            <td>{getTimeDate(app.created_at)}</td>
            <td>{getTimeDate(app.updated_at)}</td>
          </tr>
        ))}
      </tbody>
    </table>
  );
}

function MessageWithIcon({
  icon,
  message,
}: {
  icon: JSX.Element;
  message: string;
}) {
  return (
    <div className="h-64 flex justify-center items-center">
      <div className="flex items-center gap-2">
        {icon}
        <p className="text-xl">{message}</p>
      </div>
    </div>
  );
}

function getBgColorForHealth(health: string): string {
  switch (health) {
    case "healthy":
      return "border-l-4 bg-green-400/20 border-l-green-400 group-hover/app:border-l-green-300";
    case "progressing":
      return "border-l-4 bg-blue-400/20 border-l-blue-400 group-hover/app:border-l-blue-300";
    case "degraded":
      return "border-l-4 bg-yellow-400/20 border-l-yellow-400 group-hover/app:border-l-yellow-300";
    case "suspended":
      return "border-l-4 bg-red-400/20 border-l-red-400 group-hover/app:border-l-red-300";
    default:
      return "bg-inherit";
  }
}

/**
 * TODO: Change the timezone with Settings
 * use the Settings window to change the timezone and apply here
 *
 * we can use redux for global state management
 */
function getTimeDate(time: string): string {
  const date = new Date(time);
  return `${date.toLocaleTimeString()} ${date.toLocaleDateString()}`;
}
