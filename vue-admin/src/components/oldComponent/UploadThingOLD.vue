<script setup lang="ts">
import { useQuasar } from 'quasar';
import { ref } from 'vue';
import ProfileService from "@/services/profielService";

const props = defineProps({
    name: String
})
const label: string = props.name + " (8mb)"
const $q = useQuasar()
const fileEye: any = ref(null)
function onRejected(rejectedEntries: string | any[]) {
    $q.notify({
        type: 'negative',
        message: `${rejectedEntries.length} file(s) did not pass validation constraints`
    })
}
async function uploadFileEye() {
    if (fileEye.value == null) {
        $q.notify({
            type: 'negative',
            message: `select file`
        })
        return
    }
    try {
        const fd = new FormData()
        fd.append("img", fileEye.value!)
        const resp = await ProfileService.UploadThuings(fd, props.name!)
        if (resp.status == 200) {
            $q.notify({
                type: 'positive',
                message: `file upload`
            })
            console.log(resp.data)
            fileEye.value = null
        }
    } catch (error) {
        console.log(error.response?.data);
        $q.notify({
            type: 'negative',
            message: error.response?.data
        })
    }
}
</script>

<template>
    <div>
        <q-file
            style="max-width: 300px"
            v-model="fileEye"
            outlined
            accept="image/*"
            :label="label"
            max-file-size="8000000"
            @rejected="onRejected"
        />
        <q-btn label="Upload" type="submit" color="primary" @click="uploadFileEye" />
    </div>
</template>

<style>
</style>