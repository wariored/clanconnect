import { createEffect, createSignal } from "solid-js";
import { getAuthUser } from "../api/auth";
import { connectWebSocket } from "../api/client";

const Chat = () => {
  const [messages, setMessages] = createSignal([]);

  const user = getAuthUser();

  createEffect(() => {
    const socket = connectWebSocket("ws://localhost:8080/ws");

    socket.onmessage = (event) => {
      const data = JSON.parse(event.data);
      switch (data.type) {
        case "personal_message":
          setMessages((prevMessages) => [...prevMessages, data.message]);
          break;
        case "user_status":
          console.log(`User ${data.message.user_id} is ${data.message.status}`);
          break;
        case "error":
          console.error(`Error: ${data.message.text}`);
          break;
        default:
          console.error(`Unknown message type: ${data.type}`);
          break;
      }
    };

    return () => {
      console.log("Closing WebSocket connection");
      socket.close();
    };
  }, []);

  return (
    <div>
      <h1>Welcome to the chat, {user.username}!</h1>
      <ul>
        {messages().map((message) => (
          <li key={message.id}>
            {message.recipient}: {message.text}
          </li>
        ))}
      </ul>
    </div>
  );
};

export default Chat;
