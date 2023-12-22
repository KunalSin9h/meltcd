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

import { Link } from "react-router-dom";

export default function Login() {
  return (
    <div className="h-screen w-screen flex justify-center items-center">
      <div className="flex flex-col justify-center items-center">
        <Link to="/dash">
          <button className="bg-white text-black py-2 px-4 rounded font-bold border-dashed hover:bg-inherit hover:text-white border-2 border-white transition ease-in-out delay-50 hover:-translate-y-1 duration-100">
            Open Dashboard
          </button>
        </Link>
      </div>
    </div>
  );
}
