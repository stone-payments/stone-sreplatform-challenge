# SRE Platform Challenge

Bem vindo(a)! Esse desafio será importante te avaliarmos e você entender melhor como é a realidade do time no dia a dia, por isso, pensamos em um desafio bem próximo a nossa realidade.

É importante ressaltar que nenhum código produzido por você nesse desafio será utilizado na Stone, tudo que for feito será utilizado apenas para te avaliar nesse desafio.

Você não precisa entregar o desafio completo, mesmo que não implemente alguma parte, entregue-o mesmo assim e discutiremos sobre as implementações feitas!

## O desafio

Você deverá desenvolver um operator que é capaz de gerenciar o ciclo de vida de um Repositório do GitHub.

Um exemplo do manifesto Kubernetes que representa o CRD (_Custom Resource Definition_) é:

```yaml
apiVersion: repositories.platform.buy4.io/v1alpha1
kind: Repository
metadata:
  name: example
spec:
  name: golang-best-practices
  owner: stone-payments
  type: OpenSource
  credentialsRef:
    name: github-credentials
    namespace: default
    key: token
```

Os possíveis campos no spec do CRD são:
- `name` (obrigatório): nome do repositório no GitHub.
- `owner` (obrigatório): nome do owner do repositório no GitHub. Esse owner pode ser tanto um usuário como uma organização.
- `type` (obrigatório): tipo do repositório a ser criado. As definições desse campo se encontram na [subseção abaixo](#o-campo-type).
- `credentialsRef` (obrigatório): referência para uma chave de uma `Secret` que conterá o PAT (_Personal Access Token_) para se autenticar com a API do GitHub.
- `description` (opcional): a descrição do repositório.

### O campo `type`

Os possíveis valores do campo `spec.type` são: `OpenSource`, `ClosedSource` e `Template`.

A configuração do repositório que o operator criará depende diretamente de qual `type` ele é. Abaixo estão as definições de cada um dos tipos.

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

#### Template
- Repositório privado;
- Sem issues;
- Sem inicialização automática;
- Sem licença;

## Como fazer o desafio

Dentro desse repositório, existem duas principais pastas: `sdk` e `github`.

A pasta `sdk` contém todo código responsável por se comunicar com a API do GitHub.

A pasta `github` contém a implementação do operator `Repository`, que utiliza o SDK (_Software Development Kit_) da pasta `sdk`.
> É importante ressaltar que você deve utilizar o SDK que está neste repositório. A utilização, melhoria e implementação do SDK também fazem parte do desafio.

### Entregável 1

A solução já apresenta uma implementação inicial de algumas partes. Você pode modificar a implementação existente como desejar, não há restrições. 

#### SDK

Algumas tarefas que você deverá fazer no SDK:

1.  Remover o PAT que está hard-coded e recebê-lo através de parâmetros na inicialização do `Client`;
2.  Implementar o método `Update`;


#### Operator

Para a implementação do operator, utilizamos o [kubebuilder](https://kubebuilder.io/). As tarefas que deverão ser feitas no operator são:

1. Extrair as credenciais a serem utilizadas de uma `Secret` que foi referenciada no `spec` do CRD do `Repository`;

2. Adicionar o campo `spec.description` (deve ser do tipo `*string` e opcional);

3.  Adicionar o campo `status.ID` (tipo `string` e opcional) que deve ser populado durante a reconciliação do recurso;

4. Implementar a lógica de verificação se o `Repository` deve ser atualizado externamente ou não (verificar se a especificação do estado do recurso no Kubernetes bate com o estado atual do GitHub);

### Entregável 2

#### SDK
1.  Implementar o método `Archive`;

#### Operator
1. Adequar a lógica de deleção baseada no `type` do repositório;

2. Corrigir o tratamento de erros no `Get` da controller para ele retornar apenas os erros diferentes de `404 Not found`;

### Entregável 3

#### SDK
1. Adaptar o SDK para ser testável por seus clientes (nesse caso, o operator);

#### Operator
1. Implementação de testes de unidade para garantir que o `Repository` gerado está sendo configurado corretamente;

2. Criação de um pipeline de CI (_Continuous Integration_) que executa os testes automaticamente com triggers para criação de PRs (Pull Requests) ou commits na branch `main`;

## Referências

Para ajudá-lo no processo de estudos sobre os assuntos, visite as referências abaixo:
