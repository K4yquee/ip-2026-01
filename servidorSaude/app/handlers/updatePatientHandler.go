package handlers

import (
	"servidorSaude/app/utils"
	"net/http"
	"text/template"
)

// Constante contendo o template do prontuário atualizado com sucesso
const patientUpdatedProfileHTML = `<!DOCTYPE html>
<html lang="pt-br">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Clínica Saúde - Prontuário Atualizado</title>
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
            --color-success: #10b981;
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
            from { opacity: 0; transform: translateY(20px); }
            to { opacity: 1; transform: translateY(0); }
        }

        .alert-success {
            background: rgba(16, 185, 129, 0.08);
            border: 2px solid #1e293b;
            border-radius: 12px;
            padding: 12px;
            margin-bottom: 24px;
            color: #065f46;
            font-size: 0.9rem;
            font-weight: 600;
            box-shadow: 3px 3px 0px #1e293b;
        }

        h1 {
            font-size: 1.8rem;
            font-weight: 700;
            margin-bottom: 8px;
            background: linear-gradient(to right, var(--accent-red), var(--accent-glow));
            -webkit-background-clip: text;
            -webkit-text-fill-color: transparent;
        }

        .subtitle {
            color: var(--text-secondary);
            font-size: 0.95rem;
            margin-bottom: 24px;
        }

        .info-card {
            background: #ffffff;
            border: 2px solid #1e293b;
            border-radius: 16px;
            padding: 24px;
            margin-bottom: 32px;
            text-align: left;
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
        <h1>Prontuário Atualizado</h1>
        <div class="subtitle">Seus dados foram atualizados com sucesso</div>
        
        <div class="alert-success">
            Alterações gravadas no banco de dados!
        </div>

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

// UpdatePatientHandler processa a atualização cadastral do paciente de forma simplificada
func UpdatePatientHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(response, "Método não suportado", http.StatusMethodNotAllowed)
		return
	}

	// Obtém os valores do formulário simplificado
	currentEmail := request.FormValue("currentEmail")
	phone := request.FormValue("phone")
	bloodType := request.FormValue("bloodType")

	// Cria um mapa para armazenar os campos a serem atualizados
	updates := make(map[string]string)

	if phone != "" {
		updates["phone"] = phone
	}
	if bloodType != "" {
		updates["blood_type"] = bloodType
	}

	// Se houver campos a atualizar, realiza a atualização no banco
	if len(updates) > 0 {
		err := utils.UpdatePatient(currentEmail, updates)
		if err != nil {
			http.Error(response, "Erro ao atualizar os dados no banco de dados ou paciente não existe", http.StatusInternalServerError)
			return
		}
	}

	// Busca as informações completas e atualizadas do paciente
	patient, err := utils.GetPatientByEmail(currentEmail)
	if err != nil {
		http.Error(response, "Erro ao buscar informações atualizadas do paciente", http.StatusInternalServerError)
		return
	}

	// Carrega e renderiza o template de perfil atualizado
	tmpl, err := template.New("patientUpdatedProfile").Parse(patientUpdatedProfileHTML)
	if err != nil {
		http.Error(response, "Erro ao processar visualização do prontuário", http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = tmpl.Execute(response, patient)
	if err != nil {
		http.Error(response, "Erro ao gerar visualização do prontuário", http.StatusInternalServerError)
		return
	}
}
