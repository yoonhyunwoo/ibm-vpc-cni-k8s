package ipam

import (
	"log"
	"sync"
)

// IPAM 관리 구조체
type IPAM struct {
	allocatedIPs map[string]bool
	mu           sync.Mutex
}

// 새 IP 할당
func (i *IPAM) AllocateIP() string {
	i.mu.Lock()
	defer i.mu.Unlock()

	// 임시적으로 더미 IP 할당 (IBM VPC API와 연동 필요)
	ip := "10.1.0.100"
	if _, exists := i.allocatedIPs[ip]; exists {
		return ""
	}

	i.allocatedIPs[ip] = true
	log.Printf("Allocated IP: %s", ip)
	return ip
}

// IP 해제
func (i *IPAM) ReleaseIP(ip string) {
	i.mu.Lock()
	defer i.mu.Unlock()

	delete(i.allocatedIPs, ip)
	log.Printf("Released IP: %s", ip)
}
