### Justification des décisions d’architecture (ADR)
#### ADR Choix de la platform:
  - Titre : Choix de Native pour l'application de la caisse.
  - Status : Le choix de rendre l'apllication executable nativement est déjà implémenté.
  - Contexte : Je dois absolument choisir la platforme car c'est un logiciele utilisateur et les logiciel utilisateurs fonctionne sur une platforme en personne.
  - Décision : L'application va donc être executer nativement sur la machine.
  - Conséquences : Toute machine non fonctionnels doivent être maintenue en personne et non à distance.
  - Compliance : Ce choix n'affecte pas le finctionnement du système ni des autres caisses connecté.

#### ADR Choix de mécanisme de base de donnée:
  - Titre : Choix d'une base de donnée PosgreSQL en serveur.
  - Status : Le choix de rendre d'une base de donnée PosgeSQL est déjà implémenté.
  - Contexte : Pour persisté les données à travers les caisses, il faut choisir une base de donnée externe.
  - Décision : Le système va donc utilié une base de donnée SQL en serveur.
  - Conséquences : Le système dépend de l'internet et doit utiliser SQL ou un ORM.
  - Compliance : Ce choix n'affecte pas le finctionnement du système ni des autres caisses connecté.