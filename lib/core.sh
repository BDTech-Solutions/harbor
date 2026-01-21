#!/usr/bin/env bash

main() {
  case "${1:-}" in
    init)
      shift
      handle_init "$@"
      ;;
    --version)
      cat "$HARBOR_ROOT/VERSION"
      ;;
    --help|"")
      show_help
      ;;
    *)
      echo "Comando desconhecido: $1"
      exit 1
      ;;
  esac
}

show_help() {
  cat <<EOF
Harbor - Development Environment Bootstrapper

Uso:
  harbor init laravel
  harbor init wordpress

Opções:
  --version     Mostra a versão
  --help        Mostra esta ajuda
EOF
}
