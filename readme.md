# Como utilizar?
#### Nesta breve documentação, segue o guia de como utilizar o sistema de geração de recibos e borderôs.

### 1. Baixe ou Clone o projeto para seu computador
```sh
	git clone https://github.com/ogabrielrodrigues/rentalgen
```

ou, baixe a versão pré compilada no Github, e insirá na pasta do projeto para continuar.

### 2. Gere o binário (caso tenha clonado)
Windows:
```sh
	GOOS=windows go build -o rentalgen.exe cmd/main.go
```

Linux:
```sh
	go build -o rentalgen cmd/main.go
```

### 3. Execute o binário
Ao executar o binário, você precisa indicar um arquivo `.json`. Use como exemplo o arquivo `exemplo.jsonc` que está no projeto.
Windows:
```sh
	.\rentalgen -file .\alugueis.json # exemplo de arquivo
```

Linux:
```sh
	./rentalgen -file ./alugueis.json # exemplo de arquivo
```

### 4. Acesse o sistema
Se o arquivo informado for válido e for processado pelo sistema, o mesmo exibirá um aviso, indicando que estará disponível em: _[http://localhost:8080](http://localhost:8080)_
```
✓ arquivo de alugueres selecionado: alugueres.json
✓ para realizar a impressão dos recibos, acesse http://localhost:8080

Para encerrar, pressione Ctrl + C
```
