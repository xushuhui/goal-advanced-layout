package conf

import (
	"google.golang.org/protobuf/types/known/durationpb"
)

type Bootstrap struct {
	Server *Server `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	Data   *Data   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	App    *App    `protobuf:"bytes,4,opt,name=app,proto3" json:"app,omitempty"`
}

type App struct {
	Env string `protobuf:"bytes,1,opt,name=env,proto3" json:"env,omitempty"`
}

type Server struct {
	Http *Server_HTTP `protobuf:"bytes,1,opt,name=http,proto3" json:"http,omitempty"`
	Jwt  *Server_Jwt  `protobuf:"bytes,3,opt,name=jwt,proto3" json:"jwt,omitempty"`
}

type Data struct {
	Database *Data_Database `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	Redis    *Data_Redis    `protobuf:"bytes,2,opt,name=redis,proto3" json:"redis,omitempty"`
}

type Server_HTTP struct {
	Addr    string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Timeout uint32 `protobuf:"varint,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

type Server_Jwt struct {
	Secret  string `protobuf:"bytes,1,opt,name=secret,proto3" json:"secret,omitempty"`
	Timeout uint32 `protobuf:"varint,2,opt,name=timeout,proto3" json:"timeout,omitempty"`
	Refresh uint32 `protobuf:"varint,3,opt,name=refresh,proto3" json:"refresh,omitempty"`
}

type Data_Database struct {
	Driver string `protobuf:"bytes,1,opt,name=driver,proto3" json:"driver,omitempty"`
	Source string `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
}

type Data_Redis struct {
	Addr         string               `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Password     string               `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	ReadTimeout  *durationpb.Duration `protobuf:"bytes,3,opt,name=read_timeout,json=readTimeout,proto3" json:"read_timeout,omitempty"`
	WriteTimeout *durationpb.Duration `protobuf:"bytes,4,opt,name=write_timeout,json=writeTimeout,proto3" json:"write_timeout,omitempty"`
	Db           int32                `protobuf:"varint,5,opt,name=db,proto3" json:"db,omitempty"`
}
