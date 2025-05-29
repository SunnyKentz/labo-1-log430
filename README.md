## Labo 0

### Explication de l'application:

l'application s'execute sur le terminal

J'ai choisi Go comme language pour ce projet, car j'aime ce language et je veux en apprendre plus dessus.<br>
J'ai choisi make comme build tool pour me facilité la vie, car je peux exectuter plusieurs commandes rapidement.<br>
J'ai choisi fiber comme server, car c'est simililaire avec expresseJS qui est un server basé sur des middleware.

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