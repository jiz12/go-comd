package gen

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate template files",
	RunE:  GenerateTemplate,
}

func init() {
	Cmd.Flags().StringVar(&VarStringDir, "dir", ".", "The dir of the generated file")
	Cmd.Flags().StringVar(&VarStringFile, "file", "", "The name of the generated file")
	Cmd.Flags().StringVar(&VarStringTemplate, "template", "migrate", "which template to use. migrate or config or shortcut")
}
