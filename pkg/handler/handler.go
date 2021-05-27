package handler

import (
	"fmt"
	"io"
	"os"

	"github.com/LKarlon/converter/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/metrics", h.convert)

	return router
}

func (h *Handler) convert(c *gin.Context){
	file, err := os.Open("./file.yaml")
    if err != nil{
        fmt.Println(err) 
        os.Exit(1) 
    }
    defer file.Close() 
     
    data := make([]byte, 64)
    str := ""
    for{
        n, err := file.Read(data)
        if err == io.EOF{   
            break           
		}
		str = str + string(data[:n])

	}
	c.String(200, str)
}