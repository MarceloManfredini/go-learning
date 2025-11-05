O QUE FAZ A API:
API para atender a requisições CRUD de um cadastro de tarefas.

ONDE OS REGISTROS SÃO GRAVADOS:
Em base de dados SQLite local.

QUAIS OS REQUESTS DA API:
- POST - criar tarefa com dados do json
  ex: http://localhost:8080/api/tasks
  body json:
      {
        "title": "comprar banana",
        "detail": "1 dúzia"
    }
  retorno: 201-created
{
    "id": 5,
    "title": "comprar MAÇA",
    "detail": "1 dúzia",
    "done": false,
    "created_at": "2025-11-05T13:42:05.6068984-03:00",
    "updated_at": "2025-11-05T13:42:05.6068984-03:00"
}

- GET - consultar tarefa com o ID informado
ex: http://localhost:8080/api/tasks/5
retorno:200 - ok
{
    "id": 5,
    "title": "comprar MAÇA",
    "detail": "1 dúzia",
    "done": false,
    "created_at": "2025-11-05T13:42:05.6068984-03:00",
    "updated_at": "2025-11-05T13:42:05.6068984-03:00"
}
* se não existir o ID o retorno será 404 - not found
{
    "error": "task not found"
}

- GET - listar todas as tarefas
ex.: http://localhost:8080/api/tasks
retorno: 200 - ok
[
    {
        "id": 6,
        "title": "comprar banana",
        "detail": "1 dúzia",
        "done": false,
        "created_at": "2025-11-05T13:54:15.1141828-03:00",
        "updated_at": "2025-11-05T13:54:15.1141828-03:00"
    },
    {
        "id": 7,
        "title": "comprar maça",
        "detail": "1/2 dúzia",
        "done": false,
        "created_at": "2025-11-05T13:55:50.2765107-03:00",
        "updated_at": "2025-11-05T13:55:50.2765107-03:00"
    }
]
*se não existirem registros o retorno será 200 - ok
[]

- PUT - atualizar tarefa (ID informado) com dados do json
ex: http://localhost:8080/api/tasks/5
retorno: 200 - ok
{
    "message": "updated"
}
* se não existir o ID o retorno será 404 - not found
{
    "error": "record not found"
}

- DELETE - excluir tarefa com ID informado
ex.: http://localhost:8080/api/tasks/5
retorno: 204 - no content
{
    
}
* se não existir o ID o retorno será 404 - not found
{
    "error": "task not found"
}

COMO EXECUTAR NO TERMINAL VSCODE OU NO PROMPT DE COMANDO:
go run main.go
ou pelo menu RUN do VSCODE

COMO RODAR O CONTAINER NO TERMINAL VSCODE OU NO PROMPT DE COMANDO:
docker build -t task-api .
docker run -p 8080:8080 task-api