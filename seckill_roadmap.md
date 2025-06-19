
# 🛠️ Go 秒杀系统增强功能开发路线图

本文件用于跟踪和管理 `seckill` 项目的增强功能开发进度。

---

## ✅ 已完成的基础功能

- [x] 秒杀接口（Gin）
- [x] Redis + Lua 脚本控制库存
- [x] 同用户只抢一次
- [x] 项目目录结构搭建完毕

---

## 🔧 增强功能清单（建议优先级）

### 🥇 一、接口限流

- [x] 防止恶意用户/IP 高频请求
- [x] 使用 Redis 实现限流（滑动窗口/固定窗口）
- [x] 添加 Gin 中间件拦截过载请求
- 📂 建议文件：`middleware/rate_limit.go`

---

### 🥇 二、异步下单（削峰填谷）

- [ ] 将秒杀请求写入 Kafka（或 Redis Stream）
- [ ] 后台消费者异步下单 + 写入数据库
- 📂 模块建议：
  - `mq/producer.go`
  - `mq/consumer.go`
  - `service/order.go`

---

### 🥈 三、登录鉴权

- [ ] 用户通过 JWT 登录认证
- [ ] 秒杀接口中提取 token 验证 user_id
- 📂 模块建议：
  - `auth/jwt.go`
  - `middleware/auth.go`

---

### 🥈 四、性能监控与日志

- [ ] 打印或记录接口响应耗时
- [ ] 可集成 Prometheus + Grafana
- 📂 可放置：`middleware/metrics.go`

---

### 🥉 五、结果持久化

- [ ] 将成功秒杀用户记录持久化到数据库
- [ ] MySQL 表设计：`orders`（user_id, product_id, create_time）
- 📂 模块建议：
  - `model/order.go`
  - `dao/order.go`

---

### 🥉 六、IP 黑名单机制

- [ ] 使用 Redis Set 管理黑名单
- [ ] 接口请求时校验 IP 是否在封禁列表中
- 📂 模块建议：
  - `middleware/ip_block.go`

---

### 🥉 七、排行榜系统

- [ ] 使用 Redis ZSet 记录用户抢购时间戳
- [ ] 接口返回前 N 名（如 `/rank`）
- 📂 建议扩展：`handler/rank.go`

---

### 🥉 八、后台管理功能

- [ ] 简单库存调整 API
- [ ] 秒杀结果查询 API
- [ ] 可选 Vue 页面或 Postman 管理

---

## 🧱 建议目录结构补充

```bash
seckill/
├── middleware/
│   ├── rate_limit.go
│   ├── auth.go
│   └── metrics.go
├── mq/
│   ├── producer.go
│   └── consumer.go
├── model/
│   └── order.go
├── service/
│   └── order.go
├── auth/
│   └── jwt.go
├── dao/
│   └── order.go
```

---

_维护人：ChatGPT 生成于 2025，基于 Gitee 仓库 https://gitee.com/tenken/seckill_
