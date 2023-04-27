package main

import (
	"context"
	"customer/config"
	"customer/ent"
	"customer/internal/logger"
	"customer/middleware"
	"customer/resolver"
	"errors"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
	"strings"
	"time"
)

const configFile = "config.yaml"

// initEnv loads the. env file
func initEnv() {
	if _, err := os.Stat(".env"); errors.Is(err, os.ErrNotExist) {
		fmt.Println("skip load .env file")
		return
	}
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("init env has failed failed with error: %v\n", err)
		os.Exit(1)
	}
}

// initLogger creates a new zap. Logger
func initLogger() *zap.Logger {
	return logger.NewLogger()
}

// initConfig initializes the config.
func initConfig() *config.Configurations {
	viper.SetConfigType("yaml")

	// Expand environment variables inside the config file
	b, err := os.ReadFile(configFile)
	if err != nil {
		fmt.Printf("read config has failed failed with error: %v\n", err)
		os.Exit(1)
	}

	expand := os.ExpandEnv(string(b))
	configReader := strings.NewReader(expand)

	viper.AutomaticEnv()

	if err := viper.ReadConfig(configReader); err != nil {
		fmt.Printf("read config has failed with error: %v\n", err)
		os.Exit(1)
	}

	configs := config.Configurations{}
	if err := viper.Unmarshal(&configs); err != nil {
		fmt.Printf("read config has failed failed with error: %v\n", err)
		os.Exit(1)
	}

	return &configs
}

func main() {
	initEnv()
	initLogger()
	// Connect to postgresql database
	configs := initConfig()
	client, err := ent.Open("postgres", configs.Postgres.ConnectionString)
	if err != nil {
		logger.NewLogger().Error("Getting error connect to postgresql database", zap.Error(err))
		os.Exit(1)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		logger.NewLogger().Error("Failed to creating db schema from the automation migration tool", zap.Error(err))
		os.Exit(1)
	}
	// Create a Gin router instance
	r := gin.Default()

	// Use middlewares
	r.Use(
		middleware.AuthMiddleware(),
		middleware.CorsMiddleware(),
		middleware.RequestCtxMiddleware(),
		ginzap.Ginzap(logger.NewLogger(), time.RFC3339, true),
		ginzap.RecoveryWithZap(logger.NewLogger(), true),
	)

	srv := handler.NewDefaultServer(resolver.NewSchema(client))

	// Define the GraphQL endpoint
	r.POST("/query", gin.WrapH(srv))

	// Define the GraphQL Playground endpoint
	r.GET("/playground", func(c *gin.Context) {
		playground.Handler("GraphQL", "/query").ServeHTTP(c.Writer, c.Request)
	})

	logger.NewLogger().Info("Listening on port: 8000")
	if err := r.Run(":8080"); err != nil {
		logger.NewLogger().Error("Get error from run server", zap.Error(err))
	}

}
