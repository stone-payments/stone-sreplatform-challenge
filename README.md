# SRE Platform Challenge

Bem vindo(a), e obrigado pelo seu interesse na Stone! Esse desafio será importante te avaliarmos e para você entender melhor como é a realidade do time no dia a dia, por isso, pensamos em um desafio bem próximo a nossa realidade.

É importante ressaltar que nenhum código produzido por você nesse desafio será utilizado na Stone, tudo que for feito será utilizado apenas para te avaliar nesse desafio.

Você não precisa entregar o desafio completo, mesmo que não implemente alguma parte, entregue-o mesmo assim e discutiremos sobre as implementações feitas!

## O desafio

O produto do nosso time é uma plataforma interna para desenvolvedores, a plataforma é capaz de provisionar recursos para aplicações, como repositórios, pipelines de CD, e databases.  
A plataforma é contruída extendendo a API do kubernetes usando o padrão `Operator`, assim ela pode ser consumida com uma abordagem de IaC (Infra as Code), ou integrada como uma API HTTP.

Você deverá implementar algumas funcionalidades em um operator que deve ser capaz de gerenciar o ciclo de vida de um Repositório do GitHub.

Um exemplo do manifesto Kubernetes que representa o CRD (_Custom Resource Definition_) é:

```yaml
apiVersion: repositories.platform.buy4.io/v1alpha1
kind: Repository
metadata:
  name: example
spec:
  name: golang-best-practices
  owner: stone-payments
  type: OpenSource # ClosedSource
  credentialsRef:
    name: github-credentials
    namespace: default
    key: token
```

Os possíveis campos no spec do CRD são:
- `name` (obrigatório): nome do repositório no GitHub.
- `owner` (obrigatório): nome do owner do repositório no GitHub. Esse owner pode ser tanto um usuário como uma organização.
- `type` (obrigatório): tipo do repositório a ser criado.
- `credentialsRef` (obrigatório): referência para uma chave de um `Secret` que conterá o PAT (_Personal Access Token_) para se autenticar com a API do GitHub.
- `description` (opcional): a descrição do repositório.

## Como fazer o desafio

A solução já apresenta uma implementação inicial incompleta. Você deve implementar as tarefas descritas no entregáveis que os avaliadores te indicarem, não é necessário completar os outros entregáveis.

Dentro desse repositório, existem duas principais pastas: `client` e `controllers`.

A pasta `client` contém todo código responsável por se comunicar com a API do GitHub.

A pasta `controller` contém a implementação do operator `Repository`, que utiliza o pacote `client`.
> É importante ressaltar que você deve utilizar o client que está neste repositório e não um sdk externo. A utilização, melhoria e implementação dele também fazem parte do desafio.

Os testes devem ser adicionados em arquivos `_test.go` junto aos arquivos sendo testados.  
Para a implementação do operator, utilizamos o [kubebuilder](https://kubebuilder.io/).

### Entregável 1

Suporte a credenciais vindas de um kubernetes `Secret`  

1. Adaptar a controller para recuperar o PAT no secret referenciado no resource e repassalo ao client.

### Entregável 2

Corrigir bug onde a controller tenta criar um novo repo independentemente do status do recurso externamente.

1. Implementar o método Get no client;
2. Invocar a criação do repo na controller somente quando o erro retornado for `404 Not found`;

### Entregável 3

Suporte a rotina de atualização. Quando o custom resource é alterado no kubernetes o recurso externo correspondente deve ser atualizado de acordo.

1. Implementar o método Update no client;
2. Adicionar o campo `spec.description` (deve ser do tipo `*string` e opcional);
3. Implementar na controller a lógica de verificação se o `Repository` deve ser atualizado externamente ou não (verificar se a especificação do estado do recurso no Kubernetes bate com o estado atual do GitHub);
4. Adicionar o campo status.ID (tipo string e opcional) que deve ser populado durante a reconciliação do recurso;

### Entregável 4

Adicionar suporte ao `type` 'ClosedSource'. 

1. Implementar o método `Archive` no client;
2. Adequar a lógica de deleção baseada no `type` do repositório;
3. Adequar a configuração de criação do repositório baseada no `type`;

Os possíveis valores do campo `spec.type` são: `OpenSource` ou `ClosedSource`.

#### OpenSource

- Repositório público;
- Possui issues;
- Inicializado automaticamente com o arquivo `README.md`;
- Licença `Apache License 2.0`;

#### ClosedSource

- Repositório privado;
- Sem issues;
- Inicializado automaticamente com o arquivo `README.md`;
- Sem licença;
- Arquivar o repositório no processo de deleção (arquivar ao invés de deletar);

### Entregável 5

Configuração do CI do repo de desafio

1. Criação de um pipeline de CI (_Continuous Integration_) usando github actions que executa os testes automaticamente, o gatilho para execução do CI devem ser criação de PRs (Pull Requests) ou commits na branch `main`;

## Avaliação

Você será avaliado e acordo com os seguintes critérios.

### Funcionalidade: até 5 pontos:

1. Todos os entregáveis **designados** concluídos, implementados corretamente, e com testes. *5 pts*
1. Todos os entregáveis **designados** concluídos, implementados corretamente. *3 pts*
1. Todos os entregáveis **designados** concluídos, possivelmente com erros, mas executando. *2 pts*
1. Ao menos um dos entregáveis implementado, possivelmente com erros, mas executando. *1 pt*
1. Qualquer coisa diferente disso. *0 pts*

### Estilo de código e convenções: até 4 pontos:

1. Codigo logicamente organizado e com comentários claros. Estilo no código e na documentação é claro e consistente. Tratamento adequado de erros quando necessário. *4 pts*
1. Codigo logicamente organizado e com comentários claros. Estilo no código e na documentação é claro e consistente. *3 pts*
1. Codigo logicamente organizado, mas documentação é inconsistente ou confusa. *2 pts*
1. Codigo desorganizado e difícil de acompanhar. Estilo arbitrário e inconsistente. *1 pt*
1. Qualquer coisa diferente disso. *0 pts*

## Enviando sua solução para avaliação

Você pode forkar esse repo, mas não recomendamos fazer isso diretamente, já que assim qualquer um poderá ver no que você está trabalhando.
Você pode trabalhar em um repo privado e nos dar acesso quando estiver pronto, ou nos enviar um zip (contendo também o .git) para o email fornecido pelos avaliadores.

## Referências

Para ajudá-lo no processo de estudos sobre os assuntos, visite as referências abaixo:
