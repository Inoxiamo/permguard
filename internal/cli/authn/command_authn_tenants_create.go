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

package authn

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	azcli "github.com/permguard/permguard/pkg/cli"
	azconfigs "github.com/permguard/permguard/pkg/configs"
	aziclicommon "github.com/permguard/permguard/internal/cli/common"
)

const (
	// commandNameForTenantsCreate is the command name for tenants create.
	commandNameForTenantsCreate = "tenants.create"
)

// runECommandForCreateTenant runs the command for creating a tenant.
func runECommandForCreateTenant(deps azcli.CliDependenciesProvider, cmd *cobra.Command, v *viper.Viper) error {
	return runECommandForUpsertTenant(deps, cmd, v, commandNameForTenantsCreate, true)
}

// createCommandForTenantCreate creates a command for managing tenantcreate.
func createCommandForTenantCreate(deps azcli.CliDependenciesProvider, v *viper.Viper) *cobra.Command {
	command := &cobra.Command{
		Use:   "create",
		Short: "Create a tenant",
		Long: fmt.Sprintf(aziclicommon.CliLongTemplate, `This command creates a tenant.

Examples:
  # create a tenant with name tenant1 and account 301990992055
  permguard authn tenants create --account 301990992055 --name tenant1
		`),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runECommandForCreateTenant(deps, cmd, v)
		},
	}
	command.Flags().String(aziclicommon.FlagCommonName, "", "tenant name")
	v.BindPFlag(azconfigs.FlagName(commandNameForTenantsCreate, aziclicommon.FlagCommonName), command.Flags().Lookup(aziclicommon.FlagCommonName))
	return command
}