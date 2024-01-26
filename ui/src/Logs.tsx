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
import setTitle from "./lib/setTitle";
import { MessageWithIcon } from "./Apps/AllApplications";
import { WarningIcon } from "./lib/icon";

export default function Logs() {
  const [logs, setLogs] = useState<string>("");

  useEffect(() => {
    document.title = setTitle("Logs");

    const liveLogsURL = "/api/logs/live";

    const source = new EventSource(liveLogsURL);

    source.onopen = (e) => {
      console.log(e)
    }

    source.onerror = (e) => {
      console.log(e)
    }

    source.onmessage = (e) => {
      console.log(e.data)
      setLogs(e.data)
    }

    return () => {
      source.close()
    }
  });

  if (logs.length === 0) {
    return <MessageWithIcon icon={<WarningIcon />} message="No Log record found" />;
  }

  return (
    <div className="h-full p-8">
      <div className="flex justify-between items-center">
        <p className="text-2xl">Logs</p>
      </div>
      <div className="bg-sidebar h-[96%] my-4 rounded overflow-auto">
        { logs }
      </div>
    </div>
  );
}
