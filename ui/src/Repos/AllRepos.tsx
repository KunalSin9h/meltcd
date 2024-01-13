import { useQuery } from "@tanstack/react-query";
import { useEffect } from "react";
import { NavigateFunction, useNavigate } from "react-router-dom";
import { MessageWithIcon } from "../Apps/AllApplications";
import { ErrorIcon, Spinner, WarningIcon } from "../lib/icon";

type respData = {
  data: string[];
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
    return <MessageWithIcon icon={<WarningIcon />} message="No Application" />;
  }

  return (
    <ul className="xl:w-[70%] mx-auto">
      {data.data.map((repo, index) => (
        <li
          key={index}
          className="p-2 md:p-4 my-2 md:my-4 rounded bg-[#373d49]/30 hover:bg-[#373d49]/80 cursor-pointer"
        >
          <div>{repo}</div>
        </li>
      ))}
    </ul>
  );
}
