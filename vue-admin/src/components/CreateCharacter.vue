<script setup lang="ts">
import CreateCharacterSelecter from "@/components/CreateCharacterSelecter.vue";
import { onMounted, ref } from 'vue';
import ProfileService from "@/services/profielService";

const hair = ref(null)
const eyes = ref(null)
const mouth = ref(null)
const nose = ref(null)
const top = ref(null)
const bottom = ref(null)
const shoes = ref(null)
const weapon = ref(null)
const other = ref(null)
const user = ref(null)
const users = ref([] as any)
const isLoad = ref(false)
const items = ref([] as any)

onMounted(async () => {
    try {
        const resp = await ProfileService.GetItems()
        items.value = resp.data.items
        isLoad.value = true
    } catch (error) {
        console.log(error.response?.data);
    }
    try {
        const resp = await ProfileService.GetUsers();
        if (resp.status == 200) {
            users.value = resp.data.users
        }
    } catch (error) {
        console.log(error.response?.data);
    }
});

</script>

<template>
    <div class="col q-ml-md">
        <CreateCharacterSelecter
            :items="items['hair']"
            name="hair"
            v-model="hair"
            class="col-3 q-mr-md"
        />
        <CreateCharacterSelecter
            :items="items['eyes']"
            name="eye"
            v-model="eyes"
            class="col-2 q-mr-md"
        />
        <CreateCharacterSelecter
            :items="items['mouth']"
            name="mouth"
            v-model="mouth"
            class="col-2 q-mr-md"
        />
        <CreateCharacterSelecter
            :items="items['nose']"
            name="nose"
            v-model="nose"
            class="col-2 q-mr-md"
        />
        <CreateCharacterSelecter
            :items="items['top']"
            name="top"
            v-model="top"
            class="col-2 q-mr-md"
        />
    </div>
    <div class="col q-ml-md">
        <CreateCharacterSelecter
            :items="items['bottom']"
            name="bottom"
            v-model="bottom"
            class="col-2 q-mr-md"
        />
        <CreateCharacterSelecter
            :items="items['shoes']"
            name="shoe"
            v-model="shoes"
            class="col-2 q-mr-md"
        />
        <CreateCharacterSelecter
            :items="items['weapons']"
            name="weapon"
            v-model="weapon"
            class="col-2 q-mr-md"
        />
        <CreateCharacterSelecter
            :items="items['other']"
            name="other"
            v-model="other"
            class="col-2 q-mr-md"
        />
        <q-select
            option-label="Login"
            filled
            :options="users"
            use-chips
            stack-label
            label="Select users"
            v-model="user"
            class="col-2 q-mr-md"
        />
    </div>
    <q-btn color="primary" label="Создать" />
</template>