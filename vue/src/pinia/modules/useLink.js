import {defineStore} from "pinia"
import {ref} from 'vue'
import {linkListApi} from "@/api/link.js";


export const useLinkStoreV2 = defineStore("useLinkStoreV2", () => {
    const linkOptions = ref([])

    const isInitialized = ref(false)

    async function init() {
        if (isInitialized.value) {
            return
        }
        await fetchLinks()
    }

    async function fetchLinks() {
        const res = await linkListApi({page: 1, limit: Number.MAX_SAFE_INTEGER, key: ""})
        if (res["code"] === 0) {
            linkOptions.value = res.data.list?.map(item => ({
                    'name': item.lineName,
                    'value': item.lineVlan
                })
            )
            isInitialized.value = true

            console.log('linkOptions', linkOptions.value)
            console.log('isInitialized linkStore', isInitialized.value)
        }
    }

    return {
        linkOptions,
        init
    }
})
