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
import setTitle from "./lib/setTitle";

export default function Logs() {
  useEffect(() => {
    document.title = setTitle("Logs");
  }, []);

  return (
    <div className="h-screen p-8">
      <div className="flex justify-between items-center">
        <p className="text-2xl">Logs</p>
      </div>
    </div>
  );
}
