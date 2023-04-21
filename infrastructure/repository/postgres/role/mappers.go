package role

import (
	domainRole "hacktiv/final-project/domain/role"
)

// toDomainMapper function to convert role repo to role domain
func (role *Role) toDomainMapper() *domainRole.Role {
	return &domainRole.Role{
		ID:        role.ID,
		Name:      role.Name,
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
	}
}

// fromDomainMapper function to convert role domain to role repo
func fromDomainMapper(role *domainRole.Role) *Role {
	return &Role{
		ID:        role.ID,
		Name:      role.Name,
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
	}
}

// arrayToDomainMapper function to convert list role domain to list role repo
func arrayToDomainMapper(roles *[]Role) *[]domainRole.Role {
	rolesDomain := make([]domainRole.Role, len(*roles))
	for i, role := range *roles {
		rolesDomain[i] = *role.toDomainMapper()
	}

	return &rolesDomain
}
