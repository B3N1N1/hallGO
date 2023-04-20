package main

import (
	"fmt"
    "bufio"
	"os"
    "net/http"
	"os/exec"
	"github.com/fatih/color"
    "net"
    "strconv"
    "sync"
    "time"
)



func main() {
	c := color.New(color.FgRed)
	c.Add(color.FgYellow)
	c.Add(color.FgGreen)

    var opcao int

    for {
		clear()
        c.Println("MENU:")
        c.Println("1. DIRBUSTER")
        c.Println("2. PORTSCAN")
        c.Println("3. Opção 3")
        c.Println("0. Sair")
        c.Print("Escolha uma opção: ")
        fmt.Scanln(&opcao)
		
        switch opcao {
        case 0:
            fmt.Println("Saindo do programa...")
            return
        case 1:
            dirgo()
        case 2:
            port()
        case 3:
            fmt.Println("Você escolheu a opção 3")
        default:
            fmt.Println("Opção inválida. Tente novamente.")
        }
    }
}
func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
func dirgo() {
    scanner := bufio.NewScanner(os.Stdin)

    fmt.Print("Digite a URL alvo: ")
    scanner.Scan()
    targetURL := scanner.Text()

    fmt.Print("Digite o caminho da wordlist: ")
    scanner.Scan()
    wordlistPath := scanner.Text()

    wordlistFile, err := os.Open(wordlistPath)
    if err != nil {
        fmt.Println("Erro ao abrir o arquivo:", err)
        return
    }
    defer wordlistFile.Close()

    scanner = bufio.NewScanner(wordlistFile)

    for scanner.Scan() {
        dir := scanner.Text()

        url := targetURL + "/" + dir
        resp, err := http.Get(url)

        if err != nil {
            fmt.Println("Erro ao acessar o diretório", dir, err)
        } else if resp.StatusCode == 200 {
            fmt.Println("Diretório encontrado:", dir)
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Erro ao ler o arquivo:", err)
        return
    }
}
func port() {
    var target string
    fmt.Print("Digite o endereço IP do alvo: ")
    fmt.Scanln(&target)

    var wg sync.WaitGroup
    for port := 1; port <= 65535; port++ {
        wg.Add(1)
        go func(port int) {
            defer wg.Done()
            address := target + ":" + strconv.Itoa(port)
            conn, err := net.DialTimeout("tcp", address, 100*time.Millisecond)

            if err != nil {
                return
            }

            conn.Close()
            fmt.Printf("Porta %d aberta\n", port)
        }(port)
    }

    wg.Wait()
}