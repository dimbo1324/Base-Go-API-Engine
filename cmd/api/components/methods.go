package components

import (
	"net/http"

)

func (a *Application) Run() error {
	mux := http.NewServeMux()
	srv := &http.Server{
		Addr:    a.Config.Addr,
		Handler: mux,
	}
	return srv.ListenAndServe()
}
