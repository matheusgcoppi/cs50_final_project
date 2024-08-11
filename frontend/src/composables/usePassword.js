import { ref } from "vue";
import axios from "axios";

export const usePassword = () => {
    const error = ref(null);
    const successMessage = ref(null);

    const forgotPassword = async (email) => {
        try {
            const url = "http://localhost:8080/forgot-password";
            const response = await axios.post(url, { email });
            successMessage.value = response.data.message;
            error.value = null;
        } catch (err) {
            error.value = err.response?.data?.error || "An error occurred";
            successMessage.value = null;
        }
    };

    const resetPassword = async (token, password) => {
        try {
            const url = `http://localhost:8080/reset-password/${token}`;
            const response = await axios.post(url, { password });
            successMessage.value = response.data.message;
            error.value = null;
        } catch (err) {
            error.value = err.response?.data?.error || "An error occurred";
            successMessage.value = null;
        }
    };

    return {
        forgotPassword,
        resetPassword,
        error,
        successMessage,
    };
};
