package main

import (
    "flag"
    "fmt"
    "net"
    "strconv"
    "strings"
    "sync"
    "time"
)

var (
    host     string
    portRange string
    tcpScan  bool
    udpScan  bool
)

func init() {
    flag.StringVar(&host, "host", "127.0.0.1", "Хост для сканирования")
    flag.StringVar(&portRange, "range", "1-1024", "Диапазон портов (например: 1-1000)")
    flag.BoolVar(&tcpScan, "tcp", false, "Сканировать TCP-порты")
    flag.BoolVar(&udpScan, "udp", false, "Сканировать UDP-порты")
}

func parsePortRange(r string) (int, int) {
    parts := strings.Split(r, "-")
    if len(parts) != 2 {
        fmt.Println("Неверный формат диапазона портов")
        return 1, 1024
    }
    from, _ := strconv.Atoi(parts[0])
    to, _ := strconv.Atoi(parts[1])
    return from, to
}

func scanTCPPort(host string, port int, wg *sync.WaitGroup) {
    defer wg.Done()
    addr := fmt.Sprintf("%s:%d", host, port)
    conn, err := net.DialTimeout("tcp", addr, 300*time.Millisecond)
    if err == nil {
        fmt.Printf("[TCP] Порт %d открыт\n", port)
        conn.Close()
    }
}

func scanUDPPort(host string, port int, wg *sync.WaitGroup) {
    defer wg.Done()
    addr := net.UDPAddr{
        IP:   net.ParseIP(host),
        Port: port,
    }

    conn, err := net.DialUDP("udp", nil, &addr)
    if err != nil {
        return
    }
    defer conn.Close()

    conn.SetDeadline(time.Now().Add(500 * time.Millisecond))
    conn.Write([]byte{})

    buf := make([]byte, 1024)
    _, _, err = conn.ReadFrom(buf)

    if err != nil {
        if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
            fmt.Printf("[UDP] Порт %d возможно открыт (нет ответа)\n", port)
        }
    } else {
        fmt.Printf("[UDP] Порт %d откликнулся\n", port)
    }
}

func main() {
    flag.Parse()
    from, to := parsePortRange(portRange)

    if !tcpScan && !udpScan {
        fmt.Println("Укажите хотя бы один протокол: -tcp или -udp")
        return
    }

    var wg sync.WaitGroup

    for port := from; port <= to; port++ {
        if tcpScan {
            wg.Add(1)
            go scanTCPPort(host, port, &wg)
        }
        if udpScan {
            wg.Add(1)
            go scanUDPPort(host, port, &wg)
        }
    }

    wg.Wait()
}
