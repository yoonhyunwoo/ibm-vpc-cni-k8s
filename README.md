# IBM VPC CNI for Kubernetes

CNI 연습용 IBM CMI

**IBM VPC CNI for Kubernetes** is a Container Network Interface (CNI) plugin that integrates Kubernetes with IBM Cloud VPC networking. It dynamically allocates and manages network resources such as Elastic Network Interfaces (ENIs) and IP addresses for Pods, ensuring seamless communication within IBM Cloud Kubernetes Service (IKS), OpenShift, and self-managed Kubernetes clusters.

## ✨ Features

✅ **Pod Network Interface Management**
- Assigns IBM Cloud VPC IP addresses dynamically to each Pod.
- Manages ENIs (Elastic Network Interfaces) and scales efficiently.
- Supports multiple IPs per ENI for enhanced scalability.

✅ **IP Address Management (IPAM)**
- Dynamically allocates IP addresses within IBM Cloud VPC subnets.
- Implements subnet and VPC-specific IP policies.
- Pre-allocates IPs to speed up Pod scaling.

✅ **Network Performance Optimization**
- Leverages IBM Cloud VPC’s high-speed network (up to 100 Gbps).
- Prevents bottlenecks by efficiently managing ENIs.
- Enables direct Pod-to-Pod communication within the VPC.

✅ **Security & Network Policies**
- Supports Kubernetes `NetworkPolicy` for traffic control.
- Integrates with IBM Cloud VPC Security Groups.
- Implements role-based access control (RBAC) using IBM IAM.

✅ **Automatic ENI & Pod Scaling**
- Adds new ENIs dynamically as Pod count increases.
- Releases unused ENIs automatically to optimize resource usage.
- Scales network resources based on traffic demands.

✅ **Logging & Monitoring**
- Integrates with IBM Cloud Monitoring, Prometheus, and Grafana.
- Logs critical network events.
- Provides real-time monitoring of Pod network states.

---

## 🚀 Installation

### **1. Prerequisites**
- **Kubernetes 1.21+** (IBM Cloud Kubernetes Service, OpenShift, or self-managed)
- **Container runtime:** `containerd` or `CRI-O`
- **IBM Cloud CLI** installed (`ibmcloud`)
- **Helm 3+** (for deploying using Helm)

### **2. Deploy via Helm**
```sh
helm repo add ibm-vpc-cni https://github.com/yoonhyunwoo/ibm-vpc-cni-k8s/charts
helm install ibm-vpc-cni ibm-vpc-cni/ibm-vpc-cni-k8s
```

### **3. Manual Deployment**
If you prefer deploying manually:
```sh
kubectl apply -f deploy/deployment.yaml
```

---

## ⚙️ Configuration

You can customize the plugin’s behavior using the `config.yaml` file. Below is an example configuration:

```yaml
cni:
  network_name: "ibm-vpc-cni"
  log_level: "info"
  eni_pool_size: 5
  ipam:
    subnet: "10.1.0.0/24"
    allocate_ips: true
```

To apply a new configuration:
```sh
kubectl apply -f config/config.yaml
```

---

## 🛠 Usage

### **Check ENI Status**
```sh
kubectl get pods -n kube-system | grep ibm-vpc-cni
```

### **Manually Allocate an ENI**
```sh
ibm-vpc-cni eni-manager
```

### **Check CNI Logs**
```sh
kubectl logs -l app=ibm-vpc-cni -n kube-system
```

---

## 📊 Monitoring & Logging

This project integrates with Prometheus & Grafana for monitoring. You can expose CNI metrics by applying:

```sh
kubectl apply -f monitoring/prometheus.yaml
kubectl apply -f monitoring/grafana.yaml
```

Access Grafana dashboard:
```sh
kubectl port-forward svc/grafana 3000:3000 -n monitoring
```

---

## 🧪 Testing

To run unit and integration tests:

```sh
make test
```

Or run E2E tests in a real cluster:

```sh
kubectl apply -f tests/e2e.yaml
```

---

## 📦 Deployment

### **Deploy on Kubernetes**
Use Helm or `kubectl` to deploy:
```sh
helm install ibm-vpc-cni ibm-vpc-cni/ibm-vpc-cni-k8s
```

Or manually:
```sh
kubectl apply -f deploy/deployment.yaml
```

### **Upgrade**
```sh
helm upgrade ibm-vpc-cni ibm-vpc-cni/ibm-vpc-cni-k8s
```

### **Uninstall**
```sh
helm uninstall ibm-vpc-cni
```

Or manually:
```sh
kubectl delete -f deploy/deployment.yaml
```

---
