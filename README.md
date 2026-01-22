# Harbor

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

**Harbor** é um bootstrapper de ambientes de desenvolvimento para Laravel e WordPress. Ele facilita a inicialização de projetos, gerencia dependências via Docker e cria scripts de bootstrap como `harbor.sh`, garantindo idempotência e boas práticas sem depender de instalações locais.

---

## Funcionalidades

- Bootstrap de projetos Laravel e WordPress
- Integração com Docker para Composer e ambientes isolados
- Criação de `harbor.sh` para Laravel, permitindo rodar projetos clonados sem depender de PHP/Composer local
- Suporte a múltiplas stacks planejado para versões futuras (Python, Go, etc.)

---

## Instalação

Clone o repositório:

```bash
git clone https://github.com/BDTech-Solutions/harbor.git
cd harbor
chmod +x ./bin/harbor
