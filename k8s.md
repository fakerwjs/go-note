# Kubernetes 面试题及参考答案（完整 0-102 题）

---

## 一、基础概念

### Q0：Kubernetes 是什么？
**答案:** Kubernetes 是开源的容器编排平台，用于自动化部署、扩展和管理容器化应用。

### Q1：Pod 是什么？
**答案:** Pod 是 Kubernetes 的最小调度单元，一个 Pod 可以包含一个或多个容器，它们共享网络和存储。

### Q2：Deployment 与 StatefulSet 区别？
**答案:** Deployment 适合无状态应用，支持滚动更新；StatefulSet 适合有状态应用，保证 Pod 顺序和持久化卷。

### Q3：DaemonSet 有什么作用？
**答案:** DaemonSet 确保每个节点上运行一个 Pod，常用于日志收集和监控。

### Q4：Job 和 CronJob 区别？
**答案:** Job 用于一次性任务，CronJob 用于定时任务。

### Q5：Service 类型有哪些？
**答案:** ClusterIP、NodePort、LoadBalancer、ExternalName。

### Q6：Kubernetes 架构组成？
**答案:** Master 节点（kube-apiserver、etcd、kube-scheduler、kube-controller-manager）和 Node 节点（kubelet、kube-proxy、Container runtime）。

### Q7：etcd 的作用？
**答案:** etcd 用于存储集群的状态信息，保证一致性。

### Q8：Pod 生命周期状态？
**答案:** Pending、ContainerCreating、Running、Succeeded、Failed、Unknown。

### Q9：Pod 调度约束有哪些？
**答案:** NodeSelector、Node Affinity、Pod Affinity/AntiAffinity、Taints & Tolerations、资源 Requests/Limits。

### Q10：蓝绿发布与灰度发布区别？
**答案:** 蓝绿发布：旧版本和新版本同时存在，通过 Service 切换流量；灰度发布：逐步将部分流量导向新版本。

---

## 二、调度与扩缩容

### Q11：HPA 原理？
**答案:** 基于 Metrics Server 或自定义指标动态调整 Pod 副本数。

### Q12：VPA 原理？
**答案:** 自动调整 Pod 的 CPU/内存 Requests，优化资源分配。

### Q13：滚动更新策略？
**答案:** Deployment 通过新 ReplicaSet 替换旧 ReplicaSet，参数 maxUnavailable、maxSurge 控制更新速度。

### Q14：Pod 优先级和抢占？
**答案:** 可以为 Pod 设置优先级，低优先级 Pod 可被高优先级 Pod 驱逐。

### Q15：Pod 亲和与反亲和？
**答案:** Affinity/AntiAffinity 控制 Pod 调度到指定节点或靠近/远离其他 Pod。

---

## 三、网络与服务

### Q16：CNI 插件有哪些？
**答案:** Flannel、Calico、Cilium 等。

### Q17：Ingress 与 Service 区别？
**答案:** Service 用于内部/外部负载均衡；Ingress 负责 HTTP/HTTPS 路由、TLS、虚拟主机，需要 Ingress Controller。

### Q18：Service Mesh 作用？
**答案:** 管理微服务通信，提供流量控制、熔断、灰度发布、可观察性，常用 Istio、Linkerd、Consul。

### Q19：NetworkPolicy 是什么？
**答案:** NetworkPolicy 控制 Pod 间通信，基于命名空间、标签和端口实现微分段。

### Q20：DNS 原理？
**答案:** CoreDNS 提供 Service 名解析，Pod 通过 DNS 查找 Service IP。

---

## 四、存储与卷

### Q21：Pod 挂载卷类型？
**答案:** emptyDir、hostPath、PVC + PV、ConfigMap/Secret。

### Q22：PVC 与 PV 区别？
**答案:** PV 是实际存储资源，PVC 是 Pod 请求的存储声明。

### Q23：StorageClass 作用？
**答案:** 定义存储策略，实现动态卷创建。

### Q24：卷访问模式？
**答案:** ReadWriteOnce (RWO)、ReadOnlyMany (ROX)、ReadWriteMany (RWX)。

### Q25：CSI 插件作用？
**答案:** 容器存储接口，支持第三方存储动态挂载。

---

## 五、安全与访问控制

### Q26：RBAC 组成？
**答案:** Role/ClusterRole 定义权限，RoleBinding/ClusterRoleBinding 绑定用户或 ServiceAccount。

### Q27：Secret 与 ConfigMap 区别？
**答案:** Secret 保存敏感信息，ConfigMap 保存非敏感配置。

### Q28：ServiceAccount 作用？
**答案:** Pod 使用 ServiceAccount 访问 API Server，可绑定权限。

### Q29：Pod SecurityContext？
**答案:** 设置容器用户、权限、只读根文件系统等安全策略。

### Q30：NetworkPolicy 安全策略？
**答案:** 限制 Pod ingress/egress 流量，实现隔离。

---

## 六、CI/CD 与 GitOps

### Q31：Helm 作用？
**答案:** K8s 应用包管理，支持模板化、环境化部署。

### Q32：ArgoCD 作用？
**答案:** GitOps 自动化部署，管理多集群同步。

### Q33：Flux 作用？
**答案:** GitOps 工具，自动同步 K8s 配置。

### Q34：Kustomize 作用？
**答案:** 原生配置定制化工具，Overlay 管理环境差异。

### Q35：CI/CD 流程？
**答案:** 代码提交 → 测试 → 镜像构建 → Push registry → K8s 部署 → 流量切换 → 监控回滚。

---

## 七、监控与日志

### Q36：常用监控工具？
**答案:** Prometheus + Grafana, metrics-server。

### Q37：日志收集工具？
**答案:** Fluentd/Fluent Bit/Logstash + Elasticsearch/Loki + Grafana。

### Q38：Pod 日志收集方法？
**答案:** stdout/stderr 输出 + DaemonSet 收集。

### Q39：节点维护流程？
**答案:** cordon → drain → 升级 → uncordon。

### Q40：Pod 优雅终止？
**答案:** SIGTERM → preStop hook → terminationGracePeriodSeconds → 删除 Pod。

---

## 八、高可用与集群管理

### Q41：Master 高可用？
**答案:** 多 Master 节点 + etcd 多副本 + 外部 VIP/LoadBalancer。

### Q42：Node 高可用？
**答案:** 多节点冗余、Pod 分布式部署、自动扩容和恢复。

### Q43：多集群管理方式？
**答案:** KubeFed、Rancher、ArgoCD multi-cluster。

### Q44：etcd 高可用方案？
**答案:** 奇数节点集群、Leader 选举、定期备份。

### Q45：集群升级流程？
**答案:** Master 升级 → Node 升级 → kube-proxy/CNI 升级 → 验证 → 回滚策略。

---

## 九、CRD 与 Operator

### Q46：CRD 是什么？
**答案:** CustomResourceDefinition，用于扩展 Kubernetes API。

### Q47：Operator 原理？
**答案:** 控制循环管理 CRD 资源，实现自动化应用运维。

### Q48：CRD 与 Operator 关系？
**答案:** CRD 定义自定义资源，Operator 通过 Controller 管理 CRD 实例。

### Q49：Operator 使用场景？
**答案:** 数据库、Kafka、Redis 等有状态应用自动化管理。

### Q50：如何开发 Operator？
**答案:** 定义 CRD → 编写 Controller → 监听事件 → 管理资源生命周期。

---

## 十、故障排查与优化

### Q51：Pod Pending 排查？
**答案:** 资源不足、节点污点、亲和性不满足。

### Q52：CrashLoopBackOff 排查？
**答案:** 镜像拉取失败、命令错误、配置问题。

### Q53：节点 NotReady 排查？
**答案:** kubelet 日志、网络插件、资源压力、重启 kubelet。

### Q54：readinessProbe 与 livenessProbe 区别？
**答案:** readinessProbe 控制流量接入，livenessProbe 控制容器重启。

### Q55：Pod 优化调度？
**答案:** 资源 Requests/Limits、Affinity、Taints & Tolerations、Topology Spread。

---

## 十一、资源管理与 QoS

### Q56：资源 Requests 与 Limits？
**答案:** Requests 保证最低资源，Limits 限制最大资源。

### Q57：Pod QoS 分类？
**答案:** Guaranteed、Burstable、BestEffort。

### Q58：容器镜像拉取策略？
**答案:** Always、IfNotPresent、Never。

### Q59：ConfigMap/Secret 在 CI/CD 中管理？
**答案:** 注入环境变量或挂载卷，支持版本化和更新触发滚动更新。

### Q60：Pod 与 API Server 访问限制？
**答案:** ServiceAccount + RBAC + NetworkPolicy。

---

## 六十一至一百零二题（61-102）

### Q61：StatefulSet 的数据持久化机制？
**答案:** StatefulSet 为每个 Pod 创建独立 PVC，保证 Pod 重启或迁移后数据不丢失。

### Q62：滚动更新中 maxUnavailable 和 maxSurge 的作用？
**答案:** maxUnavailable：更新过程中最多不可用 Pod 数；maxSurge：更新过程中可额外创建的 Pod 数。

### Q63：Pod 调度失败可能原因？
**答案:** 资源不足、亲和性/反亲和性限制、节点污点未容忍、调度策略冲突。

### Q64：Pod 被驱逐原因？
**答案:** 节点资源不足、优先级低于其他 Pod、节点不可用或 Pod 违反限制。

### Q65：Horizontal Pod Autoscaler 限制？
**答案:** 依赖 Metrics Server 指标，仅水平扩缩容 Pod，响应短期突发流量有限。

### Q66：Vertical Pod Autoscaler 限制？
**答案:** 调整容器资源 Requests，适用于有状态服务，不能自动增加 Pod 副本数。

### Q67：Node Affinity 与 NodeSelector 区别？
**答案:** NodeSelector 简单标签匹配；Node Affinity 支持硬性/软性调度规则和表达式。

### Q68：Pod Affinity/AntiAffinity 用途？
**答案:** 控制 Pod 调度到靠近/远离其他 Pod，实现服务拓扑分布或隔离。

### Q69：Taints 与 Tolerations 用途？
**答案:** Taints 标记节点不接受某些 Pod，Tolerations 允许 Pod 忽略污点被调度到节点。

### Q70：Topology Spread Constraints 用途？
**答案:** 控制 Pod 在拓扑域（如节点、机架）间分布，避免集中部署单点故障。

### Q71：Ingress Controller 类型？
**答案:** NGINX、Traefik、HAProxy、Istio IngressGateway。

### Q72：TLS 在 Ingress 中配置方法？
**答案:** 创建 Secret 保存证书，然后在 Ingress resource 中指定 tls 字段。

### Q73：Service 类型选择原则？
**答案:** ClusterIP：集群内部服务；NodePort：简单对外暴露；LoadBalancer：云环境 LB；ExternalName：外部域名映射。

### Q74：NetworkPolicy 使用注意事项？
**答案:** 需要 CNI 插件支持，定义 ingress/egress 时覆盖默认允许策略。

### Q75：Pod 与外部通信方式？
**答案:** NodePort、LoadBalancer、Ingress、Egress gateway/NAT。

### Q76：CNI 插件选择原则？
**答案:** 集群规模、性能需求、网络策略支持、容器互通与安全性。

### Q77：IPVS 与 iptables 区别？
**答案:** iptables：链表匹配，性能一般；IPVS：基于内核虚拟服务器，高性能、支持多调度算法。

### Q78：Pod 网络通信模型？
**答案:** 每 Pod 拥有唯一 IP，Pod 间可直接通信，Service 做负载均衡，CNI 插件实现 Overlay 或 L2/L3 网络。

### Q79：Service Mesh 核心组件？
**答案:** 数据平面（Sidecar 代理，如 Envoy）、控制平面（配置与流量管理，如 Istio Pilot）。

### Q80：StatefulSet 使用 PVC 原因？
**答案:** 每 Pod 独立卷，Pod 重建后数据不丢失，支持顺序启动和更新。

### Q81：动态卷申请流程？
**答案:** PVC 创建 → StorageClass 调用 Provisioner 创建 PV → PVC 与 PV 绑定 → Pod 挂载。

### Q82：卷回收策略？
**答案:** Retain：保留 PV 数据；Delete：删除 PV 和数据；Recycle：废弃（不推荐）。

### Q83：卷共享访问方式？
**答案:** ReadWriteOnce (RWO)、ReadOnlyMany (ROX)、ReadWriteMany (RWX)。

### Q84：CSI 插件作用？
**答案:** 容器存储接口，支持第三方存储动态挂载，可扩展和标准化操作。

### Q85：蓝绿发布与灰度发布区别？
**答案:** 蓝绿：切流量到新版本；灰度：逐步放量。

### Q86：滚动更新与 Recreate 更新区别？
**答案:** 滚动更新：逐步替换 Pod；Recreate：删除旧 Pod 再创建新 Pod。

### Q87：readinessProbe 在发布中作用？
**答案:** 确保 Pod 准备就绪才接流量，避免流量切到未准备 Pod。

### Q88：Pod 调度策略优化？
**答案:** 优先级与抢占、Node Affinity/Pod Affinity、Taints & Tolerations、Topology Spread Constraints。

### Q89：资源限制及 QoS？
**答案:** Requests 保证资源，Limits 设置上限；QoS 分类：Guaranteed/Burstable/BestEffort。

### Q90：Pod 日志轮转与持久化？
**答案:** stdout/stderr 输出，结合 logrotate 或 Fluentd 集中收集，存储到持久卷或外部系统。

### Q91：Pod 安全上下文最佳实践？
**答案:** 限制 root 权限，readOnlyRootFilesystem，设置 runAsUser/runAsGroup。

### Q92：HPA 与 VPA 联合使用场景？
**答案:** HPA 控制副本数，VPA 优化单 Pod 资源，结合实现资源高效使用。

### Q93：Master 高可用实现？
**答案:** 多 Master 节点、etcd 多副本、外部 VIP 或 LoadBalancer 对外访问。

### Q94：Node 高可用策略？
**答案:** 多节点冗余、Pod 分布式部署、自动扩容和恢复。

### Q95：多集群管理方式？
**答案:** KubeFed 联邦集群、Rancher、ArgoCD multi-cluster。

### Q96：etcd 集群高可用方案？
**答案:** 奇数节点、Leader 选举、定期备份与监控。

### Q97：集群升级步骤？
**答案:** Master 升级 → Node 升级 → kube-proxy/CNI 升级 → 回滚策略与验证。

### Q98：CRD 是什么？
**答案:** CustomResourceDefinition，用于扩展 K8s API，自定义资源类型。

### Q99：Operator 原理？
**答案:** 监听 CRD 事件，通过控制循环管理资源生命周期，实现应用自动化运维。

### Q100：CRD 与 Operator 的关系？
**答案:** CRD 定义资源，Operator 控制 CRD 实例并执行逻辑。

### Q101：如何开发简单 Operator？
**答案:** 定义 CRD → 编写 Controller（Go/Kubebuilder）→ 处理事件 → 自动化部署和管理。

### Q102：常见 Operator 使用场景？
**答案:** 数据库（MySQL/Postgres）、Kafka、Redis 集群、有状态应用自动化管理。

