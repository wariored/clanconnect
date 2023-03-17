import { apiBaseURL } from '../config';

let authToken;

export const setAuthToken = (token) => {
  authToken = token;
};

export const request = async (path, options = {}) => {
  const response = await fetch(`${apiBaseURL}/${path}`, {
    headers: {
      'Content-Type': 'application/json',
      Authorization: authToken ? `Bearer ${authToken}` : '',
    },
    ...options,
  });
  if (!response.ok) {
    const error = await response.json();
    throw new Error(error.message);
  }
  return response.json();
};

