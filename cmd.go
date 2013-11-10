package runsass

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Settings struct {
	SourceDir, DestinationDir string
}

type commander interface {
	Parse([]string)
	Run(*Settings) error
}

var (
	commands = make(map[string]commander)
)

/*func (s *Settings) RunCommand(params ...string) {
	cmdName := "runsass"
	if len(params) > 0 {
		cmdName = params[0]
	}

	if len(os.Args) < 2 || os.Args[1] != cmdName {
		return
	}

	args := argString(os.Args[2:])
	name := args.Get(0)

	if name == "help" {
		printHelp()
		return
	}
	if len(s.SourceDir) == 0 || len(s.DestinationDir) == 0 {
		printHelp("Please specify the source and destination directories for runsass.")
		return
	}

	if cmd, ok := commands[name]; ok {
		cmd.Parse(os.Args[3:])
		cmd.Run(s)
		os.Exit(0)
	} else {
		if name == "" {
			printHelp()
		} else {
			printHelp(fmt.Sprintf("unknown command %s", name))
		}
	}
}*/

func printHelp(errs ...string) {
	content := `runsass command usage:

    allinfolder     - run sass on all sass/scss files in folder
`

	if len(errs) > 0 {
		fmt.Println(errs[0])
	}
	fmt.Println(content)
	os.Exit(2)
}

const cDefaultStyle = "compressed"

type RunSassAll struct {
	Style   string // nested, expanded, compact, compressed
	Cache   bool
	Verbose bool
}

func (d *RunSassAll) Parse(args []string) {
	flagSet := flag.NewFlagSet("runsass command: allinfolder", flag.ExitOnError)
	flagSet.StringVar(&d.Style, "style", cDefaultStyle, "output file style")
	flagSet.BoolVar(&d.Cache, "cache", false, "allow sass to cache its files")
	flagSet.BoolVar(&d.Verbose, "v", false, "verbose info")
	flagSet.Parse(args)

	d.Style = strings.ToLower(d.Style)
}

func (d *RunSassAll) Run(s *Settings) error {
	fromDir := SanitizePath(s.SourceDir)
	toDir := SanitizePath(s.DestinationDir)

	if d.Style != "nested" && d.Style != "expanded" && d.Style != "compact" && d.Style != "compressed" {
		d.Style = cDefaultStyle
	}

	var out []byte
	var err error
	fmt.Println("--------------------------")
	if d.Cache {
		fmt.Println(fmt.Sprintf("Running SASS WITH cache on source '%s' to destination '%s'", fromDir, toDir))
		cmd := exec.Command("sass", "--update", "--style", d.Style, fmt.Sprintf("'%s':'%s'", fromDir, toDir))
		out, err = cmd.CombinedOutput()
	} else {
		fmt.Println(fmt.Sprintf("Running SASS WITHOUT cache on source '%s' to destination '%s'", fromDir, toDir))
		cmd := exec.Command("sass", "--update", "--no-cache", "--style", d.Style, fmt.Sprintf("'%s':'%s'", fromDir, toDir))
		out, err = cmd.CombinedOutput()
	}
	if err != nil {
		fmt.Println("Error: " + err.Error())
		fmt.Println("Output: " + string(out))
		fmt.Println("--------------------------")
		return err
	} else {
		fmt.Println("SASS ran successfully.")
		fmt.Println("--------------------------")
		return nil
	}
}

func init() {
	commands["allinfolder"] = &RunSassAll{}
}
