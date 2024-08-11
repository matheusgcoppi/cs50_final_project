import axios from "axios";
import { ref } from "vue";

const GetIncomes = () => {
    const error = ref(null);
    const incomeList = ref({})
    const incomes = async () => {
        try {
            const id = localStorage.getItem('userId');
            const accountId = localStorage.getItem('accountId');
            const url = `http://localhost:8080/incomes/${accountId}`;
            const config = {
                headers: { "Content-Type": "application/json" },
            };
            const result = await axios.get(url, config, {
                withCredentials: true,
            });
            incomeList.value = result.data
        } catch (err) {
            console.error(err.response.data.error);
            error.value = err.response.data.error;
        }
    };
    return { incomeList, error, incomes };
};

export default GetIncomes;
