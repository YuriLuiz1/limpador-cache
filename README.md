# Chrome Cache Cleaner 🧹

Um utilitário de linha de comando simples e eficiente para limpar automaticamente todos os caches do Google Chrome sem a necessidade de abrir o navegador ou fazer configurações manuais.

## 📋 Descrição

O Chrome Cache Cleaner é uma ferramenta desenvolvida em Go que permite limpar rapidamente todos os arquivos de cache do Google Chrome diretamente através do sistema de arquivos. Ideal para usuários que precisam liberar espaço em disco ou resolver problemas relacionados ao cache corrompido.

## ✨ Características

- 🚀 **Execução rápida**: Limpeza completa em segundos
- 🎯 **Precisão**: Remove apenas arquivos de cache, mantendo configurações e dados importantes
- 💻 **Multi-plataforma**: Funciona no Windows, Linux e macOS
- 🔧 **Zero configuração**: Não requer instalação ou configurações adicionais
- 📦 **Portátil**: Executável único, funciona em qualquer lugar
- 🛡️ **Seguro**: Verifica se o Chrome está fechado antes da limpeza

## 🚀 Como Usar

### Windows
1. Dê um clone no projeto
2. Vá até a pasta que utilizou o clone
3. Execute o binário

## 📁 O que é limpo

O programa remove os seguintes tipos de cache:

- **Cache de navegação** - Arquivos temporários de páginas web
- **Cache de aplicações** - Dados em cache de extensões e aplicativos
- **Cache de GPU** - Cache de processamento gráfico
- **Cache de código** - Bytecode JavaScript compilado
- **Arquivos temporários** - Dados temporários diversos

## ⚠️ Requisitos

- Google Chrome instalado no sistema
- Chrome deve estar **completamente fechado** durante a execução
- Permissões de leitura/escrita na pasta do perfil do Chrome

## 🔍 Detecção Automática

O programa detecta automaticamente:
- Sistema operacional (Windows/Linux/macOS)
- Localização da pasta do perfil do Chrome
- Perfis múltiplos (se existirem)
- Status do Chrome (aberto/fechado)

## 💡 Exemplo de Uso

```bash
$ ./chrome-cleaner

Chrome Cache Cleaner v1.0
==========================

✓ Detectando sistema operacional: Windows
✓ Localizando pasta do Chrome...
✓ Encontrados 2 perfis de usuário
⚠️ Verificando se o Chrome está fechado...
✓ Chrome fechado, prosseguindo com a limpeza

Limpando perfil: Default
✓ Cache de navegação: 245 MB removidos
✓ Cache de aplicações: 67 MB removidos
✓ Cache GPU: 12 MB removidos

Limpando perfil: Profile 1
✓ Cache de navegação: 156 MB removidos
✓ Cache de aplicações: 34 MB removidos
✓ Cache GPU: 8 MB removidos

==========================
✅ Limpeza concluída!
📊 Total liberado: 522 MB
⏱️ Tempo: 3.2 segundos
```

## 🛠️ Compilação do Código Fonte

Se você deseja compilar o programa a partir do código fonte:

```bash
# Clone o repositório
git clone https://github.com/YuriLuiz1/limpador-cache.git

# Compile para o seu sistema
go build main.go

# Ou compile para Windows (a partir de Linux/macOS)
GOOS=windows GOARCH=amd64 go build -o main.go

# Ou compile para Linux (a partir de Windows/macOS)
GOOS=linux GOARCH=amd64 go build -o main.go
```

## 🔧 Tecnologias Utilizadas

- **Linguagem**: Go (Golang) 1.21+
- **Bibliotecas padrão**: 
  - `os` - Manipulação de arquivos e diretórios
  - `filepath` - Tratamento de caminhos
  - `runtime` - Detecção de sistema operacional
  - `path/filepath` - Operações multiplataforma

## 📊 Estatísticas do Projeto

![GitHub release (latest by date)](https://img.shields.io/github/v/release/seu-usuario/chrome-cache-cleaner)
![GitHub](https://img.shields.io/github/license/seu-usuario/chrome-cache-cleaner)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/seu-usuario/chrome-cache-cleaner)

## 🤝 Contribuindo

Contribuições são bem-vindas! Para contribuir:

1. Faça um fork do projeto
2. Crie uma branch para sua feature (`git checkout -b feature/MinhaFeature`)
3. Commit suas mudanças (`git commit -m 'Adiciona MinhaFeature'`)
4. Push para a branch (`git push origin feature/MinhaFeature`)
5. Abra um Pull Request

## 📝 Changelog

### v1.0.0
- Lançamento inicial
- Suporte para Windows, Linux e macOS
- Detecção automática de perfis
- Interface de linha de comando intuitiva

## ⚖️ Licença

Este projeto está licenciado sob a Licença MIT - veja o arquivo [LICENSE](LICENSE) para detalhes.

## ❓ FAQ

**P: O programa funciona com outros navegadores baseados em Chromium?**
R: Atualmente suporta apenas Google Chrome. Suporte para Edge, Brave e outros navegadores Chromium está planejado para versões futuras.

**P: É seguro usar este programa?**
R: Sim, o programa apenas remove arquivos de cache temporários e não afeta configurações, senhas, histórico ou dados pessoais.

**P: Posso usar com o Chrome aberto?**
R: Não, o Chrome deve estar completamente fechado para evitar corrupção de dados.

**P: O programa remove extensões ou configurações?**
R: Não, apenas arquivos de cache são removidos. Suas extensões, configurações e dados permanecem intactos.

---

## 🙏 Agradecimentos

Obrigado por usar o Chrome Cache Cleaner! Esperamos que esta ferramenta torne sua experiência de navegação mais rápida e eficiente.

Se este projeto foi útil para você, considere:
- ⭐ Dar uma estrela no repositório
- 🐛 Reportar bugs ou sugerir melhorias
- 📢 Compartilhar com outros usuários
- 💝 Contribuir com código ou documentação

Desenvolvido com ❤️ em **Go** para a comunidade de desenvolvedores e usuários que valorizam ferramentas simples e eficazes.

**Mantenha seu Chrome sempre limpo e rápido!** 🚀
