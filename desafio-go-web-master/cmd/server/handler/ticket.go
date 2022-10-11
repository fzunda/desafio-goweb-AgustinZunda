package handler

import (
	"net/http"

	"github.com/fzunda/desafio-goweb-AgustinZunda/internal/tickets"
	"github.com/gin-gonic/gin"
)

// esta deberia ser struct product no service
type Ticket struct {
	service tickets.Service
}

func NewTicket(s tickets.Service) *Ticket {
	return &Ticket{
		service: s,
	}
}

func (s *Ticket) GetTicketsByCountry() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		tickets, err := s.service.GetTotalTickets(c, destination)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.JSON(200, len(tickets))
	}
}

func (s *Ticket) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res, err := s.service.GetAll(ctx)
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}
		ctx.JSON(http.StatusOK, res)
		return
	}
}

/*func (s *Ticket) AverageDestination() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		avg, err := s.service.AverageDestination(c, destination)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.JSON(200, avg)
	}
}*/
