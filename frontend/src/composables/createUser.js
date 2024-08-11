import { ref } from "vue";
import axios from "axios";
import router from "@/router";

const createUser = () => {
  const user = ref(null);
  const error = ref(null);

  const create = async (username, email, password) => {
    try {
      const url = "http://localhost:8080/user";
      const config = {
        headers: { "Content-Type": "application/json" },
      };
      const data = {
        type: 2,
        username: username.value,
        email: email.value,
        password: password.value,
      };
      const response = await axios.post(url, data, config);
      console.log(response.data);

      router.push('/login');
    } catch (err) {
      console.error(err.response.data.error);
      error.value = err.response.data.error;

    }
  };
  return { create, user, error };
};

export default createUser;
