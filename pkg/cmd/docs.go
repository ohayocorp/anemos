package cmd

import (
	"fmt"
	"io/fs"
	"net/http"

	"github.com/ohayocorp/anemos/docs"
	"github.com/spf13/cobra"
)

func getDocsCommand() *cobra.Command {
	var host string
	var port int

	command := &cobra.Command{
		Use:   "docs",
		Short: "Serve documentation",
		RunE: func(cmd *cobra.Command, args []string) error {
			return serveDocs(host, port)
		},
	}

	command.Flags().IntVarP(&port, "port", "p", 9974, "Port to serve the documentation on")
	command.Flags().StringVarP(&host, "listen", "l", "localhost", "Host to listen on")

	return command
}

func serveDocs(host string, port int) error {
	docsFs, err := fs.Sub(docs.DocsFs, "build")
	if err != nil {
		return err
	}

	fmt.Println("Serving documentation at http://" + host + ":" + fmt.Sprint(port))

	http.Handle("/", http.FileServer(http.FS(docsFs)))
	return http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), nil)
}
