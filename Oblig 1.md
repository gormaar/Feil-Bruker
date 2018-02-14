<h1>      Obligatorisk oppgave 1 </h1>
<h2>      Gorm-Erik, Kevin, Marius </h2>

<h3>      1. Fyll ut manglende tall i tabell </h3>

Binære tall|Hexadesimal|Desimaltall
-|-|-
1101|0xD|13
1101 1110 1010|0xDEA |3562
1010 1111 0011 0100|0xAF34| 44852
1111 1111 1111 1111 | 0xFFFF | 65535
0001 0001 0111 1000 1010 | 0x1178A | 71562

<h4>     Oppgave A </h4>
<b>Beskriv kort metode for å gå fra binære tall til hexadesimale tall og motsatt. Beskriv kort metoden for å gå fra binære tall til desimaltall og motsatt.</b>

<p>I det vanlig tallsystemet som vi kaller "base 10" har vi totalt 10 tall vi kan bruke før vi må bruke tall om igjen.For "base 2" som er binær og "base 16" som er hexadecimaler har vi henholdsvis 2 og 16 tall. For å gå fra base 2 til base 16 deler vi opp i 4 og 4 rekker med binære tall. Etter hvor tallet er i rekke opphøyer man med 8,4,2 og 1. F.eks: 1101 blir 1^8, 1^4, 0^2 og 1^1 som blir= 8,4,0,1.
  Summerer man disse sammen blir det 13, som tilsvarer D i hex.
Motsatt tar man en og en fra hexadesimal og ser i tabellen.

Ganger sifferet med 2 opphøyd i posisjonen den står i rekka. summerer alt.
fra desimal til binær: deler tallet på 2, hvis rest skriv 1, hvis ikke skriv 0. Deler tallet du fikk på nytt og skriver 1 hvis rest osv. helt til det bare er 1 igjen.</p>
<h4>      Oppgave B </h4>
<b>Beskriv kort metoden for å gå fra hexadesimale tall til desimaltall og motsatt.</b>

<p>bruker tabellen og deler opp hexadesimalene. skriver posisjonen til hexadesimalet (fra 0-15) og ganger med 16 opphøyd i posisjonen til tallet. F.eks. DEA - Dx16^2, Ex16^1, Ax16^0</p>

<h3> 2. Forstå algoritmer og utføre “benchmark”-tester på koden <h3>

<h4> Oppgave A </h4>
<a href="https://github.com/gormaar/Feil-Bruker/blob/Under-arbeid/sorting.go">Sorting.go</a>

<h4> Oppgave B </h4>
<a href="https://github.com/gormaar/Feil-Bruker/blob/Under-arbeid/sorting_test.go">Sorting_test.go</a>


<h4> Oppgave C </h4>
<a href="https://github.com/gormaar/Feil-Bruker/blob/Under-arbeid/benchmark.md">Benchmark.md</a>

<h3> 3. Forstå prosessadministrajon på et platform </h3>

<img src="https://github.com/gormaar/Feil-Bruker/blob/Under-arbeid/Screen%20Shot%202018-02-14%20at%2023.41.03.png">
<img src="https://github.com/gormaar/Feil-Bruker/blob/Under-arbeid/Screen%20Shot%202018-02-14%20at%2023.41.10.png">

<a href="https://github.com/gormaar/Feil-Bruker/blob/Under-arbeid/oppg3.go">Oppg3.go
  
<h3> 4. Typografiske symboler </h3>

<p>(Når vi kjørte koden "ascii.go" på mac fikk vi utskrift av alle tegnene. På windows manglet vi de første 27 tegnene. Ved nærmere analyse fant vi ut at eurosymbolet ikke printer når vi kjører programmet gjennom terminal, men printer når det blir kjørt gjennom Goland)</b>

<a href="https://github.com/gormaar/Feil-Bruker/blob/Under-arbeid/ascii.go">ascii.go

<a href="https://github.com/gormaar/Feil-Bruker/blob/Under-arbeid/iso_test.go">iso_test.go
