package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "jobfinder",
    Short: "Buscador de vagas para iniciantes em TI",
    Long:  "Um CLI simples que consome a API Python e exibe vagas de help desk, suporte, monitoramento e estágios.",
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
    }
}
