package handlers

// Importa os pacotes necessários para o funcionamento do handler
import (
    "log"
    "servidorSaude/app/utils" // Importa funções utilitárias, como criptografia e banco de dados
    "net/http"              // Usado para lidar com requisições e respostas HTTP
)

// FormHandler é responsável por processar os dados enviados pelo formulário de criação de conta do paciente
func FormHandler(response http.ResponseWriter, request *http.Request) {
    // Verifica se o método da requisição é POST
    if request.Method != http.MethodPost {
        // Retorna um erro caso o método não seja suportado
        http.Error(response, "Método não suportado", http.StatusMethodNotAllowed)
        return
    }

    // Obtém os valores enviados pelo formulário
    name := request.FormValue("name")           // Nome do paciente
    email := request.FormValue("email")         // E-mail do paciente
    birthDate := request.FormValue("birthDate") // Data de nascimento do paciente (birth_date)
    bloodType := request.FormValue("bloodType") // Tipo sanguíneo (blood_type)
    phone := request.FormValue("phone")         // Telefone (phone)
    password := request.FormValue("password")   // Senha do paciente

    // Criptografa a senha para armazená-la de forma segura no banco de dados
    encryptedPassword, err := utils.HashPassword(password)
    if err != nil {
        log.Printf("Erro ao processar a senha: %v", err)
        http.Error(response, "Erro ao processar a senha do paciente", http.StatusInternalServerError)
        return
    }

    // Insere os dados do paciente diretamente no banco de dados incluindo os novos campos
    query := `
        INSERT INTO patients (name, email, birth_date, blood_type, phone, password) 
        VALUES ($1, $2, $3, $4, $5, $6)`
    
    _, err = utils.DB.Exec(query, name, email, birthDate, bloodType, phone, encryptedPassword)
    if err != nil {
        log.Printf("Erro ao salvar paciente no banco de dados: %v", err) // LOG DA CAUSA REAL DO ERRO
        // Retorna um erro caso ocorra falha ao salvar os dados no banco de dados
        http.Error(response, "Erro ao salvar os dados no banco de dados: E-mail já cadastrado ou dados inválidos", http.StatusInternalServerError)
        return
    }

    // Redireciona o paciente para a página inicial após o sucesso
    http.Redirect(response, request, "/index.html", http.StatusSeeOther)
}
