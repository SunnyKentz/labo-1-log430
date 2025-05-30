## Labo 1

### Explication de l'application:

l'application s'execute sur le terminal

### Comment run :
```
    docker-compose up
    make run
```
vous pouvez ouvrir d'autres caisses sur d'autres terminals avec : `make run`

### Comment tester :
```
    make test
```

### Explication du CI
Après avoir fait un push, Github action check le linting du push,<br> execute les testes et si tous passe, le push vers dockerhub