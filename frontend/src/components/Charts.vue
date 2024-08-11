<template>
  <div class="chart-container">
    <canvas id="myChartBar"></canvas>
    <canvas id="myChartLine"></canvas>
  </div>
</template>

<script>
import Chart from "chart.js/auto";
import { onMounted, ref } from "vue";
import GetIncomes from "@/composables/getIncomes";
import GetExpenses from "@/composables/getExpenses";

export default {
  name: "FinancialCharts",
  setup() {
    const days = ref(initializeDays());
    const year = ref(initializeYear());

    onMounted(async () => {
      try {
        const { incomeList, error: incomeError, incomes } = GetIncomes();
        await incomes();
        handleApiError(incomeError);

        const { expenseList, error: expenseError, expenses } = GetExpenses();
        await expenses();
        handleApiError(expenseError);

        updateData(incomeList.value, days.value, year.value, "income");
        updateData(expenseList.value, days.value, year.value, "expense");

        renderBarChart(days.value);
        renderLineChart(year.value);
      } catch (err) {
        console.error("Error loading data:", err);
      }
    });

    function initializeDays() {
      return {
        sunday: 0,
        monday: 0,
        tuesday: 0,
        wednesday: 0,
        thursday: 0,
        friday: 0,
        saturday: 0,
      };
    }

    function initializeYear() {
      return {
        january: 0,
        february: 0,
        march: 0,
        april: 0,
        may: 0,
        june: 0,
        july: 0,
        august: 0,
        september: 0,
        october: 0,
        november: 0,
        december: 0,
      };
    }

    function handleApiError(error) {
      if (error && error.value) {
        console.error("API error:", error.value);
      }
    }

    function updateData(list, days, year, type) {
      const weekDays = ["sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"];
      const monthsOfYear = ["january", "february", "march", "april", "may", "june", "july", "august", "september", "october", "november", "december"];

      verifyWeekList(list, weekDays, days, type);
      verifyYearList(list, monthsOfYear, year, type);
    }

    function renderBarChart(days) {
      const ctx = document.getElementById("myChartBar").getContext("2d");
      new Chart(ctx, {
        type: "bar",
        data: {
          labels: ["Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"],
          datasets: [
            {
              label: "Net Income by Day",
              data: [days.sunday, days.monday, days.tuesday, days.wednesday, days.thursday, days.friday, days.saturday],
              backgroundColor: "rgba(75, 192, 192, 0.2)",
              borderColor: "rgba(75, 192, 192, 1)",
              borderWidth: 1,
            },
          ],
        },
        options: {
          responsive: true,
          scales: {
            y: {
              beginAtZero: true,
            },
          },
        },
      });
    }

    function renderLineChart(year) {
      const ctx = document.getElementById("myChartLine").getContext("2d");
      new Chart(ctx, {
        type: "line",
        data: {
          labels: ["January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"],
          datasets: [
            {
              label: "Net Income by Month",
              data: [year.january, year.february, year.march, year.april, year.may, year.june, year.july, year.august, year.september, year.october, year.november, year.december],
              backgroundColor: "rgba(153, 102, 255, 0.2)",
              borderColor: "rgba(153, 102, 255, 1)",
              borderWidth: 1,
              fill: true,
            },
          ],
        },
        options: {
          responsive: true,
          scales: {
            y: {
              beginAtZero: true,
            },
          },
        },
      });
    }

    return {};
  },
};

function verifyWeekList(list, weekDays, days, type) {
  if (!list) return;

  Object.values(list).forEach((val) => {
    const when = new Date(val.when);
    const start = new Date();
    const sevenDaysAgo = new Date(start.setDate(start.getDate() - 7));

    if (sevenDaysAgo <= when) {
      const dayOfWeek = weekDays[when.getDay()];
      days[dayOfWeek] += type === "income" ? val.price : -val.price;
    }
  });
}

function verifyYearList(list, months, year, type) {
  if (!list) return;

  Object.values(list).forEach((val) => {
    const when = new Date(val.when);
    const start = new Date();
    const yearAgo = new Date(start.setDate(start.getDate() - 365));

    if (yearAgo <= when) {
      const month = months[when.getMonth()];
      year[month] += type === "income" ? val.price : -val.price;
    }
  });
}
</script>

<style scoped>
.chart-container {
  display: flex;
  justify-content: space-around;
  align-items: flex-start; /* Align items to the start of the container */
  width: 100vw;
  height: calc(100vh - 100px); /* Adjust height to leave room for the header */
  background-color: #f8f9fa;
  padding: 10px; /* Reduced padding to bring charts higher */
  box-sizing: border-box;
}

canvas {
  max-width: 45%;
  max-height: 80%;
  box-shadow: 0px 4px 10px rgba(0, 0, 0, 0.1);
  border-radius: 10px;
}
</style>

