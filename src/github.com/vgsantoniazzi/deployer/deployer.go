package main

import (
  "log"
  "os"
  "strings"
  "github.com/vgsantoniazzi/deployer/parser"
  "golang.org/x/crypto/ssh"
)

func main() {
  RunRemote(Commands())
}

func RunRemote(commands []string) {
  if err := Session().Run(strings.Join(commands, " && ")); err != nil {
    panic("Failed to run: " + err.Error())
  }
}

func Session() (*ssh.Session) {
  session, err := Client().NewSession()
  if err != nil {
    panic("Failed to create session: " + err.Error())
  }

  session.Stdout = os.Stdout
  session.Stderr = os.Stderr
  return session
}

func Client() (*ssh.Client) {
  client, err := ssh.Dial("tcp", Auth().Host +":"+ Auth().Port, ClientConfig())
  if err != nil {
    panic("Failed to dial: " + err.Error())
  }
  return client
}

func ClientConfig() (*ssh.ClientConfig) {
  return &ssh.ClientConfig{
    User: Auth().Username,
    Auth: []ssh.AuthMethod{
      ssh.Password(Auth().Password),
    },
  }
}

func Auth()(*parser.Auth) {
  return parser.Access(os.Args[1])
}

func Commands() ([]string) {
  if len(os.Args) < 2 || !strings.Contains(os.Args[1], ".yml") {
    log.Fatalln("usage: deployer path/to/file.yml")
  }
  return parser.Commands(os.Args[1])
}


