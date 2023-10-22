package app

import (
	"github.com/zhigulin-io/user-service/internal/transport"
	"github.com/zhigulin-io/user-service/internal/user"
)

func Run() {
	userStorage := user.NewInMemStorage()
	userService := user.NewService(userStorage)
	restServer := transport.NewRestServer(userService)
	restServer.Start()
}
