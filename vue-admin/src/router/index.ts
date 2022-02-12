import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";
import Login from "../views/Login.vue";
import HomePage from "../views/HomePage.vue";
import RolesPage from "../views/RolesPage.vue";
import UserPage from "../views/UserPage.vue";
import ItemsPage from "../views/ItemsPage.vue";
import StorageLocal from "@/services/storage-handler";
import CharacterPage from "@/views/CharacterPage.vue"
import CreateCharacterPage from "@/components/CreateCharacter.vue"

const routes: Array<RouteRecordRaw> = [
    {
        path: "/login",
        name: "Login",
        component: Login,
    },
    {
        path: "/",
        name: "Home",
        component: HomePage,
    },
    {
        path: "/roles",
        name: "Roles",
        component: RolesPage,
    },
    {
        path: "/user/:userID",
        name: "User",
        component: UserPage,
    },
    {
        path:"/items",
        name:"items",
        component:ItemsPage,
    },
    {
        path:"/characters",
        name:"character",
        component:CharacterPage,
    },
    {
        path:"/character/create",
        name:"characterCreate",
        component:CreateCharacterPage,
    },
];

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes,
});
router.beforeEach((to, from, next) => {
    if(to.fullPath == "/login"){
        console.log("not auth");
        next();
    }
    if (to.fullPath != "/login" && StorageLocal.getData("token")) {
        next();
    }else{
        next({name:"Login"});
    }
    
});
export default router;
