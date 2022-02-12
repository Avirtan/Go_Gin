import { createStore, Store} from "vuex";
import { InjectionKey } from 'vue'
import types from "./user-mutation-types";
import ProfileService from "@/services/profielService";
import AuthService from "@/services/auth-service";
import Login from "@/types/login";
import UserType from "@/types/user";
import StorageLocal from "@/services/storage-handler";

export interface User {
    isAuth: boolean;
    Users: Array<UserType>;
}

export const userKey: InjectionKey<Store<User>> = Symbol()

export const userStore = createStore<User>(
    {
    state: {
        isAuth: false,
        Users: [],
    },
    getters: {
        getUsers(state) {
            return state.Users;
        },
    },
    mutations: {
        [types.Login](state) {
            state.isAuth = true;
        },
        [types.Logout](state) {
            state.isAuth = false;
        },
        [types.Users](state, users: Array<UserType>) {
            state.Users = users;
        },
    },
    actions: {
        async [types.Login](context, data: Login) {
            try {
                const resp = await AuthService.login(data);
                if (resp.status == 200) {
                    context.commit(types.Login, true);
                    StorageLocal.setData("token", resp.data.token);
                    return Promise.resolve(resp);
                }
            } catch (error) {
                return Promise.reject(error);
            }
        },
        [types.Logout](context) {
            StorageLocal.delKey("token");
            context.commit(types.Logout);
        },
        async [types.Users](context) {
            try {
                const resp = await ProfileService.GetUsers();
                if (resp.status == 200) {
                    context.commit(types.Users, resp.data.users);
                    return Promise.resolve(resp);
                }
            } catch (error) {
                return Promise.reject(error);
            }
        },
    },
});
