package clinic

import (
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
	"gocourse22/internal/interface/http"
	httpInt "net/http"
)

func NewClinicHandler(_ *do.Injector) (*ClinicHandler, error) {
	return &ClinicHandler{}, nil
}

type ClinicHandler struct {
}

func (h *ClinicHandler) RegisterRoutes(g *gin.Engine) {
	api := g.Group("/api")
	{
		group := api.Group("/clinic")
		{
			group.GET("",
				h.viewAll,
			)

			group.POST("",
				h.create,
			)

			group.GET("/:clinic_id",
				h.view,
			)

			group.PUT("/:clinic_id",
				h.edit,
			)

			group.DELETE("/:clinic_id",
				h.delete,
			)

			group.GET("/:clinic_id/config",
				h.config,
			)

			group.GET("/visits/",
				h.groupVisits,
			)
		}
	}
}

func (h *ClinicHandler) groupVisits(c *gin.Context) {
	s := &Service{}
	res := s.GroupPatientsVisits()

	http.NewResponse(c, res)
}

func (h *ClinicHandler) viewAll(c *gin.Context) {
	// do view all

	http.NewResponse(c, []byte{})
}

func (h *ClinicHandler) view(c *gin.Context) {
	// do view

	http.NewResponse(c, []byte{})
}

func (h *ClinicHandler) create(c *gin.Context) {
	// do create

	http.NewResponse(c, []byte{}, http.WithStatusCode(httpInt.StatusCreated))
}

func (h *ClinicHandler) edit(c *gin.Context) {
	// do edit

	http.NewResponse(c, []byte{})
}

func (h *ClinicHandler) delete(c *gin.Context) {
	// do delete

	c.Status(httpInt.StatusNoContent)
}

func (h *ClinicHandler) config(c *gin.Context) {
	// do config creation

	c.FileAttachment(`/tmp/`, "tmp.zip")
}
