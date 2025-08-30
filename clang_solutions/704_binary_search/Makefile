bld:
	gcc -Wall -Wextra -Wpedantic -Wconversion -Wshadow test_imp.c imp.c -o program
run:
	./program
cln:
	rm -f program

all: cln bld run cln

.PHONY: bld run cln all