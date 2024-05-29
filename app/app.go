package app

import (
	"simple-api/server"
)

func Run() {
	(&server.Server{}).InstanceRun()
}
