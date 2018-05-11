<h1>      Obligatorisk oppgave 4 </h1>
<h2>      Gorm-Erik, Kevin, Marius </h2>

<h3> Systemspesifikasjon </h3>
<h4>Applikasjon og brukergrensesnitt</h4>
<p>Vi har utviklet en applikasjon som lokaliserer dagligvarebutikker ut i fra brukeren sin posisjon.
Applikasjonen finner brukeren sin posisjon, i tillegg til et antall butikker i en bestemt radius rundt brukeren. 
Dette blir representert i et kart i GUI med markører over brukeren sin posisjon, samt dagligvarebutikker i området rundt brukeren.
Navnet på butikkene vises ved å klikke på markørene. 
Kartet er interaktivt, så brukeren kan fritt zoome inn og ut, eller bevege seg rundt på kartet. 
Til høyre vises det en liste over de dagligvarebutikkene som applikasjonen har registrert rundt brukeren sin posisjon, med navn på butikkene. 
Dette gir en rask tilbakemelding om hvor brukeren befinner seg, samt hvor og hvilke dagligvarebutikker som finnes i området. 
Applikasjonen bruker Google Maps sin API for å visualisere kartet, finne brukeren sin posisjon og lokalisere butikker i nærheten av brukeren.
</p>

  
  
<h4> Hensikt og nytteverdi </h4>
<p> Hensikten med applikasjonen er å raskt og enkelt kunne se hvor dagligvarebutikker befinner seg i forhold til brukerens egen posisjon. 
Det er noe som vil kunne bli nyttig dersom en bruker skulle befinne seg på et helt nytt sted, uten kjennskap til hvor man kan handle dagligvarer. 
Det kan brukes til å finne nye butikker, alternative butikker, eller for å se over utvalget av butikker i sitt område. 
Applikasjonen er rask, og så enkel at omtrent hvem som helst vil kunne bruke den. 
For å kunne dra nytte av applikasjonen kreves det en enhet med tilgang til internett, nettleser og GPS. 
I tillegg er det viktig å tillate at nettsiden bruker din posisjon, ellers vil det ikke være mulig å få opp dagligvarebutikker i nærheten. 
Applikasjonen vil gi samfunn og brukere en økt oversikt over matbutikker i sitt område, i tillegg til et bedre geografisk bilde av sin egen posisjon. 
Brukere vil kunne effektivt lokalisere og velge foretrukne matbutikker ut i fra det utvalget de har rundt seg. 
Utover dette vil applikasjonen være kjekk for dem som er interesserte i butikker.
</p>

<h3>Systemarkitektur</h3>
<p>I denne applikasjonen har vi tatt utgangspunkt i at vi ønsker å lage noe som er oversiktlig og enkelt å bruke. Vi har utviklet en applikasjon som prosesserer data og gir bruker output med liten til ingen behov for interaksjon fra brukeren selv. 
Applikasjonen er hovedsakelig skrevet i HTML, CSS og JavaScript. Den bruker Go til håndtering av nettverk, og for å skrive applikasjonen til serveren. Vi har tatt en vurdering på at dette vil være hensiktsmessig ettersom at dette inneholdt omtrent alt vi trengte for å fremstille det vi ønsket. Det gir god tilgang til styling, struktur, og funksjonalitet, i tillegg til at det er bra å jobbe med sammen med Google Maps sin API, verktøyene som ligger i API-en og Google sin dokumentasjon.


<br> Når  applikasjonen kjøres settes det opp en lokal server på port 8080 med path “/”. For å sette opp serveren brukes det en Handler som parser og eksekverer HTML-filen. Handleren bruker en ResponseWriter for å skrive innholdet i HTML-filen til serveren. 
Filene i applikasjonen vil måtte bli lagret lokalt på et block device på en server-maskin, men når brukeren aksesserer applikasjonen gjennom en URL og nettleser vil dataen ligge i minnet.
HTML-filen består av HTML, CSS og JavaScript. HTML har blitt brukt for blant annet overskrifter, og generell struktur. CSS blir brukt for å style siden med blant annet bakgrunnsfarge.
JavaScript blir brukt for den overordnede funksjonaliteten på siden. Applikasjonen har blitt utviklet på både Windows 10 og macOS.
 

<br> Verktøyene fra Google Maps API-en som blir anvendt i applikasjonen består blant annet av Google Geolocation, Google Places, og Google Maps.
Google Maps blir brukt for å fremkalle selve kartet. Det består av Google sitt eget brukergrensesnitt og deres egen engine. Google Geolocation blir brukt for å få tilgang til brukerens posisjon. Google Places benyttes for å gjøre et søk over matbutikker i en gitt radius rundt brukerens posisjon.
I applikasjonen brukes det også andre verktøy, som for eksempel “Markers”, for å sette en markør over hvert element på kartet.

<br> Nettverksprotokollen som blir brukt i applikasjonen er HTTP. Det vil bli initialisert en three-way-handshake med serveren for å opprette forbindelse. Klienten sender forespørsel til serveren om å åpne nettsiden, og får deretter opp et kart over Norge i tillegg til en forespørsel om tilgang til klientens posisjon. Gis det tilgang, sendes brukerens posisjon til server som deretter sentrerer kartet til brukerens posisjon. Samtidig sendes det forespørsel om å vise alle dagligvarebutikker innenfor en gitt radius rundt brukeren. Serveren svarer med å legge butikkene inn på kartet og skrive ut en liste til høyre for kartet over alle butikker i området. 
Dersom det ikke blir gitt tillatelse til å benytte brukerens posisjon vil det bli behandlet med en feilmelding.


<br> Med tanke på testing har vi laget tester for Go-filen ved hjelp av testing-pakken fra Go-biblioteket. Denne pakken gjør det mulig å teste de filene som ligger i samme package som test-filen, ved å skrive kommando “go test” i terminal eller Command Prompt. I Go filen har vi laget test for både nettverksoppsettet, og parsing av HTML-filen, men siden Go-filen hovedsakelig blir brukt for å sette opp serveren er det den som står i fokus. Ettersom at HTML-filer ikke kan testes på denne måten har vi heller vært grundige med feilhåndtering i HTML-filen. Siden applikasjonen ikke tar imot brukerinput er den største kilden for potensielle feil være ved kall på Geolocation og Google Places.
Her kan det skje feil med tilgangen på brukerens posisjon, som også vil gå ut over om Google Places vil finne matbutikker. En annen ting som er verdt å merke seg er at Google Places vil bare ta med en håndfull butikker. Selv om APIen finner butikkene, ser det ut til at den har en øvre grense. </p>


<p>Bildet 1 viser kartet når bruker har godkjent bruk av stedsstjenester, med markers på butikkene i området rundt, og liste til høyre.</p>
<img src="https://github.com/gormaar/Feil-Bruker/blob/Under-arbeid/Oblig4/bilder/bilde1.png">
<p> Bildet 2 viser kartet når bruker ikke har godkjent bruk av stedsstjenester.</p>
<img src="https://github.com/gormaar/Feil-Bruker/blob/Under-arbeid/Oblig4/bilder/bilde2.png">
