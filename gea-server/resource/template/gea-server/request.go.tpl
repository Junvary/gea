package request

import "gin-element-admin/model"

type {{.StructName}}Search struct{
    model.{{.StructName}}
    PageInfo
}