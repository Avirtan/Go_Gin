<script setup lang="ts">
import { useStore } from 'vuex'
import UserMutatuinTypes from "@/store/user/user-mutation-types"
import { onMounted, ref } from 'vue'
import User from "@/types/user"
import { date } from "quasar"
import {userKey} from "@/store/user/user-store"
import { useRouter } from 'vue-router'
import CreateAxios from '@/services/HttpCommon'

const router = useRouter()
const store = useStore(userKey)
let users = ref([] as User[])
let selected = ref([])
const columns = ref([
    { name: 'Id', field: "Id", label: 'Id' },
    { name: 'Login', label: 'Login', field: 'Login' },
    { name: 'Email', field: "Email", label: 'Email' },
    { name: 'Verivication', label: 'Verivication', field: 'Verivication', format: (val: any) => val==undefined?"-":"+" },
    {
        name: 'Create', label: 'Create', field: 'Create', format: (val: any, row: any) => {
            val = date.formatDate(val, 'DD-MM-YYYY HH:mm')
            return val
        }
    },
])
onMounted(async () => {
    try {
        await store.dispatch(UserMutatuinTypes.Users)
        users.value = store.getters.getUsers
    } catch (error) {
        console.log(error.response?.data);
    }
});
function ClickRow(evt: any, row: any, index: any) {
    const user = (row as User)
    router.push({ name: 'User', params: { userID: user.Id } })
}
</script>

<template>
    <div>
        <div class="q-pa-md">
            <q-table
                title="Users"
                :rows="users"
                :columns="columns"
                row-key="Id"
                v-model:selected="selected"
                selection="multiple"
                @row-click="ClickRow"
            />
        </div>
    </div>
</template>

<style>
</style>