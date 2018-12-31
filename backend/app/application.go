package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/gustavohenrique/e2e-tests-using-docker-cypress-vuejs-golang-postgres/backend/app/database"
	"github.com/gustavohenrique/e2e-tests-using-docker-cypress-vuejs-golang-postgres/backend/app/todolist"
)

type Application struct {
	Router  *mux.Router
	Handler *todolist.TaskHandler
}

func New() *Application {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "It is working.")
	})

	app := &Application{}
	app.Router = router
	return app
}

func (app *Application) EnableCORS() http.Handler {
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "OPTIONS"})
	app.Router.Use(mux.CORSMethodMiddleware(app.Router))
	return handlers.CORS(originsOk, headersOk, methodsOk)(app.Router)
}

func (app *Application) Initialize(params map[string]string) {
	databaseURL := params["databaseURL"]
	db := database.NewDB(databaseURL)
	service := todolist.NewService(db)
	handler := app.Handler
	if handler == nil {
		h := todolist.NewHandler(service)
		handler = &h
	}
	app.Router.HandleFunc("/todos", handler.FindAll).Methods("GET")
	app.Router.HandleFunc("/todos", handler.Create).Methods("POST", "OPTIONS")
	app.Router.HandleFunc("/todos/{id}", handler.Done).Methods("PUT", "OPTIONS")
}
