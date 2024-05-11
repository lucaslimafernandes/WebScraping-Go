# WebScraping-Go
Web Scraping project written in Go, based on WebGlobal evaluation.

## About evaluation

### 📌 Avaliação Python
Avaliação de seleção de candidatos ao cargo de programador Python *(Ago/2021)*

### 🕷 Spider / Web Crawler

*Um [`Web Crawler`](https://pt.wikipedia.org/wiki/Rastreador_web) ou [`Spider`](https://pt.wikipedia.org/wiki/Rastreador_web), é uma programa de computador, ou robô, que navega por sites da internet de forma metódica e automatizada.*

*O principal propósito de um Web Crawler é fazer o rastreamento de novas páginas e indexá-las. Em geral, ele começa com uma lista de URLs para visitar (páginas-chave ou sementes), e à medida que o *crawler* visita essas URLs, ele identifica e extrai todos os *links* contidos da página e os armazena em uma lista.*

**O seu primeiro desafio consiste na construção de um web crawler implementando em Python, que seja capaz de identificar e indexar um mínimo de <span style="color: red;">75</span> URLs dos produtos ofertados pelo site da [`DrogaRaia`](http://www.drogaraia.com.br/):**                                                                                                         

```
http://www.drogaraia.com.br/
```

Sugestões de Páginas Chave:

```
https://www.drogaraia.com.br/medicamentos.html
https://www.drogaraia.com.br/beleza.html
https://www.drogaraia.com.br/cabelo.html
https://www.drogaraia.com.br/bem-estar.html
https://www.drogaraia.com.br/mamae-e-bebe.html
```

As URLs devem ser gravadas em um arquivo texto no seguinte formato:

```
https://www.drogaraia.com.br/raia-multi-50-60-capsulas.html
https://www.drogaraia.com.br/maracugina-90mg-ml-solucao-com-100ml.html
https://www.drogaraia.com.br/catarinense-cloreto-de-magnesio-com-100-comprimidos.html
https://www.drogaraia.com.br/fertisop-com-30-saches-4g-cada.html
https://www.drogaraia.com.br/colflex-complet-60-comprimidos.html
https://www.drogaraia.com.br/regenesis-pre-30-capsulas.html
https://www.drogaraia.com.br/zero-cal-adocante-po-sucralose-com-50-sache-600mg-cada.html
https://www.drogaraia.com.br/bio-c-vitamina-1g-30-comprimidos-efervecentes.html
https://www.drogaraia.com.br/omega-3-kit-catarinense-nutricacao-1000mg-2-fracos-com-120-capsulas-cada-1-frasco-com-60-capsulas-gratis.html
```

### 🕸 Web Scrapping

*[`Web Scrapping`](https://pt.wikipedia.org/wiki/Coleta_de_dados_web), é uma forma de [`mineração de dados`](https://pt.wikipedia.org/wiki/Minera%C3%A7%C3%A3o_de_dados) que permite a extração de dados de sites da web convertendo-os em informação estruturada. O tipo mais básico de coleta é o download manual das páginas, copiando e colando o conteúdo, e isso pode ser feito por qualquer pessoa. Contudo, essa técnica geralmente é feita através de um software que simula uma navegação humana e extraindo as informações de interesse.*

**Seu segundo desafio é construir um Web Scrapper em linguagem Python capaz de extrair a `Descrição`, o `Preço` e o código `SKU` de cada um dos produtos indexados no desafio anterior.**

Exemplo de Saída Esperada:

```
Nome="Suplemento Alimentar Cloreto de Magnésio P.A. Catarinense Nutrição com 100 comprimidos" Preco=R$39.90 SKU=26602
Nome="Suplemento Vitamínico em Pó Myralis FertiSop com 30 sachês" Preco=R$127.99 SKU=33072
Nome="Colflex Complet 40mg Colágeno Tipo II Não Hidrolisado com 60 Comprimidos" Preco=R$159.99 SKU=72575
Nome="Adoçante em Pó Zero-Cal Sucralose com 50 sachês de 600mg cada" Preco=R$8.59 SKU=113154
Nome="Suplemento Alimentar ReGenesis Pré com 30 cápsulas" Preco=R$99.99 SKU=71533
Nome="Vitamina C Bio-C 1g Sabor Laranja com 30 comprimidos efervescentes" Preco=R$28.99 SKU=33285
Nome="Kit Complexo Vitamínico Ômega 3 1000mg " Preco=R$98.49 SKU=73113
Nome="Choco Soy Pops Banana Passa Coberta com Chocolate com 40g" Preco=R$8.05 SKU=33851
```

🍀 *Boa Sorte!* 🍀
