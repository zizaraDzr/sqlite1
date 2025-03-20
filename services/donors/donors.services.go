package donors

import (
	"mvc/domain/donors"
	"mvc/utils/errors"
)

func CreateDonor(donor donors.Donor) (*donors.Donor, *errors.RestErr) {
	if err := donor.Validate(); err != nil {
		return nil, err
	}
	if err := donor.Save(); err != nil {
		return nil, err
	}

	return &donor, nil
}
func GetDonor(donorId int64) (*donors.Donor, *errors.RestErr) {
	if donorId <= 0 {
		return nil, errors.NewBadRequestError("invalid donor id")
	}

	donor := &donors.Donor{DonorId: donorId}
	if err := donor.Get(); err != nil {
		return nil, err
	}

	return donor, nil
}
