package utils

import (
    "log"
)

// DeletePatient apaga a ficha de um paciente do banco de dados pelo e-mail
func DeletePatient(email string) error {
    query := `DELETE FROM patients WHERE email = $1`
    _, err := DB.Exec(query, email)
    if err != nil {
        log.Printf("Erro ao apagar paciente do banco de dados: %v", err)
        return err
    }
    log.Println("Paciente apagado com sucesso!")
    return nil
}
