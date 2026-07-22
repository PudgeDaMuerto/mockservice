package app

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

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
			message := pattern
			delay := route.Response.Delay
			if delay != nil {
				message = fmt.Sprintf("%s (delay: %s)", pattern, time.Duration(*delay))
			}
			log.Println(message)

			json, err := json.Marshal(route.Response.Body)
			if err != nil {
				return err
			}

			mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
				defer r.Body.Close()

				body, err := io.ReadAll(r.Body)
				if err != nil {
					log.Printf("ERROR: %s", err)
					return
				}

				requestLog := "nil"
				if len(body) != 0 {
					requestLog = string(body)
				}

				log.Printf("%s called with request: %s", pattern, requestLog)

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(route.Response.Status)

				if d := route.Response.Delay; d != nil {
					select {
					case <-time.After(time.Duration(*d)):
					case <-r.Context().Done():
						log.Println("request canceled")
						return
					}
				}

				if _, err := w.Write(json); err != nil {
					log.Printf("ERROR: %s", err)
					return
				}
			})
		}
	}

	log.Printf("app started on %s", a.addr)
	return http.ListenAndServe(a.addr, mux)
}
