import { Routes, Route, A } from "@solidjs/router";
import { createStore } from "solid-js/store";
import { createSignal, createEffect, lazy } from "solid-js";
import { isLoggedIn } from "./api/auth";

import { setAuthToken } from "./api/client";

const Home = lazy(() => import("./components/Home"));
const Login = lazy(() => import("./components/Login"));
//const Signup = lazy(() => import('./components/Signup'));
//const Profile = lazy(() => import('./components/Profile'));
const Chat = lazy(() => import("./components/Chat"));

function App() {
  const [store, setStore] = createStore({
    user: null,
  });
  const [loading, setLoading] = createSignal(false);

  createEffect(() => {
    const token = localStorage.getItem("token");
    if (token) {
      setAuthToken(token);
      isLoggedIn()
        .then((user) => {
          setStore(user);
        })
        .catch((error) => {
          console.error("Error checking authentication:", error);
          localStorage.removeItem("token");
          setAuthToken(null);
          setStore("user", null);
        })
        .finally(() => {
          setLoading(false);
        });
    } else {
      setLoading(false);
    }
  });

  if (loading()) {
    return <p>Loading...</p>;
  }

  return (
    <>
      <Routes>
        <Route path="/" component={Home} />
        <Route path="/login" component={Login} />
        <Route path="/chat" component={Chat} />
      </Routes>
    </>
  );
}

export default App;
