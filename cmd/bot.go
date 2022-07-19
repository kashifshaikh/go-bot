package main

import (
	"bot"
	"bot/profile"
	"bot/task"
	"bot/user"
	"bot/ws"
	"fmt"
	"sort"

	"github.com/asdine/storm/v3"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type Bot struct {
	// Configuration path and parsed config data.
	Config         *bot.Config
	Log            *zap.SugaredLogger
	DB             *storm.DB
	Echo           *echo.Echo
	ProfileRoutes  *profile.Routes
	TaskRoutes     *task.Routes
	UserRoutes     *user.Routes
	WSRoutes       *ws.Routes
	ProfileService *profile.Service
}

func registerServices(b *Bot) {
	b.ProfileService = profile.NewService(b.Log, b.DB)
}

func registerRoutes(b *Bot) {

	b.ProfileRoutes = profile.RegisterRoutes(b.Log, b.Echo, b.ProfileService)
	b.TaskRoutes = task.RegisterRoutes(b.Log, b.Echo)
	b.UserRoutes = user.RegisterRoutes(b.Log, b.Echo)
	b.WSRoutes = ws.RegisterRoutes(b.Log, b.Echo)

	s := "Route dump:\n"
	routes := b.Echo.Routes()

	sort.Slice(routes, func(a, b int) bool {
		return routes[a].Path < routes[b].Path
	})

	for _, r := range routes {
		s = s + fmt.Sprintf("\t%s %s -> %s\n", r.Method, r.Path, r.Name)
	}
	b.Log.Infof(s)
}

func main() {

	cfg := bot.NewConfig()
	log := bot.NewLogger(cfg)
	db := bot.NewDB(cfg)
	defer db.Close()
	e := bot.NewEcho(cfg)
	b := &Bot{
		Config: cfg,
		Log:    log,
		DB:     db,
		Echo:   e,
	}
	registerServices(b)
	registerRoutes(b)

	bot.StartEchoServer(b.Echo)
}
