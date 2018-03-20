package api

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/clyhs/goweb/controller/models"
	"github.com/clyhs/goweb/controller/setting"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/mailgun/oxy/forward"

	_ "github.com/go-sql-driver/mysql"
)

type (
	Api struct {
		listenAddr string
		dUrl       string
		fwd        *forward.Forwarder
	}
	ApiConfig struct {
		ListenAddr string
	}
)

func writeCorsHeaders(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	w.Header().Add("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, OPTIONS")
}

func NewApi(config ApiConfig) (*Api, error) {
	return &Api{
		listenAddr: config.ListenAddr,
	}, nil
}

func (a *Api) Run() error {

	setting.LoadConfig()

	models.Init(false)

	globalMux := http.NewServeMux()

	apiRouter := mux.NewRouter()
	apiRouter.HandleFunc("/module/hello", a.Hello).Methods("GET")
	apiRouter.HandleFunc("/module/user", a.User).Methods("GET")
	globalMux.Handle("/module/", apiRouter)

	globalMux.Handle("/", http.FileServer(http.Dir("static")))

	log.Infof("controller listening on %s", a.listenAddr)
	s := &http.Server{
		Addr:    a.listenAddr,
		Handler: context.ClearHandler(globalMux),
	}
	var runErr error
	runErr = s.ListenAndServe()
	return runErr
}
