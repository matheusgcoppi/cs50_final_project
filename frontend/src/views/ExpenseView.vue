<template>
  <Header :active-page="activePage"></Header>
  <div class="container">
    <form @submit.prevent="handleSubmit">
      <h2>New Expense</h2>
      <div class="form-group">
        <label for="description">Description</label>
        <input
          type="text"
          id="description"
          required
          v-model="description"
          placeholder="Add a description"
        />
      </div>
      <div class="form-group">
        <label for="price">Price</label>
        <CurrencyInput v-model="value" :options="{ currency: 'BRL' }" />
      </div>
      <div class="form-group">
        <label for="date">Date</label>
        <input type="datetime-local" id="date" required v-model="date" />
      </div>
      <div class="form-group">
        <label for="payment">Payment Method</label>
        <select v-model="selected">
          <option disabled value="">Please select one</option>
          <option v-for="option in options" :value="option.value">
            {{ option.text }}
          </option>
        </select>
      </div>
      <div class="button-container">
        <button class="btn-circle">
          <font-awesome-icon :icon="['fas', 'check']" style="font-size: 26px" />
        </button>
      </div>
    </form>
  </div>
</template>

<script>
import { ref } from "vue";
import Header from "@/components/Header.vue";
import CurrencyInput from "./CurrencyInput";
import expense from "@/composables/Expense";
import router from "@/router";
import Expense from "@/composables/Expense";

export default {
  components: { Header, CurrencyInput },
  name: "App",
  setup() {
    const activePage = ref("expense");
    const value = ref(0.0);
    const description = ref("");
    const date = ref(null);
    const selected = ref("");
    const options = ref([
      { text: "Cash", value: 1 },
      { text: "Credit Card", value: 2 },
      { text: "Debit Card", value: 3 },
      { text: "Pix", value: 4 },
      { text: "Bill", value: 5 },
      { text: "Other", value: 6 },
    ]);

    const handleSubmit = async () => {
      const { expense, error } = Expense();
      const formattedDate = new Date(date.value).toISOString();
      console.log(date.value)
      console.log(formattedDate)
      await expense(value, description, formattedDate, selected);
      if (error.value) {
        // Handle the error (e.g., display an error message)
        console.error("Income error:", error.value);
      } else {
        await router.push({ name: "home" });
      }
    };
    return {
      activePage,
      value,
      options,
      description,
      date,
      selected,
      handleSubmit,
    };
  },
};
</script>

<style scoped>
.container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 80vh;
}

form {
  display: block;
  max-height: 550px;
  height: 100vh;
  max-width: 600px;
  width: 100vw;
  background-color: white;
  border-radius: 5px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
  justify-content: start;
}

form h2 {
  margin-top: 18px;
  margin-bottom: 10px;
}

.form-group {
  margin-bottom: 20px;
}

label {
  font-weight: bold;
  margin-left: 30px;
  display: flex;
  justify-content: flex-start;
  margin-bottom: 5px;
}

input,
select {
  width: 90%;
  padding: 10px;
  border: 1px solid #ccc;
  border-radius: 5px;
}

.btn-circle {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 80px;
  height: 80px;
  border-radius: 50%;
  background-color: #007bff;
  color: white;
  border: none;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.btn-circle:hover {
  background-color: #0056b3;
}

.button-container {
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>