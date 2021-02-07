package wingPi

import (
	"fmt"
	"github.com/commander-cli/cmd"
	"github.com/wingofsky/wingPi/utils"
)

func TakePicture(fileName, rootDir string) {
	// 质量为5的文件，不显示预览
	cl := "raspistill -o "
	cl += rootDir + "/"
	cl += fileName
	cl += "."
	cl += utils.PointNow()
	cl += ".jpg -q 5 -n"

	fmt.Println(cl)

	c := cmd.NewCommand(cl)
	utils.CheckERR(c.Execute())
	//fmt.Println(c.Stdout())
	//fmt.Println(c.Stderr())
}
