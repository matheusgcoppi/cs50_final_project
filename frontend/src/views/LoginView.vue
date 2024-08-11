<template>
  <div class="container">
    <form @submit.prevent="handleSubmit">
      <h2>Login</h2>
      <label for="email">Email:</label>
      <input
        type="email"
        id="email"
        required
        v-model="email"
        placeholder="Email"
      />
      <label for="password">Password:</label>
      <input
        type="password"
        id="password"
        required
        v-model="password"
        placeholder="Password"
      />
      <div class="forgot-password">
        <a v-bind:href="'/password-recovery'">Forgot Password?</a>
      </div>
      <div>
        <button>Log In</button>
      </div>
      <hr class="divider" />
      <router-link to="/signup">
        <div class="signup">
          <button>Sign Up</button>
        </div>
      </router-link>
    </form>
  </div>
</template>

<script>
import { ref } from "vue";
import LoginUser from "@/composables/LoginUser";

export default {
  setup() {
    const email = ref("");
    const password = ref("");
    const isAuthenticated = ref(false);

    const handleSubmit = async () => {
      const { login } = LoginUser();

      let result = await login(email, password);
      console.log(result);
      if (result) isAuthenticated.value = true;
    };
    return { email, password, handleSubmit };
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

.signup button {
  background-color: white;
  color: #3498db;
  border-radius: 5px;
  border: 2px solid #3498db;
}

.divider {
  margin-top: 20px;
  border: 0;
  border-top: 1px solid #ccc;
}
</style>