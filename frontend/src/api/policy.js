import service from '@/utils/request'

//业务流量告警分页查询
export const alarmConfigPageApi=(data)=>{
    return service({
        url:'/policy/alarmConfig/page',
        method:'post',
        data
    })
}

//业务流量告警删除
export const deleteAlarmConfigApi=(data)=>{
    return service({
        url:'/policy/alarmConfig/delete',
        method:'post',
        data
    })
}

//业务流量告警新增或修改
export const saveOrUpdateAlarmConfigApi=(data)=>{
    return service({
        url:'/policy/alarmConfig/saveOrUpdate',
        method:'post',
        data
    })
}

// 优先策略
//优先策略分页查询
export const White_policyListApi = (data) => {
    return service({
        url: '/policy/whitePolicy/page',
        method: 'post',
        data
    })
}
//优先策略更新或保存
export const White_createPolicyApi = (data) => {
    return service({
        url: '/policy/whitePolicy/saveOrUpdate',
        method: 'post',
        data
    })
}

//优先策略删除
export const White_deletePolicyApi = (data) => {
    return service({
        url: '/policy/whitePolicy/delete',
        method: 'post',
        data
    })
}

// 策略管控
//策略管控分页查询
export const control_policyListApi = (data) => {
    return service({
        url: '/policy/controlPolicy/page',
        method: 'post',
        data
    })
}
//策略管控更新或保存
export const control_createPolicyApi = (data) => {
    return service({
        url: '/policy/controlPolicy/saveOrUpdate',
        method: 'post',
        data
    })
}

//策略管控删除
export const control_deletePolicyApi = (data) => {
    return service({
        url: '/policy/controlPolicy/delete',
        method: 'post',
        data
    })
}