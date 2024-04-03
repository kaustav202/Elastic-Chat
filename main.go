package main

import (
    "fmt"
    "net/http"
    "sync"

    "github.com/go-redis/redis/v8"
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

type Client struct {
    ID   string
    Conn *websocket.Conn
}

var clients = make(map[string]*Client)
var clientsMutex sync.Mutex

var rdb *redis.Client

func initRedis() {
    rdb = redis.NewClient(&redis.Options{
        Addr:     "localhost:6379", 
        Password: "",              
        DB:       0,              
    })
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        fmt.Println("Error upgrading to WebSocket:", err)
        return
    }

    userID := r.URL.Query().Get("userID")
    if userID == "" {
        userID = generateUserID()
    }

    client := &Client{
        ID:   userID,
        Conn: conn,
    }

    clientsMutex.Lock()
    clients[userID] = client
    clientsMutex.Unlock()

    err = rdb.Publish(ctx, "websocket_channel", userID).Err()
    if err != nil {
        fmt.Println("Error publishing to Redis channel:", err)
    }

    defer func() {
        conn.Close()

        clientsMutex.Lock()
        delete(clients, userID)
        clientsMutex.Unlock()
    }()

    for {
        messageType, p, err := conn.ReadMessage()
        if err != nil {
            fmt.Println("Error reading message:", err)
            break
        }

        fmt.Printf("Received message from %s: %s\n", userID, p)


		err = conn.WriteMessage(messageType, p)
        if err != nil {
            fmt.Println("Error writing message:", err)
            break
        }
    }
}


func generateUserID() string {

	return "user-" + "randomID"
}

func main() {
    initRedis()


	pubsub := rdb.Subscribe(ctx, "websocket_channel")
    defer pubsub.Close()

    go func() {
        for {
            msg, err := pubsub.ReceiveMessage(ctx)
            if err != nil {
                fmt.Println("Error receiving message from Redis:", err)
                continue
            }
            fmt.Println("Received message from Redis channel:", msg.Payload)
        }
    }()

    http.HandleFunc("/ws", handleWebSocket)


	fmt.Println("Server listening on :8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println("Error starting server:", err)
    }
}
