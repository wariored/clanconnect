import { createStore } from 'solid-js/store';

const [userStoreValue, setUserStoreValue] = createStore({
  id: '',
  username: '',
  email: '',
  role: ''
});
const userStore = {
  value: userStoreValue,
  set: setUserStoreValue
};

export default userStore;
