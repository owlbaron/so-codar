# Só codar

Só codar ou `sc` é o meu parser.

## Como rodar

É só rodar o seguinte comando na raiz do projeto:

```
./bin/sc run ./tests/parser/$FILE_NAME.sc
```

Onde `$FILE_NAME` é um dos arquivos de teste que se encontra dentro do path `./tests/parser`, então fique a vontade para rodar eles ou criar um novo seguindo o mesmo modelo.

## EBNF

```
program              ::= "program" block
block                ::= "{" statement-sequence "}"
statement-sequence   ::= statement*
statement            ::= assignment-statement | structured-statement
variable-declaration ::= "var" identifier ":=" expression ";"
assignment-statement ::= identifier ":=" expression ";"
structured-statement ::= while-statement | if-statement
while-statement      ::= "while" expression "do" block
if-statement         ::= "if" expression block ("else" block)?
expression           ::= simple-expression (relational-operator simple-expression)*
simple-expression    ::= (sign)? term (lower-precedence-operator term)*
term                 ::= factor (higher-precendence-operator factor)*
factor               ::= number |
                         string |
                         identifier |
                         "(" expression ")"
```

reference: http://www.cs.kent.edu/~durand/CS43101Fall2004/resources/Pascal-EBNF.html#simple-expression