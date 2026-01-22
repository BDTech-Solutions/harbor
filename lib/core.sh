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
  harbor laravel init
  harbor laravel bootstrap
  harbor wordpress init

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
      init_wordpress
      ;;
    *)
      echo "Comando WordPress inválido. Use: init"
      exit 1
      ;;
  esac
}

main "$@"
