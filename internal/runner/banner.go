package runner

import (
	"fmt"
	"github.com/projectdiscovery/gologger"
)

func ShowBanner() {
	v := `0.0.1`
	gologger.Print().Msgf(`
   ______    ___________ __             __           
  / ____/___/_  __/ ___// /_____ ______/ /____  _____
 / / __/ __ \/ /  \__ \/ __/ __ */ ___/ __/ _ \/ ___/
/ /_/ / /_/ / /  ___/ / /_/ /_/ / /  / /_/  __/ /
\____/\____/_/  /____/\__/\__,_/_/   \__/\___/_/     %s`, v)
	fmt.Println()
}
