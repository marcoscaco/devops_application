# devops_application

## Sobre
    Aplicação desenvolvida para demonstração de conceitos de DevOps

    O desafio exige:
        - A crição de um POD micro serviço
        - A criação de um POD de persistencia

### Para alcançar tal objetivo foi escolhido as seguinte stack:
- **[Helm](https://helm.sh/)** como gestor de pacotes para kubernetes
- **[Golang](https://go.dev/)** como linguagem para o micro serviço
- **[Redis](https://redis.io/)** como base de persistencia
- **[Terraform](https://www.terraform.io/)** para gestão da infra as code
- **[AWS](https://aws.amazon.com/)** como cloud kubernetes provider provider
- **[Minikube](https://minikube.sigs.k8s.io/docs/)** como local kubernetes provider
- **[Homebrew](https://brew.sh/)** para a gestão de dependencias do ambiente de desenvolvimento
- **[Postman/Newman](https://learning.postman.com/docs/running-collections/using-newman-cli/command-line-integration-with-newman/)** para os testes da api REST


## Organização

#### o Projeto está organizado da seguinte maneira:
- /main
    - Contém o necessário (helm) para a aplicação CORE
- /redis
    - Contém o necessário (helm) para a database (REDIS)
        
- /src/main
    - Contém o necessário (GoLang source code + Dockerfile) para a aplicação CORE
- /src/persistence
    - Contém o necessário (Dockerfile) para a aplicação de persistencia

- /tests/ -> Contém os "testes" das aplicações
  - construidos para rodar com **[Postman/Newman](https://learning.postman.com/docs/running-collections/using-newman-cli/command-line-integration-with-newman/)**

#### o Projeto contém também os seguintes arquivos utilitários
- Brewfile
  - carregando o conjunto de pacotes necessários para o setup do ambiente de desenvolvimento
  - sendo elas:
    - docker
    - helm
    - kubernetes-cli
    - minikube

- run.sh
  - Script auxiliar para o ambiente de desenvolvimento e para o desenvolvedor

## Como executar
    Disclaimer[1]: Buscando manter a sanidade mental do autor as instruções a seguir são aplicaveis a usuários do Mac Os (e por sua vez replicaveis com alguns ajustes no linux)
    Disclaimer[2]: Minimamente faz-se necessário o Homebrew instalado e disponivel para o usuário atual

 - Devemos utilizar a opção **0** do script **run.sh** para instalar os componentes necessários no ambiente de desenvolvimento
 - Devemos utilizar a opção **1** do script **run.sh** para executar a aplicação utilizando o mini cluster (minikube)
 - Devemos utilizar a opção **7** do script **run.sh** para executar os testes da aplicação REST