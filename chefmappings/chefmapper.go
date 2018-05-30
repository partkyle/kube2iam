package chefmappings

import (
	"log"

	"github.com/jtblin/kube2iam/iam"
)

// RoleMapper handles relevant logic around associating IPs with a given IAM role
type RoleMapper struct {
	defaultRoleARN string
	iam            *iam.Client
}

// RoleMappingResult represents the relevant information for a given mapping request
type RoleMappingResult struct {
	Role string
	IP   string
}

// GetRoleMapping returns the normalized iam RoleMappingResult based on IP address
func (r *RoleMapper) GetRoleMapping(IP string) (*RoleMappingResult, error) {
	log.Printf("ip=%s", IP)
	return &RoleMappingResult{
		Role: r.defaultRoleARN,
		IP:   IP,
	}, nil

	// return nil, fmt.Errorf("Role requested not valid for namespace of pod at %s", IP)
}

// NewRoleMapper returns a new RoleMapper for use.
func NewRoleMapper(defaultRole string, iamInstance *iam.Client) *RoleMapper {
	return &RoleMapper{
		defaultRoleARN: iamInstance.RoleARN(defaultRole),
		iam:            iamInstance,
	}
}
