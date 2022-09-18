# Aufgaben

1. **Erstellen Sie einen Fork von diesem Repository!**
2. **Klonen Sie Ihren Fork, nicht das Original-Repository!**
3. **Reichen Sie Ihre Lösungen per Pull Request ein!**

Die folgenden Aufgaben können Sie direkt in die angegebene Datei lösen. Beachten
Sie hierzu die `// TODO: `-Kommentare im Code und die folgenden Instruktionen:

## 1) Sternzeichen aus Geburtsdatum ermitteln (`if`/`else`)

`ex1/main.go`: Die zwölf Tierkreiszeichen (Sternzeichen) werden für das
Geburtsdatum innerhalb eines Datumsbereiches vergeben, z.B. Krebs für ein
Geburtsdatum vom 22. Juni bis am 22. Juli. Auf Wikipedia finden Sie die
[komplette Tabelle](https://de.wikipedia.org/wiki/Tierkreiszeichen#Die_zw%C3%B6lf_Tierkreiszeichen_des_Zodiaks)
mit Sternzeichen und Datumsbereichen.

In der Funktion `outputWithZodiacSign` soll das Sternzeichen einer gegebenen
Person `p` ermittelt werden. (Die Ausgabe wurde schon ausprogrammiert.)
Bestimmen Sie das Sternzeichen mithilfe von `if` und `else` anhand des
Geburtsdatum der Person (genauer: `p.Day` und `p.Month`). Die Sternzeichen sind
oben bereits als Konstanten vordefiniert.

Starten Sie das Programm mit `go run ex1/main.go` und überprüfen Sie, ob die
korrekten Sternzeichen ausgegeben werden. Die Ausgabe sollte folgendermassen
aussehen:

```
Grace Hopper, born on 09.12.1906, has the zodiac sign ♐.
Dennis Ritchie, born on 09.09.1941, has the zodiac sign ♍.
Rick Astley, born on 06.02.1966, has the zodiac sign ♒.
Edsger Dijkstra, born on 11.05.1930, has the zodiac sign ♉.
Alan Turing, born on 23.06.1912, has the zodiac sign ♋.
```

### Zusatzaufgabe

Definieren Sie in der Funktion `main()` weiter unten noch drei weitere Personen,
deren Geburtsdatum und Sternzeichen Sie kennen. Rufen Sie die Funktion
`outputWithZodiacSign()` für diese Personen ebenfalls auf. Überprüfen Sie, ob
die Sternzeichen korrekt ausgegeben werden. (Funktionsaufrufe schauen wir erst
später an, Sie sollten aber die bestehenden Aufrufe einfach kopieren und
anpassen können.)

## 2) Datumsbereich aus Sternzeichen ermitteln (`switch`/`case`)

`ex2/main.go`: Dieses Programm soll die [Tabelle der
Tierkreiszeichen](https://de.wikipedia.org/wiki/Tierkreiszeichen#Die_zw%C3%B6lf_Tierkreiszeichen_des_Zodiaks)
ausgeben. Die Funktion `outputDateRange()` nimmt eine `rune` namens `zodiacSign`
entgegen. Zu diesem Sternzeichen sollen Sie den dazugehörigen Datumsbereich
ausgeben, z.B. `"21.03. - 20.04."` für das Sternzeichen Widder (_Aries_).

Es wurde bereits mit den ersten beiden Fällen begonnen. Dabei hat sich die
Lösungsvariante mit `if`, `else if` und `else` als umständlich herausgestellt.
Sie sollen den Code mittels `switch`/`case` umschreiben und vervollständigen.

Überprüfen Sie Ihre Lösung anschliessend anhand der verlinkten Tabelle.

### Zusatzaufgabe

Rufen Sie die Funktion `outputDateRange()` mit einem weiteren Zeichen (`rune`)
auf, das _kein_ Tierkreiszeichen ist. Behandeln Sie diesen Fall mit einem
`default`-Fall im zuvor geschriebenen `switch`/`case`-Konstrukt.

## 3) FizzBuzz (`for`)

`ex3/main.go`: FizzBuzz ist ein Spiel, das in einer Runde von Leuten gespielt
wird. Es beginnt bei der Zahl 1; und der Spieler an der Reihe spricht diese Zahl
aus. Jeder weitere Spieler erhöht diese Zahl um 1 und spricht sie aus ‒ ausser
wenn die Zahl durch 3 oder durch 5 restlos teilbar ist: Dann spricht er "Fizz"
(durch 3 teilbar) bzw. "Buzz" (durch 5 teilbar) aus. Beispiel:

1. 1
2. 2
3. Fizz
4. 4
5. Buzz
6. Fizz
7. 7
8. 8
9. Fizz
10. Buzz
11. 11

usw.

Ist eine Zahl durch 3 _und_ 5 teilbar, spricht man "FizzBuzz" aus.

Schreiben Sie eine `for`-Schleife, die von `Lower` bis `Upper` (vordefinierte
Konstanten) geht, und die Zählervariable gemäss den Spielregeln ausgibt:

- durch 3 teilbar: `"Fizz"`
- durch 5 teilbar: `"Buzz"`
- durch 3 _und_ 5 teilbar: `"FizzBuzz"`
- sonst: Zählervariable ausgeben

### Zusatzaufgabe

Falls Sie die Regeln mittels `if`/`else` umgesetzt haben, schreiben Sie diese
per `switch`/`case` um. Möglicherweise können Sie hierzu `fallthrough`
verwenden!

## 4) Spielkarten (`for`/`range`)

`ex4/main.go`: Für viele Kartenspiele verwendet man ein Kartenset bestehend aus
vier Farben (engl. _suits_) und neun Werten (engl. _ranks_). Die
[Deutschschweizer Jasskarten](https://jassverzeichnis.ch/deutsche-jasskarten/)
haben neun Werte von sechs bis neun ("nell"), sowie "Banner", "Under", "Ober",
"König", "Ass" und die vier Farben "Eichel", "Schelle", "Schilte" und "Rose".

Die [Französischen Karten](https://jassverzeichnis.ch/das-sind-die-franzoesischen-jasskarten/)
haben die Werte von sechs bis zehn, sowie "Bube" (engl.  _Jack_), "Dame" (engl.
_Queen_), König (engl. _King_) und Ass (engl. _Ace_) und die vier Farben "Karo"
(engl. _Diamonds_), "Pik" (engl. _Spades_), "Kreuz" (engl. _Clubs_) und "Herz
(engl. _Hearts_). Da der Unicode-Zeichensatz entsprechende Symbole definiert,
sollen hier die französischen Karten verwendet werden. (Am Stammtisch im Gasthof
Rössli in Escholzmatt verwenden Sie aber beim Jassen gefälligst weiterhin an die
Deutschschweizer Karten, sonst gibt es Ärger.)

Die Farben (`suits`) und Werte (`ranks`) sind als zwei `rune`-Slices
vordefiniert. Schreiben Sie eine Schleife, welche alle 36 Spielkarten in vier
Spalten (eine Spalte pro Farbe, mittels Tabs `\t` getrennt) ausgibt. Die Ausgabe
sollte etwa so aussehen:

```
◆⑥	♠⑥	♣⑥	♥⑥	
◆⑦	♠⑦	♣⑦	♥⑦	
◆⑧	♠⑧	♣⑧	♥⑧	
◆⑨	♠⑨	♣⑨	♥⑨	
◆⑩	♠⑩	♣⑩	♥⑩	
◆J	♠J	♣J	♥J	
◆Q	♠Q	♣Q	♥Q	
◆K	♠K	♣K	♥K	
◆A	♠A	♣A	♥A	
```

Tipp: Sie benötigen verschachtelte Schleifen mit `for` und `range` und die
Funktion `fmt.Printf()`.

### Zusatzaufgabe

Definieren Sie eine Map mit dem Namen `cards` und dem Typ `map[rune][]string`
und legen Sie zu jeder Farbe die Liste der dazugehörigen neun Spielkarten ab.
Wenn Sie die Funktion `fmt.Sprintf()` verwenden, können Sie viel vom bestehenden
Code wiederverwenden.
