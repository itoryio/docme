package agent

import (
	"fmt"

	"github.com/itoryio/docme/storage"

	"github.com/itoryio/docme/driver"
	"github.com/spf13/cobra"
)

//Config describe plugin configuration
type Config struct {
	storagePath  string
	multihosMode bool
}

var config = &Config{}

//AgentCmd define daemon command
var AgentCmd = &cobra.Command{
	Use:   "agent",
	Short: "`agent` initialization docker volume driver",
	Run: func(cmd *cobra.Command, args []string) {
		//cmd.Usage()
		initialization()
	},
}

func initialization() {

	// 2. создать tcp или unix socket
	var h *driver.Handler

	if config.multihosMode {
		d := &driver.DocmeMultihostDriver{Storage: storage.Init()}
		h = driver.NewHandler(d)
	} else {
		d := &driver.DocmeLocalDriver{Storage: storage.Init()}
		h = driver.NewHandler(d)
	}

	err := h.ServeTCP("docme", "localhost:8080", nil)
	//err := h.ServeUnix("docme", os.Getegid())
	fmt.Printf("Data: %v", h)
	if err != nil {
		fmt.Println(err.Error())
	}

}

func init() {
	AgentCmd.Flags().StringVar(&config.storagePath, "storage-path", "/Volumes/docme", "Path where agent will mount volumes")
	AgentCmd.Flags().BoolVar(&config.multihosMode, "multihost", false, "Set agent work mode multihost or local. Default local")
	/*
		DaemonCmd.Flags().BoolVarP(&dryRun, "dry-run", "d", false, "do everything except remove the blobs")
		DaemonCmd.Flags().StringVar(&conf.path, "config", ".", "config file")
		DaemonCmd.Flags().IntVar(&conf.port, "port", 8080, "listen port for api requests")
	*/
}
