# DevOps & CI/CD 面试题与知识图谱（进阶详解版）

---

## 目录
1. [DevOps 理论与核心理念](#devops-理论与核心理念)
2. [CI/CD 流程与架构设计](#cicd-流程与架构设计)
3. [工具深度题](#工具深度题)
4. [云原生与 Kubernetes 实践](#云原生与-kubernetes-实践)
5. [DevSecOps 与自动化安全](#devsecops-与自动化安全)
6. [SRE 与可观测性](#sre-与可观测性)
7. [实战场景与问题分析](#实战场景与问题分析)
8. [实战练习题与参考答案](#实战练习题与参考答案)

---

## DevOps 理论与核心理念

### 题 1：什么是 DevOps？
- **定义 / 概念**：  
  DevOps 是 Development（开发）与 Operations（运维）的结合，强调文化、协作、自动化和持续改进。  
- **流程 / 核心方法**：  
  - 建立跨团队协作机制  
  - 自动化构建、测试、部署  
  - 持续监控与反馈  
- **涉及方面 / 技术点**：CI/CD pipeline、自动化测试、监控、日志系统  
- **注意点 / 风险**：  
  - 团队文化落地困难  
  - 自动化不完善容易引入新问题  
- **示例**：使用 Jenkins/GitLab CI 自动化代码构建和部署  

### 题 2：DevOps 生命周期包含哪些阶段？
- **定义 / 概念**：计划→开发→构建→测试→发布→部署→运行→监控  
- **流程 / 操作步骤**：  
  1. Plan：需求分析与版本计划  
  2. Code：开发、版本控制  
  3. Build：编译、打包  
  4. Test：自动化单元测试、集成测试  
  5. Release：生成发布包/镜像  
  6. Deploy：部署到环境  
  7. Operate：运行、日志收集、监控  
  8. Monitor：监控指标、异常报警、反馈到 Plan  
- **涉及方面 / 技术点**：CI/CD、自动化测试、日志监控、告警系统  
- **注意点 / 风险**：各阶段闭环不完善可能导致质量问题  
- **示例图占位**：`![DevOps 生命周期](./images/devops-lifecycle.png)`  

### 题 3：DevOps 与传统运维区别
- **定义 / 概念**：传统运维偏人工，DevOps 强调自动化、持续交付  
- **涉及方面**：部署频率、反馈周期、协作模式、自动化程度  
- **注意点**：文化落地是核心难点  
- **表格**：

| 特性 | 传统运维 | DevOps |
|------|-----------|--------|
| 部署频率 | 低，手动 | 高，自动化持续交付 |
| 反馈周期 | 长 | 短，快速迭代 |
| 协作模式 | 分工明确 | 跨团队协作 |
| 自动化程度 | 低 | 高，CI/CD、IaC |

### 题 4：什么是 Infrastructure as Code（IaC）？
- **定义 / 概念**：用代码管理基础设施，支持版本控制、可复现、可审计  
- **流程 / 操作步骤**：写模块化配置 → 执行 plan/apply → 管理 state → 监控变更  
- **涉及方面 / 技术点**：Terraform、Ansible、CloudFormation、CI/CD 集成  
- **注意点**：state 管理不当可能导致资源冲突或丢失  
- **示例**：Terraform 管理 AWS VPC 和 EC2  

---

## CI/CD 流程与架构设计

### 题 1：什么是 CI/CD？区别是什么？
- **定义 / 概念**：  
  - CI：持续集成 → 自动构建与测试代码  
  - CD：持续交付/部署 → 自动发布  
- **流程 / 核心方法**：  
  CI: Checkout → Build → Test  
  CD: Package → Deploy → Verify → Switch/Release  
- **涉及方面 / 技术点**：Git、Jenkins/GitLab CI、Artifact Repo、测试框架  
- **注意点 / 风险**：缺少自动化测试会导致上线风险  
- **示例**：Jenkins pipeline 自动执行 CI/CD  

### 题 2：零停机部署策略（蓝绿 / 滚动 / 金丝雀）
#### 蓝绿部署（Blue-Green Deployment）
- **定义 / 概念**：生产环境同时保留两套服务环境，一套服务线上，另一套部署新版本  
- **流程**：部署新版本 → 测试 → 切流量 → 原环境保留用于回滚  
- **涉及方面**：负载均衡、CI/CD pipeline、健康检查  
- **注意点**：数据库迁移兼容、资源成本高  
- **示例**：Nginx/Ingress 切换 blue/green 服务  

#### 滚动升级（Rolling Update）
- **定义 / 概念**：逐步替换 Pod，保证服务持续可用  
- **流程**：设置 Deployment 更新策略 → 分批替换 Pod → 每批健康检查  
- **涉及方面**：K8s Deployment、服务发现  
- **注意点**：Pod 启动慢可能影响流量，依赖服务需兼容  
- **示例**：K8s `maxUnavailable=1 maxSurge=1`  

#### 金丝雀发布（Canary Release）
- **定义 / 概念**：新版本仅对部分流量或用户开放，降低发布风险  
- **流程**：部署少量实例 → 监控指标 → 逐步扩大流量 → 完全切换  
- **涉及方面**：流量控制（Istio/K8s Service）、监控系统、CI/CD pipeline  
- **注意点**：灰度比例过大风险高，数据库兼容性需保证，必须有回滚机制  
- **示例**：Istio VirtualService 流量分配 10% → 50% → 100%  

---

## 工具深度题（Jenkins / GitLab CI / ArgoCD / Terraform / Ansible）

### Jenkins
#### Jenkinsfile 与普通脚本区别
- **定义 / 概念**：声明式 pipeline，可版本控制、多分支支持  
- **流程**：Checkout → Build → Test → Deploy  
- **涉及方面**：Pipeline stage、agent 节点、artifact  
- **注意点**：脚本错误易导致整个 pipeline 停止  
- **示例**：`pipeline { agent any; stages { stage('Build') { steps { sh 'make' } } } }`  

#### 构建优化
- 缓存依赖（Docker layer / Maven cache）  
- 并行 stage  
- 分配合适节点（agent 标签）  

### GitLab CI
- YAML pipeline 结构：stages、jobs、scripts、artifacts、cache  
- 多环境部署：变量 + 分支策略  
- 并行化：parallel、needs、dependencies  

### ArgoCD / GitOps
- Git 为唯一可信源  
- 自动同步，支持回滚  
- 健康状态监控  

### Terraform
- State 管理，远程 backend  
- 模块化设计，环境隔离  
- Plan / Apply / Destroy 流程  

### Ansible
- 幂等性保证重复执行结果一致  
- Inventory 与 roles 管理  
- Playbook 管理多环境部署  

---

## 云原生与 Kubernetes 实践

### K8s CI/CD 实践
- Helm/Kustomize 管理资源  
- GitOps 自动同步  
- readinessProbe/livenessProbe 健康检查  
- 注意 Pod 启动慢、依赖兼容性  

### 部署策略
- 蓝绿、滚动、金丝雀  
- Service/Ingress 切流量  
- 注意数据库兼容、监控、回滚  

### 常见故障排查
- 502/503：Pod 未就绪、Ingress 配置错误  
- 镜像拉取慢：使用私有 registry  
- Secret 配置缺失  

---

## DevSecOps 与自动化安全

- 安全左移：尽早发现漏洞  
- 自动化扫描：SAST、DAST、依赖漏洞  
- CI/CD 集成安全：构建失败阻止部署  
- 注意：Secrets 管理、镜像安全、自动化合规  

---

## SRE 与可观测性

- 核心理念：可靠性指标 SLO/SLA/SLI，自动化运维  
- Metrics：Prometheus  
- Tracing：Jaeger/Zipkin  
- Logging：ELK/Loki  
- Alerting：Prometheus Alertmanager  

---

## 实战练习题与参考答案（示例）

### 练习题 1：设计多环境 CI/CD pipeline
- **要求**：dev/staging/prod 分支，自动化测试，失败回滚  
- **流程**：分支策略 → 构建 → 测试 → Deploy → Verify → Rollback  
- **涉及方面**：Jenkinsfile/GitLab CI YAML、Artifact repo、K8s / Helm  
- **注意点**：环境隔离、失败回滚、监控告警  

### 练习题 2：K8s 金丝雀发布实践
- **要求**：灰度发布，监控指标判断是否全量  
- **流程**：部署少量 Pod → 监控 → 分批扩大 → 全量切换  
- **涉及方面**：Istio VirtualService、Prometheus/Grafana  
- **注意点**：流量比例控制、数据库兼容、回滚机制  

### 练习题 3：Terraform 管理多云基础设施
- **要求**：AWS + GCP，模块化  
- **流程**：模块拆分 → plan → apply → verify  
- **涉及方面**：远程 state 管理、模块化、环境变量  
- **注意点**：资源冲突、状态丢失、权限控制  

---

> ⚠️ 提示：文档中所有 `![图名](路径)` 为占位，可根据实际情况添加架构图或流程图。  

