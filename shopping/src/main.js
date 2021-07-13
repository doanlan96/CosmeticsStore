import { createApp } from "vue";
import App from "./App.vue";
import router from "@/router/index";
import store from "@/store/index";
import axios from "axios";
import { VueCookieNext } from 'vue-cookie-next'


axios.defaults.baseURL = "http://localhost:8000/api/"

const app = createApp(App);
app.use(router);
app.use(store);
app.use(VueCookieNext)
app.mount("#app");
