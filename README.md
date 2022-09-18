# Verzweigungen und Schleifen

Programme bestehen aus Daten und Logik. Im
[ersten](https://code.frickelbude.ch/m346/go-1-vars-types-output) und
[zweiten](https://code.frickelbude.ch/m346/go-2-structs-slices-maps) Teil haben
wir primitive und zusammengesetzte Datentypen kennengelernt. In diesem Teil
wollen wir diese Daten mithilfe von Verzweigungen und Schleifen verarbeiten.

Go kennt nur relativ wenige Schlüsselwörter
([Spec](https://go.dev/ref/spec#Keywords)), 25 um genau zu sein. In diesem Teil
werden Sie zehn davon kennenlernen, nämlich (in alphabetischer Reihenfolge)
`break`, `case`, `continue`, `default`, `else`, `fallthrough`, `for`, `if`,
`range`, und `switch` ‒ `goto` lassen wir weg. Dementsprechend ist dieser
Kursteil auch anhand dieser Schlüsselwörter organisiert.

## Verzweigungen (_Branching_) 

Go kennt zwei Arten der Verzwegungen: `if`/`else` und `switch`/`case`. Die
Funktionsweise ist vergleichbar mit den entsprechenden Konstrukten aus anderen
Programmiersprachen, jedoch gibt es auch einige wichtige Unterschiede.

### `if`/`else`

Ein bedingter Codeblock kann mit `if` umgesetzt werden, wobei die Bedingung
_nicht_ in Klammern stehen muss:

```go
maxPoints := 100.0
scored := 75.0
ratio := scored / maxPoints

if ratio > 0.8 {
	fmt.Println("a ratio of", ratio, "is excellent")
}
if ratio > 0.6 {
	fmt.Println("a ratio of", ratio, "is good")
}
```

Es wird immer ein Codeblock benötigt, der auf der gleichen Zeile wie die
Bedingung anfangen muss.

Ausgabe:

```
a ratio of 0.75 is good
```

Alternativbedingungen können mit `else` formuliert werden. Mehrere
Alternativbedingungen lassen sich mittels `else if` definieren:

```go
if ratio > 0.8 {
	fmt.Println("a ratio of", ratio, "is excellent")
} else if ratio > 0.6 {
	fmt.Println("a ratio of", ratio, "is good")
} else if ratio > 0.4 {
	fmt.Println("a ratio of", ratio, "is acceptable")
} else {
	fmt.Println("a ratio of", ratio, "is poor")
}
```

### `switch`/`case`

Oftmals wird dieselbe Variable in allen Bedingungen mit einem konstanten Wert
verglichen:

```go
if grade == 6 {
	fmt.Println("very good")
} else if grade == 5 {
	fmt.Println("good")
} else if grade == 4 {
	fmt.Println("sufficient")
} else if grade == 3 {
	fmt.Println("insufficient")
} else if grade == 2 {
	fmt.Println("bad")
} else if grade == 1 {
	fmt.Println("very bad")
}
```

In diesem Fall ist ein `switch`/`case`-Konstrukt besser lesbar:

```go
case 6:
	fmt.Println("very good")
case 5:
	fmt.Println("good")
case 4:
	fmt.Println("sufficient")
case 3:
	fmt.Println("insufficient")
case 2:
	fmt.Println("bad")
case 1:
	fmt.Println("very bad")
}
```

Im Gegensatz zu anderen Programmiersprachen aus der C-Sprachfamilie ist _kein_
`break` nötig um einen Arm zu terminieren! (Das Gegenteilige Verhalten kann
jedoch mit `fallthrough` erreicht werden; siehe weiter unten.)

#### `default`

Beim `if`/`else`-Konstrukt fängt der `else`-Block alle Fälle ab, auf die keine
der vorhergehenden `if`- oder `else if`-Bedingungen gepasst hat. Beim
`switch`/`case`-Konstrukt gibt es hierfür den `default`-Block:

```go
switch grade {
case 6:
	fmt.Println("very good")
case 5:
	fmt.Println("good")
case 4:
	fmt.Println("sufficient")
case 3:
	fmt.Println("insufficient")
case 2:
	fmt.Println("bad")
case 1:
	fmt.Println("very bad")
default:
	fmt.Println("unknown grade")
}
```

#### `fallthrough`

Soll nach zutreffender Bedingung nicht nur der aktuelle Block ausgeführt werden,
sondern auch der nächste Block, kann dies mit `fallthrough` erreicht werden:

```go
switch grade {
case 6:
	fallthrough
case 5:
	fallthrough
case 4:
	fmt.Println("passed the exam")
case 3:
	fallthrough
case 2:
	fallthrough
case 1:
	fmt.Println("failed the exam")
default:
	fmt.Println("unknown result")
}
```

Für die Werte `4`, `5` und `6` von `grade` gilt die Prüfung als bestanden, für
die Werte `1`, `2` und `3` als nicht bestanden.

## Schleifen (_Loops_)

### `for`

#### `continue`

#### `break`

### `range`
