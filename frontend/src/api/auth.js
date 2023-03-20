import { createStore } from "solid-js/store";
import userStore from '../stores/userStore';

export const getAuthToken = () => localStorage.getItem("token");

const [authStore, setAuthStore] = createStore({
  isAuthenticated: !!getAuthToken()
});

export const login = (token) => {
  localStorage.setItem("token", token);
  setAuthStore({"isAuthenticated": true});
};

export const logout = () => {
  localStorage.removeItem("token");
  setAuthStore({"isAuthenticated": false});
  userStore.set("user", null);
};

export const getAuthUser = () => userStore.value;
export const setAuthUser = (data) => {
  userStore.set(data)
}
// Check if the user is authenticated
export const isAuthenticated = () => authStore.isAuthenticated;

