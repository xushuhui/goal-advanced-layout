package conf

type Bootstrap struct {
	Server *Server `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	Data   *Data   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	App    *App    `protobuf:"bytes,4,opt,name=app,proto3" json:"app,omitempty"`
}
type App struct {
	Env string `json:"env,omitempty"`
}
type Server struct {
	HTTP *HTTP `protobuf:"bytes,1,opt,name=http,proto3" json:"http,omitempty"`
	Jwt  *Jwt  `protobuf:"bytes,3,opt,name=jwt,proto3" json:"jwt,omitempty"`
}

type Data struct {
	Database *Database `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	Redis    *Redis    `protobuf:"bytes,2,opt,name=redis,proto3" json:"redis,omitempty"`
}

type HTTP struct {
	Addr    string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Timeout uint32 `protobuf:"varint,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

type Jwt struct {
	Secret  string `protobuf:"bytes,1,opt,name=secret,proto3" json:"secret,omitempty"`
	Timeout uint32 `protobuf:"varint,2,opt,name=timeout,proto3" json:"timeout,omitempty"`
	Refresh uint32 `protobuf:"varint,3,opt,name=refresh,proto3" json:"refresh,omitempty"`
}

type Database struct {
	Source string `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
}

type Redis struct {
	Addr     string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Db       int32  `protobuf:"varint,5,opt,name=db,proto3" json:"db,omitempty"`
}
