package main

import (
	"fmt"
    "bufio"
	"os"
    "net/http"
	"os/exec"
    "net"
    "strconv"
    "sync"
    "time"
    "golang.org/x/crypto/ssh"
    "log"
)



func main() {
    var opcao int

    for {
		clear()
        fmt.Println("MENU:")
        fmt.Println("1. DIRBUSTER")
        fmt.Println("2. PORTSCAN")
        fmt.Println("3. BRUTE FORCE FTP/SSH")
        fmt.Println("0. Sair")
        fmt.Print("Escolha uma opção: ")
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
            brutemenu()
        default:
            fmt.Println("Opção inválida. Tente novamente.")
        }
    }
}
func brutemenu() {
    var opcao int

    for {
		clear()
        fmt.Println("MENU:")
        fmt.Println("1. SSH")
        fmt.Println("2. FTP")
        fmt.Println("0. Voltar")
        fmt.Print("Escolha uma opção: ")
        fmt.Scanln(&opcao)
		
        switch opcao {
        case 0:
            main()
            return
        case 1:
            sshb()
        case 2:
            port()
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
func sshb() {
    var host string
    var wordlist string

    // Lendo endereço do host e a lista de senhas
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Print("Digite o endereço do host: ")
    scanner.Scan()
    host = scanner.Text()

    fmt.Print("Digite o caminho da wordlist: ")
    scanner.Scan()
    wordlist = scanner.Text()

    // Abrindo o arquivo da wordlist
    file, err := os.Open(wordlist)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner = bufio.NewScanner(file)
    // Lendo a senha da wordlist e tentando se conectar com ela
    for scanner.Scan() {
        password := scanner.Text()

        err = sshConnect(host, "22", "root", password)
        if err == nil {
            fmt.Printf("Senha encontrada: %s\n", password)
            return
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}

func sshConnect(host, port, user, password string) error {
    config := &ssh.ClientConfig{
        User: user,
        Auth: []ssh.AuthMethod{
            ssh.Password(password),
        },
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }

    client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", host, port), config)
    if err != nil {
        return err
    }

    // Teste de conexão
    session, err := client.NewSession()
    if err != nil {
        return err
    }
    defer session.Close()

    return nil
}