import { useQuery } from "@tanstack/react-query";
import { useEffect, useState } from "react";
import { NavigateFunction, useNavigate } from "react-router-dom";
import { MessageWithIcon } from "../Apps/AllApplications";
import {
  CloseIcon,
  ErrorIcon,
  Spinner,
  TickIcon,
  TrashIcon,
  WarningIcon,
} from "../lib/icon";
import Tooltip from "../lib/Tooltip";
import toast from "react-hot-toast";

type repoData = {
  url: string;
  reachable: boolean;
};

type respData = {
  data: repoData[];
};

const fetchRepos = (navigate: NavigateFunction): Promise<respData> =>
  fetch("/api/repo").then(async (resp) => {
    if (resp.status === 401) {
      navigate("/login");
      return;
    }
    return await resp.json();
  });

interface AllReposProps {
  refresh: boolean;
}

export default function AllRepos(props: AllReposProps) {
  const navigate = useNavigate(); // react router dom navigator for programmatically
  // navigate, used here to go to specific application

  const { data, isLoading, isError, refetch } = useQuery({
    queryKey: ["GET /api/repo", "GET_ALL_REPOS"],
    queryFn: () => fetchRepos(navigate),
  });

  // fetching the current status of application on regular interval
  useEffect(() => {
    const refreshing = setInterval(() => {
      refetch();
    }, 5000);

    return () => {
      clearInterval(refreshing);
    };
  }, [refetch]);

  // when adding a new application
  // this refresh will be updated by other component NewApplication.tsx
  // so this will also be updated
  if (props.refresh === true) {
    refetch();
  }

  if (isError) {
    return (
      <MessageWithIcon
        icon={<ErrorIcon />}
        message="Something wend wrong while fetching repositories"
      />
    );
  }

  if (isLoading || data === undefined) {
    return <MessageWithIcon icon={<Spinner />} message="Loading" />;
  }

  if (data.data === null || data.data.length === 0) {
    return <MessageWithIcon icon={<WarningIcon />} message="No Repositories" />;
  }

  return (
    <ul className="xl:w-[70%] mx-auto">
      {data.data.map((repo, index) => {
        return (
          <li
            key={index}
            className="p-2 md:p-4 my-2 md:my-4 rounded bg-[#373d49]/30 hover:bg-[#373d49]/80 cursor-pointer"
          >
            <div className="flex items-center justify-between gap-2">
              <div>
                <span className="font-semibold mr-4">{repo.url}</span>
                <span>
                  {repo.reachable ? (
                    <Tooltip
                      className="text-green-400 bg-green-300/10"
                      content="Repository is reachable"
                    >
                      <span>
                        <TickIcon className="text-green-400 inline" />
                      </span>
                    </Tooltip>
                  ) : (
                    <Tooltip
                      className="text-red-400 bg-red-300/10"
                      content="Repository is not reachable"
                    >
                      <span>
                        <CloseIcon className="text-red-400 inline" />
                      </span>
                    </Tooltip>
                  )}
                </span>
              </div>
              <EditRepo repoURL={repo.url} refetch={refetch} />
            </div>
          </li>
        );
      })}
    </ul>
  );
}

function EditRepo(props: { repoURL: string; refetch: () => void }) {
  const [deleteOpen, setDeleteOpen] = useState(false);
  return (
    <div className="flex items-center gap-4">
      <div className="relative">
        <span
          onClick={() => {
            setDeleteOpen(!deleteOpen);
          }}
        >
          <TrashIcon />
        </span>
        <div
          className={`absolute top-2 right-8 bg-sidebar/80 rounded p-2 w-48
          ${deleteOpen ? "" : "hidden"}
        `}
        >
          <p>Are you sure you want to delete repo?</p>

          <div className="flex items-center justify-around my-2">
            <button
              className="px-2 py-1 rounded bg-green-600 hover:bg-green-500"
              onClick={() => {
                const api = "/api/repo";

                const req = fetch(api, {
                  method: "DELETE",
                  headers: {
                    "Content-Type": "application/json",
                  },
                  body: JSON.stringify({
                    repo: props.repoURL,
                  }),
                });

                toast.promise(req, {
                  loading: "Deleting repository",
                  success: (res) => {
                    if (res.status === 200) {
                      toast.success("Successfully deleted repository");
                      setDeleteOpen(false);
                      props.refetch();
                    } else {
                      toast.success("something went wrong, try again");
                    }
                    return "Executing task";
                  },
                  error: "Failed to delete repository",
                });
              }}
            >
              Yes
            </button>
            <button
              className="px-2 py-1 rounded border hover:bg-gray-500/80"
              onClick={() => setDeleteOpen(false)}
            >
              Cancel
            </button>
          </div>
        </div>
      </div>
    </div>
  );
}
