package upload

import (
	"github.com/HackIllinois/api/common/apiserver"
	"github.com/HackIllinois/api/services/upload/config"
	"github.com/HackIllinois/api/services/upload/controller"
	"github.com/HackIllinois/api/services/upload/service"
	"github.com/gorilla/mux"
	"log"
)

func Initialize() error {
	err := config.Initialize()

	if err != nil {
		return err

	}

	err = service.Initialize()

	if err != nil {
		return err
	}

	return nil
}

func Entry() {
	err := Initialize()

	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()
	controller.SetupController(router.PathPrefix("/upload"))

	log.Fatal(apiserver.StartServer(config.UPLOAD_PORT, router, "upload", Initialize))
}
