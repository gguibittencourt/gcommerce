package server

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func StartHTTPServer(l fx.Lifecycle) {
	l.Append(fx.Hook{
		OnStart: func(c context.Context) error {
			go func() {
				router := gin.Default()
				if err := router.Run(); err != nil {
					panic(err)
				}
			}()
			return nil
		},
	})
}
