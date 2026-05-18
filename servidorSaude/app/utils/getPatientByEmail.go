package utils

import (
    "log"
)

// Patient representa a estrutura de dados expandida de um paciente clínico
type Patient struct {
    ID        int
    Name      string
    Email     string
    BirthDate string
    BloodType string
    Phone     string
    CreatedAt string
}

// GetPatientByEmail busca as informações completas de um paciente cadastrado a partir de seu e-mail
func GetPatientByEmail(email string) (*Patient, error) {
    // Seleciona todas as colunas com tratamento COALESCE para campos opcionais que podem conter NULL
    query := `
        SELECT 
            id, 
            name, 
            email, 
            birth_date::text, 
            COALESCE(blood_type, ''), 
            COALESCE(phone, ''), 
            created_at::text 
        FROM patients 
        WHERE email = $1`
    
    var patient Patient
    err := DB.QueryRow(query, email).Scan(
        &patient.ID,
        &patient.Name,
        &patient.Email,
        &patient.BirthDate,
        &patient.BloodType,
        &patient.Phone,
        &patient.CreatedAt,
    )
    if err != nil {
        log.Printf("Erro ao buscar paciente no banco de dados: %v", err)
        return nil, err
    }
    return &patient, nil
}
