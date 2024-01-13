import { useState } from "react";
import normalizeInput from "../utils/normalizeInput";
import toast from "react-hot-toast";
import { Spinner } from "../lib/icon";

interface NewRepositoryProps {
  newRepoOpen: boolean;
  closeNewRepoOpen: (b: boolean) => void;
  setRefreshSignal: (b: boolean) => void;
}

export default function NewRepository(props: NewRepositoryProps) {
  const [repoURL, setRepoURL] = useState("");
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [processing, setProcessing] = useState(false);

  return (
    <div
      className={`absolute top-16 bg-white right-0 p-4 flex flex-col gap-4 text-black w-96 rounded
            ${!props.newRepoOpen ? "hidden" : ""}
          `}
    >
      <label className="flex flex-col">
        <span className="font-bold">Repository Url</span>
        <input
          placeholder="https://github.com/k9exp/infra-test"
          className="bg-white border p-2 rounded mt-1"
          onChange={(e) => {
            setRepoURL(
              normalizeInput(e.target.value, [":", "/", ".", "_", "-"], true)
            );
          }}
          value={repoURL}
        />
      </label>
      <div className="text-center">
        <span className="rounded-full px-2 py-1 border border-green-400 bg-green-200 ">
          &darr; Basic Auth
        </span>
      </div>
      <label className="flex flex-col">
        <span className="font-bold">Username</span>
        <input
          placeholder="beff_jezos"
          type="text"
          className="bg-white border p-2 rounded mt-1"
          onChange={(e) => setUsername(e.target.value)}
          value={username}
        />
      </label>
      <label className="flex flex-col">
        <span className="font-bold">Password</span>
        <input
          placeholder="••••••••"
          type="password"
          className="bg-white border p-2 rounded mt-1"
          onChange={(e) => setPassword(e.target.value)}
          value={password}
        />
      </label>
      <button
        className="text-black py-2 px-4 rounded font-bold bg-green-400 hover:bg-green-500 cursor-pointer flex justify-center items-center"
        onClick={() => {
          if (username === "" || password === "") {
            toast.error("Input must not be empty!");
            return;
          }

          setProcessing(true);
          const api = "/api/repo";

          const req = fetch(api, {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify({
              username,
              password,
              url: repoURL,
            }),
          });

          toast.promise(req, {
            loading: "Adding new Repository",
            success: (res) => {
              if (res.status === 202) {
                toast.success("Successfully added repository");
                setProcessing(false);
                props.closeNewRepoOpen(false);
                props.setRefreshSignal(true);
              } else {
                try {
                  res.json().then((d) => {
                    setProcessing(false);
                    props.closeNewRepoOpen(false);
                    toast.error(d.message);
                  });
                } catch (error) {
                  toast.error(
                    "Failed to sent api request, something went wrong"
                  );
                  console.log(error);
                }
              }

              return "Executing task";
            },
            error: "Failed to add new Repository",
          });
        }}
      >
        {processing ? <Spinner black /> : "Add"}
      </button>
    </div>
  );
}
