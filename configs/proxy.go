package configs

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"

	"golang.org/x/term"
)

func StartInteractiveProxyConfigurationSession() {
	fmt.Println("Starting interactive session")
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Proxy type [httpsProxy, httpProxy]: ")
	proxyType, _ := reader.ReadString('\n')
	proxyType = strings.Replace(proxyType, "\n", "", -1)

	fmt.Print("Hostname (without http/https): ")
	hostname, _ := reader.ReadString('\n')
	hostname = strings.Replace(hostname, "\n", "", -1)

	fmt.Print("Port: ")
	port, _ := reader.ReadString('\n')
	port = strings.Replace(port, "\n", "", -1)

	fmt.Print("Username (Press enter if auth is not required): ")
	username, _ := reader.ReadString('\n')
	username = strings.Replace(username, "\n", "", -1)

	fmt.Print("Password: ")
	password, _ := term.ReadPassword(int(syscall.Stdin))
	SetupProxyInConfiguration(proxyType, hostname, port, username, string(password))
}

func SetupProxyInConfiguration(proxyType string, hostname string, port string, username string, password string) {
	fmt.Println("\nSetting up proxy")
	if !contains([]string{"httpsProxy", "httpProxy"}, proxyType) {
		log.Fatal("proxyType is invalid or empty")
	}
	if hostname == "" {
		log.Fatal("Hostname is required but empty")
	}
	if port == "" {
		log.Fatal("Port is required but empty")
	}

	if username == "" || password == "" {
		log.Println("Username or password is not passed. So no authentication will be used")
	}
	fmt.Println(password)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
