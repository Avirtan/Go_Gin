<script setup lang="ts">import { onMounted, ref } from 'vue';
import ProfileService from "@/services/profielService";
import Role from "@/types/role"
import { useQuasar } from 'quasar'

const $q = useQuasar()
let roles = ref([] as Role[])
let selected = ref([])
onMounted(async () => {
    try {
        const resp = await ProfileService.GetRoles()
        if (resp.status == 200) {
            roles.value = resp.data.roles
        }
    } catch (error) {
        console.log(error)
    }
});

const columns = ref([
    { name: 'ID', field: "ID", label: "ID" },
    { name: 'Name', label: 'Name', field: 'Name' },
])

const name = ref(null)

async function CreateRole() {
    const data = {
        "name": name.value,
    }
    try {
        const resp = await ProfileService.CreateRole(data)
        if (resp.status == 200) {
            $q.notify({
                type: 'positive',
                message: 'Role Add'
            })
            console.log(resp)
            roles.value.push(resp.data.role)
            name.value = null
        }
    } catch (error) {
        console.log(error.response?.data);
        $q.notify({
            type: 'negative',
            message: 'Error create role'
        })
    }
}

async function DeleteRole() {
    for (let role of selected.value as Role[]) {
        const data = {
            "id": role.ID,
        }
        try {
            const resp = await ProfileService.DeleteRole(data)
            if (resp.status == 200) {
                roles.value = roles.value.filter((value) => value.ID != role.ID);
            }
        } catch (error) {
            console.log(error.response);
        }
    }
}

function onReset() {
    name.value = null
}
</script>

<template>
    <div class="row">
        <div class="q-pa-md col-4">
            <q-table
                title="Roles"
                :rows="roles"
                :columns="columns"
                v-model:selected="selected"
                row-key="ID"
                selection="multiple"
            >
                <template v-slot:top>
                    <q-btn class="q-ml-sm" color="primary" label="Remove Role" @click="DeleteRole" />
                </template>
            </q-table>
        </div>
        <div class="q-pa-md col-3">
            <q-form @submit="CreateRole" @reset="onReset" class="q-gutter-md">
                <q-input
                    filled
                    v-model="name"
                    label="Name role"
                    lazy-rules
                    :rules="[val => val && val.length > 0 || 'Please type name']"
                />
                <div>
                    <q-btn label="Create" type="submit" color="primary" />
                    <q-btn label="Reset" type="reset" color="primary" flat class="q-ml-sm" />
                </div>
            </q-form>
        </div>
    </div>
</template>

<style>
</style>