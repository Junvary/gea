import service from '@/utils/request'

export const create{{.StructName}} = (data) => {
    return service({
        url: "/{{.Abbreviation}}/create{{.StructName}}",
        method: 'post',
        data
    })
}

export const delete{{.StructName}} = (data) => {
    return service({
        url: "/{{.Abbreviation}}/delete{{.StructName}}",
        method: 'delete',
        data
    })
}

export const delete{{.StructName}}ByIds = (data) => {
    return service({
        url: "/{{.Abbreviation}}/delete{{.StructName}}ByIds",
        method: 'delete',
        data
    })
}

export const update{{.StructName}} = (data) => {
    return service({
        url: "/{{.Abbreviation}}/update{{.StructName}}",
        method: 'put',
        data
    })
}

export const find{{.StructName}} = (params) => {
    return service({
        url: "/{{.Abbreviation}}/find{{.StructName}}",
        method: 'get',
        params
    })
}

export const get{{.StructName}}List = (params) => {
    return service({
        url: "/{{.Abbreviation}}/get{{.StructName}}List",
        method: 'get',
        params
    })
}