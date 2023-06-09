Aqui está um exemplo de README.md para o seu projeto CLI do ClickUp em Golang:

```markdown
# ClickUp CLI

Este é um CLI (Command-Line Interface) para interagir com a API do ClickUp usando Golang.

## Pré-requisitos

Certifique-se de ter o Golang instalado no seu sistema.

## Instalação

1. Clone este repositório para o seu diretório local:

   ```bash
   git clone https://github.com/seu-usuario/clickup-cli.git
   ```

2. Navegue até o diretório do projeto:

   ```bash
   cd clickup-cli
   ```

3. Compile o código fonte:

   ```bash
   go build
   ```

4. Execute o CLI:

   ```bash
   ./clickup-cli
   ```

## Configuração

Antes de executar o CLI, você precisa configurar as seguintes variáveis de ambiente:

- `CLICKUP_BASE_URL`: A URL base da API do ClickUp.
- `CLICKUP_TOKEN`: O token de autenticação para acessar a API do ClickUp.
- `CLICKUP_SPACE_ID`: O ID do espaço no ClickUp.
- `CLICKUP_TEAM_ID`: O ID da equipe no ClickUp.
- `CLICKUP_FOLDER_ID`: O ID da pasta no ClickUp.
- `CLICKUP_LIST_ID`: O ID da lista no ClickUp.

Certifique-se de definir essas variáveis de ambiente antes de executar o CLI. Você pode fazer isso manualmente ou usando um arquivo `.env`.

## Comandos Disponíveis

- `get-user`: Obtém informações do usuário autorizado.
- `get-folders`: Obtém todas as pastas do espaço atual.
- `get-teams`: Obtém todas as equipes do usuário.
- `get-spaces`: Obtém todos os espaços da equipe atual.
- `get-lists`: Obtém todas as listas da pasta atual.
- `get-tasks`: Obtém todas as tarefas da lista atual.

Para executar um comando, use o seguinte formato:

```bash
./clickup-cli <comando>
```

## Exemplo de Uso

```bash
# Obtém informações do usuário autorizado
./clickup-cli get-user

# Obtém todas as pastas do espaço atual
./clickup-cli get-folders

# Obtém todas as equipes do usuário
./clickup-cli get-teams

# Obtém todos os espaços da equipe atual
./clickup-cli get-spaces

# Obtém todas as listas da pasta atual
./clickup-cli get-lists

# Obtém todas as tarefas da lista atual
./clickup-cli get-tasks
```

## Contribuição

Contribuições são bem-vindas! Sinta-se à vontade para abrir uma issue ou enviar um pull request.

## Licença

Este projeto está licenciado sob a [MIT License](LICENSE).
```

Lembre-se de substituir `seu-usuario` no link do repositório e personalizar o README.md conforme necessário. Certifique-se de fornecer informações adicionais, como a instalação de dependências, se houver.
