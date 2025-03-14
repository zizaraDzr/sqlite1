package donors

import (
	"mvc/domain/donors"
	"mvc/utils/errors"
)

func CreateDonor(donor donors.Donor) (*donors.Donor, *errors.RestErr) {
	// restErr := errors.RestErr{
	// 	Message: "invali json body",
	// 	Error:   "bad_request",
	// 	Status:  http.StatusBadRequest,
	// }
	return &donor, nil
}
