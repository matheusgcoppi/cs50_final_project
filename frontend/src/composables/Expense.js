import axios from "axios";
import { ref } from "vue";

const Expense = () => {
  const error = ref(null);
  const expense = async (value, description, date, selected) => {
    try {
      const url = "http://localhost:8080/expense";
      const config = {
        headers: { "Content-Type": "application/json" },
      };
      const data = {
        price: value.value,
        description: description.value,
        when: date,
        payment: selected.value,
      };
      await axios.post(url, data, config, {
        withCredentials: true,
      });
    } catch (err) {
      console.error(err.response.data.error);
      error.value = err.response.data.error;
    }
  };
  return { error, expense };
};

export default Expense;
