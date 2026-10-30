# Como utilizar?
#### Nesta breve documentação, segue o guia de como utilizar o sistema de geração de recibos e borderôs.

### 1. Baixe ou Clone o projeto para seu computador
```sh
	git clone https://github.com/ogabrielrodrigues/rentalgen
```

ou, baixe a versão pré-compilada no GitHub. O binário já inclui os templates e pode ser executado de qualquer pasta.

### 2. Gere o binário (caso tenha clonado)
Windows:
```sh
	go build -o rentalgen.exe ./cmd
```

Linux:
```sh
	go build -o rentalgen ./cmd
```

Para gerar o binário de Windows a partir do Linux:
```sh
	GOOS=windows go build -o rentalgen.exe ./cmd
```

### 3. Execute o binário
Ao executar o binário, você precisa indicar um arquivo `.json`. Sem a opção `-file`, o sistema procura por `alugueres.json` na pasta atual.

Use o arquivo `exemplo.json` como ponto de partida. O arquivo `exemplo.jsonc` explica cada campo, mas serve apenas como referência: o sistema não aceita arquivos com comentários.

Windows:
```sh
	.\rentalgen -file .\alugueres.json # exemplo de arquivo
```

Linux:
```sh
	./rentalgen -file ./alugueres.json # exemplo de arquivo
```

Antes de iniciar, o sistema valida o arquivo. Caso existam problemas, todos são listados e nada é gerado:
```
✗ arquivo inválido:
  - aluguel 2: campo obrigatório city não preenchido
  - aluguel 3, owner: campo obrigatório cpf não preenchido
```

- Os meses devem ser escritos por extenso, em português (ex: `Janeiro`).
- Os campos desconhecidos são rejeitados, para evitar que erros de digitação passem despercebidos.
- Os alugueres com `"skip": true` não são validados.
- Os alugueres com `"onlyVoucher": true` não precisam dos campos usados somente no borderô (`guarantee`, `amountWithoutTax`, `signedIn`, `extendedIn` e `expiresIn`).
- O texto sai no feminino somente quando `nationality` é exatamente `brasileira`.
- Se `dueDay` não existir no mês de vencimento, a data do recibo usa o último dia do mês (ex: dia 31 em Fevereiro de 2026 vira dia 28).

### 4. Acesse o sistema
Se o arquivo informado for válido e for processado pelo sistema, o mesmo exibirá um aviso, indicando que estará disponível em: _[http://localhost:8080](http://localhost:8080)_
```
✓ arquivo de alugueres selecionado: alugueres.json
✓ para realizar a impressão dos recibos, acesse http://localhost:8080

Para encerrar, pressione Ctrl + C
```

O sistema é acessível somente pelo próprio computador. Se a porta 8080 estiver em uso, ele exibirá o erro e será encerrado.

### 5. Imprima os recibos
Na página, imprima com `Ctrl + P`. O layout é preparado para papel A4, com cada documento iniciando em uma nova página.

# Alterando os templates
Os templates ficam dentro do binário, portanto é necessário gerar o binário novamente após alterá-los.

O CSS em `templates/styles.css` é gerado a partir de `templates/tailwind.css`. Se adicionar novas classes do Tailwind aos templates, gere o arquivo novamente (necessário Node.js):
```sh
	npm install --no-save tailwindcss@4.3.3 @tailwindcss/cli@4.3.3
	npx tailwindcss -i templates/tailwind.css -o templates/styles.css --minify
```
