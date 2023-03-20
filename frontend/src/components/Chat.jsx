import { createStore } from "solid-js/store"; 
import { createSignal } from 'solid-js';

const Chat = () => {
  const [messages, setMessages] = createSignal([]);

  useEffect(() => {
    const socket = new WebSocket(`ws://localhost:8080/ws`);
    
    socket.onopen = () => {
      console.log('WebSocket connected');
    };

    socket.onclose = () => {
      console.log('WebSocket disconnected');
    };

    socket.onmessage = (event) => {
      const message = JSON.parse(event.data);
      switch (message.type) {
        case 'personal_message':
          setMessages((prevMessages) => [...prevMessages, message.message]);
          break;
        case 'user_status':
          console.log(`User ${message.message.user_id} is ${message.message.status}`);
          break;
        case 'error':
          console.error(`Error: ${message.message.message}`);
          break;
        default:
          console.error(`Unknown message type: ${message.type}`);
          break;
      }
    };

    return () => {
      console.log('Closing WebSocket connection');
      socket.close();
    };
  }, []);

  return (
    <div>
      <h1>Welcome to the chat, {user.username}!</h1>
      <ul>
        {messages().map((message) => (
          <li key={message.id}>{message.sender_id}: {message.text}</li>
        ))}
      </ul>
    </div>
  );
};

export default Chat;

