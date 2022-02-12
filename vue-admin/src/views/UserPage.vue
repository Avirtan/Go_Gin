<script setup lang="ts">
import { onMounted, reactive, ref, toRefs } from 'vue';
import { useRouter } from 'vue-router'
import { useStore } from 'vuex'
import { userKey } from "@/store/user/user-store"
import ProfileService from "@/services/profielService";
import Role from "@/types/role"
interface User {
    ID: string;
    Login: string;
    Email: string;
    Password: string;
    Verification: boolean;
    CreatedAt: Date;
    IconID?: any;
    Resume?: any;
    Statistics?: any;
    Roles?: any;
}

const store = useStore(userKey)
const router = useRouter()
const state = reactive<any>({ user: {} })
const ID = ref(router.currentRoute.value.params.userID as string)
const Roles = ref([] as Role[])

onMounted(async () => {
    try {
        const resp = await ProfileService.GetUserById(ID.value)
        if (resp.status == 200) {
            console.log(resp.data.user)
            state.user = resp.data.user
        }
    } catch (error) {
        console.log(error.response?.data);
    }
    try {
        const resp = await ProfileService.GetRoles()
        if (resp.status == 200) {
            //console.log(resp.data.roles)
            Roles.value = resp.data.roles
        }
    } catch (error) {
        console.log(error.response?.data);
    }
});
async function onSubmit() {
    interface Map {
        [key: string]: string | undefined | unknown
    }
    let data = {} as Map
    for (let [key, value] of Object.entries(state.user)) {
        data[key] = value;
    }
    try {
        const resp = await ProfileService.UpdateUser(data)
        if (resp.status == 200) {
            console.log(resp.data)
            state.user = resp.data.user
        }
    } catch (error) {
        console.log(error.response?.data);
    }
}
</script>

<template>
    <div>
        <div class="q-pa-md" style="max-width: 400px">
            <q-form @submit="onSubmit">
                <q-input
                    filled
                    v-model="state.user.login"
                    label="Your Login"
                    lazy-rules
                    :rules="[val => val && val.length > 0 || 'Please type something']"
                />
                <q-input
                    filled
                    v-model="state.user.email"
                    label="Your Email"
                    type="email"
                    lazy-rules
                    :rules="[val => val && val.length > 0 || 'Please type something']"
                />
                <q-select
                    filled
                    v-model="state.user.roles"
                    multiple
                    :options="Roles"
                    label="Roles"
                    option-label="Name"
                />
                <div>
                    <q-btn label="Update" type="submit" color="primary" class="q-mt-md" />
                </div>
            </q-form>
        </div>
    </div>
</template>

<style>
</style>