package main

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
)

func main() {

	exibeIntroducao()
	for {
		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs")
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço esse comando")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {

	var name = "Luan"
	var version float32 = 1.1
	age := 29
	fmt.Println("Hello, Mr. ", name)
	fmt.Println("This program is in version ", version)
	fmt.Println("Your age is ", age)

	fmt.Println("Type of variable 'name' is ", reflect.TypeOf(name))
}

func leComando() int {
	var comando int
	fmt.Scan(&comando)
	fmt.Println("Comando escolhido foi ", comando)
	return comando
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do programa")
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := []string{"https://random-status-code.herokuapp.com", "https://www.alura.com.br", "https://www.caelum.com.br"}

	for _, site := range sites {
		// fmt.Println("Posicao:", i, "site:", site)
		monitoraSite(site)
	}
}

func monitoraSite(site string) {
	resp, _ := http.Get(site)
	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso")
	} else {
		fmt.Println("Site:", site, "está com problemas. StatusCode:", resp.StatusCode)
	}
}
