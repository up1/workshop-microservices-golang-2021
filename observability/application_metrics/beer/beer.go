package beer

import (
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
)

func NewRoutes(r *gin.Engine) {
	b := r.Group("/beer")
	b.GET("/", getAllBeer)
	b.GET("/:id", getBeerById)

	gaugeMetric := &ginmetrics.Metric{
		Type:        ginmetrics.Counter,
		Name:        "get_beer_by_id",
		Description: "get_beer_by_id",
		Labels:      []string{"status"},
	}

	// Add metric to global monitor object
	_ = ginmetrics.GetMonitor().AddMetric(gaugeMetric)
}

type BeerResponse struct {
	Message string `json:"message"`
}

func getAllBeer(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get all beer",
	})
}

func getBeerById(c *gin.Context) {
	id := c.Param("id")
	if id == "0" {
		_ = ginmetrics.GetMonitor().GetMetric("get_beer_by_id").Inc([]string{"not_found"})
		c.JSON(404, gin.H{
			"message": "Beer not found",
		})
		return
	}
	_ = ginmetrics.GetMonitor().GetMetric("get_beer_by_id").Inc([]string{"found"})
	c.JSON(200, gin.H{
		"message": "Found beer",
	})
	return
}
