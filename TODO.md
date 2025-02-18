### 🚀 **IBM VPC CNI Plugin 개발을 위한 TODO 리스트**

---

## ✅ **1. 프로젝트 기본 구조 정리**
- [x] 프로젝트 디렉토리 구조 설정 (`cmd/`, `pkg/`, `deploy/`, `config/`, `tests/` 등)
- [x] `main.go` 작성 및 Cobra CLI 초기화
- [x] `cmd/server.go` 작성 (CNI 서버 실행)
- [x] `cmd/eni_manager.go` 작성 (ENI 자동 관리)

---

## ✅ **2. CNI 기능 개발 (`pkg/cni/`)**
- [x] CNI 서버 시작 및 기본 엔드포인트 추가 (`/cni/add`, `/cni/del`)
- [x] CNI 요청을 처리할 구조체 (`CNIRequest`) 작성
- [ ] IBM Cloud VPC API 연동하여 ENI와 IP 자동 할당 구현
- [ ] `/cni/add` 엔드포인트에서 Pod에 ENI 및 IP를 동적으로 할당
- [ ] `/cni/del` 엔드포인트에서 Pod의 ENI 및 IP를 해제

---

## ✅ **3. IP 주소 관리 (IPAM, `pkg/ipam/`)**
- [x] `IPAM` 구조체 정의 및 동시성 처리 (`sync.Mutex` 사용)
- [ ] IBM Cloud API 연동하여 실제 VPC Subnet에서 IP 할당
- [ ] 특정 서브넷과 VPC 설정을 반영하여 IP 정책 적용
- [ ] 사전 IP 확보 기능 추가 (빠른 Pod 스케일링 지원)

---

## ✅ **4. ENI 관리 (`pkg/eni/`)**
- [x] `ManageENI()` 함수 정의하여 주기적으로 ENI 상태 확인
- [ ] IBM Cloud API 연동하여 ENI 동적 할당 및 해제 구현
- [ ] Pod 개수 증가 시 ENI 자동 추가
- [ ] 사용하지 않는 ENI 자동 해제 로직 구현
- [ ] 트래픽 증가 시 네트워크 리소스 자동 확장

---

## ✅ **5. 네트워크 성능 최적화**
- [ ] IBM Cloud VPC의 고속 네트워크 (최대 100 Gbps) 활용 검토
- [ ] ENI를 효율적으로 활용하여 네트워크 병목 방지
- [ ] Pod 간 직접 VPC 네트워크 통신 최적화

---

## ✅ **6. 네트워크 정책 및 보안**
- [ ] Kubernetes `NetworkPolicy` 지원 추가
- [ ] IBM Cloud VPC의 Security Group과 연동하여 트래픽 제어
- [ ] IBM Cloud IAM 역할 기반 접근 제어 (RBAC) 적용

---

## ✅ **7. 로깅 및 모니터링**
- [ ] `logrus` 또는 `zap`을 사용하여 구조적 로깅 추가
- [ ] IBM Cloud Monitoring, Prometheus, Grafana 연동
- [ ] 주요 네트워크 이벤트 로깅
- [ ] Pod 네트워크 상태 모니터링 메트릭 추가

---

## ✅ **8. 배포 및 운영 자동화**
- [ ] `deploy/helm-chart/` 디렉토리에 Helm Chart 작성
- [ ] Kubernetes `Deployment` 및 `DaemonSet` 매니페스트 작성
- [ ] CI/CD 파이프라인 설정 (GitHub Actions, Tekton 등)
- [ ] 컨테이너 이미지 빌드 및 IBM Cloud Container Registry (ICR) 푸시

---

## ✅ **9. 테스트 및 안정성 확보**
- [ ] Unit Test 추가 (각 모듈별 기능 검증)
- [ ] Integration Test 추가 (실제 Kubernetes 클러스터와 연동 테스트)
- [ ] E2E (End-to-End) 테스트 추가 (IBM Cloud 환경에서 전체 기능 검증)
- [ ] Kubernetes 클러스터에서 성능 및 부하 테스트 수행

---

## ✅ **10. 문서화 및 최종 정리**
- [ ] `README.md` 작성 (설치 방법, 사용법, 아키텍처 설명 포함)
- [ ] IBM Cloud API 사용 가이드 추가
- [ ] 기여 가이드 (`CONTRIBUTING.md`) 및 코드 스타일 문서화
- [ ] 이슈 및 PR 템플릿 설정 (GitHub Issues & Pull Requests)

---

### 🎯 **최종 목표**
✔ IBM Cloud VPC와 완벽히 연동되는 **Kubernetes CNI 플러그인** 완성  
✔ 확장성과 보안성을 갖춘 **네트워크 관리 솔루션** 구축  
✔ **Helm Chart** 기반의 손쉬운 배포 지원  
✔ **Prometheus & Grafana** 연동하여 실시간 모니터링 가능  
✔ **CI/CD 자동화 및 테스트 통과율 90% 이상** 유지

---

### **🚀 진행 방식**
1. 우선순위를 고려하여 CNI 기능 개발 → IPAM → ENI 관리 → 최적화 → 보안 및 모니터링 순으로 진행
2. 핵심 기능이 완성되면 IBM Cloud API와 연동하여 실환경 테스트 진행
3. 배포 및 운영 자동화를 통해 프로덕션 수준의 안정성을 확보
4. 문서화 및 테스트 커버리지 강화 후 정식 릴리즈

이 TODO 리스트를 기반으로 개발을 진행하면 IBM VPC 네트워크와 완벽히 연동되는 **고성능 Kubernetes CNI 플러그인**을 구축할 수 있습니다. 🚀
