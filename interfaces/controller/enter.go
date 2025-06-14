package controller

import (
	fileAppsrv "server/application/service/file"
	userAppSrv "server/application/service/user"
	"server/interfaces/controller/file"
	"server/interfaces/controller/sys"
	"server/interfaces/controller/user"
)

type apiGroup struct {
	Sys  sys.EndpointCtl
	User user.EndpointCtl
	File file.EndpointCtl
}

var APIs *apiGroup

func InitSrvInject(userSrv userAppSrv.Service, fileSrv fileAppsrv.Service) {
	APIs = &apiGroup{
		Sys:  sys.EndpointCtl{Srv: userSrv},
		User: user.EndpointCtl{Srv: userSrv},
		File: file.EndpointCtl{Srv: fileSrv},
	}
}
