// Copyright 2024 Nitro Agility S.r.l.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

// Package mocks implements mocks for testing.
package mocks

import (
	mock "github.com/stretchr/testify/mock"

	azmodels "github.com/permguard/permguard/pkg/agents/models"
)

// GrpcAAPClientMock is a mock type for the CliDependencies type.
type GrpcAAPClientMock struct {
	mock.Mock
}

// CreateAccount creates a new account.
func (m *GrpcAAPClientMock) CreateAccount(name string) (*azmodels.Account, error) {
	args := m.Called(name)
	var r0 *azmodels.Account
	if val, ok := args.Get(0).(*azmodels.Account); ok {
		r0 = val
	}
	return r0, args.Error(1)
}

// UpdateAccount updates an account.
func (m *GrpcAAPClientMock) UpdateAccount(account *azmodels.Account) (*azmodels.Account, error) {
	args := m.Called(account)
	var r0 *azmodels.Account
	if val, ok := args.Get(0).(*azmodels.Account); ok {
		r0 = val
	}
	return r0, args.Error(1)
}

// DeleteAccount deletes an account.
func (m *GrpcAAPClientMock) DeleteAccount(accountID int64) (*azmodels.Account, error) {
	args := m.Called(accountID)
	var r0 *azmodels.Account
	if val, ok := args.Get(0).(*azmodels.Account); ok {
		r0 = val
	}
	return r0, args.Error(1)
}

// FetchAccounts fetches accounts.
func (m *GrpcAAPClientMock) FetchAccounts(page int32, pageSize int32) ([]azmodels.Account, error) {
	args := m.Called(page)
	var r0 []azmodels.Account
	if val, ok := args.Get(0).([]azmodels.Account); ok {
		r0 = val
	}
	return r0, args.Error(1)
}

// FetchAccountsByID fetches accounts by ID.
func (m *GrpcAAPClientMock) FetchAccountsByID(page int32, pageSize int32, accountID int64) ([]azmodels.Account, error) {
	args := m.Called(page, pageSize, accountID)
	var r0 []azmodels.Account
	if val, ok := args.Get(0).([]azmodels.Account); ok {
		r0 = val
	}
	return r0, args.Error(1)
}

// FetchAccountsByName fetches accounts by name.
func (m *GrpcAAPClientMock) FetchAccountsByName(page int32, pageSize int32, name string) ([]azmodels.Account, error) {
	args := m.Called(page, pageSize, name)
	var r0 []azmodels.Account
	if val, ok := args.Get(0).([]azmodels.Account); ok {
		r0 = val
	}
	return r0, args.Error(1)
}

// FetchAccountsBy fetches accounts by.
func (m *GrpcAAPClientMock) FetchAccountsBy(page int32, pageSize int32, accountID int64, name string) ([]azmodels.Account, error) {
	args := m.Called(page, pageSize, accountID, name)
	var r0 []azmodels.Account
	if val, ok := args.Get(0).([]azmodels.Account); ok {
		r0 = val
	}
	return r0, args.Error(1)
}

// CreateIdentity creates a new identity.
func (m *GrpcAAPClientMock) CreateIdentity(accountID int64, identitySourceID string, kind string, name string) (*azmodels.Identity, error) {
	args := m.Called(accountID, identitySourceID, kind, name)
	var r0 *azmodels.Identity
	if val, ok := args.Get(0).(*azmodels.Identity); ok {
		r0 = val
	}
	return r0, args.Error(1)
}

// UpdateIdentity updates an identity.
func (m *GrpcAAPClientMock) UpdateIdentity(identity *azmodels.Identity) (*azmodels.Identity, error) {
	args := m.Called(identity)
	var r0 *azmodels.Identity
	if val, ok := args.Get(0).(*azmodels.Identity); ok {
		r0 = val
	}
	return r0, args.Error(1)
}

// DeleteIdentity deletes an identity.
func (m *GrpcAAPClientMock) DeleteIdentity(accountID int64, identityID string) (*azmodels.Identity, error) {
	args := m.Called(accountID, identityID)
	var r0 *azmodels.Identity
	if val, ok := args.Get(0).(*azmodels.Identity); ok {
		r0 = val
	}
	return r0, args.Error(1)
}

// FetchIdentities returns all identities.
func (m *GrpcAAPClientMock) FetchIdentities(page int32, pageSize int32, accountID int64) ([]azmodels.Identity, error) {
	args := m.Called(page, pageSize, accountID)
	var r0 []azmodels.Identity
	if val, ok := args.Get(0).([]azmodels.Identity); ok {
		r0 = val
	}
	return r0, args.Error(1)
}

// FetchIdentitiesByID returns all identities filtering by identity id.
func (m *GrpcAAPClientMock) FetchIdentitiesByID(page int32, pageSize int32, accountID int64, identityID string) ([]azmodels.Identity, error) {
	args := m.Called(page, pageSize, accountID, identityID)
	var r0 []azmodels.Identity
	if val, ok := args.Get(0).([]azmodels.Identity); ok {
		r0 = val
	}
	return r0, args.Error(1)
}

// FetchIdentitiesByEmail returns all identities filtering by name.
func (m *GrpcAAPClientMock) FetchIdentitiesByEmail(page int32, pageSize int32, accountID int64, name string) ([]azmodels.Identity, error) {
	args := m.Called(page, pageSize, accountID, name)
	var r0 []azmodels.Identity
	if val, ok := args.Get(0).([]azmodels.Identity); ok {
		r0 = val
	}
	return r0, args.Error(1)
}

// FetchIdentitiesBy returns all identities filtering by all criteria.
func (m *GrpcAAPClientMock) FetchIdentitiesBy(page int32, pageSize int32, accountID int64, identitySourceID string, identityID string, kind string, name string) ([]azmodels.Identity, error) {
	args := m.Called(page, pageSize, accountID, identitySourceID, identityID, kind, name)
	var r0 []azmodels.Identity
	if val, ok := args.Get(0).([]azmodels.Identity); ok {
		r0 = val
	}
	return r0, args.Error(1)
}

// CreateIdentitySource creates a new identity source.
func (m *GrpcAAPClientMock) CreateIdentitySource(accountID int64, name string) (*azmodels.IdentitySource, error) {
	args := m.Called(accountID, name)
	var r0 *azmodels.IdentitySource
	if val, ok := args.Get(0).(*azmodels.IdentitySource); ok {
		r0 = val
	}
	return r0, args.Error(1)
}

// UpdateIdentitySource updates an identity source.
func (m *GrpcAAPClientMock) UpdateIdentitySource(identitySource *azmodels.IdentitySource) (*azmodels.IdentitySource, error) {
	args := m.Called(identitySource)
	var r0 *azmodels.IdentitySource
	if val, ok := args.Get(0).(*azmodels.IdentitySource); ok {
		r0 = val
	}
	return r0, args.Error(1)
}

// DeleteIdentitySource deletes an identity source.
func (m *GrpcAAPClientMock) DeleteIdentitySource(accountID int64, identitySourceID string) (*azmodels.IdentitySource, error) {
	args := m.Called(accountID, identitySourceID)
	var r0 *azmodels.IdentitySource
	if val, ok := args.Get(0).(*azmodels.IdentitySource); ok {
		r0 = val
	}
	return r0, args.Error(1)
}

// FetchIdentitySources returns all identity sources.
func (m *GrpcAAPClientMock) FetchIdentitySources(page int32, pageSize int32, accountID int64) ([]azmodels.IdentitySource, error) {
	args := m.Called(page, pageSize, accountID)
	var r0 []azmodels.IdentitySource
	if val, ok := args.Get(0).([]azmodels.IdentitySource); ok {
		r0 = val
	}
	return r0, args.Error(1)
}

// FetchIdentitySourcesByID returns all identity sources filtering by identity source id.
func (m *GrpcAAPClientMock) FetchIdentitySourcesByID(page int32, pageSize int32, accountID int64, identitySourceID string) ([]azmodels.IdentitySource, error) {
	args := m.Called(page, pageSize, accountID, identitySourceID)
	var r0 []azmodels.IdentitySource
	if val, ok := args.Get(0).([]azmodels.IdentitySource); ok {
		r0 = val
	}
	return r0, args.Error(1)
}

// FetchIdentitySourcesByName returns all identity sources filtering by name.
func (m *GrpcAAPClientMock) FetchIdentitySourcesByName(page int32, pageSize int32, accountID int64, name string) ([]azmodels.IdentitySource, error) {
	args := m.Called(page, pageSize, accountID, name)
	var r0 []azmodels.IdentitySource
	if val, ok := args.Get(0).([]azmodels.IdentitySource); ok {
		r0 = val
	}
	return r0, args.Error(1)
}

// FetchIdentitySourcesBy returns all identity sources filtering by identity source id and name.
func (m *GrpcAAPClientMock) FetchIdentitySourcesBy(page int32, pageSize int32, accountID int64, identitySourceID string, name string) ([]azmodels.IdentitySource, error) {
	args := m.Called(page, pageSize, accountID, identitySourceID, name)
	var r0 []azmodels.IdentitySource
	if val, ok := args.Get(0).([]azmodels.IdentitySource); ok {
		r0 = val
	}
	return r0, args.Error(1)
}

// CreateTenant creates a tenant.
func (m *GrpcAAPClientMock) CreateTenant(accountID int64, name string) (*azmodels.Tenant, error) {
	args := m.Called(accountID, name)
	var r0 *azmodels.Tenant
	if val, ok := args.Get(0).(*azmodels.Tenant); ok {
		r0 = val
	}
	return r0, args.Error(1)
}

// UpdateTenant updates a tenant.
func (m *GrpcAAPClientMock) UpdateTenant(tenant *azmodels.Tenant) (*azmodels.Tenant, error) {
	args := m.Called(tenant)
	var r0 *azmodels.Tenant
	if val, ok := args.Get(0).(*azmodels.Tenant); ok {
		r0 = val
	}
	return r0, args.Error(1)
}

// DeleteTenant deletes a tenant.
func (m *GrpcAAPClientMock) DeleteTenant(accountID int64, tenantID string) (*azmodels.Tenant, error) {
	args := m.Called(accountID, tenantID)
	var r0 *azmodels.Tenant
	if val, ok := args.Get(0).(*azmodels.Tenant); ok {
		r0 = val
	}
	return r0, args.Error(1)
}

// FetchTenants returns all tenants.
func (m *GrpcAAPClientMock) FetchTenants(page int32, pageSize int32, accountID int64) ([]azmodels.Tenant, error) {
	args := m.Called(page, pageSize, accountID)
	var r0 []azmodels.Tenant
	if val, ok := args.Get(0).([]azmodels.Tenant); ok {
		r0 = val
	}
	return r0, args.Error(1)
}

// FetchTenantsByID returns all tenants filtering by tenant id.
func (m *GrpcAAPClientMock) FetchTenantsByID(page int32, pageSize int32, accountID int64, tenantID string) ([]azmodels.Tenant, error) {
	args := m.Called(page, pageSize, accountID, tenantID)
	var r0 []azmodels.Tenant
	if val, ok := args.Get(0).([]azmodels.Tenant); ok {
		r0 = val
	}
	return r0, args.Error(1)
}

// FetchTenantsByName returns all tenants filtering by name.
func (m *GrpcAAPClientMock) FetchTenantsByName(page int32, pageSize int32, accountID int64, name string) ([]azmodels.Tenant, error) {
	args := m.Called(page, pageSize, accountID, name)
	var r0 []azmodels.Tenant
	if val, ok := args.Get(0).([]azmodels.Tenant); ok {
		r0 = val
	}
	return r0, args.Error(1)
}

// FetchTenantsBy returns all tenants filtering by tenant id and name.
func (m *GrpcAAPClientMock) FetchTenantsBy(page int32, pageSize int32, accountID int64, tenantID string, name string) ([]azmodels.Tenant, error) {
	args := m.Called(page, pageSize, accountID, tenantID, name)
	var r0 []azmodels.Tenant
	if val, ok := args.Get(0).([]azmodels.Tenant); ok {
		r0 = val
	}
	return r0, args.Error(1)
}

// NewGrpcAAPClientMock creates a new GrpcAAPClientMock.
func NewGrpcAAPClientMock() *GrpcAAPClientMock {
	return &GrpcAAPClientMock{}
}