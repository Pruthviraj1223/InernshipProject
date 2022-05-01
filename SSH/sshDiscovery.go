package SSH

import (
	_ "encoding/json"
	"fmt"
	_ "fmt"
	"golang.org/x/crypto/ssh"
	_ "golang.org/x/crypto/ssh"
	"strings"
	"time"
)

func SshDiscovery(data map[string]interface{}) string {

	sshUser := (data["name"]).(string)

	sshPassword := (data["password"]).(string)

	sshHost := (data["ip"]).(string)

	sshPort := int((data["port"]).(float64))

	config := &ssh.ClientConfig{

		Timeout: 10 * time.Second,

		User: sshUser,

		HostKeyCallback: ssh.InsecureIgnoreHostKey(),

		Config: ssh.Config{Ciphers: []string{

			"aes128-ctr", "aes192-ctr", "aes256-ctr",
		}},
	}

	config.Auth = []ssh.AuthMethod{ssh.Password(sshPassword)}

	addr := fmt.Sprintf("%s:%d", sshHost, sshPort)

	sshClient, err := ssh.Dial("tcp", addr, config)

	if err != nil {
		return err.Error()
	}

	defer sshClient.Close()

	return "true"
}

func standardizeSpaces(s string) string {

	return strings.Join(strings.Fields(s), " ")

}
