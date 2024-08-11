import axios from "axios";
import { ref } from "vue";

const GetExpenses = () => {
    const errorExpense = ref(null);
    const expenseList = ref({})
    const expenses = async () => {
        try {
            const id = localStorage.getItem('userId');
            const accountId = localStorage.getItem('accountId');
            console.log(id);
            const url = `http://localhost:8080/expenses/${accountId}`;
            const config = {
                headers: { "Content-Type": "application/json" },
            };
            const result = await axios.get(url, config, {
                withCredentials: true,
            });
            expenseList.value = result.data
        } catch (err) {
            console.error(err.response.data.error);
            errorExpense.value = err.response.data.error;
        }
    };
    return { expenseList, errorExpense, expenses };
};

export default GetExpenses;
