import { createApp } from "vue";
import router from "./router";
import {userStore,userKey} from "./store/user/user-store";
import { Quasar, Notify } from "quasar";
// Import icon libraries
import "@quasar/extras/material-icons/material-icons.css";

// Import Quasar css
import "quasar/src/css/index.sass";

import App from "./App.vue";

const app = createApp(App);
app.use(router);
app.use(userStore,userKey);
app.use(Quasar, {
    plugins: { Notify },
});
app.mount("#app");
