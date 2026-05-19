package handlers

// Importa os pacotes necessários para o funcionamento do handler
import (
    "fmt"                  // Usado para depuração
    "servidorSaude/app/utils" // Importa funções utilitárias do paciente
    "net/http"            // Usado para lidar com requisições e respostas HTTP
    "text/template"       // Usado para renderizar templates HTML
)

// Template HTML inline para exibir o perfil do paciente com todos os novos campos clínicos
const patientProfileHTML = `<!DOCTYPE html>
<html lang="pt-br">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Clínica Saúde - Prontuário Digital</title>
    <link href="https://fonts.googleapis.com/css2?family=Outfit:wght@300;400;600;700&display=swap" rel="stylesheet">
    <style>
        :root {
            --bg-gradient: radial-gradient(circle at top right, #fff1f2 0%, #fafafa 70%);
            --card-bg: rgba(255, 255, 255, 0.88);
            --border-color: rgba(225, 29, 72, 0.08);
            --accent-red: #be123c; /* Vermelho carmim profundo, elegante */
            --accent-glow: #e11d48; /* Vermelho cereja vibrante */
            --text-primary: #1e293b; /* Slate escuro */
            --text-secondary: #64748b; /* Slate cinza */
            --shadow-soft: 0 15px 35px rgba(190, 18, 60, 0.05), 0 1px 3px rgba(0, 0, 0, 0.02);
        }

        * {
            box-sizing: border-box;
            margin: 0;
            padding: 0;
            font-family: 'Outfit', sans-serif;
            transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
        }

        body {
            background: var(--bg-gradient);
            min-height: 100vh;
            display: flex;
            justify-content: center;
            align-items: center;
            color: var(--text-primary);
            padding: 20px;
        }

        .container {
            background: var(--card-bg);
            backdrop-filter: blur(16px);
            -webkit-backdrop-filter: blur(16px);
            border: 1px solid var(--border-color);
            border-radius: 24px;
            padding: 40px;
            width: 100%;
            max-width: 500px;
            box-shadow: var(--shadow-soft);
            text-align: center;
            position: relative;
            overflow: hidden;
            animation: fadeIn 0.8s cubic-bezier(0.16, 1, 0.3, 1);
        }

        @keyframes fadeIn {
            from {
                opacity: 0;
                transform: translateY(20px);
            }
            to {
                opacity: 1;
                transform: translateY(0);
            }
        }

        h1 {
            font-size: 1.8rem;
            font-weight: 700;
            margin-bottom: 8px;
            background: linear-gradient(to right, var(--accent-red), var(--accent-glow));
            -webkit-background-clip: text;
            -webkit-text-fill-color: transparent;
            position: relative;
            z-index: 1;
        }

        .subtitle {
            color: var(--text-secondary);
            font-size: 0.95rem;
            margin-bottom: 30px;
            position: relative;
            z-index: 1;
        }

        .info-card {
            background: #ffffff;
            border: 2px solid #1e293b;
            border-radius: 16px;
            padding: 24px;
            margin-bottom: 30px;
            text-align: left;
            position: relative;
            z-index: 1;
            display: flex;
            flex-direction: column;
            gap: 16px;
            box-shadow: 4px 4px 0px #1e293b;
        }

        .info-item {
            border-bottom: 1px solid rgba(225, 29, 72, 0.05);
            padding-bottom: 12px;
        }

        .info-item:last-child {
            border-bottom: none;
            padding-bottom: 0;
        }

        .info-label {
            font-size: 0.75rem;
            text-transform: uppercase;
            letter-spacing: 0.05em;
            color: var(--accent-red);
            font-weight: 600;
            margin-bottom: 4px;
        }

        .info-value {
            font-size: 1.05rem;
            font-weight: 400;
            color: var(--text-primary);
        }

        .btn-portal {
            display: inline-block;
            width: 100%;
            padding: 14px;
            background: linear-gradient(135deg, var(--accent-red) 0%, #9f1239 100%);
            border: 2px solid #1e293b;
            border-radius: 12px;
            color: white;
            font-size: 1rem;
            font-weight: 600;
            text-decoration: none;
            transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
            box-shadow: 4px 4px 0px #1e293b;
            cursor: pointer;
            position: relative;
            z-index: 1;
        }

        .btn-portal:hover {
            transform: translate(-3px, -3px);
            box-shadow: 7px 7px 0px #1e293b;
            background: linear-gradient(135deg, var(--accent-glow) 0%, var(--accent-red) 100%);
        }

        .btn-portal:active {
            transform: translate(2px, 2px);
            box-shadow: 2px 2px 0px #1e293b;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>Prontuário Digital</h1>
        <div class="subtitle">Bem-vindo à sua área de saúde</div>
        
        <div class="info-card">
            <div class="info-item">
                <div class="info-label">Nome Completo</div>
                <div class="info-value">{{.Name}}</div>
            </div>
            <div class="info-item">
                <div class="info-label">E-mail de Cadastro</div>
                <div class="info-value">{{.Email}}</div>
            </div>
            <div class="info-item">
                <div class="info-label">Data de Nascimento</div>
                <div class="info-value">{{.BirthDate}}</div>
            </div>
            <div class="info-item">
                <div class="info-label">Tipo Sanguíneo</div>
                <div class="info-value">{{if .BloodType}}{{.BloodType}}{{else}}Não informado{{end}}</div>
            </div>
            <div class="info-item">
                <div class="info-label">Telefone de Contato</div>
                <div class="info-value">{{if .Phone}}{{.Phone}}{{else}}Não informado{{end}}</div>
            </div>
            <div class="info-item">
                <div class="info-label">Data de Registro no Sistema</div>
                <div class="info-value">{{.CreatedAt}}</div>
            </div>
        </div>

        <a href="/index.html" class="btn-portal">Voltar ao Portal</a>
    </div>
</body>
</html>`

// LoginHandler é responsável por processar os dados enviados pelo formulário de login
func LoginHandler(response http.ResponseWriter, request *http.Request) {

    // Verifica se o método da requisição é POST
    if request.Method != http.MethodPost {
        // Retorna um erro caso o método não seja suportado
        http.Error(response, "Método não suportado", http.StatusMethodNotAllowed)
        return
    }

    // Obtém os valores enviados pelo formulário
    email := request.FormValue("email")       // E-mail do paciente
    password := request.FormValue("password") // Senha do paciente

    // Verifica se o paciente existe no banco de dados e se a senha está correta
    isValidPatient, err := utils.ValidatePatient(email, password)
    if err != nil {
        // Retorna um erro caso ocorra falha ao validar o paciente
        http.Error(response, "Erro ao validar o paciente", http.StatusInternalServerError)
        return
    }

    // Verifica se as credenciais são inválidas
    if !isValidPatient {
        // Retorna um erro caso as credenciais estejam incorretas
        http.Error(response, "Credenciais inválidas", http.StatusUnauthorized)
        fmt.Println("Erro de login: Credenciais incorretas para", email) // Mensagem de debug
        return
    }

    // Busca as informações do paciente no banco de dados
    patient, err := utils.GetPatientByEmail(email)
    if err != nil {
        // Retorna um erro caso ocorra falha ao buscar as informações do paciente
        http.Error(response, "Erro ao buscar informações do paciente", http.StatusInternalServerError)
        return
    }

    // Cria um template a partir do HTML inline
    tmpl, err := template.New("patientProfile").Parse(patientProfileHTML)
    if err != nil {
        // Retorna um erro caso ocorra falha ao analisar o template
        http.Error(response, "Erro ao analisar o prontuário", http.StatusInternalServerError)
        return
    }

    // Define o cabeçalho de resposta e renderiza o template com os dados do paciente
    response.Header().Set("Content-Type", "text/html; charset=utf-8")
    err = tmpl.Execute(response, patient)
    if err != nil {
        // Retorna um erro caso ocorra falha ao renderizar o template
        http.Error(response, "Erro ao gerar o prontuário do paciente", http.StatusInternalServerError)
        return
    }
}
