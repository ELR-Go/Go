/*
 * Servidor daytime baseado em:
 * https://github.com/nbeyer/trifling-servers/blob/master/daytime-server.go
 */
package main

import (
	"fmt"
	"net"
	"time"
)

func checaErro(err error) {
	if err != nil {
		panic(err)
	}
}

func geraDaytime() string {
	t := time.Now()
	return t.Format(time.RFC3339) + "\n"
}

func gerenciaConexao(conn net.Conn) {
	daytime := geraDaytime()

	// Send current daytime to client
	conn.Write([]byte(daytime))
	conn.Close()
}

func main() {
	// Create listening socket
	listen, err := net.Listen("tcp", "127.0.0.1:13")
	checaErro(err)
	defer listen.Close()

	fmt.Println("** Servidor daytime em funcionamento! **")
	fmt.Printf("Escutando por conexões em %s\n", listen.Addr().String())

	// Listen for incoming connections
	for {
		// Create a new socket for the connection
		conn, err := listen.Accept()
		checaErro(err)

		fmt.Printf("  Conexão aceita de %s\n", conn.RemoteAddr().String())
		go gerenciaConexao(conn)
	}
}
