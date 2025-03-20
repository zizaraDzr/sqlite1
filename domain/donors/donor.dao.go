package donors

import (
	"fmt"
	"mvc/utils/errors"
)

var (
	donorDB = make(map[int64]*Donor)
)

func (donor *Donor) Save() *errors.RestErr {
	current := donorDB[donor.DonorId]
	if current != nil {
		if current.FullName != donor.FullName {
			return errors.NewBadRequestError(fmt.Sprintf("donor %s already exists", donor.FullName))
		}
		return errors.NewBadRequestError(fmt.Sprintf("donor %d already exists", donor.DonorId))
	}
	donorDB[donor.DonorId] = donor

	return nil
}
func (donor *Donor) Get() *errors.RestErr {
	result := donorDB[donor.DonorId]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("donor %d not found", donor.DonorId))
	}
	donor.DonorId = result.DonorId
	donor.FullName = result.FullName
	donor.Phone = result.Phone
	donor.BloodGroup = result.BloodGroup
	donor.RhesusFactor = result.RhesusFactor
	donor.DonationsPerYear = result.DonationsPerYear

	return nil
}
