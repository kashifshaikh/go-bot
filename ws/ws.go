package ws

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type Routes struct {
	Echo *echo.Echo
	Log  *zap.SugaredLogger
}

func RegisterRoutes(log *zap.SugaredLogger, e *echo.Echo) *Routes {
	r := &Routes{Log: log, Echo: e}
	r.Echo.GET("/ping", r.ping)
	r.Echo.GET("/ws", r.ws)
	return r
}

var upgrader = websocket.Upgrader{} // use default options

func wshandler(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Print("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Print("write:", err)
			break
		}
	}
}

// Handler
func (r *Routes) ping(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}

func (r *Routes) ws(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		mt, message, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error(err)
			return nil
		}
		log.Printf("recv: %s", message)
		err = ws.WriteMessage(mt, message)
		if err != nil {
			c.Logger().Error(err)
			return nil
		}
	}
}
