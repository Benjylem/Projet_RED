# 🏢 Pôle Emploi Quest

> Un RPG textuel humoristique où vous incarnez un demandeur d'emploi naviguant dans les méandres administratifs de Pôle Emploi !

## 📖 Description

**Pôle Emploi Quest** est un jeu de rôle textuel développé en Go qui transforme l'expérience de la recherche d'emploi en aventure épique. Incarnez un personnage aux origines régionales diverses et survivez aux défis bureaucratiques tout en gérant vos indemnisations !

## ✨ Fonctionnalités

### 🎭 Création de Personnage
- **3 Classes régionales** avec des caractéristiques uniques :
  - **Sud** : Équipé d'un béret, chemise et mocassins (350 jours d'indemnisation)
  - **Est** : Casquette, sweat et baskets (300 jours d'indemnisation)
  - **Nord** : Bob, marcel et claquettes-chaussettes (400 jours d'indemnisation)

### ⚔️ Système de Missions
- **10 missions thématiques** inspirées de la réalité Pôle Emploi :
  - Remplir son CV
  - Entretiens simulés
  - Ateliers de motivation
  - Formations courtes
  - Et bien d'autres défis administratifs !

- **3 niveaux de difficulté** :
  - 🟢 **Facile** : 75% de réussite
  - 🟡 **Moyenne** : 45% de réussite
  - 🔴 **Difficile** : 15% de réussite

### 🎒 Gestion d'Inventaire
- **Inventaire limité** à 10 objets
- **Système de craft** pour créer des équipements
- **Équipements évolutifs** qui améliorent vos caractéristiques

### 🛠️ Système de Craft
Fabriquez des équipements d'aventurier avec des matériaux obtenus en mission :
- **Bob** : +10 jours d'indemnisation
- **Marcel** : +25 jours d'indemnisation  
- **Claquette chaussette** : +15 jours d'indemnisation

### 🏪 Commerce
- **Épicerie (Forgeron)** : Créez vos équipements
- **CAF (Marchand)** : Achetez des consommables avec votre XP

### 💊 Système de Consommables
- **Oasis** : Restaure 40 jours d'indemnisation
- **Café réchauffé** : Effet de poison mais bonus XP

## 🚀 Installation et Lancement

### Prérequis
- Go 1.16 ou supérieur

### Installation
```bash
git clone [votre-repo]
cd pole-emploi-quest
go mod init pole-emploi-quest
go run main.go
```

## 🎮 Comment Jouer

### 1. Création du Personnage
```
Entrez le nom de votre personnage (lettres uniquement) : Jean
Choisissez une classe (sud / est / nord) : nord
```

### 2. Menu Principal
```
🏢 Bienvenue à Pôle Emploi Quest 🏢
══════════════════════════════════════
[1] 📜 Afficher les informations du personnage
[2] 🎒 Accéder à l'inventaire
[3] ⚔️ Chercher une mission
[4] 📅 Consulter les jours d'indemnisation
[5] 🏪 Epicerie
[6] 🏪 Aller à la CAF (Marchand)
[7] 🎉 Qui sont les artistes?
[8] ❌ Quitter le jeu
```

### 3. Stratégie de Jeu
- **Gérez vos jours d'indemnisation** : C'est votre barre de vie !
- **Collectez des matériaux** en réussissant des missions
- **Craftez des équipements** pour augmenter vos jours max
- **Utilisez l'XP** intelligemment au marchand

## 🎯 Mécaniques de Jeu

### 💰 Système d'Expérience
- Gagnez **100 XP** par mission réussie
- Dépensez vos XP pour acheter des consommables
- Utilisez **5 XP** pour crafter des équipements

### ❤️ Jours d'Indemnisation (Points de Vie)
- Diminuent de **10 jours** en cas d'échec de mission
- Se restaurent avec des consommables
- Augmentent définitivement avec de meilleurs équipements

### 🎲 Système de Chances
Les probabilités de réussite dépendent de la difficulté :
- Missions **Faciles** : Idéales pour débuter
- Missions **Moyennes** : Risk/reward équilibré  
- Missions **Difficiles** : Hautes récompenses, hauts risques

## 🎨 Easter Eggs

Le jeu contient plusieurs références humoristiques :
- **Équipements régionaux** stéréotypés avec humour
- **Missions** inspirées de situations réelles de demandeurs d'emploi

## 🛠️ Architecture Technique

### Structure du Projet
```
├── main.go              # Point d'entrée
├── code/
│   ├── character.go     # Gestion des personnages
│   ├── menu.go         # Interface principale
│   ├── mission.go      # Système de missions
│   ├── inventory.go    # Gestion inventaire
│   ├── forge.go        # Système de craft
│   ├── health.go       # Gestion mort/résurrection
│   └── potion.go       # Système consommables
```

### Types Principaux
```go
type Character struct {
    Name            string
    Class           string
    Level           int
    MaxCompDay      int
    CurrentCompDay  int
    Experience      int
    Inventory       []string
    Equipment       Equipment
}

type Mission struct {
    Nom         string
    Exp         int
    Difficulte  string
    EchecEffet  string
    Loot        string
}
```

## 🎮 Conseils de Survie

1. **Commencez par des missions faciles** pour accumuler XP et matériaux
2. **Craftez rapidement** des équipements pour augmenter vos jours max
3. **Gardez toujours une potion** en cas d'urgence
4. **Attention à l'inventaire plein** : libérez de la place régulièrement
5. **Le café réchauffé** est dangereux mais rentable en XP !

## 🏆 Objectifs

- **Survivre** le plus longtemps possible
- **Équiper** votre personnage au maximum
- **Explorer** toutes les mécaniques du jeu
- **S'amuser** avec l'humour administratif français !

## 👥 Crédits

**Développé avec amour et café réchauffé par :**
- Benjamin
- Jérémi
- Camille
- Alex

---

*"Dans Pôle Emploi Quest, même les formulaires sont une aventure !"* 📋✨