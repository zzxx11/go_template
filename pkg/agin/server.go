package agin

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type Server struct {
	logger      *zap.Logger
	controllers []Controller
	server      *http.Server
	engine      *gin.Engine
	options     Options
}

func NewServer(logger *zap.Logger, controllers []Controller) *Server {
	return &Server{
		logger:      logger,
		controllers: controllers,
	}
}

func (s *Server) WithOptions(opt Options) *Server {
	s.options = opt
	return s
}

func (s *Server) Serve(ctx context.Context) error {
	go func() {
		<-ctx.Done()
		s.logger.Info("ready to shutdown http server...", zap.String("", ""))
		_ = s.server.Shutdown(ctx)
		s.logger.Info("http server shutdown")
	}()
	if s.options.ListenAddr == "" {
		panic("server listen address can not be empty.")
	}

	gin.SetMode(gin.ReleaseMode)
	s.engine = gin.New()
	s.engine.Use(GinLogger(s.logger), GinRecovery(s.logger, true))
	// 注册中间件
	for _, middleware := range s.options.middlewares {
		s.engine.Use(middleware)
	}
	// 注册路由
	for _, ctrl := range s.controllers {
		ctrl.InitRouter(s.engine)
	}
	//for _, handler := range s.options.handlers {
	//	s.engine.Handle(handler.method, handler.path, handler.handlerFunc)
	//}

	//routes := s.engine.Routes()
	//s.engine.GET("/__routes", func(c *gin.Context) {
	//	type resp struct {
	//		Endpoint string `json:"endpoint"`
	//		Method   string `json:"method"`
	//		Backends []struct {
	//			Url    string `json:"url"`
	//			Method string `json:"method"`
	//		} `json:"backends"`
	//	}
	//	result := []resp{}
	//	for _, r := range routes {
	//		if r.Path == "/healthz" || r.Path == "/metrics" {
	//			continue
	//		}
	//		path := r.Path
	//		regex, _ := regexp.Compile(":[a-zA-Z_-]+")
	//		path = regex.ReplaceAllStringFunc(path, func(s string) string {
	//			return "{" + s[1:] + "}"
	//		})
	//
	//		result = append(result, resp{
	//			Endpoint: path,
	//			Method:   r.Method,
	//			Backends: []struct {
	//				Url    string `json:"url"`
	//				Method string `json:"method"`
	//			}{
	//				{
	//					Url:    path,
	//					Method: r.Method,
	//				},
	//			},
	//		})
	//	}
	//	c.JSONP(200, map[string]interface{}{"routes": result, "host": "jinjiang-backend-svc"})
	//})

	s.server = &http.Server{
		Addr:    s.options.ListenAddr,
		Handler: s.engine,
	}
	s.logger.Debug("start listen http on " + s.options.ListenAddr)
	// 启动服务监听
	return s.server.ListenAndServe()
}
