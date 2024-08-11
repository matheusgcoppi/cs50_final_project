<template>
  <div class="container">
    <form @submit.prevent="handleSubmit">
      <h2>Sign Up</h2>
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
        required
        v-model="password"
        placeholder="Password"
      />
      <label for="password">Confirm Password:</label>
      <input
        type="password"
        required
        v-model="repeatPassword"
        placeholder="Repeat Your Password"
      />
      <p class="error-post"></p>
      <div>
        <button>Sign Up</button>
      </div>
      <hr class="divider" />
      <router-link to="/login">
        <div class="signup">
          <button>Login</button>
        </div>
      </router-link>
    </form>
  </div>
</template>

<script>
import { ref } from "vue";
import createUser from "@/composables/createUser";
import { useRouter } from "vue-router";

export default {
  setup() {
    const username = ref("");
    const email = ref("");
    const password = ref("");
    const repeatPassword = ref("");
    const user = ref(null);
    const error = ref(null);
    const passwordNotMatch = ref(false);
    const router = useRouter();

    const handleSubmit = async () => {
      error.value = null;
      const {
        create,
        user: createUserResult,
        error: createUserError,
      } = createUser();
      user.value = createUserResult;
      error.value = createUserError;
      passwordNotMatch.value = false;

      if (password.value !== repeatPassword.value) {
        const errorPost = document.querySelector(".error-post");
        errorPost.textContent = "Passwords do not match. Please try again.";
        password.value = "";
        repeatPassword.value = "";
        return;
      }

      username.value = email.value.substring(0, email.value.indexOf("@"));

      await create(username, email, password);
      if (error.value !== null) {
        const errorPost = document.querySelector(".error-post");
        errorPost.textContent = error.value._value;
        email.value = "";
        password.value = "";
        repeatPassword.value = "";
        return;
      }

      await router.push({ name: "login" });
    };

    return {
      username,
      email,
      password,
      handleSubmit,
      user,
      error,
      repeatPassword,
      passwordNotMatch,
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
  max-width: 600px; /* Adjust the max-width as needed */
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
  margin-top: 15px;
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

.signup button {
  background-color: white;
  color: #3498db;
  border-radius: 5px;
  border: 2px solid #3498db;
}

.divider {
  margin-top: 15px;
  border: 0;
  border-top: 1px solid #ccc; /* Set the color of the divider */
}

.error-post {
  color: red;
  font-weight: bold;
  display: flex;
  align-content: flex-start;
}
</style>