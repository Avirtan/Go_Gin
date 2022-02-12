<script setup lang="ts">
import { ref } from 'vue';


const select = ref("" as string | number | null )
const props = defineProps({
    name:String,
    items:Array
})

const label = ref(props.name + " select")
</script>

<template>
    <q-select
        option-label="Name"
        filled
        :options="props.items"
        use-chips
        stack-label
        :label="label"
        :modelValue="label"
        @update:modelValue="select = $event"
    />
</template>