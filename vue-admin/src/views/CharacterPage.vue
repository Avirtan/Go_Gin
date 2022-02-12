<script setup lang="ts">
import { onMounted, ref } from 'vue';
import ProfileService from "@/services/profielService";
import { useRouter } from 'vue-router'
import  CreateCharacter from "@/components/CreateCharacter.vue";

const arr = [] as any
const router = useRouter()
const columns = ref([
    { name: 'Login', field: "Login", label: "Login", headerStyle: 'max-width: 100px', style: 'max-width: 10px', },
    { name: 'ID', label: 'ID', field: 'ID' },
])
let selected = ref([])
const characters = ref([] as any)
onMounted(async () => {
    try {
        const resp = await ProfileService.GetCharacters()
        if (resp.status == 200) {
            for (let char of resp.data.characters) {
                arr.push({
                    "ID": char.ID,
                    "Login": "",
                    "UserID":char.UserID
                })
            }
            for (let item of arr) {
               const user = resp.data.users.filter((user:any) => user.id == item.UserID)
               item.Login = user[0].login
            }
            characters.value = arr
        }
    } catch (error) {
        console.log(error)
    }
});


//@row-click="ClickRow"/
</script>

<template>
    <div class="row">
        <div class="col-4">
            <q-table
                title="Characters"
                :rows="characters"
                :columns="columns"
                row-key="ID"
                v-model:selected="selected"
                selection="multiple"
            >
            </q-table>
        </div>
        <CreateCharacter />
    </div>
</template>

<style>
</style>