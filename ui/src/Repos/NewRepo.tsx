import { useState } from "react";
import normalizeInput from "../utils/normalizeInput";
import toast from "react-hot-toast";
import { Spinner } from "../lib/icon";

interface NewRepositoryProps {
  newRepoOpen: boolean;
  closeNewRepoOpen: (b: boolean) => void;
  setRefreshSignal: (b: boolean) => void;
}

enum RepoInputType {
  GitRepo,
  ContainerImage
}

export default function NewRepository(props: NewRepositoryProps) {
  const [inputType, setInputType] = useState<RepoInputType>(RepoInputType.GitRepo)
  const [repoInput, setRepoInput] = useState("");
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [processing, setProcessing] = useState(false);

  return (
    <div
      className={`absolute top-16 bg-white right-0 p-4 flex flex-col gap-4 text-black w-96 rounded z-10
            ${!props.newRepoOpen ? "hidden" : ""}
          `}
    >
      <div className="flex items-center justify-around">
        <button className={`px-4 py-2 hover:bg-green-300 rounded 
          ${inputType === RepoInputType.GitRepo && "shadow bg-green-400/40"}
        `}
          onClick={() => setInputType(RepoInputType.GitRepo)}>Git Repository</button>

        <button className={`px-4 py-2 hover:bg-green-300 rounded 
          ${inputType === RepoInputType.ContainerImage && "shadow bg-green-400/40"}
        `}
          onClick={() => setInputType(RepoInputType.ContainerImage)}>Container Image</button>
      </div>
      <label className="flex flex-col">
        <span className="font-bold">{inputType === RepoInputType.GitRepo ? "Repository URL" : "Container Image"}</span>
        <input
          placeholder={`${ inputType === RepoInputType.GitRepo ? "https://github.com/k9exp/infra-test" : "ghcr.io/meltred/meltcd"}`}
          className="bg-white border p-2 rounded mt-1"
          onChange={(e) => {
            setRepoInput(
              normalizeInput(e.target.value, [":", "/", ".", "_", "-"], true)
            );
          }}
          value={repoInput}
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
              url: inputType === RepoInputType.GitRepo ? repoInput : "",
              image_ref: inputType === RepoInputType.ContainerImage ? repoInput : ""
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
