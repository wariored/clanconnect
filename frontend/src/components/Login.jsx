import { createSignal } from "solid-js";
import { useNavigate } from "@solidjs/router";
import { request } from "../api/client";
import { getAuthUser, login, setAuthUser } from "../api/auth";

function Login() {
  const [username, setUsername] = createSignal("");
  const [password, setPassword] = createSignal("");
  const [error, setError] = createSignal("");
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();
try {
      const response = await request("POST", "/login", {
        username: username(),
        password: password(),
      });
      login(response.token);
      console.log(response);

      // Fetch user data and update store
      const userData = await request("GET", `/users/${response.id}`);
      console.log(userData);
      setAuthUser(userData);
      console.log(getAuthUser());

      navigate("/chat", { replace: true });
    } catch (err) {
      console.log(err);
      setError(err?.message);
    }
  };

  return (
    <div class="h-screen">
      <form onSubmit={handleSubmit}>
        <div class="flex flex-col items-center justify-center lg:justify-start">
          <div class="my-5">
            <h1 class="text-lg font-medium">Login</h1>
          </div>
          {error() && <div>{error()}</div>}
          <div class="mb-4">
            <div>
              <label class="font-medium" for="username">
                Username
              </label>
              <input
                type="text"
                id="username"
                class="form-input mx-4 rounded-full"
                value={username()}
                onInput={(e) => setUsername(e.target.value)}
              />
            </div>
            <div>
              <label class="font-medium" htmlFor="password">
                Password
              </label>
              <input
                type="password"
                id="password"
                class="form-input mx-4 my-3 rounded-full"
                value={password()}
                onInput={(e) => setPassword(e.target.value)}
              />
            </div>
          </div>
          <div class="flex items-center">
            <button
              class="mt-3 mx-4 bg-sky-400 rounded-full p-3 w-40 font-medium text-white"
              type="submit"
            >
              Submit
            </button>
          </div>
        </div>
      </form>
    </div>
  );
}

export default Login;
