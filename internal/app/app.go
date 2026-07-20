package app

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/pudgedamuerto/mockservice/internal/config"
)

type App struct {
	cnf  config.ServiceConfig
	addr string
}

func NewApp(cnf config.ServiceConfig, addr string) App {
	return App{
		cnf,
		addr,
	}
}

func (a App) Serve() error {
	mux := http.NewServeMux()

	for endpoint, routes := range a.cnf.Routes {
		for _, route := range routes {
			pattern := fmt.Sprintf("%s %s", route.Method, endpoint)
			log.Println(pattern)

			mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
				body, err := io.ReadAll(r.Body)
				if err != nil {
					log.Printf("ERROR: %s", err)
				}

				var requestLog string = "nil"
				if len(body) != 0 {
					requestLog = string(body)
				}

				log.Printf("%s called with request: %s", pattern, requestLog)
				json, err := json.Marshal(route.Response.Body)
				if err != nil {
					log.Printf("ERROR: %s", err)
				}

				w.Header().Set("content-type", "application/json")
				w.WriteHeader(route.Response.Status)

				if _, err := w.Write(json); err != nil {
					log.Printf("ERROR: %s", err)
				}
			})
		}
	}

	log.Printf("app started on %s", a.addr)
	return http.ListenAndServe(a.addr, mux)
}
