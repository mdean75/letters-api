compile: scanner.o parser.o listing.o
	g++ -o compile scanner.o parser.o listing.o
	
scanner.o: scanner.c listing.h tokens.h
	g++ -c scanner.c

scanner.c: scanner.l
	flex scanner.l
	mv lex.yy.c scanner.c

parser.o: parser.c listing.h 
	g++ -c parser.c

parser.c tokens.h: parser.y
	bison -d -v parser.y
	mv parser.tab.c parser.c
	mv parser.tab.h tokens.h

listing.o: listing.cc listing.h
	g++ -c listing.cc

runall:
	./compile < test1.txt
	./compile < test2.txt
	./compile < test3.txt
	./compile < test4.txt
	./compile < test5.txt
	./compile < test6.txt
	./compile < test7.txt
	./compile < test8.txt
	./compile < test9.txt
	./compile < test10.txt
	./compile < test11.txt
	./compile < test12.txt
	./compile < test13.txt
	./compile < test14.txt
	./compile < syntax1.txt
	./compile < syntax2.txt
	./compile < syntax3.txt
	./compile < syntax4.txt
	./compile < syntax5.txt
