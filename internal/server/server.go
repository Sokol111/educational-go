package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"log"
	"net"
	"net/http"
	"strconv"
)

type Handler interface {
	BindRoutes(*gin.Engine)
}

type Server struct {
	httpServer  *http.Server
	baseContext context.Context
}

func NewServer(port int, baseContext context.Context, handlers ...Handler) *Server {
	engine := gin.Default()
	engine.Use(errorHandler)

	for _, h := range handlers {
		h.BindRoutes(engine)
	}

	s := &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: engine,
		BaseContext: func(_ net.Listener) context.Context {
			return baseContext
		},
	}

	return &Server{s, baseContext}
}

func (s *Server) Start() {
	g, gCtx := errgroup.WithContext(s.baseContext)

	g.Go(func() error {
		log.Println("Listening and serving HTTP on", s.httpServer.Addr)
		return s.httpServer.ListenAndServe()
	})

	g.Go(func() error {
		<-gCtx.Done()
		return s.httpServer.Shutdown(context.Background())
	})

	if err := g.Wait(); err != nil {
		log.Printf("exit reason: %s ", err)
	}
}
