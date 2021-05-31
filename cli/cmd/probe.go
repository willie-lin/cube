package cmd

import (
	"cube/cli"
	"cube/log"
	Plugins "cube/plugins"
	"os"
	"strings"

	//"cube/log"
	"cube/model"
	"fmt"
	"github.com/spf13/cobra"
)

var probeCmd *cobra.Command

func runProbe(cmd *cobra.Command, args []string) {
	globalopts, opt, _ := parseProbeOptions()
	_, key := Plugins.ProbeFuncMap[opt.ScanPlugin]
	if !key {
		log.Fatalf("Available Plugins: %s", strings.Join(Plugins.ProbeKeys, ","))
		os.Exit(2)
	}
	cli.StartProbeTask(opt, globalopts)
}

func parseProbeOptions() (*model.GlobalOptions, *model.ProbeOptions, error) {
	globalOpts, err := parseGlobalOptions()
	if err != nil {
		return nil, nil, err
	}
	probeOption := model.NewProbeOptions()

	probeOption.ScanPlugin, err = probeCmd.Flags().GetString("plugin")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for plugin: %v", err)
	}

	probeOption.Port, err = probeCmd.Flags().GetInt("port")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for scan port: %v", err)
	}

	probeOption.Target, err = probeCmd.Flags().GetString("target-ip")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for target-ip: %w", err)
	}
	probeOption.TargetFile, err = probeCmd.Flags().GetString("target-file")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for target-file: %w", err)
	}
	return globalOpts, probeOption, nil
}

func init() {
	probeCmd = &cobra.Command{
		Use:   "probe",
		Short: "collect pentest environment information",
		Run:   runProbe,
	}

	probeCmd.Flags().IntP("port", "p", 135, "target port")
	probeCmd.Flags().StringP("plugin", "x", "", "plugin to scan(e.g. OXID)")
	probeCmd.Flags().StringP("target-ip", "i", "", "ip range to probe for(e.g. 192.168.1.1/24)")
	probeCmd.Flags().StringP("target-file", "I", "", "File to probe for(e.g. ip.txt)")

	//if err := probeCmd.MarkPersistentFlagRequired("plugin"); err != nil {
	//	log.Fatalf("on marking flag as required: %v", err)
	//	//log.Fatalf("error on marking flag as required: %v", err)
	//}

	//probeCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
	//
	//}
	rootCmd.AddCommand(probeCmd)
}