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

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	azmodels "github.com/permguard/permguard/pkg/agents/models"
	azcli "github.com/permguard/permguard/pkg/cli"
	azconfigs "github.com/permguard/permguard/pkg/configs"
	aziclicommon "github.com/permguard/permguard/internal/cli/common"
)

const (
	// commandNameForIdentity is the command name for identity.
	commandNameForIdentitiesDelete = "identities.delete"
)

// runECommandForDeleteIdentity runs the command for creating an identity.
func runECommandForDeleteIdentity(deps azcli.CliDependenciesProvider, cmd *cobra.Command, v *viper.Viper) error {
	ctx, printer, err := aziclicommon.CreateContextAndPrinter(deps, cmd, v)
	if err != nil {
		color.Red(aziclicommon.ErrorMessageCliBug)
		return aziclicommon.ErrCommandSilent
	}
	aapTarget := ctx.GetAAPTarget()
	client, err := deps.CreateGrpcAAPClient(aapTarget)
	if err != nil {
		printer.Error(fmt.Errorf("invalid aap target %s", aapTarget))
		return aziclicommon.ErrCommandSilent
	}
	accountID := v.GetInt64(azconfigs.FlagName(commandNameForIdentity, aziclicommon.FlagCommonAccountID))
	identityID := v.GetString(azconfigs.FlagName(commandNameForIdentitiesDelete, flagIdentityID))
	identity, err := client.DeleteIdentity(accountID, identityID)
	if err != nil {
		printer.Error(err)
		return aziclicommon.ErrCommandSilent
	}
	output := map[string]any{}
	if ctx.IsTerminalOutput() {
		identityID := identity.IdentityID
		identityName := identity.Name
		output[identityID] = identityName
	} else if ctx.IsJSONOutput() {
		output["identities"] = []*azmodels.Identity{identity}
	}
	printer.Print(output)
	return nil
}

// createCommandForIdentityDelete creates a command for managing identitydelete.
func createCommandForIdentityDelete(deps azcli.CliDependenciesProvider, v *viper.Viper) *cobra.Command {
	command := &cobra.Command{
		Use:   "delete",
		Short: "Delete an identity",
		Long: fmt.Sprintf(aziclicommon.CliLongTemplate, `This command deletes an identity.

Examples:
  # delete an identity with id 19159d69-e902-418e-966a-148c4d5169a4 and account 301990992055
  permguard authn identities delete --account 301990992055 --identityid 19159d69-e902-418e-966a-148c4d5169a4
		`),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runECommandForDeleteIdentity(deps, cmd, v)
		},
	}
	command.Flags().String(flagIdentityID, "", "identity id")
	v.BindPFlag(azconfigs.FlagName(commandNameForIdentitiesDelete, flagIdentityID), command.Flags().Lookup(flagIdentityID))
	return command
}