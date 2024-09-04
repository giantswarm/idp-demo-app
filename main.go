package main

import (
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/gin-gonic/gin"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func setupPrometheusExporter() {
	// Create a new Prometheus gauge metric
	heartbeatMetric := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "my_webservice_heartbeat",
		Help: "Web Service heartbeat metric",
	})

	// Set the value of the heartbeat metric to 1
	heartbeatMetric.Set(1)

	// Register the metric with the Prometheus default registry
	prometheus.MustRegister(heartbeatMetric)
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong\n")
	})

	router.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello world!\n")
	})

	router.POST("/hello", func(c *gin.Context) {
		val, found := c.GetPostForm("name")
		if !found || val == "" {
			c.String(400, "I need the 'name' parameter")
			return
		}
		c.String(http.StatusOK, "Hello %s!\n", val)
	})

	return router
}

func main() {
	log.Printf("Web service starting, version: '%s', commit: '%s', build date: '%s'", version, commit, date)
	listeningPort := os.Getenv("LISTEN_PORT")
	if listeningPort == "" {
		listeningPort = "8080"
	}

	router := setupRouter()
	setupPrometheusExporter()
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	// Listen and Server on the LocalHost:Port
	err := router.Run(":" + listeningPort)
	if err != nil {
		log.Panicf("Critical error: %s", err)
	}
}
