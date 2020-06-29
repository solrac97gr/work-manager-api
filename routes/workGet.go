package routes

import (
	"github.com/gofiber/fiber"
	work2 "github.com/solrac97gr/yendoapi/database/work"
	"github.com/solrac97gr/yendoapi/utilities"
	"net/http"
)

func GetWorks(c *fiber.Ctx) {
	ID := c.Query("id")
	/*If the api call come to the endpoint with out a query ID search for all works of the user*/
	if len(ID) > 1 {
		work, err := work2.SearchWork(ID)
		if err != nil {
			c.Send("Error Occurred" + err.Error())
			c.SendStatus(http.StatusBadRequest)
			return
		}
		if err := c.JSON(work); err != nil {
			c.Status(http.StatusInternalServerError).Send(err)
			return
		}
		c.Accepts("application/json")
		c.SendStatus(http.StatusAccepted)
		return
	}

	works, err := work2.GetAll(utilities.UserID)
	if err != nil {
		c.Send("Error Occurred " + err.Error())
		c.SendStatus(http.StatusInternalServerError)
		return
	}
	err = c.JSON(works)
	if err != nil {
		c.Status(http.StatusInternalServerError).Send(err)
		return
	}
	c.Accepts("application/json")
	c.SendStatus(http.StatusAccepted)

}
