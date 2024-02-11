//package main
//
//import (
//	"github.com/labstack/echo/v4"
//	"github.com/labstack/echo/v4/middleware"
//	"log"
//	"net/http"
//)
//
//func main() {
//	AllRooms.Init()
//
//	e := echo.New()
//	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
//		AllowOrigins:     []string{"*"},
//		AllowHeaders:     []string{"GET", "POST", "HEAD", "OPTIONS", "PUT", "PATCH", "DELETE"},
//		AllowCredentials: false,
//	}))
//
//	e.GET("/create", func(ctx echo.Context) error {
//		roomID := AllRooms.CreateRoom()
//
//		type resp struct {
//			RoomID string `json:"room_id"`
//		}
//
//		log.Println(AllRooms.Map)
//		return ctx.JSON(http.StatusOK, resp{RoomID: roomID})
//	})
//
//	e.GET("/join", func(ctx echo.Context) error {
//		roomID := ctx.QueryParam("roomID")
//
//		ws, err := upgrader.Upgrade(ctx.Response().Writer, ctx.Request(), nil)
//		if err != nil {
//			log.Fatal("Web Socket Upgrade Error", err)
//		}
//
//		AllRooms.InsertIntoRoom(roomID, false, ws)
//
//		go broadcaster()
//
//		for {
//			var msg broadcastMsg
//
//			err := ws.ReadJSON(&msg.Message)
//			if err != nil {
//				log.Fatal("Read Error: ", err)
//			}
//
//			msg.Client = ws
//			msg.RoomID = roomID
//
//			log.Println(msg.Message)
//
//			broadcast <- msg
//		}
//	})
//
//	e.Logger.Fatal(e.Start(":8000"))
//}

package main

import (
	"log"
	"net/http"
)

func main() {
	AllRooms.Init()

	http.HandleFunc("/create", CreateRoomRequestHandler)
	http.HandleFunc("/join", JoinRoomRequestHandler)

	log.Println("Starting Server on Port 8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
