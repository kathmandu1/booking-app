package faciltycontroller

import (
	models "dinning/app/modals"
	"dinning/config"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// Controller defines a struct for controller methods
type FacilityController struct {
	db *gorm.DB
}

func NewController(db *gorm.DB) *FacilityController {
	return &FacilityController{db: db}
}

// Index handles the root path of your application
func (c *FacilityController) Index(ctx echo.Context) error {
	// Access the request body to get user data

}

// Save handles the root path of your application
func (c *FacilityController) Save(ctx echo.Context) error {
	facilityData := models.Facility{
		ID:               1,
		FacilityName:     "Jinzhu",
		Status:           true,
		FacilityCategory: 1}

	if err := ctx.Bind(&facilityData); err != nil {
		return echo.ErrBadRequest
	}
	db := config.DB()

	result := db.Create(&facilityData)
	if err := result.Error; err != nil {
		// Handle database error appropriately
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	// 4. Return successful response (optional)
	return ctx.JSON(http.StatusCreated, facilityData)
}

// Show handles the root path of your application
func (c *FacilityController) Show(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello from  Show the controller!")
}

// Update handles the root path of your application
func (c *FacilityController) Update(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello Update from the controller!")
}

// Delete handles the root path of your application
func (c *FacilityController) Delete(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello from the controller!")
}

// RegisterRoutes mounts the controller methods to Echo routes
func RegisterRoutes(e *echo.Echo) {
	c := &FacilityController{}

	// Mount routes with appropriate HTTP methods
	e.GET("facilities", c.Index)
	e.POST("facilities", c.Save)
	e.PATCH("facilities", c.Update)
	e.DELETE("facilities", c.Delete)
	// Add other routes for your application endpoints here
}
