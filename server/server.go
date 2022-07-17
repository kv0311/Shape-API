package server

import "shape-api/config"

func Init() {
	r := NewRouter()
	r.Run(":" + config.GetConfig("port"))
}
