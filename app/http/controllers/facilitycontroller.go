package faciltycontroller

import (
	models "dinning/app/modals"
	"dinning/config"
	"net/http"
	"strconv"

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

// Index method  return the all the list of facilty
func (c *FacilityController) Index(ctx echo.Context) error {

	var facilities []models.Facility
	db := config.DB()

	result := db.Find(&facilities)
	if err := result.Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, facilities)

}

// func ValidateFacility(facility models.Facility) error {
// 	validate := validator.New()
// 	facilityFormValidation := requests.FacilityFormValidation{
// 		FacilityName:     facility.FacilityName,
// 		Status:           facility.Status,
// 		FacilityCategory: facility.FacilityCategory,
// 	}

// 	err := validate.Struct(facilityFormValidation)
// 	if err != nil {
// 		// Extract validation errors and create a custom error response
// 		var ve validator.ValidationErrors
// 		if errors.As(err, &ve) {
// 			out := make([]map[string]interface{}, len(ve))
// 			for i, fe := range ve {
// 				out[i] = map[string]interface{}{
// 					"field": fe.Field(),
// 					"error": fe.Tag(), // Use custom message from tag if available
// 				}
// 			}
// 			return echo.NewHTTPError(403, out)
// 		}
// 		return err // Handle non-validation errors
// 	}
// 	return nil
// }

// Save handles the root path of your application
func (c *FacilityController) Save(ctx echo.Context) error {

	facilityCategoryFromRequest := ctx.FormValue("FacilityCategory")
	facilityCategoryInt, inputerr := strconv.Atoi(facilityCategoryFromRequest)
	if inputerr != nil {
		return echo.ErrBadRequest

	}
	status := (ctx.FormValue("Status"))
	statusToInt, booleanDataerr := strconv.ParseBool(status)
	if booleanDataerr != nil {
		return echo.ErrBadRequest

	}
	facilityData := models.Facility{
		FacilityName:     ctx.FormValue("FacilityName"),
		Status:           statusToInt,
		FacilityCategory: facilityCategoryInt,
	}

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

// Show handles the facility model of your application
func (c *FacilityController) Show(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.ErrBadRequest // Handle invalid ID format
	}

	var facilities = models.Facility{ID: id}

	db := config.DB()

	result := db.First(&facilities)
	if err := result.Error; err != nil {
		// Handle database error appropriately
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, facilities)
}

// Update handles the  facility model of your application
func (c *FacilityController) Update(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello Update from the controller!")
}

// Delete handles the facility model of your application
func (c *FacilityController) Delete(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return echo.ErrBadRequest // Handle invalid ID format
	}

	var facility models.Facility
	db := config.DB()
	result := db.First(&facility, id)
	if result.Error != nil {
		return echo.ErrInternalServerError // Handle other database errors
	}

	err = db.Delete(&facility).Error
	if err != nil {
		return echo.ErrInternalServerError
	}

	return ctx.NoContent(http.StatusNoContent)
}

// RegisterRoutes mounts the controller methods to Echo routes
func RegisterRoutes(e *echo.Echo) {
	c := &FacilityController{}

	// Mount routes with appropriate HTTP methods
	e.GET("facilities", c.Index)
	e.POST("facilities", c.Save)
	e.GET("facilities/:id", c.Show)
	e.PATCH("facilities/:id", c.Update)
	e.DELETE("facilities/:id", c.Delete)
	// Add other routes as per need
}
