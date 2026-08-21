package main

import (
	"flag"
	"log/slog"
	"net"
	"net/http"
	"os"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"

	signaling "QUIADSA-ScreenSharing/signaling"
	"QUIADSA-ScreenSharing/static"
)

func main() {
	port := flag.String("port", "8080", "Puerto en el que se ejecuta el servidor")
	logLevel := flag.String("log-level", "info", "Nivel de log (debug, info, warn, error)")

	flag.Parse()

	var level slog.Level
	switch *logLevel {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		slog.Warn("Nivel de log no válido, usando info por defecto", "log-level", *logLevel)
		level = slog.LevelInfo
	}
	slog.SetLogLoggerLevel(level)

	server := &signaling.Signaler{
		SR: &signaling.StreamerRegistry{
			Entries: map[string]*signaling.Streamer{},
			Mu:      sync.RWMutex{},
		},
		VR: &signaling.ViewerRegistry{
			Entries: map[uuid.UUID]*signaling.Viewer{},
			Mu:      sync.RWMutex{},
		},
	}

	r := chi.NewRouter()

	r.Use(middleware.RealIP)
	r.Use(requestLogger)
	r.Use(middleware.Recoverer)

	r.Get("/stream", server.HandleStreamerWS)
	r.Get("/watch/{streamerCode}", server.HandleViewerWS)

	r.Get("/", static.Serve(static.IndexHTML, "text/html; charset=utf-8"))
	r.Get("/style.css", static.Serve(static.StyleCSS, "text/css; charset=utf-8"))
	r.Get("/main.js", static.Serve(static.MainJS, "application/javascript; charset=utf-8"))
	r.Get("/thumbnail.png", static.Serve(static.ThumbnailPNG, "image/png"))
	r.Get("/img/background.png", static.Serve(static.BackgroundPNG, "image/png"))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusPermanentRedirect)
	})

	listener, err := net.Listen("tcp", ":"+*port)
	if err != nil {
		slog.Error("No se pudo escuchar en el puerto", "port", *port, "error", err)
		os.Exit(1)
	}

	slog.Info("Servidor escuchando", "addr", listener.Addr().String())
	http.Serve(listener, r)
}
