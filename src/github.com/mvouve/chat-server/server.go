/*******************************************************************************
 * SOURCE FILE: server.go - A simple chat server.
 *
 * PROGRAM: basic go chat
 *
 * FUNCTIONS:
 *      readmsgs(Listener conn)
 *      main()
 *
 * DATE: June 9, 2015
 *
 * REVISIONS: none
 *
 * DESIGNER: Marc Vouve
 *
 * PROGRAMMER: Marc Vouve
 *
 * NOTES: The program acts as the server in a simple client server application.
 ******************************************************************************/
package main

import (
    "fmt"
    "net"
    "os"
    "bufio"
)

var connections []net.Conn

/*******************************************************************************
 * FUNCTION:    readmsgs
 *
 * DATE:        June 9, 2015
 *
 * REVISIONS:   NONE
 *
 * PARAMS:      Listener conn - an established connection to listen on.
 *
 * DESIGNER:    Marc Vouve
 *
 * PROGRAMMER:  Marc Vouve
 *
 * INTERFACE:   func readmsgs(conn net.Conn)
 *
 * RETURN:      VOID
 *
 * NOTES:
 *          This function is intended to be called as a goroutine, needs an
 *          active connection, will read from that connection, and forward
 *          the messages to other users.
 ******************************************************************************/
func readmsgs(conn net.Conn) {
    reader := bufio.NewReader(conn)
    for {
        msg, _ := reader.ReadString('\n')
        if len(msg) < 1 {
            msg = conn.RemoteAddr().String() + " disconnected\n"
            fmt.Printf(msg)
            for idx, val := range connections {
                if val != conn {
                    connections = connections[:idx+copy(connections[:idx], connections[idx+1:])]
                }
            }
            return
        }

        msg = "[" + conn.RemoteAddr().String() + "] " + msg
        fmt.Printf(msg)
        for _, val := range connections {
            if val != conn {
                fmt.Fprintf(val, msg)
            }
        }
    }
}

/*******************************************************************************
 * FUNCTION:    main
 *
 * DATE:        June 9, 2015
 *
 * REVISIONS:   NONE
 *
 * PARAMS:      NONE
 *
 * DESIGNER:    Marc Vouve
 *
 * PROGRAMMER:  Marc Vouve
 *
 * INTERFACE:   func main()
 *
 * RETURN:      VOID
 *
 * NOTES:
 *          This function is intended to be called as a goroutine, needs an
 *          active connection, will read from that connection, and forward
 *          the messages to other users.
 ******************************************************************************/
func main() {
    if len(os.Args) < 2{
        fmt.Printf("usage: %s [Port]\n", os.Args[0])

        return
    }
    ln, err := net.Listen("tcp", ":" + os.Args[1])
    if err != nil {
        fmt.Printf("Unexpected error occured!\n")

        return
    }

    for {
        conn, err := ln.Accept()
        if err != nil {
            fmt.Printf("Connection received error initializing %d", err)
        } else {
            connections = append(connections, conn);
            fmt.Printf(conn.RemoteAddr().String() + " has connected to the chat.\n")
            go readmsgs(conn)
            for _, val := range connections {
                msg := conn.RemoteAddr().String() + " has connected to the chat.\n"
                val.Write([]byte(msg))
            }
        }
    }

}
