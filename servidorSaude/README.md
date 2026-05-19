# Servidor de Saúde com Go (servidorSaude) 🩺

Este projeto é um servidor HTTP robusto desenvolvido em **GoLang** que permite o gerenciamento completo de prontuários de pacientes (CRUD: Criação, Autenticação, Atualização e Exclusão). Ele utiliza o **PostgreSQL** como banco de dados relacional e fornece uma interface web moderna, responsiva e estilizada para interação dos usuários e profissionais de saúde.

---

## 🚀 Diferenciais de Segurança e UX
*   **Criptografia Bcrypt**: Diferente do MD5 básico, este servidor utiliza o algoritmo de segurança industrial **Bcrypt** para realizar o hashing e a validação de senhas dos pacientes, protegendo contra ataques de força bruta e colisões.
*   **Controle Manual do Esquema de Banco**: Maior controle e segurança operacional, delegando a criação do banco de dados `clinica` e da tabela `patients` ao administrador através de comandos SQL simples e diretos antes de subir o servidor.
*   **Design Clínico Hospitalar Premium (Vanilla CSS)**: Interfaces modernas e exclusivas em tons sofisticados de vermelho carmim e branco, conceitos elegantes de *glassmorphism* leve e total responsividade para dispositivos móveis e desktops.

---

## 📂 Estrutura de Pastas e Arquivos

O projeto está organizado exatamente seguindo os padrões recomendados de arquitetura limpa:

```
servidorSaude/
├── .env.example
├── docker-compose.yml.example
├── go.mod
├── go.sum
├── .gitignore
├── README.md
├── app/
│   ├── main.go
│   ├── handlers/
│   │   ├── formHandler.go
│   │   ├── loginHandler.go
│   │   ├── updatePatientHandler.go
│   │   └── deletePatientHandler.go
│   └── utils/
│       ├── connectToDB.go
│       ├── encrypt.go
│       ├── validatePatient.go
│       ├── updatePatient.go
│       ├── deletePatient.go
│       └── getPatientByEmail.go
└── static/
    ├── index.html
    ├── forms/
    │   ├── createPatient.html
    │   ├── login.html
    │   ├── updatePatient.html
    │   └── deletePatient.html
    └── styles/
        ├── index.style.css
        ├── createPatient.style.css
        ├── login.style.css
        ├── updatePatient.style.css
        └── deletePatient.style.css
```

### Descrição dos Componentes Principais

*   **`app/main.go`**: Inicializador do sistema. Conecta ao banco de dados e registra os middlewares e as rotas HTTP de controle.
*   **`app/handlers/`**: Controladores que interceptam as requisições HTTP, processam formulários, criptografam senhas e respondem com os fluxos correspondentes do paciente.
*   **`app/utils/`**: Utilitários centrais de banco de dados (conexão, queries CRUD dinâmicas e verificação do banco) e segurança (algoritmos Bcrypt).
*   **`static/`**: Camada visual com páginas estáticas HTML e estilizações CSS puro organizadas individualmente por funcionalidade.

---

## 🛠️ Pré-requisitos
1.  **GoLang**: Certifique-se de ter o Go instalado (Recomendado v1.18+).
2.  **PostgreSQL**: Banco de dados relacional (instalado nativamente ou via Docker).

---

## ⚙️ Configuração e Instalação

### 1. Clonar ou Baixar o Projeto
```bash
git clone <URL_DO_REPOSITORIO>
cd servidorSaude
```

### 2. Configurar Variáveis de Ambiente (`.env`)
Crie um arquivo chamado **`.env`** na raiz do projeto e configure as credenciais do seu banco de dados local:
```ini
DB_USER=postgres
DB_PASSWORD=sua_senha_do_postgres
DB_NAME=clinica
DB_HOST=localhost
DB_PORT=5432
```

### 3. Banco de Dados (PostgreSQL)

Antes de iniciar o servidor Go, você precisa criar manualmente o banco de dados e a tabela dos pacientes no seu PostgreSQL:

1. Abra o terminal do seu PostgreSQL (`psql`) ou acesse a ferramenta gráfica (como o **pgAdmin**) e crie o banco de dados do projeto:
   ```sql
   CREATE DATABASE clinica;
   ```

2. Conecte-se à base de dados recém-criada (`\c clinica` no console do `psql`) e execute o script SQL abaixo para estruturar a tabela `patients`:
   ```sql
   CREATE TABLE patients(
       id SERIAL PRIMARY KEY,
       name VARCHAR(100) NOT NULL,
       email VARCHAR(150) UNIQUE NOT NULL,
       birth_date DATE NOT NULL,
       blood_type VARCHAR(3),
       phone VARCHAR(20),
       password VARCHAR(255),
       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
   );
   ```

### 4. Instalar Dependências e Iniciar o Servidor
No diretório raiz do projeto, execute os comandos no terminal:
```bash
# Limpa e instala as dependências externas do Go
go mod tidy

# Executa o servidor HTTP
go run app/main.go
```

A aplicação estará disponível em: **`http://localhost:3000/`**

---

## 🌐 Rotas Principais Mapeadas

*   `http://localhost:3000/` -> Serve o arquivo portal de entrada [`index.html`].
*   `http://localhost:3000/create` -> Processa a inserção (POST) de um novo paciente a partir de `forms/createPatient.html`.
*   `http://localhost:3000/login` -> Autentica (POST) o paciente via Bcrypt a partir de `forms/login.html` e carrega o prontuário.
*   `http://localhost:3000/update` -> Executa a alteração dinâmica (POST) do telefone ou tipo sanguíneo a partir de `forms/updatePatient.html`.
*   `http://localhost:3000/delete` -> Remove permanentemente (POST) o prontuário do paciente do banco a partir de `forms/deletePatient.html`.

---

## 🔒 Licença e Observações
*   O projeto foi estruturado com fins acadêmicos e profissionais de desenvolvimento.
*   **Segurança**: O arquivo `.env` que armazena as suas senhas físicas de banco local já está devidamente listado no arquivo `.gitignore`, garantindo que suas credenciais pessoais **nunca** sejam expostas publicamente no GitHub!
