import { createRouter, createWebHistory } from "vue-router";
import HomeView from "../views/HomeView.vue";
import SignupView from "../views/SignupView.vue";
import LoginView from "../views/LoginView.vue";
import passwordRecoveryView from "../views/PasswordRecoveryView.vue";
import axios from "axios";
import PageNotFoundView from "@/views/PageNotFoundView.vue";
import IncomeView from "@/views/IncomeView.vue";
import ExpenseView from "@/views/ExpenseView.vue";

const routes = [
  {
    path: "/",
    name: "home",
    component: HomeView,
    meta: {
      requiresAuth: true,
    },
  },
  {
    path: "/signup",
    name: "signup",
    component: SignupView,
  },
  {
    path: "/login",
    name: "login",
    component: LoginView,
  },
  {
    path: "/password-recovery",
    name: "password-recovery",
    component: passwordRecoveryView,
  },
  {
    path: "/income",
    name: "income",
    component: IncomeView,
    meta: {
      requiresAuth: true,
    },
  },
  {
    path: "/expense",
    name: "expense",
    component: ExpenseView,
    meta: {
      requiresAuth: true,
    },
  },
  {
    path: "/:pathMatch(.*)*",
    name: "page-not-found",
    component: PageNotFoundView,
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

router.beforeEach(async (to, from, next) => {
  if (to.matched.some((record) => record.meta.requiresAuth)) {
    try {
      const response = await axios.get("http://localhost:8080/validate");
      // console.log(response);
      if (response.data.message === true) {
        next();
      } else {
        next({ name: "login" });
      }
    } catch (e) {
      console.log(e);
      next({ name: "login" });
    }
  } else {
    next();
  }
});

export default router;
