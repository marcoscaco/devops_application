# devops_application

## Sobre
    Aplicação desenvolvida para demonstração de conceitos de DevOps

    O desafio exige:
        - A crição de um POD micro serviço
        - A criação de um POD de persistencia

    Para alcançar tal objetivo foi escolhido as seguinte stack:
        - [Helm](https://helm.sh/) como gestor de pacotes para kubernetes
        - [Golang](https://go.dev/) como linguagem para o micro serviço
        - [Redis](https://redis.io/) como base de persistencia
        - [Terraform](https://www.terraform.io/) para gestão da infra as code
        - [AWS](https://aws.amazon.com/) como cloud kubernetes provider provider
        - [Minikube](https://minikube.sigs.k8s.io/docs/) como local kubernetes provider


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