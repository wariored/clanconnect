import { A } from "@solidjs/router";
import { createSignal, Show } from "solid-js";
import { isAuthenticated } from "../api/auth";

const Home = () => {
  const [count, setCount] = createSignal(0);

  const handleIncrement = () => {
    setCount(count() + 1);
  };

  const handleDecrement = () => {
    setCount(count() - 1);
  };

  return (
    <div>
      <h1>Home</h1>
      <p>Current count: {count()}</p>
      <button onClick={handleIncrement}>+</button>
      <button onClick={handleDecrement}>-</button>
      <Show when={!isAuthenticated()}>
        <p>
          <A href="/login">Login</A> | <A href="/signup">Signup</A>
        </p>
      </Show>
    </div>
  );
};

export default Home;
