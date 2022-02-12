<script setup lang="ts">
import { useStore } from 'vuex'
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import UserMutatuinTypes from "@/store/user/user-mutation-types"
import {userKey} from "@/store/user/user-store"


const store = useStore(userKey)
const login = ref("")
const password = ref("")
const isPwd = ref(true)
const router = useRouter()
async function onSubmit() {
    const data = {
        "login": login.value,
        "password": password.value
    }
    try {
        const resp = await store.dispatch(UserMutatuinTypes.Login, data)
        router.push("/")
    } catch (error) {
        console.log(error.response?.data);
    }
}

function onReset() {
    login.value = ""
    password.value = ""
}
</script>

<template>
    <div class="fit row items-center justify-center" style="min-height: inherit;">
        <q-form @submit="onSubmit" @reset="onReset" class="col-xs-8 col-sm-4">
            <q-input
                class="q-mb-sm"
                filled
                v-model="login"
                label="Your login *"
                lazy-rules
                :rules="[val => val && val.length > 0 || 'Please type login']"
            />
            <q-input
                class="q-mb-sm"
                :type="isPwd ? 'password' : 'text'"
                filled
                v-model="password"
                label="Your password *"
                lazy-rules
                :rules="[val => val && val.length > 0 || 'Please type password']"
            >
                <template v-slot:append>
                    <q-icon
                        :name="isPwd ? 'visibility_off' : 'visibility'"
                        class="cursor-pointer"
                        @click="isPwd = !isPwd"
                    />
                </template>
            </q-input>

            <div>
                <q-btn label="Войти" type="submit" color="primary" />
                <q-btn label="Сброс" type="reset" color="negative" class="q-ml-sm" />
            </div>
        </q-form>
    </div>
</template>

<style scoped>
</style>