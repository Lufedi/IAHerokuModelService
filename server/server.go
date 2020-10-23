package server

import "ModelsService/config"

func Init() {
	config := config.GetConfig()
	r := NewRouter()
	r.MaxMultipartMemory = 8 << 20
	r.Run(config.GetString("server.port"))
}