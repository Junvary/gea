package request

import "gin-element-admin/model"

type DevopsServerSearch struct{
    model.DevopsServer
    PageInfo
}