package raft

import (
	"bytes"
	"fmt"
	"testing"
)

func BenchmarkAppendEntriesRequestEncoding(b *testing.B) {
	req, tmp := createTestAppendEntriesRequest(2000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		req.Encode(&buf)
	}
	b.SetBytes(int64(len(tmp)))
}

func Test_AppendEntriesRequestEncoding(t *testing.T) {
	req, _ := createTestAppendEntriesRequest(1)
	//fmt.Printf("%+v\n", req)
	var buf bytes.Buffer
	req.Encode(&buf)
	//fmt.Printf("%s", string(buf))
	req.Decode(bytes.NewReader(buf.Bytes()))
	//fmt.Printf("%+v", req)
	//createTestAppendEntriesRequest(2)
}

func BenchmarkAppendEntriesRequestDecoding(b *testing.B) {
	req, buf := createTestAppendEntriesRequest(2000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req.Decode(bytes.NewReader(buf))
	}
	b.SetBytes(int64(len(buf)))
}

func BenchmarkAppendEntriesResponseEncoding(b *testing.B) {
	req, tmp := createTestAppendEntriesResponse(2000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		req.Encode(&buf)
	}
	b.SetBytes(int64(len(tmp)))
}

func BenchmarkAppendEntriesResponseDecoding(b *testing.B) {
	req, buf := createTestAppendEntriesResponse(2000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req.Decode(bytes.NewReader(buf))
	}
	b.SetBytes(int64(len(buf)))
}

func createTestAppendEntriesRequest(entryCount int) (*AppendEntriesRequest, []byte) {
	entries := make([]*LogEntry, 0)
	for i := 0; i < entryCount; i++ {
		command := &DefaultJoinCommand{Name: "localhost:1000"}
		entry, _ := newLogEntry(nil, nil, 1, 2, command)
		entries = append(entries, entry)
		fmt.Printf("entries[i].pb %s", entries[i].pb)
	}
	req := newAppendEntriesRequest(1, 1, 1, 1, "leader", entries)

	var buf bytes.Buffer
	req.Encode(&buf)
	//fmt.Printf("%+v\n%s", req, buf.String())

	return req, buf.Bytes()
}

func createTestAppendEntriesResponse(entryCount int) (*AppendEntriesResponse, []byte) {
	resp := newAppendEntriesResponse(1, true, 1, 1)

	var buf bytes.Buffer
	resp.Encode(&buf)

	return resp, buf.Bytes()
}
