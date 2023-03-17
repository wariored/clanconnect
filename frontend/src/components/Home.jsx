import { A } from "@solidjs/router";
import { createSignal } from 'solid-js';

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
      <p>
        <A href="/login">Login</A> | <A href="/signup">Signup</A>
      </p>
    </div>
  );
};

export default Home;
