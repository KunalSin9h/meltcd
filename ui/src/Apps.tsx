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

import getTitle from "./lib/getTitle";
import { useEffect } from "react";

export default function Apps() {
  useEffect(() => {
    document.title = getTitle("Applications");
  }, []);

  return (
    <div className="h-screen p-8">
      <div className="flex justify-between items-center">
        <p className="text-2xl">Applications</p>
        <button
          onClick={(e) => {
            e.preventDefault();
          }}
          className="bg-white text-black py-2 px-4 rounded font-bold border-dashed hover:bg-inherit hover:text-white border-2 border-white transition ease-in-out delay-50 hover:-translate-y-1 duration-100"
        >
          New Application
        </button>
      </div>
    </div>
  );
}
