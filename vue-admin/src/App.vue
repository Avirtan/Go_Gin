<script setup lang="ts">
import Header from "@/components/Header.vue"
import LeftNavMenu from "@/components/LeftNavMenu.vue"
import Footer from "@/components/Footer.vue"
import { ref, computed } from 'vue'
import StorageLocal from "@/services/storage-handler"
import { useRouter } from 'vue-router'
import { useStore } from 'vuex'
import UserMutatuinTypes from "@/store/user/user-mutation-types"
import {userKey} from "@/store/user/user-store"

const store = useStore(userKey)
const router = useRouter()
const token = ref("")
const leftDrawerOpen = ref(true)
if (StorageLocal.getData("token")) {
  token.value = StorageLocal.getData("token")
  store.commit(UserMutatuinTypes.Login, true)
} else {
  router.push("/login")
}
let auth = computed(() => store.state.isAuth)
function ShowLeftBar() {
  leftDrawerOpen.value = !leftDrawerOpen.value
}

</script>

<template>
  <q-layout view="hHh Lpr fFf">
    <Header @show="ShowLeftBar" :isAuth="auth" />
    <LeftNavMenu :show="leftDrawerOpen" :isAuth="auth" />
    <q-page-container>
      <q-page>
        <router-view />
      </q-page>
    </q-page-container>
    <Footer />
  </q-layout>
</template>

<style>
</style>
