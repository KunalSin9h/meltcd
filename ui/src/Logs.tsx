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

import { useEffect, useState, useRef } from "react";
import setTitle from "./lib/setTitle";
import { MessageWithIcon } from "./Apps/AllApplications";
import { WarningIcon } from "./lib/icon";

type LogData = {
  time: string;
  level: string;
  msg: string;
  [k: string]: string;
}

export default function Logs() {
  const listRef = useRef<HTMLUListElement>(null);
  const [logs, setLogs] = useState<LogData[]>([]);

  useEffect(() => {
    document.title = setTitle("Logs");

    const liveLogsURL = "/api/logs/live";

    const source = new EventSource(liveLogsURL);

    source.addEventListener("log", (e) => {
      const ld = JSON.parse(e.data) as LogData
      setLogs([...logs, ld])

      listRef.current?.lastElementChild?.scrollIntoView()
    })

    return () => {
      source.close()
    }
  });

  if (logs.length === 0){
    return <MessageWithIcon icon={<WarningIcon />} message="No Logs Records" />;
  }
 
  return (
    <div className="h-full p-8">
      <div className="flex justify-between items-center">
        <p className="text-2xl">Logs</p>
      </div>

      <div
        id="logs-list"
        className="bg-sidebar h-[96%] my-4 rounded overflow-auto transition-all"
       >
        <ul className="h-full"  ref={listRef}>
          {logs.map((ld, index) => (
            <li
              key={index}
              className="flex px-2 py-1 items-center gap-2 font-mono cursor-pointer hover:bg-sidebarLite/20"
            >
              <span className="text-slate-500">{getTime(ld.time)}</span>
              <span className={`font-bold ${getStyleForLevel(ld.level)}`}>
                {ld.level}
              </span>
              <span>{ld.msg}</span>
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
}

function getTime(time: string): string {
  const d = new Date(time)
  return d.toLocaleString()
}

function getStyleForLevel(level: string): string {
  switch (level) {
    case "INFO":
      return "text-cyan-500";
    case "WARN":
      return "text-yellow-500";
    case "ERROR":
      return "text-red-500";
  }

  return "text-slate-200";
}