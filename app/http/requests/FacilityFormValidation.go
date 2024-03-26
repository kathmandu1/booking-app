package requests

type FacilityFormValidation struct {
	FacilityName     string `validate:"required"`
	Status           bool   `validate:"required"`
	FacilityCategory int    `validate:"required"`
}
