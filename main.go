package main

import (
    "fmt"
)

func main() {
    var opcao int

    for {
        fmt.Println("MENU:")
        fmt.Println("1. Opção 1")
        fmt.Println("2. Opção 2")
        fmt.Println("3. Opção 3")
        fmt.Println("0. Sair")
        fmt.Print("Escolha uma opção: ")
        fmt.Scanln(&opcao)

        switch opcao {
        case 0:
            fmt.Println("Saindo do programa...")
            return
        case 1:
            direc()
        case 2:
            fmt.Println("Você escolheu a opção 2")
        case 3:
            fmt.Println("Você escolheu a opção 3")
        default:
            fmt.Println("Opção inválida. Tente novamente.")
        }
    }
}
func direc() {
    var opcao int

    for {
        fmt.Println("MENU:")
        fmt.Println("1. scan")
        fmt.Println("2. Opção 2")
        fmt.Println("3. Opção 3")
        fmt.Println("0. Sair")
        fmt.Print("Escolha uma opção: ")
        fmt.Scanln(&opcao)

        switch opcao {
        case 0:
            fmt.Println("Saindo do programa...")
            return
        case 1:
            fmt.Println("Você escolheu a opção 1")
        case 2:
            fmt.Println("Você escolheu a opção 2")
        case 3:
            fmt.Println("Você escolheu a opção 3")
        default:
            fmt.Println("Opção inválida. Tente novamente.")
        }
    }
}
