package apiserver

import (
	"net/http"

	"github.com/wberdnik/CICD_HTTP_Agent/internal/config"
	"github.com/wberdnik/CICD_HTTP_Agent/internal/pkg/services"
	"github.com/sirupsen/logrus"
)

// Start
func Start() error {
	logger := logrus.New()
	level, err := logrus.ParseLevel(logrus.InfoLevel.String())
	if err != nil {
		panic(err)
	}
	logger.SetLevel(level)

	logger.Info(`Start CICD Agent server on port ` + config.HTTP_PORT)
	defer logger.Info(`Stopped CICD Agent server`)

	http.HandleFunc("/"+config.YOUR_NATIVE_TOKEN, services.HandlerAgent)
	return http.ListenAndServe("127.0.0.1:"+config.HTTP_PORT, nil)
}
