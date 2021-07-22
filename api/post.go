package orderapi

import (
	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
)

func Post(c *gin.Context) {
	var order Order
	c.Bind(&order)

	queueURL := "http://in-memory-queue:8080/queue/" + order.Product

	request := gorequest.New()

	resp, body, err := request.Post(queueURL).Send(order).End()

	var res QResponse

	res.Response = resp
	res.Body = body
	res.Error = err

	c.JSON(200, body)

}

type QResponse struct {
	Response gorequest.Response `json:"response"`
	Body     string             `json:"body"`
	Error    []error            `json:"error"`
}

type Order struct {
	Product  string `json:"product"`
	Quantity int    `json:"quantity"`
}
