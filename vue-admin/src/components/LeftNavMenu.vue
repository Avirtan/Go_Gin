<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router'

const router = useRouter()
const miniState = ref(true)
const props = defineProps({show:Boolean, isAuth:Boolean})

function showUser(){
    router.push("/")
}
function showRoles(){
    router.push("/roles")
}
function showItems(){
    router.push("/items")
}
function showCharacter(){
    router.push("/characters")
}
</script>

<template>
    <q-drawer
        v-show="props.isAuth"
        show-if-above
        v-model="props.show"
        :mini="miniState"
        @mouseover="miniState = false"
        @mouseout="miniState = true"
        :width="200"
        :breakpoint="500"
        bordered
        class="bg-grey-3"
    >
        <q-scroll-area class="fit">
            <q-list padding>
                <q-item clickable v-ripple @click="showUser">
                    <q-item-section avatar>
                        <q-icon name="face" />
                    </q-item-section>

                    <q-item-section>Users</q-item-section>
                </q-item>
                <q-item clickable v-ripple @click="showRoles">
                    <q-item-section avatar>
                        <q-icon name="portrait" />
                    </q-item-section>

                    <q-item-section>Roles</q-item-section>
                </q-item>
                <q-item clickable v-ripple @click="showItems">
                    <q-item-section avatar>
                        <q-icon name="image" />
                    </q-item-section>

                    <q-item-section>Items</q-item-section>
                </q-item>
                <q-item clickable v-ripple @click="showCharacter">
                    <q-item-section avatar>
                        <q-icon name="man" />
                    </q-item-section>

                    <q-item-section>Characters</q-item-section>
                </q-item>
            </q-list>
        </q-scroll-area>
    </q-drawer>
</template>

<style>
</style>