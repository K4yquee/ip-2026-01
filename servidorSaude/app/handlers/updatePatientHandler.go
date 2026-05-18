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
        * {
            box-sizing: border-box;
            margin: 0;
            padding: 0;
            font-family: 'Outfit', sans-serif;
        }

        body {
            background: linear-gradient(135deg, #0f172a 0%, #1e1b4b 50%, #0d9488 100%);
            min-height: 100vh;
            display: flex;
            justify-content: center;
            align-items: center;
            color: #f8fafc;
            padding: 20px;
        }

        .container {
            background: rgba(255, 255, 255, 0.05);
            backdrop-filter: blur(16px);
            -webkit-backdrop-filter: blur(16px);
            border: 1px solid rgba(255, 255, 255, 0.1);
            border-radius: 24px;
            padding: 40px;
            width: 100%;
            max-width: 500px;
            box-shadow: 0 20px 40px rgba(0, 0, 0, 0.3), inset 0 1px 0 rgba(255, 255, 255, 0.1);
            text-align: center;
            position: relative;
            overflow: hidden;
            animation: fadeIn 0.8s cubic-bezier(0.16, 1, 0.3, 1);
        }

        @keyframes fadeIn {
            from { opacity: 0; transform: translateY(20px); }
            to { opacity: 1; transform: translateY(0); }
        }

        .avatar-glow {
            width: 80px;
            height: 80px;
            border-radius: 50%;
            background: linear-gradient(135deg, #10b981, #059669);
            margin: 0 auto 24px;
            display: flex;
            align-items: center;
            justify-content: center;
            font-size: 2rem;
            color: white;
            box-shadow: 0 0 25px rgba(16, 185, 129, 0.6);
            border: 2px solid rgba(255, 255, 255, 0.2);
        }

        .alert-success {
            background: rgba(16, 185, 129, 0.15);
            border: 1px solid rgba(16, 185, 129, 0.3);
            border-radius: 12px;
            padding: 12px;
            margin-bottom: 24px;
            color: #34d399;
            font-size: 0.9rem;
            font-weight: 600;
        }

        h1 {
            font-size: 1.8rem;
            font-weight: 700;
            margin-bottom: 8px;
            background: linear-gradient(to right, #38bdf8, #34d399);
            -webkit-background-clip: text;
            -webkit-text-fill-color: transparent;
        }

        .subtitle {
            color: #94a3b8;
            font-size: 0.95rem;
            margin-bottom: 24px;
        }

        .info-card {
            background: rgba(15, 23, 42, 0.4);
            border: 1px solid rgba(255, 255, 255, 0.05);
            border-radius: 16px;
            padding: 24px;
            margin-bottom: 32px;
            text-align: left;
            display: flex;
            flex-direction: column;
            gap: 16px;
        }

        .info-item {
            border-bottom: 1px solid rgba(255, 255, 255, 0.05);
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
            color: #34d399;
            font-weight: 600;
            margin-bottom: 4px;
        }

        .info-value {
            font-size: 1.05rem;
            font-weight: 400;
            color: #e2e8f0;
        }

        .btn-portal {
            display: inline-block;
            width: 100%;
            padding: 14px;
            background: linear-gradient(135deg, #0d9488 0%, #0f766e 100%);
            border: none;
            border-radius: 12px;
            color: white;
            font-size: 1rem;
            font-weight: 600;
            text-decoration: none;
            transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
            box-shadow: 0 4px 12px rgba(13, 148, 136, 0.3);
            cursor: pointer;
        }

        .btn-portal:hover {
            transform: translateY(-2px);
            box-shadow: 0 8px 20px rgba(13, 148, 136, 0.5);
            background: linear-gradient(135deg, #14b8a6 0%, #0d9488 100%);
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="avatar-glow">✅</div>
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
                <div class="info-value{{if .Phone}}{{.Phone}}{{else}}Não informado{{end}}"></div>
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
