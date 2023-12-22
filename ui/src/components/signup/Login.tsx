import { Link } from "react-router-dom";

export default function Login() {
  return (
    <div className="h-screen w-screen flex justify-center items-center">
      <div className="flex flex-col justify-center items-center">
        <Link to="/dashboard">
          <button className="bg-white text-black py-2 px-4 rounded font-bold border-dashed hover:bg-inherit hover:text-white border-2 border-white transition ease-in-out delay-50 hover:-translate-y-1 hover:scale-110 duration-200">
            Open Dashboard
          </button>
        </Link>
      </div>
    </div>
  );
}
