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

### Choix technologiques
- J'ai choisi Go comme language pour ce projet, car j'aime ce language et je veux en apprendre plus dessus.
- J'ai choisi make comme build tool pour me facilité la vie, car je peux exectuter plusieurs commandes rapidement.
- J'ai choisi PosgreSQL car je trouve que c'est une base de donnée plus rapide et plus facile que MySQL, et est plus convenable à l'application que NoSQL.
- J'ai choisi GORM pour le ORM car c'est le meilleur pour go
- J'ai choisi BubbleTea pour mon ui du terminal car c'est le plus répendu et le mieux maintenu.
- J'ai choisi Docker car docker est mieux connue que les autres outils de conteneurs.