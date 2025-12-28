# Groupie Tracker

Groupie Tracker est une application web qui permet de découvrir et d'explorer des informations sur des groupes et artistes musicaux. Les utilisateurs peuvent filtrer, rechercher et afficher des détails pertinents tels que le nom, les albums, dates de création, membres du groupe, etc. Toutes les données sont récupérées dynamiquement depuis une API tierce.

## Fonctionnalités

### Backend
- **Gestion des routes :**
  - `/` : Accueil de l'application.
  - `/artist/` : Liste des artistes ou détail d’un artiste spécifique.
  - `/search` : Recherche dynamique dans les données des artistes.
  - `/filter` : Filtrage des artistes basé sur plusieurs critères.
  - `/healthz` : Vérification de l'état du serveur.
  
- **Récupération des données :**
  - Appels API parallèles pour des performances améliorées.
  - Données synchronisées concernant les artistes, leurs relations, leurs emplacements, et plus.

- **Gestion des erreurs :**
  - Pages personnalisées générées pour les erreurs (404, 500, etc.).

### Frontend
- **Interface utilisateur élégante et réactive (CSS) :**
  - Grille des artistes avec un affichage optimisé pour desktop et mobile.
  - Effets visuels comme des animations au survol (hover effect).
  - Thème visuel agréable et minimaliste.

- **JavaScript :**
  - Interaction utilisateur simple (affichage conditionnel des artistes, etc.).

### Fonctionnalités utilisateur
- **Recherche :** Trouvez rapidement des artistes par nom, album, date de création, etc.
- **Filtrage :** Filtrez les artistes par critères comme le nombre de membres, l'année de création, ou le premier album.

## Aperçu

### Capture d'écran de l'accueil :
![Accueil de Groupie Tracker](assets/static/image/homepage_screenshot.png) *(Ajoutez une capture d'écran du site ici)*

### Page détaillée d'un artiste :
![Page artiste de Groupie Tracker](assets/static/image/artist_screenshot.png)

### Exemple de grille d'artistes filtrés :
![Grille des artistes](assets/static/image/filtered_artists_screenshot.png)

## Technologies utilisées

### Backend
- Langage principal : **Go (Golang)**
- Architecture modulaire.
- Frameworks/utilitaires :
  - `net/http` pour la gestion des requêtes HTTP.
  - Gestion des templates HTML Go.

### Frontend
- **HTML** / **CSS** : Génération de modèles dynamiques.
- **JavaScript** : Interaction client minimaliste.
- **Rendu dynamique** : Les données sont converties et injectées dans les pages HTML côté serveur.

### API tierce
- L’application utilise l'API suivante pour récupérer des informations :  
  **[Groupie Trackers API](https://groupietrackers.herokuapp.com/api)**

#### Points de terminaison API utilisés :
- `/artists` : Liste des artistes.
- `/locations` : Lieux associés aux artistes.
- `/relation` : Relations des artistes avec des événements.
- `/dates` : Dates importantes des concerts et événements.

### Configuration
#### Ports et connectivité
- Adresse de connexion locale : `http://localhost:8080`
- Port serveur : **8080**

#### Dépendances Go
N'oubliez pas d'initialiser les modules dans votre environnement de développement avec la commande :
```bash
go mod init groupie-tracker
```

## Installation

### Prérequis
- **Go (Golang)** : Version 1.18 ou supérieure.
- Un navigateur web moderne.

### Étapes
1. Clonez ce repository :
   ```bash
   git clone https://github.com/IronBeagle404/groupie-tracker.git
   cd groupie-tracker
   ```
2. Téléchargez les dépendances Go :
   ```bash
   go mod tidy
   ```
3. Lancez le serveur :
   ```bash
   go run main.go
   ```
4. Accédez à votre navigateur à l'adresse suivante :
   ```
   http://localhost:8080
   ```

## Contribution

Les contributions sont les bienvenues ! Voici comment contribuer :
1. Forkez le repository.
2. Créez une branche pour vos modifications :
   ```bash
   git checkout -b feature/ma-fonctionnalite
   ```
3. Effectuez vos changements et poussez la branche :
   ```bash
   git push origin feature/ma-fonctionnalite
   ```
4. Ouvrez une Pull Request pour revue !

## Problèmes connus
- Les filtres multi-critères peuvent être ralentis lorsque de grands ensembles de données sont récupérés.
- La gestion des sessions utilisateurs et des cookies n’est pas encore implémentée.

## Licence

Ce projet n'a pas encore de licence explicite. Si ajouté, veillez à respecter les règles associées.

---

## Auteur

Créé par [IronBeagle404](https://github.com/IronBeagle404).  
