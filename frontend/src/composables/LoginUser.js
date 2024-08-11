import { ref } from "vue";
import axios from "axios";
import router from "@/router";

const LoginUser = () => {
  const user = ref(null);
  const error = ref(null);

  const login = async (email, password) => {
    try {
      const url = "http://localhost:8080/login";
      const config = {
        headers: { "Content-Type": "application/json" },
      };
      console.log(email.value)
      console.log(password.value)
      const data = {
        email: email.value,
        password: password.value,
      };
      const result = await axios.post(url, data, config, {
        withCredentials: true,
      });
      // Set the user information in the ref and in localStorage
      user.value = result.data.user;
      localStorage.setItem('userId', String(result.data.user.id));
      localStorage.setItem('accountId', String(result.data.accountId));
      localStorage.setItem('username', result.data.user.username);

      await router.push('/');
    } catch (err) {
      console.error(err.response.data.error);
      error.value = err.response.data.error;
    }
  };
  return { login, user, error };
};

export default LoginUser;
