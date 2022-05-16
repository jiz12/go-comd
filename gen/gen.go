package gen

import (
	"errors"
	"github.com/jiz12/go-comd/core/logger"
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

	templateFile, ok := template.Templates[VarStringTemplate]
	if !ok {
		logger.Logger.Error(errors.New("can't find template"))
		return nil
	}

	if VarStringFile == "" {
		VarStringFile = templateFile.DefaultName
	}

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

	err = template.LoadContentFromTemplate(f, VarStringDir, templateFile.Content)
	if err != nil {
		logger.Logger.Error(err)
	}
	return err
}
