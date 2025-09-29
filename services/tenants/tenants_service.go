package tenants

import "mvc/domain/tenants"

func GetAll() ([]tenants.Tenant, error) {
	return tenants.GetAll()
}

func GetByID(id string) (tenants.Tenant, error) {
	return tenants.GetByGuid(id)
}

func Create(user *tenants.Tenant) error {
	return tenants.Create(user)
}

func Update(id string, user *tenants.Tenant) error {
	return tenants.Update(id, user)
}

func Delete(id string) error {
	return tenants.Delete(id)
}
