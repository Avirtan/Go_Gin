<script setup lang="ts">
import { useStore } from 'vuex'
import UserMutatuinTypes from "@/store/user/user-mutation-types"
import { useRouter } from 'vue-router'
import { userKey } from "@/store/user/user-store"

const router = useRouter()
const store = useStore(userKey)
const emits = defineEmits(["show"])
const props = defineProps({ isAuth: Boolean })
function Show() {
    emits("show")
}

function Exit() {
    store.dispatch(UserMutatuinTypes.Logout)
    router.push("/login")
}

function Home() {
    if (props.isAuth) {
        router.push("/")
    } else {
        router.push("/login")
    }
}

</script>

<template>
    <q-header elevated class="bg-primary text-white" id="header">
        <q-toolbar>
            <q-btn dense flat round icon="menu" @click="Show" v-show="props.isAuth" />
            <q-toolbar-title class="text-center" @click="Home">
                <a href="#">Work-Craft Administration</a>
            </q-toolbar-title>
            <q-btn flat dense v-show="props.isAuth" @click="Exit">Logout</q-btn>
        </q-toolbar>
    </q-header>
</template>

<style scoped>
a {
    color: inherit;
    text-decoration: none;
}
</style>