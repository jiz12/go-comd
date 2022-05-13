package gen

import (
	"fmt"
	"github.com/urfave/cli"
	"gocomd/core/logger"
	"gocomd/utils"
	"os"
)

// GoCommand generate mysql template
func GoCommand(c *cli.Context) error {

	dir := c.String("dir")
	file := c.String("file")
	template := c.String("template")
	if len(dir) == 0 {
		dir, _ = os.Getwd() //当前路径
	}

	if len(file) == 0 {
		file = fmt.Sprintf("gen_%v", template)
	}

	f, err := utils.CreateFileIfNotExist(file)
	if err != nil {
		logger.Logger.Error(err)
		return err
	}
	defer f.Close()

	fmt.Println("我在执行了！")

	return nil
}
