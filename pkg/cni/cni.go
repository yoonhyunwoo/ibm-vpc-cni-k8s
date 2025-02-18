package cni

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// CNI 요청을 처리하는 구조체
type CNIRequest struct {
	PodName string `json:"pod_name"`
	PodIP   string `json:"pod_ip"`
	ENI     string `json:"eni"`
}

// CNI 서버 실행
func StartServer() {
	http.HandleFunc("/cni/add", handleAdd)
	http.HandleFunc("/cni/del", handleDel)
}

// Pod 네트워크 추가 처리
func handleAdd(w http.ResponseWriter, r *http.Request) {
	var req CNIRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	log.Printf("Allocating ENI for Pod: %s, IP: %s\n", req.PodName, req.PodIP)
	fmt.Fprintln(w, "ENI Allocated:", req.ENI)
}

// Pod 네트워크 삭제 처리
func handleDel(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ENI Released")
}
