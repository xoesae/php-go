ANTLR_JAR=antlr-4.13.0-complete.jar
OUTPUT_DIR=parser
PACKAGE_NAME=parser
GO_MOD=github.com/xoesae/php-go

OS := $(shell uname -s)
ifeq ($(OS),Darwin)
	OS_NAME=mac
else
	OS_NAME=linux
endif

ANTLR_CMD=java -jar $(ANTLR_JAR)

antlr:
	@if [ ! -f "$(ANTLR_JAR)" ]; then \
		curl -O https://www.antlr.org/download/$(ANTLR_JAR); \
	fi

gen-antlr: antlr
	$(ANTLR_CMD) -Dlanguage=Go -o $(OUTPUT_DIR) PhpLexer.g4
	$(ANTLR_CMD) -Dlanguage=Go -o $(OUTPUT_DIR) PhpParser.g4

test:
	go test ./...