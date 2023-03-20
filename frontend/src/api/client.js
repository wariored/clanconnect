import { apiBaseURL } from '../config';
import {getAuthToken} from './auth';



export const request = async (method, path, data = null) => {
  const url = `${apiBaseURL}${path}`
  const token = getAuthToken();
  const headers = new Headers({
    'Content-Type': 'application/json',
  });
  if (token) {
    headers.set('Authorization', `Bearer ${token}`);
  }

  const response = await fetch(url, {
    method,
    headers,
    body: data && JSON.stringify(data),
  });
  if (!response.ok) {
    const error = await response.json();
    throw new Error(error.message);
  }
  return response.json();
};
