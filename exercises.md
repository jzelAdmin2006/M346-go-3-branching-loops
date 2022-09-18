# Aufgaben

1. **Erstellen Sie einen Fork von diesem Repository!**
2. **Klonen Sie Ihren Fork, nicht das Original-Repository!**
3. **Reichen Sie Ihre Lösungen per Pull Request ein!**

Die folgenden Aufgaben können Sie direkt in die angegebene Datei lösen. Beachten
Sie hierzu die `// TODO: `-Kommentare im Code und die folgenden Instruktionen:

## 1) Sternzeichen aus Geburtsdatum ermitteln

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

## 2) Datumsbereich aus Sternzeichen ermitteln

`ex2/main.go`: Dieses Programm soll die [Tabelle der
Tierkreiszeichen](https://de.wikipedia.org/wiki/Tierkreiszeichen#Die_zw%C3%B6lf_Tierkreiszeichen_des_Zodiaks)
ausgeben. Die Funktion `outputDateRange()` nimmt eine `rune` namens `zodiacSign`
entgegen. Zu diesem Sternzeichen sollen Sie den dazugehörigen Datumsbereich
ausgeben, z.B. `"21.03. - 20.04."` für das Sternzeichen Widder (_Aries_).

Es wurde bereits mit den ersten beiden Fällen begonnen. Dabei hat sich die
Lösungsvariante mit `if`, `else if` und `else` als umständlich herausgestellt.
Sie sollen den Code mittels `switch`/`case` umschreiben und vervollständigen.

### Zusatzaufgabe

Rufen Sie die Funktion `outputDateRange()` mit einem weiteren Zeichen (`rune`)
auf, das _kein_ Tierkreiszeichen ist. Behandeln Sie diesen Fall mit einem
`default`-Fall im zuvor geschriebenen `switch`/`case`-Konstrukt.
