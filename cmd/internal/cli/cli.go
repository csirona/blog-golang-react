package cli
import(
	"flag"
	"fmt"
	"os"
	"blog.com/packages/cmd/internal/logging"
)

func usage(){
	fmt.Print(`This program runs RGB backend server.
	
	Usage:
	
	rgb [arguments]
	
	Supported arguments:
	
	`)
	flag.PrintDefaults()
	os.Exit(1)
}

func Parse() string {
	flag.Usage = usage
	env := flag.String("env","dev", `Set run enviroment. Possible values are "dev and "prod"`)
	flag.Parse()
	logging.ConfigureLogger(*env)
	if *env == "prod" {
		logging.SetGinLogToFile()
	}
	return *env
}