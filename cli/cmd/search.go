package cmd

import (
    "encoding/json"
    "fmt"
    "net/http"
    "github.com/spf13/cobra"
)

type Vaga struct {
    Titulo string `json:"titulo"`
    Link   string `json:"link"`
}

type Resultado struct {
    Termo string `json:"termo"`
    Total int    `json:"total"`
    Vagas []Vaga `json:"vagas"`
}

var searchCmd = &cobra.Command{
    Use:   "search",
    Short: "Busca vagas pelo termo informado",
    Run: func(cmd *cobra.Command, args []string) {
        termo, _ := cmd.Flags().GetString("termo")
        resp, err := http.Get("http://localhost:8000/vagas?termo=" + termo)
        if err != nil {
            fmt.Println("❌ Erro ao conectar na API:", err)
            return
        }
        defer resp.Body.Close()

        var resultado Resultado
        json.NewDecoder(resp.Body).Decode(&resultado)

        fmt.Printf("🔍 Resultados para '%s' (%d vagas)\n\n", resultado.Termo, resultado.Total)
        for _, v := range resultado.Vagas {
            fmt.Printf("- %s\n  %s\n\n", v.Titulo, v.Link)
        }
    },
}

func init() {
    rootCmd.AddCommand(searchCmd)
    searchCmd.Flags().StringP("termo", "t", "help desk", "Termo da busca")
}
