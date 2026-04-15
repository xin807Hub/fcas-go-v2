import {defineStore} from "pinia"
import {ref} from 'vue'
import {userCrowdGroupTreeApi} from "@/api/traffic"


export const useUserTreeStore = defineStore("userTreeStore", () => {
    const groupTree = ref({})
    const crowdTree = ref({})
    const userTree = ref({})

    const isInitialized = ref(false)

    async function syncTree() {
        const res = await userCrowdGroupTreeApi(3)
        if (res["code"] === 0) {
            groupTree.value = extractByDepth(res.data, 1)
            crowdTree.value = extractByDepth(res.data, 2)
            userTree.value = extractByDepth(res.data, 3)

            isInitialized.value = true

            // console.log('groupTree', groupTree.value)
            // console.log('crowdTree', crowdTree.value)
            // console.log('userTree', userTree.value)
            // console.log('isInitialized userTreeStore', isInitialized.value)
        }
    }

    async function init() {
        if (isInitialized.value) {
            return
        }
        await syncTree()
    }

    async function getUserTree() {
        await init()
        return userTree.value
    }

    async function getCrowdTree() {
        await init()
        return crowdTree.value
    }

    async function getGroupTree() {
        await init()
        return groupTree.value
    }


    return {
        syncTree,
        init,
        getUserTree,
        getCrowdTree,
        getGroupTree,
    }
})

/**
 * @param {Object} node 树节点
 * @param {number} node.id 节点ID
 * @param {string} node.label 节点标签
 * @param {Array} node.children 子节点数组
 * @param {number} depth 目标深度值，从1开始，最大为3
 * @return {Object} 处理后的树结构
 */
function extractByDepth(node, depth) {
    const traverse = (node, curDepth) => {
        if (!node || node?.children?.length === 0 || node?.label === "") return null

        if (curDepth === depth) {
            return {
                id: node.id,
                value: getLastNumber(node.id),
                label: node.label,
                children: null
            }
        }

        return {
            id: node.id,
            value: getLastNumber(node.id),
            label: node.label,
            children: node.children?.map(child => traverse(child, curDepth + 1)).filter(Boolean)
        }
    }

    return traverse(node, 0)
}

function getLastNumber(str) {
    const num = Number(str.split('-').pop());
    return isNaN(num) ? str : num;
}