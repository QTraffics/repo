package main

import (
	"fmt"

	"github.com/qtraffics/qtfra/log"
	"github.com/qtraffics/repo/qcli"

	"github.com/spf13/cobra"
)

var (
	ConfigFile string
	RuleFile   string
)

func init() {
	runCommand.Flags().StringVarP(
		&ConfigFile, "config", "c", "config.yaml", "Set the config file path.")

	runCommand.Flags().StringVarP(
		&RuleFile, "rule", "r", "rule.conf", "Set the rule file")

	mainCommand.AddCommand(runCommand)
}

var runCommand = &cobra.Command{
	Use:     "run",
	Short:   "Run a new instance",
	Example: fmt.Sprintf("%s run -c config.yaml", qcli.Name),
	Run:     runCommandRunFn,
}

func runCommandRunFn(cmd *cobra.Command, args []string) {
	if len(ConfigFile) == 0 || len(RuleFile) == 0 {
		_ = cmd.Help()
		return
	}

	create, err := qcli.Create(ConfigFile, "")
	if err != nil {
		log.Error("create new instance failed", log.AttrError(err))
		return
	}
	_ = create
}
