package utils

// Importa a biblioteca de criptografia do Go
import (
    "golang.org/x/crypto/bcrypt"
)

// HashPassword gera um hash seguro utilizando o algoritmo Bcrypt a partir de uma senha em texto plano
func HashPassword(password string) (string, error) {
    // Gera o hash com o custo padrão (DefaultCost) que equilibra segurança e velocidade
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(bytes), nil
}

// CheckPassword compara uma senha em texto plano com um hash Bcrypt salvo no banco para verificar se batem
func CheckPassword(hashedPassword, password string) bool {
    // Compara o hash com a senha digitada
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
    // Retorna true se a senha estiver correta (erro nil) ou false caso contrário
    return err == nil
}
