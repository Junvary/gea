import service from '@/utils/request'

export const createDevopsServer = (data) => {
    return service({
        url: "/devops/createDevopsServer",
        method: 'post',
        data
    })
}

export const deleteDevopsServer = (data) => {
    return service({
        url: "/devops/deleteDevopsServer",
        method: 'delete',
        data
    })
}

export const deleteDevopsServerByIds = (data) => {
    return service({
        url: "/devops/deleteDevopsServerByIds",
        method: 'delete',
        data
    })
}

export const updateDevopsServer = (data) => {
    return service({
        url: "/devops/updateDevopsServer",
        method: 'put',
        data
    })
}

export const findDevopsServer = (params) => {
    return service({
        url: "/devops/findDevopsServer",
        method: 'get',
        params
    })
}

export const getDevopsServerList = (params) => {
    return service({
        url: "/devops/getDevopsServerList",
        method: 'get',
        params
    })
}