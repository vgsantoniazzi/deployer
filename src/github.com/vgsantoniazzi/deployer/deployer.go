package deployer

import (
  "log"
  "os"
  "golang.org/x/crypto/ssh"
)

func main() {
  if len(os.Args) != 5 {
    log.Fatalln("usage: deployer <username> <password> <host>:<port> <cmd>")
  }

  clientConfig := &ssh.ClientConfig{
    User: os.Args[1],
    Auth: []ssh.AuthMethod{
      ssh.Password(os.Args[2]),
    },
  }

  client, err := ssh.Dial("tcp", os.Args[3], clientConfig)
  if err != nil {
    panic("Failed to dial: " + err.Error())
  }
  session, err := client.NewSession()
  if err != nil {
    panic("Failed to create session: " + err.Error())
  }
  defer session.Close()

  session.Stdout = os.Stdout
  session.Stderr = os.Stderr

  if err := session.Run(os.Args[4]); err != nil {
  panic("Failed to run: " + err.Error())
  }
}
