package hera_server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	v1_hera "github.com/zeus-fyi/chatgpt/api/v1"
	"github.com/zeus-fyi/zeus/pkg/hera"
)

var cfg = Config{}
var (
	env          string
	openAiApiKey string
)

func Hera() {
	cfg.Host = "0.0.0.0"
	srv := NewHeraServer(cfg)
	// Echo instance
	srv.E = v1_hera.Routes(srv.E)

	hera.InitHeraOpenAI(openAiApiKey)

	if env == "local" {
		srv.E.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     []string{"http://localhost:3000"},
			AllowMethods:     []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.OPTIONS},
			AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, echo.HeaderAccessControlAllowHeaders, "X-CSRF-Token", "Accept-Encoding"},
			AllowCredentials: true,
		}))
	}
	srv.Start()
}

func init() {
	viper.AutomaticEnv()
	Cmd.Flags().StringVar(&cfg.Port, "port", "9000", "server port")
	Cmd.Flags().StringVar(&env, "env", "local", "environment")
	Cmd.Flags().StringVar(&openAiApiKey, "openai-api-key", "", "openai api key")
}

// Cmd represents the base command when called without any subcommands
var Cmd = &cobra.Command{
	Use:   "Chatgpt server",
	Short: "Chatgpt server with markdown",
	Run: func(cmd *cobra.Command, args []string) {
		Hera()
	},
}
