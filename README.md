# es-analyzer

A command line tool for elasticsearch analyzer

## Installation

```
$ go get github.com/shogo-ma/es-analyzer
```

## Usage

```
$ es-analyzer --host [host url] --analyzer [analyzer] --query [query]
```

```
$ es-analyzer --host http://localhost:9200 --analyzer standard --query "This is a pen."
token	start_offset	end_offset	type	position
this	0	4	<ALPHANUM>	0
is	5	7	<ALPHANUM>	1
a	8	9	<ALPHANUM>	2
pen	10	13	<ALPHANUM>	3
```
