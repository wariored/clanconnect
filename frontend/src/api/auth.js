import { apiBaseURL } from '../config';

export const login = async (username, password) => {
  const response = await fetch(`${apiBaseURL}/login`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ username, password }),
  });
  if (!response.ok) {
    const error = await response.json();
    throw new Error(error.message);
  }
  const { token } = await response.json();
  return token;
};

export const isLoggedIn = async () => {
  const response = await fetch(`${apiBaseURL}/me`);
  if (!response.ok) {
    throw new Error('Unauthorized');
  }
  const user = await response.json();
  return user;
};

