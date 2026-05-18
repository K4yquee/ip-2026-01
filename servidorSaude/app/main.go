package main

import (
    "fmt"
    "log"
    "net/http"
    "servidorSaude/app/handlers"
    "servidorSaude/app/utils"
)

func main() {
    // Inicializa a conexão com o banco PostgreSQL (cria a tabela automaticamente se não existir)
    utils.ConnectToDB()

    // Cria o FileServer para servir os arquivos estáticos da pasta "./static"
    fileserver := http.FileServer(http.Dir("./static"))

    // Registra a rota raiz para servir os arquivos estáticos utilizando http.Handle()
    http.Handle("/", fileserver)

    // Registra os Handlers HTTP padrão utilizando http.HandleFunc() para cada rota
    http.HandleFunc("/create", handlers.FormHandler)
    http.HandleFunc("/login", handlers.LoginHandler)
    http.HandleFunc("/update", handlers.UpdatePatientHandler)
    http.HandleFunc("/delete", handlers.DeletePatientHandler)

    // Porta em que o servidor irá escutar
    port := "3000"
    fmt.Printf("Servidor de Saúde rodando em: http://localhost:%s/\n", port)

    // Inicia o servidor HTTP utilizando http.ListenAndServe()
    if err := http.ListenAndServe("0.0.0.0:"+port, nil); err != nil {
        log.Fatalf("Erro ao iniciar o servidor HTTP: %v", err)
    }
}
