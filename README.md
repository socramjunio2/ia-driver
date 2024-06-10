# Projeto de Machine Learning para identificação automática de motoristas.

Hoje em dia muitas empresas e desenvolvedores tem ajudado a comunidade de códigos abertos.

Decidi enviar o meu primeiro projeto open source desenvolvido em Go(Language).

Este projeto identifica motoristas automaticamente baseado em algorítimos de padrões de aceleração,
frenagens e curvas acentuadas.

1. Recebimento de dados de direção em tempo real via Kafka.
2. Cálculo de aceleração, frenagem e curvas.
3. Identificação de motoristas com base nos padrões detectados utilizando Machine Learning (KNN).
4. Geração e envio de um ID único via Kafka se um motorista correspondente for encontrado.
5. Armazenamento e consulta de informações de motoristas em um banco de dados SQLite.

Baseando-se neste modelo, criei uma interface de aprendizado de máquina com GoLearn, um pacote excepcional que contém um conjunto de soluções para construir do zero IAs generativas.

O mais interessante é que, através da linguagem Go, ficou extremamente simples a construção de um micro-serviço de identificação de motoristas. Não é brincadeira, se eu tivesse construído este microsserviço com uma linguagem diferente, como Python por exemplo, teria adicionado muito mais linhas de código.

Para construir essa inteligência artificial, fiz uso de ferramentas cruciais como:

## Ferramentas

- Golang 1.6
- Kafka (confluent-kafka-go)
- Database (go-sqlite3)

## Fluxo de desenvolvimento

Para adicionar integração com Kafka no projeto em Go, utilizei a biblioteca confluent-kafka-go, que é uma biblioteca popular para trabalhar com Apache Kafka. A seguir, vou mostrar como configurar a integração com Kafka e também explicar como calcular aceleração, frenagem e curvas a partir dos dados recebidos.

### Integração com Kafka

#### Passo 1: Instalar a Biblioteca confluent-kafka-go

Primeiro, você precisa instalar a biblioteca `confluent-kafka-go`:
```sh
  go get github.com/confluentinc/confluent-kafka-go/kafka
```

#### Passo 2: Configurar o Consumidor Kafka

Para iniciar a conexão com o Kafka, adicionei uma camada de configuração para conectar-se ao consumidor

Após a configuração, fiz uma subscrição no tópico `vehicle-data` e através da leitura das mensagens para iniciar
o processamento da IA.


### Cálculo de Aceleração, Frenagem e Curva

Para calcular aceleração, frenagem e curva, você deve ter dados de posição (ou velocidade) e tempo. Vamos supor que você tenha dados de posição em coordenadas (x, y) e timestamps (t).

#### Cálculo da Aceleração

A aceleração é a taxa de variação da velocidade. Se você tem a posição e o tempo, pode calcular a velocidade e, em seguida, a aceleração.

#### Cálculo da Frenagem

A Frenagem é a taxa de variação entre a velocidade atual e tempo de desaceleração baseado na frenagem do veículo.

#### Calculo de Curvas

O calculo de uma curva é baseado em x quantidade de mudançã de direção baseado no trajeto.

#### Instalação de Dependências:

Instale a biblioteca golearn:

```sh
go get github.com/sjwhitworth/golearn
```
Instale a biblioteca confluent-kafka-go:

```sh
go get github.com/confluentinc/confluent-kafka-go/kafka
```
Configuração do Kafka:

Certifique-se de que o Apache Kafka esteja instalado e executando em sua máquina ou servidor.
Ajuste o bootstrap.servers e outros parâmetros de configuração conforme necessário.
Recebimento de Dados:

Configure seu sistema para enviar dados de direção (posição e tempo) para o tópico Kafka especificado.
Treinamento e Avaliação do Modelo:

A função `TrainAndEvaluateModel` é um placeholder e deve ser ajustada para treinar e avaliar seu modelo com os dados reais.

## Contribuição

Quem quiser contribuir para o projeto, faça um fork e envie pull requests.

Todos os colaboradores serão listados aqui.

Obrigado por contribuir.