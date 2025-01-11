package server

import (
	"net/http"

	"morbo/context"
	"morbo/db"
	"morbo/log"
)

type Connection struct {
	db      *db.DB
	writer  http.ResponseWriter
	request *http.Request
	log     log.Log
}

func BigEndianUInt40(b []byte) uint64 {
	_ = b[4]
	return uint64(b[4]) | uint64(b[3])<<8 | uint64(b[2])<<16 | uint64(b[1])<<24 | uint64(b[0])<<32
}

func NewConnection(
	handler *baseHandler,
	writer http.ResponseWriter,
	request *http.Request,
) *Connection {
	id := handler.newConnectionID()
	log := log.NewLog(id)
	return &Connection{handler.db, writer, request, log}
}

func (conn *Connection) SendOriginHeaders() {
	if origin := conn.request.Header.Get("Origin"); origin != "" {
		conn.writer.Header().Set("Access-Control-Allow-Origin", origin)
	}
	conn.writer.Header().Set("Vary", "Origin")
}

func (conn *Connection) Error(message string, statusCode int) {
	conn.DistinctError(message, message, statusCode)
}

func (conn *Connection) DistinctError(serverMessage string, userMessage string, statusCode int) {
	conn.log.Error.Println(serverMessage)
	conn.writer.WriteHeader(statusCode)
	conn.writer.Write([]byte(userMessage))
}

func (conn *Connection) ContextAlive(ctx context.Context) bool {
	if err := ctx.Err(); err != nil {
		switch err {
		case context.ErrCanceled:
			conn.Error("the request has been canceled by the server", http.StatusServiceUnavailable)
		case context.ErrDeadlineExceed:
			conn.Error("took too long to finish the request", http.StatusGatewayTimeout)
		}
		return false
	}
	return true
}
