# devops_application

Aplicação desenvolvida para demonstração de conceitos de DevOps

### Organização:
    o Projeto está organizado da seguinte maneira:
        - /src/ -> Contém os "source code" das 2 aplicações necessárias
          - /src/main -> Contém a aplicação "CORE"
            - src/main/Dockerfile -> Contém a definição do container que carrega a aplicação "CORE"
          - /src/persistence -> Contém a aplicação "de persistencia"
      
      - /tests/ -> Contém os "testes" das aplicações