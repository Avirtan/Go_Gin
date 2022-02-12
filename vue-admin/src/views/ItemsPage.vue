<script setup lang="ts">
import { onMounted, ref } from 'vue';
import UploadThing from "@/components/UploadItems.vue"
import DeleteThing from "@/components/DeleteItems.vue"
import ProfileService from "@/services/profielService";

const items = ref([] as any)
const isLoad = ref(false)
onMounted(async () => {
    await UpdateThings()
});
async function UpdateThings() {
    try {
        const resp = await ProfileService.GetItems()
        items.value = resp.data.items
        console.log(resp)
        isLoad.value = true
    } catch (error) {
        console.log(error.response?.data);
    }
}
console.log(import.meta.env.VITE_ADDRES)
</script>

<template>
    <UploadThing />
    <div class="row q-pt-md">
        <DeleteThing class="col-2 q-pl-md" v-if="isLoad" :items="items['eyes']" type="eye" />
        <DeleteThing class="col-2 q-pl-md" v-if="isLoad" :items="items['hair']" type="hair" />
        <DeleteThing class="col-2 q-pl-md" v-if="isLoad" :items="items['mouth']" type="mouth" />
        <DeleteThing class="col-2 q-pl-md" v-if="isLoad" :items="items['nose']" type="nose" />
        <DeleteThing class="col-2 q-pl-md" v-if="isLoad" :items="items['bottom']" type="bottom" />
    </div>
    <div class="row q-pt-md">
        <DeleteThing class="col-2 q-pl-md" v-if="isLoad" :items="items['other']" type="other" />
        <DeleteThing class="col-2 q-pl-md" v-if="isLoad" :items="items['shoes']" type="shoe" />
        <DeleteThing class="col-2 q-pl-md" v-if="isLoad" :items="items['top']" type="top" />
        <DeleteThing class="col-2 q-pl-md" v-if="isLoad" :items="items['weapons']" type="weapon" />
    </div>
</template>

<style>
</style>