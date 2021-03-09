package beer

import "github.com/gin-gonic/gin"

func NewRoutes(r *gin.Engine) {
	b := r.Group("/beer")
	b.GET("/", getAllBeer)
}

type BeerResponse struct {
	Message string `json:"message"`
}

// getAllBeer  get all beers from database
// @Summary Retrieves users from mongodb
// @Description Get All Beers
// @Produce json
// @Param name query string false "Name"
// @Param age query int false "Age"
// @Success 200 {object} BeerResponse
// @Router /beer [get]
func getAllBeer(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get all beer",
	})
}
