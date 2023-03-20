import { Routes, Route, useNavigate, A } from "@solidjs/router";
import { createEffect, lazy, Show } from "solid-js";
import { isAuthenticated, logout } from "./api/auth";

const Home = lazy(() => import("./components/Home"));
const Login = lazy(() => import("./components/Login"));
//const Signup = lazy(() => import('./components/Signup'));
//const Profile = lazy(() => import('./components/Profile'));
const Chat = lazy(() => import("./components/Chat"));

function App() {
  const navigate = useNavigate();
  createEffect(() => {
    if (!isAuthenticated()) {
      navigate("/", { replace: true });
      console.log("user not authenticated");
    }
  });

  const handleLogout = () => {
    logout();
    console.log("logout");
  };

  return (
    <>
      <Show when={isAuthenticated()}>
        <button
          class="text-blue-500 hover:bg-red-700 float-right"
          onClick={handleLogout}
        >
          Logout
        </button>
      </Show>
      <Routes>
        <Route path="/" component={Home} />
        <Route path="/login" component={Login} />
        <Route path="/chat" component={Chat} />
      </Routes>
    </>
  );
}

export default App;
