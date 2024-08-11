import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import axios from "axios";
import "vuetify/styles";
import { createVuetify } from "vuetify";
import { library } from "@fortawesome/fontawesome-svg-core";
import { faCheck } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";

library.add(faCheck); // Remember to add here all icons that you want to use

axios.defaults.withCredentials = true; //it makes the cookie send from back-end

createApp(App)
  .use(router)
  .use(createVuetify())
  .component("font-awesome-icon", FontAwesomeIcon)
  .mount("#app");
