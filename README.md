# ChromeDriver Updater

ChromeDriver Updater é uma ferramenta escrita em Go que verifica, baixa e atualiza automaticamente a versão mais recente do ChromeDriver na pasta do [SeleniumBasic](https://florentbr.github.io/SeleniumBasic/ "Selenium Basic by Florentbr").

## Recursos

* Verifica a versão instalada do ChromeDriver.
* Obtém a versão mais recente do ChromeDriver disponível.
* Baixa e substitui automaticamente o arquivo na pasta do SeleniumBasic.
* Interface simples via linha de comando.

## Requisitos

* Windows
* Conexão com a internet
* SeleniumBasic instalado em `<span>C:\Users\<seu_usuario>\AppData\Local\SeleniumBasic</span>`

## Instalação

1. Clone o repositório:
   ```
   git clone https://github.com/seu-usuario/chromedriver-updater.git
   ```
2. Acesse a pasta do projeto:
   ```
   cd chromedriver-updater
   ```
3. Compile o programa (necessário ter o [Go](https://go.dev/ "Go language") instalado):
   ```
   go build -o chromedriver-updater.exe
   ```

## Uso

1. Execute o programa:
   ```
   chromedriver-updater.exe
   ```
2. Escolha uma opção no menu:
   * `<span>1</span>`: Atualizar o ChromeDriver.
   * `<span>2</span>`: Sair.

O programa buscará a versão mais recente do ChromeDriver, fará o download e substituirá automaticamente o arquivo antigo.

## Exemplo de Execução

```
Atualizador ChromeDriver

https://github.com/rRopelato

Versão disponível do ChromeDriver: 123.0.0
Versão instalada do ChromeDriver: 122.0.0

1. Atualizar ChromeDriver
2. Sair
```

## Autor

Se você tiver alguma dúvida ou sugestão, sinta-se à vontade para entrar em contato:

| Plataforma   | Contato                  |
| ------------ | ------------------------ |
| Discord User | ropelato                 |
| Discord ID   | 220701036929613825       |
| E-mail       | ropelato.dev@outlook.com |

Feito com ❤️ por [@rRopelato](https://github.com/rRopelato).
