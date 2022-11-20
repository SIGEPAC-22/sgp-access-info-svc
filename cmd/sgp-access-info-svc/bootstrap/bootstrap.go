package bootstrap

import (
	"database/sql"
	"fmt"
	"github.com/dimiro1/health"
	kitlog "github.com/go-kit/log"
	_ "github.com/go-sql-driver/mysql"
	goconfig "github.com/iglin/go-config"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"os/signal"
	handler2 "sgp-access-info-svc/internal/getDocumentType/platform/handler"
	mysql2 "sgp-access-info-svc/internal/getDocumentType/platform/storage/mysql"
	service3 "sgp-access-info-svc/internal/getDocumentType/service"
	"sgp-access-info-svc/internal/getInfoPersonal/platform/handler"
	"sgp-access-info-svc/internal/getInfoPersonal/platform/storage/mysql"
	service2 "sgp-access-info-svc/internal/getInfoPersonal/service"
	handler4 "sgp-access-info-svc/internal/getOneInfoPersonal/platform/handler"
	mysql4 "sgp-access-info-svc/internal/getOneInfoPersonal/platform/storage/mysql"
	"sgp-access-info-svc/internal/getOneInfoPersonal/service"
	handler3 "sgp-access-info-svc/internal/getPersonalSex/platform/handler"
	mysql3 "sgp-access-info-svc/internal/getPersonalSex/platform/storage/mysql"
	service4 "sgp-access-info-svc/internal/getPersonalSex/service"

	"syscall"
)

func Run() {
	config := goconfig.NewConfig("./application.yaml", goconfig.Yaml)
	port := config.GetString("server.port")

	var kitlogger kitlog.Logger
	kitlogger = kitlog.NewJSONLogger(os.Stderr)
	kitlogger = kitlog.With(kitlogger, "time", kitlog.DefaultTimestamp)

	mux := http.NewServeMux()
	errs := make(chan error, 2)
	////////////////////////////////////////////////////////////////////////
	////////////////////////CORS///////////////////////////////////////////
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodPost,
			http.MethodGet,
			http.MethodPut,
			http.MethodDelete,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})

	handlerCORS := cors.Handler(mux)
	////////////////////////CORS///////////////////////////////////////////

	db, err := sql.Open("mysql", getStrConnection())
	if err != nil {
		log.Fatalf("unable to open database connection %s", err.Error())
	}

	/////////////////////GET PERSONAL INFO/////////////////////
	repoGetOnePersonalInfo := mysql4.NewGetOneInfoPersonalRepository(db, kitlogger)
	serviceGetOneInfoPersonal := service.NewGetOneInfoPersonalSvc(repoGetOnePersonalInfo, kitlogger)
	endpointGetOneInfoPersonal := handler4.MakeGetOneInfoPersonalEndpoint(serviceGetOneInfoPersonal)
	endpointGetOneInfoPersonal = handler4.GetOneInfoPersonalTransportMiddleware(kitlogger)(endpointGetOneInfoPersonal)
	transportGetOneInfoPersonal := handler4.NewGetOneInfoPersonalHandler(config.GetString("paths.getOnePersonalInfo"), endpointGetOneInfoPersonal)
	/////////////////////GET PERSONAL INFO/////////////////////

	/////////////////////GET PERSONAL INFO/////////////////////
	repoGetPersonalInfo := mysql.NewGetInfoPersonalRepository(db, kitlogger)
	serviceGetInfoPersonal := service2.NewGetInfoPersonalSvc(repoGetPersonalInfo, kitlogger)
	endpointGetInfoPersonal := handler.MakeGetInfoPersonalEndpoint(serviceGetInfoPersonal)
	endpointGetInfoPersonal = handler.GetInfoPersonalTransportMiddleware(kitlogger)(endpointGetInfoPersonal)
	transportGetInfoPersonal := handler.NewGetInfoPersonalHandler(config.GetString("paths.getPersonalInfo"), endpointGetInfoPersonal)
	/////////////////////GET PERSONAL INFO/////////////////////

	/////////////////////GET DOCUMENT TYPE/////////////////////
	repoGetDocumentType := mysql2.NewGetDocumentTypeRepo(db, kitlogger)
	serviceGetDocumentType := service3.NewGetDocumentTypeSvc(repoGetDocumentType, kitlogger)
	endpointGetDocumentType := handler2.MakeGetDocumentTypeEndpoints(serviceGetDocumentType)
	endpointGetDocumentType = handler2.GetDocumentTypeMiddleware(kitlogger)(endpointGetDocumentType)
	transportGetDocumentType := handler2.NewGetDocumentTypeHandler(config.GetString("paths.getDocumentType"), endpointGetDocumentType)

	/////////////////////GET DOCUMENT TYPE/////////////////////
	repoGetPersonalSex := mysql3.NewGetPersonalSexRepo(db, kitlogger)
	serviceGetPersonalSex := service4.NewGetPersonalSexSvc(repoGetPersonalSex, kitlogger)
	endpointGetPersonalSex := handler3.MakeGetPersonalSexEndpoints(serviceGetPersonalSex)
	endpointGetPersonalSex = handler3.GetPersonalSexMiddleware(kitlogger)(endpointGetPersonalSex)
	transportGetPersonalSex := handler3.NewGetPersonalSexHandler(config.GetString("paths.getPersonalSex"), endpointGetPersonalSex)

	mux.Handle(config.GetString("paths.getOnePersonalInfo"), transportGetOneInfoPersonal)
	mux.Handle(config.GetString("paths.getPersonalInfo"), transportGetInfoPersonal)
	mux.Handle(config.GetString("paths.getDocumentType"), transportGetDocumentType)
	mux.Handle(config.GetString("paths.getPersonalSex"), transportGetPersonalSex)
	mux.Handle("/health", health.NewHandler())

	go func() {
		kitlogger.Log("listening", "transport", "http", "address", port)
		errs <- http.ListenAndServe(":"+port, handlerCORS)
	}()

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		signal.Notify(c, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
		db.Close()
	}()
	kitlogger.Log("terminated", <-errs)
}

func getStrConnection() string {
	config := goconfig.NewConfig("./application.yaml", goconfig.Yaml)
	host := config.GetString("datasource.host")
	user := config.GetString("datasource.user")
	pass := config.GetString("datasource.pass")
	dbname := config.GetString("datasource.dbname")
	strconn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=True", user, pass, host, dbname)
	return strconn
}
