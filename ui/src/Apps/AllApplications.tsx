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
import { useEffect, useState } from "react";
import { NavigateFunction, useNavigate } from "react-router-dom";
import { toast } from "react-hot-toast";

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

const fetchApps = (navigate: NavigateFunction): Promise<respData> =>
	fetch("/api/apps").then(async (resp) => {
		if (resp.status === 401) {
			navigate("/login");
		} else if (resp.status !== 200) {
			toast.error("Something wend wrong, server didn't respond with 200");
			return;
		}

		return await resp.json();
	});

export default function AllApplications({ refresh }: { refresh: boolean }) {
	const navigate = useNavigate(); // react router dom navigator for programmatically
	// navigate, used here to go to specific application

	const { data, isLoading, isError, refetch } = useQuery({
		queryKey: ["GET /api/apps", "GET_ALL_APPLICATIONS"],
		queryFn: () => fetchApps(navigate),
	});

	// fetching the current status of application on regular interval
	useEffect(() => {
		const refreshing = setInterval(() => {
			refetch();
		}, 2000);

		return () => {
			clearInterval(refreshing);
		};
	}, [refetch]);

	// when adding a new application
	// this refresh will be updated by other component NewApplication.tsx
	// so this will also be updated
	if (refresh === true) {
		refetch();
	}

	if (isError) {
		return (
			<MessageWithIcon
				icon={<ErrorIcon />}
				message="Something wend wrong while fetching applications"
			/>
		);
	}

	if (isLoading || data === undefined) {
		return <MessageWithIcon icon={<Spinner />} message="Loading" />;
	}

	if (data.data === null || data.data.length === 0) {
		return (
			<MessageWithIcon icon={<WarningIcon />} message="No Application" />
		);
	}

	return (
		<ul className="xl:w-[70%] mx-auto">
			{data.data.map((app, index) => (
				<li
					key={index}
					className="p-2 md:p-4 my-2 md:my-4 rounded bg-[#373d49]/30 hover:bg-[#373d49]/80 cursor-pointer"
					onClick={(e) => {
						e.preventDefault();
						navigate(`/apps/${app.name}`);
					}}
				>
					<div className="flex items-center justify-between">
						<div className="flex  items-center justify-center">
							<span className="md:font-bold md:text-xl mr-1 md:mr-4">
								{app.name}
							</span>
							<GetHealthBadge health={app.health} />
						</div>
					</div>
					<div className="md:flex md:items-center md:justify-start mt-4 text-sm md:gap-8">
						<div>
							<span className="opacity-50">Last synched: </span>
							<GetSinceTime time={app.last_synced_at} />
						</div>
						<div>
							<span className="opacity-50">Updated: </span>
							<GetSinceTime time={app.updated_at} />
						</div>{" "}
						<div>
							<span className="opacity-50">Created: </span>
							<GetSinceTime time={app.created_at} />
						</div>
					</div>
				</li>
			))}
		</ul>
	);
}

export function MessageWithIcon({
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

function GetHealthBadge({ health }: { health: string }) {
	switch (health) {
		case "healthy":
			return (
				<span className="text-xs text-green-400 font-semibold rounded-lg py-1 px-2  bg-green-400/20">
					healthy
				</span>
			);
		case "progressing":
			return (
				<span className="text-xs text-blue-400 font-semibold rounded-lg py-1 px-2  bg-blue-400/20">
					progressing
				</span>
			);
		case "degraded":
			return (
				<span className="text-xs text-yellow-400 font-semibold rounded-lg py-1 px-2  bg-yellow-400/20">
					degraded
				</span>
			);
		case "suspended":
			return (
				<span className="text-xs text-red-400 font-semibold rounded-lg py-1 px-2  bg-red-400/20">
					suspended
				</span>
			);
		default:
			return (
				<span className="text-xs text-gray-400 font-semibold rounded-lg py-1 px-2  bg-gray-400/20">
					Status: N/A
				</span>
			);
	}
}

export function GetSinceTime({ time }: { time: string }) {
	const [currentTime, setCurrentTime] = useState(Date.now());

	useEffect(() => {
		const refreshTimer = setInterval(() => {
			setCurrentTime(Date.now());
		}, 10000); // get counter by every minute

		return () => {
			clearInterval(refreshTimer);
		};
	}, []);

	// Time means time is not set by backend, or empty time
	// we we must not show time
	if (time === "0001-01-01T00:00:00Z") {
		return <span className="opacity-90">N/A</span>;
	}

	const t = new Date(time);
	const elapsed = currentTime - t.getTime();

	if (isNaN(elapsed)) {
		return "Just now";
	}

	const seconds = elapsed / 1000;
	let minutes = seconds / 60;
	let hours = minutes / 60;
	let days = hours / 24;
	let weeks = days / 7;
	let months = weeks / 4.34524;
	let year = months / 12;

	year = Math.floor(year);
	if (year > 0) {
		return `${year} year ago`;
	}

	months = Math.floor(months);
	if (months > 0) {
		return `${months} months ago`;
	}

	weeks = Math.floor(weeks);
	if (weeks > 0) {
		return `${weeks} weeks ago`;
	}

	days = Math.floor(days);
	if (days > 0) {
		return `${days} days ago`;
	}

	hours = Math.floor(hours);
	if (hours > 0) {
		return `${hours} hours ago`;
	}

	minutes = Math.floor(minutes);
	if (minutes > 0) {
		return `${minutes} minutes ago`;
	}

	return "Just now";
}
