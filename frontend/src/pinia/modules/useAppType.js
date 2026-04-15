import { defineStore } from "pinia"
import { ref } from 'vue'
import { dictInfoApi } from "@/api/traffic";


export const useAppTypeStore = defineStore("appTypeStore", () => {
    const options = ref([])

    const isInitialized = ref(false)


    async function sync() {
        const resp = await dictInfoApi('appTypeIdTree')
        if (resp.code === 0) {
            options.value = resp.data
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

    // 获取选项列表，不包含子选项
    async function getOptionExcluseChildren() {
        await init()
        return options.value?.map(item => {
            const { children, ...rest } = item;
            return rest
        })
    }


    return {
        getOptions,
        getOptionExcluseChildren,
    }
})
