package handlers

import (
    "servidorSaude/app/utils"
    "net/http"
)

// DeletePatientHandler é responsável por processar a remoção de um cadastro de paciente
func DeletePatientHandler(response http.ResponseWriter, request *http.Request) {
    if request.Method != http.MethodPost {
        http.Error(response, "Método não suportado", http.StatusMethodNotAllowed)
        return
    }

    // Lê apenas o e-mail enviado pelo formulário simplificado
    email := request.FormValue("email")

    // Remove o paciente do banco de dados diretamente
    err := utils.DeletePatient(email)
    if err != nil {
        http.Error(response, "Erro ao apagar o prontuário do paciente ou paciente não existe", http.StatusInternalServerError)
        return
    }

    // Redireciona o paciente para a página inicial (Portal) após exclusão com sucesso
    http.Redirect(response, request, "/index.html", http.StatusSeeOther)
}
