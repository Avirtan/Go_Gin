<script setup lang="ts">
import { computed, onBeforeMount, onMounted, ref } from 'vue';
import ProfileService from "@/services/profielService";
import { useQuasar } from 'quasar'

const $q = useQuasar()
const props = defineProps({
    items: Array,
    type: String
})
const select = ref([])
const items = ref()
const label = "Select " + props.type
items.value = props.items
async function DeleteThing() {
    for (let thing of select.value as any[]) {
        const data = {
            "id": thing.ID,
            "type": props.type,
            "name": thing.Name
        }
        try {
            const resp = await ProfileService.DeleteThing(data)
            if (resp.status == 200) {
                items.value = items.value.filter((value: { ID: any; }) => value.ID != thing.ID);
                console.log(resp)
                select.value = []
                $q.notify({
                    type: 'positive',
                    message: 'file delete'
                })
            }
        } catch (error) {
            console.log(error.response);
        }
    }
}
</script>

<template>
    <div>
        <q-select
            option-label="Name"
            filled
            v-model="select"
            multiple
            :options="items"
            use-chips
            stack-label
            :label="label"
        />
        <q-btn label="Delete" color="primary" class="q-mt-sm" @click="DeleteThing" />
    </div>
</template>

<style>
</style>