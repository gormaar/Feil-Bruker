<h1>      Obligatorisk oppgave 3 </h1>
<h2>      Gorm-Erik, Kevin, Marius </h2>

<h3> Oppgave 1 </h3>
<p> a) Sett opp en webserver som lytter til port 8080.
<br> b) Når klienter aksesserer webserveren på path lik "/" skal klienten få tilbake strengen "Hello, client".
<br> Strengen skal vises i nettleseren.</p>

<a href="https://github.com/gormaar/Feil-Bruker/blob/master/Oblig3/src/Oppgave1/Oppgave1.go"> Oppgave1.go
  
  
<h3> Oppgave 2 API kall og behandling av JSON og HTML templates </h3>
<p> a) Presenter data på webserveren slik at den er leselig for mennesker
<br> (Ren tekst, f.eks. "Sted: Jorenhholem, dato: 13.04.2018", ikke i curly brackets.) 
<br> når klienter aksesserer stiene /1, /2, /3, /4, /5. </p>

<p> Dataen som skal presenteres skal hentes fra fem ulike APIer, hvor alle returnerer data i JSON format. 
<br> Dere velger selv hvilke datasett dere benytter. Det er denne dataen som skal presenteres på de ulike stiene på webserveren.</p>

<p> b) Samtlige stier i oppgave 2 skal rendres til klienter ved bruk av Go templates. </p>

<a href="https://github.com/gormaar/Feil-Bruker/blob/master/Oblig3/src/Oppgave2/Oppgave2.go"> Oppgave2.go
<br> <a href="https://github.com/gormaar/Feil-Bruker/blob/master/Oblig3/src/Oppgave2/Templates"> Templates


<h3> Oppgave 3 UDP/TCP server og Internett standarder </h3>
<p> Implementer et serverprogram i henhold til RFC 865. Serveren skal svare både på UDP og TCP.</p>


<a href="https://github.com/gormaar/Feil-Bruker/blob/master/Oblig3/src/Oppgave3/server.go"> server.go
<br> <a href="https://github.com/gormaar/Feil-Bruker/blob/master/Oblig3/src/Oppgave3/clientUDP.go"> clientUDP.go
<br> <a href="https://github.com/gormaar/Feil-Bruker/blob/master/Oblig3/src/Oppgave3/clientTCP.go"> clientTCP.go
