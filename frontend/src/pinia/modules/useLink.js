import {defineStore} from "pinia"
import {ref} from 'vue'
import {linkListApi} from "@/api/link.js";


export const useLinkStoreV2 = defineStore("linkStoreV2", () => {
    const options = ref([])

    const isInitialized = ref(false)

    async function sync() {
        const res = await linkListApi({page: 1, limit: Number.MAX_SAFE_INTEGER, key: ""})
        if (res["code"] === 0) {
            options.value = res.data.list?.map(item => ({
                    label: item.lineName,
                    value: item.lineVlan
                })
            )
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
        sync,
        getOptions,
    }
})
