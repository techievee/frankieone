package apiServer

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/techievee/frankieone/customLogger/debugcore"
)

const (
	environmentProd = "prod"
)

var Debug bool

// API Server structure
type APIServer struct {
	GinFramework *gin.Engine
	appConfig    *viper.Viper
	Logger       debugcore.Logger
}

func NewRestAPI(env string, appConfig *viper.Viper, logger debugcore.Logger) *APIServer {
	gin.SetMode(env)
	ginFramework := gin.New()

	// By default gin.DefaultWriter = os.Stdout
	ginFramework.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	if appConfig.GetBool("app.service.enable_recovery") {
		ginFramework.Use(gin.Recovery())
	}

	return &APIServer{
		GinFramework: ginFramework,
		appConfig:    appConfig,
		Logger:       logger,
	}

}

// Start the TLS Server
func (s *APIServer) StartTLSServer() {

	tlsServer := s.appConfig.GetString("app.service.tls.host") + ":" + s.appConfig.GetString("app.service.tls.port")

	err := s.GinFramework.RunTLS(tlsServer, s.appConfig.GetString("app.service.tls.certificate"), s.appConfig.GetString("app.service.tls.key"))
	if err != nil {
		s.Logger.Error("Cannot start the TLS Server", "error", err)
	}
}
