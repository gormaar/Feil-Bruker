Marius-MacBook-Pro:algorithms Marius$ go test -bench=.
goos: darwin
goarch: amd64

BenchmarkBSortM100-4    | 	  100000	|     15792 ns/op
-|-|-
BenchmarkBSortM1000-4    |	    2000	 |   897538 ns/op
BenchmarkBSortM10000-4   |	      10	| 131703143 ns/op
BenchmarkBSort100-4      |	  100000	|     18951 ns/op
BenchmarkBSort1000-4     |	    2000	|   1124657 ns/op
BenchmarkBSort10000-4    |	      10	| 149146395 ns/op
BenchmarkQSort100-4      |	  300000	|      4390 ns/op
BenchmarkQSort1000-4     |	   30000	|     48590 ns/op
BenchmarkQSort10000-4    |	    2000	|    617748 ns/op

PASS
ok  	_/Users/Marius/downloads/oblig1/src/algorithms	27.962s
Marius-MacBook-Pro:algorithms Marius$ 

BenchmarkBSortM100-4     |	  100000	|     15792 ns/op
-|-|-
BenchmarkBSortM1000-4    |	    2000	|    897538 ns/op
BenchmarkBSortM10000-4   |	      10	| 131703143 ns/op

BenchmarkBSort100-4      |	  100000	 |    18951 ns/op
-|-|-
BenchmarkBSort1000-4     |	    2000	|   1124657 ns/op
BenchmarkBSort10000-4    |	      10	| 149146395 ns/op

BenchmarkQSort100-4      |	  300000	 |     4390 ns/op
-|-|-
BenchmarkQSort1000-4     |	   30000	  |   48590 ns/op
BenchmarkQSort10000-4    |	    2000	  |  617748 ns/op

Om vi ser på tallene som benchmark-testene genererer kan vi se at BenchmarkBSortM100-4 sorterte 100 tall og kjørte 100.000 looper for å finne ut at den kjører i gjennomsnitt med en fart på 15.792 nanosekund/per loop. BenchmarkBSort100-4 brukte 158.951 ns/op på like mange looper, altså noe lengre tid. Dette blir mer og mer tydelig jo flere tall som sorteres, og modified-versjonen gjør det litt bedre på alle testene. Dette er sannsynligvis fordi modified ser bort fra det bakerste tallet for hver runde, fordi det ble sortert forrige runde.

Quicksort kommer desidert best ut i denne testen, den bruker desidert kortest tid per loop for alle testene. Ved å se på Big-O ser vi at:
            Best          Average       Worst
Bubble Sort Ω(n)          Θ(n^2)	      O(n^2)
Quicksort	  Ω(n log(n))	  Θ(n log(n))	  O(n^2)
Dette viser at om tallene er så og si ferdigsortert på forhånd er bubblesort mer effektiv enn quicksort, men i snitt gjør quicksort det bedre enn bubblesort - noe vi også så i vår benchmarktest.På sitt verste er de like dårlige.
