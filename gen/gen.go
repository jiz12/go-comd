package gen

import (
	"github.com/jiz12/go-comd/core/utils"
	"github.com/jiz12/go-comd/template"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var (
	// VarStringDir describes the dir.
	VarStringDir string
	// VarStringFile describes the file name.
	VarStringFile string
	// VarStringTemplate describes which template to use.
	VarStringTemplate string
)

//GenerateTemplate generate template file
func GenerateTemplate(_ *cobra.Command, _ []string) (err error) {

	f, err := utils.CreateFileIfNotExist(VarStringDir, VarStringFile)
	if err != nil {
		return err
	}

	defer func() {
		f.Close()
		if err != nil && f != nil {
			err = os.Remove(filepath.Join(VarStringDir, VarStringFile))
		}
	}()

	return template.LoadContentFromTemplate(f, VarStringDir, VarStringTemplate)
}
