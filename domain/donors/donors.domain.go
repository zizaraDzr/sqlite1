package donors

type Donor struct {
	DonorId          int64  `json:"donor_id"`
	FullName         string `json:"full_name"`
	Phone            string `json:"phone"`
	BloodGroup       string `json:"blood_group"`
	RhesusFactor     string `json:"rhesus_factor"`
	DonationsPerYear int    `json:"donations_per_year"`
}
