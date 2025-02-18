package eni

import (
	"log"
	"time"
)

// ENI 관리 함수
func ManageENI() {
	for {
		log.Println("Checking ENI status...")
		// IBM Cloud API를 호출하여 ENI 추가/삭제 처리 (추후 구현 필요)
		time.Sleep(10 * time.Second)
	}
}
