package server

import (
	"fmt"
	"net/http"
	"path"

	"github.com/servusdei2018/hype/pkg/config"
	"github.com/servusdei2018/hype/pkg/futil"
)

// Serve serves a HTTP server.
func Serve(conf *config.Conf) (err error) {
	fs := http.FileServer(http.Dir(path.Join(futil.WD, "www")))
	http.Handle("/", fs)

	return http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", conf.Port), nil)
}
