import {defineStore} from "pinia"
import {ref} from 'vue'
import {dictInfoApi} from "@/api/traffic";


export const useIspStore = defineStore("ispStore", () => {
    const options = ref([])

    const isInitialized = ref(false)


    async function sync() {
        const resp = await dictInfoApi('ispSelect')
        if (resp.code === 0) {
            options.value = resp.data?.map((item) => (
                {
                    ...item,
                    value: item.name
                }
            ))
            isInitialized.value = true
        }
    }

    async function init() {
        if (isInitialized.value) {
            return
        }
        await sync()
    }

    async function getOptions() {
        await init()
        return options.value
    }


    return {
        getOptions,
    }
})
