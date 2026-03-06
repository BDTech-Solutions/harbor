#!/usr/bin/env bash

check_docker() {
  if ! command -v docker >/dev/null 2>&1; then
    echo "❌ Docker não encontrado."
    exit 1
  fi

  if ! docker info >/dev/null 2>&1; then
    echo "❌ Docker não está em execução."
    exit 1
  fi
}

bootstrap_wordpress() {
    echo "⏳ Aguardando a extração inicial dos arquivos do WordPress..."

    local COUNTER=0
    # Aguarda até o core do WP ser extraído para a pasta /wp
    while [ ! -f wp/wp-settings.php ] && [ $COUNTER -lt 30 ]; do
        sleep 1
        ((COUNTER++))
    done

    if [ -f wp/wp-settings.php ]; then
        echo "🐳 Ajustando permissões via Docker para o usuário $(whoami)..."
        # Executa o chown via root do container para liberar os arquivos no Host (Ubuntu)
        docker compose exec -u root wordpress chown -R $(id -u):$(id -g) /var/www/html
        echo "✅ Permissões ajustadas."
    else
        echo "⚠️  Aviso: O WordPress demorou a inicializar. Verifique 'docker compose logs'."
    fi
}

init_wordpress() {
    local UP_FLAG=false
    for arg in "$@"; do
        [[ "$arg" == "--up" ]] && UP_FLAG=true
    done

    echo "⚓ Harbor - Inicializando projeto WordPress"

    # 1. Preparação de Arquivos (Copiar templates)
    [[ ! -f docker-compose.yml ]] && cp "$HARBOR_ROOT/templates/wordpress/docker-compose.yml" .
    [[ ! -f .gitignore ]] && cp "$HARBOR_ROOT/templates/wordpress/.gitignore" .gitignore
    [[ ! -f .env ]] && cp "$HARBOR_ROOT/templates/wordpress/.env" .env
    
    # 2. Preparação de Diretórios
    mkdir -p wp/wp-content/plugins wp/wp-content/themes wp/wp-content/uploads
    
    # 3. Binário de controle local
    if [[ ! -f bin/harbor ]]; then
        mkdir -p bin
        cp "$HARBOR_ROOT/templates/wordpress/harbor" bin/harbor
        chmod +x bin/harbor
    fi

    echo "✅ Estrutura de arquivos criada."

    # 4. Provisionamento do Container (Só ocorre no Init)
    check_docker
    docker compose up -d
    
    # Chama a função de permissões apenas aqui, na criação
    bootstrap_wordpress

    echo "✅ Projeto inicializado com sucesso!"
    echo "👉 Use ./bin/harbor.sh para gerenciar o projeto."
}

up_wordpress() {
  echo "🚀 Subindo containers WordPress..."

  check_docker

  if [[ ! -f docker-compose.yml ]]; then
    echo "❌ Erro: docker-compose.yml não encontrado. Rode 'harbor wordpress init' primeiro."
    exit 1
  fi

  # No dia-a-dia, apenas sobe os containers. Sem loops, sem chown.
  docker compose up -d

  # Carrega a porta do .env para exibir a URL correta
  local PORT=$(grep WP_PORT .env | cut -d '=' -f2)
  echo "✅ Ambiente WordPress pronto!"
  echo "🌐 Acesse: http://localhost:${PORT:-8080}"
}

down_wordpress() {
  echo "🛑 Parando containers WordPress..."
  check_docker
  docker compose down
  echo "✅ Containers WordPress parados."
}
