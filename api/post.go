package orderapi

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
)

func Post(c *gin.Context) {
	var order Order
	c.Bind(&order)

	fmt.Println("Product is: ", order.Product)

	QueueURL := os.Getenv("QUEUE_URL")
	fmt.Println("QueueURL is: ", QueueURL)

	if QueueURL == "" {
		QueueURL = "http://localhost:8080/queue/"
	}

	queueURL := QueueURL + order.Product

	request := gorequest.New()

	resp, body, err := request.Post(queueURL).Send(order).End()

	if len(err) > 0 {
		c.JSON(500, gin.H{
			"message": err,
		})
		return
	}

	var res QResponse

	res.Response = resp
	res.Body = body
	res.Error = err

	var jres map[string]interface{}

	json.Unmarshal([]byte(body), &jres)

	c.JSON(200, jres)

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
