#!/usr/bin/env bash

init_wordpress() {
  echo "⚓ Harbor - Inicializando projeto WordPress"

  # 1. Verifica se estamos em um diretório Harbor (wp-content ou docker-compose.yml)
  if [[ -f docker-compose.yml ]]; then
    echo "✅ docker-compose.yml já existe"
  else
    echo "📦 Criando docker-compose.yml básico para WordPress..."
    cp "$HARBOR_ROOT/templates/wordpress/docker-compose.yml" .
  fi

  # 2. Cria diretórios wp-content/plugins e wp-content/themes se não existirem
  mkdir -p wp/wp-content/plugins wp/wp-content/themes

  # 3. Cria arquivo .env se não existir
  if [[ ! -f .env ]]; then
    echo "🌿 Criando arquivo .env..."
    cp "$HARBOR_ROOT/templates/wordpress/.env" .env
  fi

  # 4. Cria harbor.sh na raiz do projeto se não existir
  if [[ ! -f bin/harbor.sh ]]; then
    echo "📄 Criando bin/harbor.sh..."
    mkdir -p bin
    cp "$HARBOR_ROOT/templates/wordpress/harbor.sh" bin/harbor.sh
    chmod +x bin/harbor.sh
  fi

  echo "✅ Estrutura inicial do WordPress criada com sucesso!"
  echo "Use ./bin/harbor.sh para subir containers e gerenciar o projeto."
}
