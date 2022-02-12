<script setup lang="ts">
import { ref } from 'vue';
import ProfileService from "@/services/profielService";
import { useQuasar } from 'quasar'

const $q = useQuasar()
const files = ref([])
const options = ["eye", "hair", "mouth", "nose", "other", "bottom", "shoe", "top", "weapon"]
function counterLabelFn(props: { totalSize: string; filesNumber: number; maxFiles: string | number; }) {
    return `${props.filesNumber} files of ${props.maxFiles} | ${props.totalSize}`
}
const emit = defineEmits(['uploadThing'])
const type = ref(null)
async function uploadFileEye() {
    if (files.value == null || type.value == null) {
        return
    }
    try {
        const fd = new FormData()
        // console.log(files.value)
        if (files.value && files.value.length > 0) {
            for (let i = 0; i < files.value.length; i++) {
                fd.append('imgs[]', files.value[i])
            }
        }
        const resp = await ProfileService.UploadThing(fd, type.value)
        if (resp.status == 200) {
            console.log(resp.data)
            files.value = []
            $q.notify({
                type: 'positive',
                message: 'file upload'
            })
        }
    } catch (error) {
        console.log(error.response?.data);
    }
}
function Reset() {
    files.value = []
    type.value = null
}
</script>

<template>
    <div class="row q-pt-md q-pl-md">
        <q-file
            v-model="files"
            label="Pick files (max 10MB)"
            filled
            counter
            use-chips
            :counter-label="counterLabelFn"
            max-files="20"
            multiple
            class="q-mr-md col-4"
        >
            <template v-slot:prepend>
                <q-icon name="attach_file" />
            </template>
        </q-file>
        <q-select v-model="type" :options="options" label="Standard" class="col-4" />
    </div>
    <div class="row q-ml-md">
        <q-btn label="Upload" type="submit" color="primary" class="q-mr-sm" @click="uploadFileEye" />
        <q-btn label="Reset" type="submit" color="primary" @click="Reset" />
    </div>
</template>