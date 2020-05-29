package cmd

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"runtime"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	version     bool
	inFile   string
	outFile   string
	logger      = logrus.New()
	// Version shows the appendtoyml binary version.
	Version string
	// GitSHA shows the  appendtoyml binary code commit SHA on git.
	GitSHA string
)

func printVersionInfo() {
	logger.Infof("etcd-backup-restore Version: %s", Version)
	logger.Infof("Git SHA: %s", GitSHA)
	logger.Infof("Go Version: %s", runtime.Version())
	logger.Infof("Go OS/Arch: %s/%s", runtime.GOOS, runtime.GOARCH)
}

func NewAppendtoymlCommand() *cobra.Command {
	var RootCmd = &cobra.Command{
		Use:   "appendtoyml",
		Short: "command line utility for append host in host.csv to prometheus ",
		Long: `The appendtoyml, command line utility, is built to read host info in host.csv, and append to prometheus config file proemthesu.yml.`,
		Run: func(cmd *cobra.Command, args []string) {
			if version {
				printVersionInfo()
			} else {
				AppendToYaml(inFile, outFile)
			}
		},
	}
	RootCmd.Flags().BoolVarP(&version, "version", "v", false, "print version info")
	RootCmd.PersistentFlags().StringVarP(&inFile, "inFile", "i", "./hosts.csv", "input file, default is ./hosts.csv")
	RootCmd.PersistentFlags().StringVarP(&outFile, "outFile", "o", "./prometheus/prometheus.yml", "template file and also is out file, default is ./prometheus/prometheus.yml")
	return RootCmd
}

type Host struct {
	Hostname string   `json:"hostname"`
	HostIP string   `json:"hostIp"`
}

func AppendToYaml(infile string, outfile string){
	csvFile, _ := os.Open(infile)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var host []Host
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			fmt.Printf("\nerror: %v\n", error)
		}
		host = append(host, Host{
			Hostname: line[0],
			HostIP:  line[1],
		})
	}
	peopleJson, _ := json.Marshal(host)
	fmt.Println(string(peopleJson))
}