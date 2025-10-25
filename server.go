package ti84_webserver

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed server/css/*
//go:embed server/fonts/*
//go:embed server/js/*
//go:embed server/modules/native_eZ80/*
//go:embed server/index.html
var staticFiles embed.FS

func NewCeServer(baseDomain string) http.Handler {
	ce_mux := http.NewServeMux()
	serverFS, _ := fs.Sub(staticFiles, "server")
	ce_mux.Handle("/", http.FileServer(http.FS(serverFS)))
	return ce_mux
}
