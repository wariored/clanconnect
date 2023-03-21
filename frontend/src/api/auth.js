import { createStore } from "solid-js/store";

export const getAuthToken = () => localStorage.getItem("token");

const [authStore, setAuthStore] = createStore({
  isAuthenticated: !!getAuthToken(),
});
// Check if the user is authenticated
export const isAuthenticated = () => authStore.isAuthenticated;

const [userStore, setUserStore] = createStore({
  id: "",
  username: "",
  email: "",
  role: "",
});

export const getAuthUser = () => userStore;
export const setAuthUser = (data) => setUserStore(data);


export const login = (token) => {
  localStorage.setItem("token", token);
  setAuthStore({ isAuthenticated: true });
};

export const logout = () => {
  localStorage.removeItem("token");
  setAuthStore({ isAuthenticated: false });
  setAuthUser("user", null);
};

