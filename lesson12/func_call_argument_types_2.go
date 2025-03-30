// Technologie WebAssembly
//
// - rozhraní mezi jazyky Go a JavaScript
// - metoda PrintArguments vytiskne typy svých argumentů
// - lepší řešení založené na typu js.Type

package main

import (
	"fmt"
	"syscall/js"
)

// funkce, která se bude volat z HTML stránky, jakoby
// se jednalo o JavaScriptovou funkci
func PrintArguments(this js.Value, args []js.Value) any {
	fmt.Println("function PrintArguments called")
	fmt.Printf("# of parameters: %d\n", len(args))

	for i, arg := range args {
		fmt.Printf("parameter # %d with type '%s': %s\n",
			i,
			arg.Type().String(),
			arg.String())
	}

	// je nutné vrátit nějakou hodnotu
	return nil
}

func main() {
	fmt.Println("started")

	c := make(chan bool)

	// export funkce PrintArguments tak, aby byla volatelná
	// z JavaScriptu
	js.Global().Set("printArguments", js.FuncOf(PrintArguments))

	// realizace nekonečného čekání
	// (nutno provést při překladu do WebAssembly, ktežto
	// v případě použití GopherJS je možné hlavní funkci ukončit)
	<-c

	fmt.Println("finished")
}
