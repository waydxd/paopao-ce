// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/waydxd/paopao-ce/pkg/utils"
	"github.com/waydxd/paopao-ce/pkg/version"
)

func init() {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "show version information",
		Long:  "show version information",
		Run:   versionRun,
	}
	Register(versionCmd)
}

func versionRun(_cmd *cobra.Command, _args []string) {
	utils.PrintHelloBanner(version.VersionInfo())
}
