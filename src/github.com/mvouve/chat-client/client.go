package main

import(
    "fmt"
    "net"
    "os"
    "bufio"
)

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
func readmsgs(conn net.Conn){
    reader := bufio.NewReader(conn)
    for {
        msg, _ := reader.ReadString('\n')
        fmt.Printf(msg)
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
 *          This is the main server function
 ******************************************************************************/
func main(){
    if len(os.Args) < 2{
        fmt.Printf("usage: %s [Port]\n", os.Args[0])

        return
    }

    conn, err := net.Dial("tcp", os.Args[1])

    if err != nil  {
        fmt.Printf("Could not connect!\n")

        return
    }

    fmt.Printf("Connected\n")

    go readmsgs(conn)
    reader := bufio.NewReader(os.Stdin)
    for {
        msg, _ := reader.ReadBytes('\n')
        conn.Write(msg)
    }
}
