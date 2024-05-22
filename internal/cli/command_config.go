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

package cli

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// runECommandForConfig runs the command for managing config.
func runECommandForConfig(cmd *cobra.Command, args []string) error {
	return cmd.Help()
}

// createCommandForConfig for managing config.
func createCommandForConfig(v *viper.Viper) *cobra.Command {
	command := &cobra.Command{
		Use:   "config",
		Short: "Configure CLI settings",
		Long:  `This command configure CLI settings.`,
		RunE:  runECommandForConfig,
	}
	command.AddCommand(createCommandForConfigAAPGet(v))
	command.AddCommand(createCommandForConfigAAPSet(v))
	command.AddCommand(createCommandForConfigPAPGet(v))
	command.AddCommand(createCommandForConfigPAPSet(v))
	return command
}
