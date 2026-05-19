package utils

// Importa os pacotes necessários para a conexão com o banco de dados
import (
    "database/sql"          // Usado para interagir com o banco de dados
    "fmt"                   // Usado para formatar strings
    "log"                   // Usado para registrar mensagens de log e erros
    "os"                    // Usado para acessar variáveis de ambiente

    _ "github.com/lib/pq"          // Driver PostgreSQL para o pacote database/sql
    "github.com/joho/godotenv"     // Usado para carregar variáveis de ambiente de um arquivo .env
)

// Declara uma variável global para armazenar a conexão com o banco de dados
var DB *sql.DB

// Função responsável por conectar ao banco de dados e criar a tabela se não existir
func ConnectToDB() {
    // Carrega as variáveis de ambiente do arquivo .env
    err := godotenv.Load()
    if err != nil {
        log.Println("Aviso: Não foi possível carregar o arquivo .env, buscando variáveis de ambiente globais...")
    }

    // Obtém as variáveis de ambiente necessárias para a conexão
    user := os.Getenv("DB_USER")       // Usuário do banco de dados
    password := os.Getenv("DB_PASSWORD") // Senha do banco de dados
    dbname := os.Getenv("DB_NAME")     // Nome do banco de dados
    host := os.Getenv("DB_HOST")       // Host do banco de dados
    port := os.Getenv("DB_PORT")       // Porta do banco de dados

    // Caso alguma variável venha em branco, define os padrões definidos pelo usuário
    if user == "" { user = "postgres" }
    if password == "" { password = "123456" }
    if dbname == "" { dbname = "clinica" }
    if host == "" { host = "localhost" }
    if port == "" { port = "5432" }

    // Abre a conexão definitiva com a base de dados do projeto (clinica)
    connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
        user, password, dbname, host, port)

    // Abre a conexão com o banco de dados usando a string de conexão
    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
    }

    // Testa a conexão com o banco de dados
    err = DB.Ping()
    if err != nil {
        log.Fatalf("Erro ao verificar a conexão com o banco de dados. Certifique-se de que o banco '%s' existe no seu PostgreSQL e que a tabela 'patients' foi criada manualmente: %v", dbname, err)
    }

    fmt.Println("Conexão com o banco de dados estabelecida com sucesso!")
}
