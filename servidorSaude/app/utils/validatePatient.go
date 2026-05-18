package utils

import (
    "database/sql"
    "log"
)

// ValidatePatient verifica se o paciente existe e se a senha digitada está correta
func ValidatePatient(email, password string) (bool, error) {
    var hashedPassword string

    // Busca a senha criptografada (hash) do paciente correspondente ao e-mail informado
    query := "SELECT password FROM patients WHERE email = $1"
    err := DB.QueryRow(query, email).Scan(&hashedPassword)
    if err != nil {
        // Se não encontrar nenhuma linha, significa que o paciente não está cadastrado
        if err == sql.ErrNoRows {
            return false, nil
        }
        log.Printf("Erro ao buscar a senha para validação: %v", err)
        return false, err
    }

    // Compara o hash Bcrypt armazenado com a senha digitada pelo paciente
    isValid := CheckPassword(hashedPassword, password)
    
    return isValid, nil
}
