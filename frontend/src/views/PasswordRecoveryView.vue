<template>
  <div class="container">
    <form @submit.prevent="handleForgotPassword">
      <h2>Forgot Password?</h2>
      <label for="email">Email:</label>
      <input
          type="email"
          id="email"
          required
          v-model="email"
          placeholder="Email"
      />
      <div v-if="error" class="error">{{ error }}</div>
      <div v-if="successMessage" class="success">{{ successMessage }}</div>
      <div class="forgot-password">
        <a href="/login">Return to login</a>
      </div>
      <div>
        <button type="submit">Reset Password</button>
      </div>
    </form>
  </div>
</template>

<script>
import { ref } from "vue";
import { usePassword } from "@/composables/usePassword";

export default {
  setup() {
    const email = ref("");
    const { forgotPassword, error, successMessage } = usePassword();

    const handleForgotPassword = async () => {
      await forgotPassword(email.value);
    };

    return {
      email,
      handleForgotPassword,
      error,
      successMessage,
    };
  },
};
</script>

<style scoped>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

.container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
}

label {
  display: flex;
  align-items: flex-start;
  font-weight: bold;
  margin-bottom: 10px;
}

form {
  display: block;
  max-height: 450px;
  height: 100%;
  max-width: 600px;
  width: 100%;
  padding: 20px;
  background-color: white;
  border-radius: 5px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
}

input {
  display: inline-block;
  width: 100%;
  padding: 10px;
  border: 1px solid #ccc;
  border-radius: 5px;
  box-sizing: border-box;
  margin-bottom: 10px;
}

button {
  margin-top: 20px;
  width: 100%;
  padding: 13px;
  background-color: #3498db;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  transition: background-color 0.3s;
  font-weight: bold;
}

button:hover {
  background-color: #2980b9;
}

.forgot-password {
  margin-top: 5px;
  margin-bottom: 5px;
  display: flex;
  align-content: flex-start;
  font-weight: bold;
}

.forgot-password a {
  text-decoration: none;
  color: #479cd9;
}

.error {
  color: red;
  margin-bottom: 10px;
}

.success {
  color: green;
  margin-bottom: 10px;
}
</style>
