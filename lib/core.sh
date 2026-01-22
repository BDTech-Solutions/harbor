#!/usr/bin/env bash

source "$HARBOR_ROOT/lib/laravel.sh"
source "$HARBOR_ROOT/lib/wordpress.sh"

main() {
  case "${1:-}" in
    laravel)
      shift
      handle_laravel "$@"
      ;;
    wordpress)
      shift
      handle_wordpress "$@"
      ;;
    --version)
      cat "$HARBOR_ROOT/VERSION"
      ;;
    --help|"")
      show_help
      ;;
    *)
      echo "Comando desconhecido: ${1:-}"
      exit 1
      ;;
  esac
}

show_help() {
  cat <<EOF
Harbor - Development Environment Bootstrapper

Uso:
  harbor laravel init # Inicializa um novo projeto Laravel
  harbor laravel bootstrap # Gera o arquivo harbor.sh para o projeto Laravel atual
  harbor wordpress init # Inicializa um novo projeto WordPress
  harbor wordpress up   # Sobe os containers do WordPress
  harbor wordpress down # Para os containers do WordPress

Opções:
  --version     Mostra a versão
  --help        Mostra esta ajuda
EOF
}

handle_laravel() {
  case "${1:-}" in
    init)
      init_laravel
      ;;
    bootstrap)
      bootstrap_laravel
      ;;
    *)
      echo "Comando Laravel inválido. Use: init | bootstrap"
      exit 1
      ;;
  esac
}

handle_wordpress() {
  case "${1:-}" in
    init)
      init_wordpress "$@"
      ;;
    up)
      up_wordpress
      ;;
    down)
      down_wordpress
      ;;
    *)
      echo "Comando WordPress inválido. Use: init | up | down"
      exit 1
      ;;
  esac
}

main "$@"
