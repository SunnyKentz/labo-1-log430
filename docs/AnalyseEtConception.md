### Analyse des besoins fonctionnels et non-fonctionnels du système.
---
#### Besoins fonctionnels
- Rechercher un produit (par identifiant, nom ou catégories) dans une liste de produit sur une base de donnée
- Enregistrer une transaction (sélection de produits et calcul du total) dans une base de donnée, une
- Enregistrer une transaction de retour de produit.
- Consulter l'état du stock des produits mise à jours après les transactions
    

#### Besoins non-fonctionnels
- Performance et scalabilité
  - Le système doit supporter simultanément 3 caisses en opération

- Sécurité
  - Journalisation(Logging) de toutes les transactions pour audit

### Proposition d’architecture sous forme de vues UML selon le modèle 4+1
---
- Vue logique :

  | MDD                                 | Diagramme de classe                    |
  | ------------------------------------| ---------------------------------------|
  | ![Modèle du domaine](../out/MDD.svg)| ![Modèle du domaine](../out/class.svg) |

- Vue des processus :
- Vue de déploiement :s
- Vue d’implémentation :
- Vue des cas d’utilisation :
![Diagramme de Cas d'utilisation](../out/ucd.svg)