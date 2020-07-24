# GRADUEL

Un RAD pour développer un CRUD en GO dans un container DOCKER

**G Rad uel**
- **G** comme Go le langage que je vais pratiquer pour la 1ère fois
- **rad** comme "Rapid Application Development"
  pour développer un CRUD "Create Read Update Delete"
  Utilisation d'un dictionnaire JSON pour déclarer 
  - les rubriques, les vues et formulaires du CRUD
- tout ça progressivement -> grad**uel**ement

Application qui sera mise dans un container DOCKER

## Organisation des composants à gérer

- le **dictionnaire** graduel.json
- une **base de données** Sqlite pour démarrer
- le **portail** web qui présentera les **vues** et **formulaires**
- les **modèles** de page 
- la **feuille de style** bootstrap-material-design

## Références et commandes utiles

Développement sous **Debian Buster** avec **VSCodium**
- sudo apt install golang
- cd le répertoire racine du projet
- export GOPATH=$(pwd)
- mkdir -p $GOPATH/src/github.com/pbillerot/hello
- go install github.com/pbillerot/hello
- $GOPATH/bin/hello

