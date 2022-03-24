# SRE Platform Challenge

Bem vindo(a), e obrigado pelo seu interesse na Stone! Esse desafio será importante para te avaliarmos e para você entender melhor como é o nosso dia a dia. Por isso, pensamos em um desafio bem próximo a nossa realidade.

É importante ressaltar que nenhum código produzido por você nesse desafio será utilizado na Stone, tudo que for feito será utilizado apenas para te avaliar nesse desafio.

Você não precisa entregar o desafio completo. Mesmo que não implemente alguma parte, entregue-o mesmo assim e discutiremos sobre as implementações feitas!

## O desafio

O produto do nosso time é uma plataforma interna para desenvolvedores. A plataforma é capaz de provisionar recursos para aplicações, como repositórios, pipelines de CI/CD e databases.
A plataforma é contruída extendendo a API do Kubernetes usando o padrão `Operator`, assim ela pode ser consumida com uma abordagem de IaC (Infra as Code) ou integrada como uma API HTTP.

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
  type: OpenSource # or ClosedSource
  credentialsRef:
    name: github-credentials
    key: token
```

Os possíveis campos no spec do CRD são:

- `name` (obrigatório): nome do repositório no GitHub.
- `owner` (obrigatório): nome do owner do repositório no GitHub.
- `type` (obrigatório): tipo do repositório a ser criado.
- `credentialsRef` (obrigatório): referência para uma chave de um `Secret` que conterá um PAT (_Personal Access Token_) para se autenticar com a API do GitHub.
- `description` (opcional): a descrição do repositório.

## Como fazer o desafio

O código neste repositório já apresenta uma implementação inicial incompleta. Você deve implementar as tarefas descritas nos entregáveis que os avaliadores te indicarem, não é necessário completar os outros entregáveis.

Dentro deste repositório, existem duas principais pastas: `client` e `controllers`.

A pasta `client` contém todo código responsável por se comunicar com a API do GitHub.

A pasta `controllers` contém a implementação do operator `Repository`, que utiliza o pacote `client`.
> É importante ressaltar que você deve utilizar o client que está neste repositório e não um sdk externo. A utilização, melhoria e implementação dele também fazem parte do desafio.

Os testes devem ser adicionados em arquivos `*_test.go` junto aos arquivos sendo testados. Os testes presentes neste repositório utilizam o framework [Ginkgo](https://onsi.github.io/ginkgo/).

Para a implementação do operator, utilizamos o [kubebuilder](https://kubebuilder.io/), um dos frameworks mais populares para a construção de operators Kubernetes.

### Entregável 1

Suporte a credenciais vindas de um kubernetes `Secret`

1. Adaptar a controller para recuperar o PAT do `Secret` referenciado no resource e repassá-lo ao client.

### Entregável 2

Corrigir bug onde a controller tenta criar um novo repositório independentemente do status do recurso externamente.

1. Implementar o método `Get` no client;
2. Invocar a criação do repositório na controller somente quando o erro retornado for `404 Not found`;

### Entregável 3

Suporte à rotina de atualização. Quando o custom resource é alterado no Kubernetes o recurso externo correspondente deve ser atualizado de acordo.

Os campos `spec.name`, `spec.owner` e `spec.type` devem ser imutáveis.

1. Implementar o método `Update` no client;
2. Adicionar o campo `spec.description` (deve ser do tipo `*string` e opcional);
3. Implementar na controller a lógica de verificação se o `Repository` deve ser atualizado externamente ou não (verificar se a especificação do estado do recurso no Kubernetes bate com o estado atual do GitHub);
4. Adicionar o campo `status.ID` (tipo `int64` e opcional) que deve ser o valor do `ID` do repositório no GitHub;

### Entregável 4

Adicionar suporte ao `type` 'ClosedSource'.

_Os possíveis valores do campo `spec.type` são: `OpenSource` ou `ClosedSource`._

O parâmetro type representa uma lógica de negócio da empresa implementada no operator. Cada _type_ implica em um conjunto diferente de características no repo gerenciado pelo operator:

OpenSource:

- Repositório público;
- Possui issues;
- Inicializado automaticamente com o arquivo `README.md`;
- Licença `Apache License 2.0`;

ClosedSource:

- Repositório privado;
- Sem issues;
- Inicializado automaticamente com o arquivo `README.md`;
- Sem licença;
- Arquivar o repositório no processo de deleção (arquivar ao invés de deletar);

Tarefas:

1. Implementar o método `Archive` no client;
2. Adequar a lógica de deleção baseada no `type` do repositório;
3. Adequar a configuração de criação do repositório baseada no `type`;

### Entregável 5

Configuração do CI do repo de desafio

1. Criação de um pipeline de CI (_Continuous Integration_) usando Github Actions que executa os testes automaticamente. Os gatilhos para execução do CI devem ser `pull_request` e `push` para a branch `main`;

## Avaliação

Você será avaliado e acordo com os seguintes critérios.

### Funcionalidade: até 5 pontos

1. Todos os entregáveis **designados** concluídos, implementados corretamente e com testes. _5 pts_
1. Todos os entregáveis **designados** concluídos e implementados corretamente. _3 pts_
1. Todos os entregáveis **designados** concluídos, possivelmente com erros, mas executando. _2 pts_
1. Ao menos um dos entregáveis implementado, possivelmente com erros, mas executando. _1 pt_
1. Qualquer coisa diferente disso. _0 pts_

### Estilo de código e convenções: até 4 pontos

1. Código logicamente organizado e com comentários claros. Estilo no código e na documentação é claro e consistente. Tratamento adequado de erros quando necessário. _4 pts_
1. Código logicamente organizado e com comentários claros. Estilo no código e na documentação é claro e consistente. _3 pts_
1. Código logicamente organizado, mas documentação é inconsistente ou confusa. _2 pts_
1. Código desorganizado e difícil de acompanhar. Estilo arbitrário e inconsistente. _1 pt_
1. Qualquer coisa diferente disso. _0 pts_

## Enviando sua solução para avaliação

Você pode forkar esse repositório, mas não recomendamos fazer isso diretamente, já que assim qualquer um poderá ver no que você está trabalhando.
Você pode trabalhar em um repositório privado e nos dar acesso quando estiver pronto, ou nos enviar um zip (contendo também o `.git`) para o e-mail fornecido pelos avaliadores.

## Setup de desenvolvimento local

Durante o ciclo de desenvolvimento, recomendamos que utilize um cluster local. Isso evita faturas em cloud providers e torna o ciclo de desenvolvimento mais simples. Para isso, recomendamos que utilize qualquer uma das ferramentas abaixo:
- [Minikube](https://minikube.sigs.k8s.io/docs/start/)
- [Kind](https://kind.sigs.k8s.io/)
- [K3D](https://k3d.io/)

## Referências

Para ajudá-lo no processo de estudos sobre os assuntos, separamos alguns materiais de estudos:

### Go

- [A Tour of Go](https://go.dev/tour/)
- [Aprenda Go com Testes](https://larien.gitbook.io/aprenda-go-com-testes/)
- [Effective Go](https://go.dev/doc/effective_go)

### Kubernetes
- [Secrets](https://kubernetes.io/pt-br/docs/concepts/configuration/secret/)
- [Custom Resources](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/)

### Kubernetes Operator/Kubebuilder

- [Kubernetes Operator simply explained in 10 mins](https://youtu.be/ha3LjlD6g7g)
- [The Kubebuilder Book](https://kubebuilder.io/)
- [Tutorial: Deep Dive into the Operator Framework for... Melvin Hillsman, Michael Hrivnak, & Matt Dorn](https://youtu.be/8_DaCcRMp5I) - (até os 37 minutos)
- [Writing a Kubernetes Operator from Scratch Using Kubebuilder - Dinesh Majrekar](https://youtu.be/LLVoyXjYlYM)

### Testes
- [Testing framework Ginkgo](https://onsi.github.io/ginkgo/)
- [Testing Kubernetes CRDs - Christie Wilson, Google](https://youtu.be/T4EB0KB1-fc)

### Boas práticas

- [Boas práticas na Stone](https://github.com/stone-payments/stoneco-best-practices/blob/master/README_pt.md)
