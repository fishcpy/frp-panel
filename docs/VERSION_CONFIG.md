# 版本配置说明

## 配置文件位置

版本配置文件位于 `etc/version.yaml`

## 配置项说明

```yaml
# 当前系统版本号（必填）
# 客户端和服务器版本低于此版本时会显示升级提示
version: "v0.1.0"

# 是否启用版本升级检查
# true: 启用升级提示功能
# false: 禁用升级提示
enable_upgrade_check: true
```

## 使用场景

### 1. 设置系统版本

配置文件中的 `version` 字段有两个作用：
- 作为服务器的显示版本号
- 作为客户端升级检查的基准版本

```yaml
version: "v1.2.0"
enable_upgrade_check: true
```

当客户端版本号不是 `v1.2.0` 时，会显示升级提示。

### 2. 禁用升级提示

如果不想显示升级提示：

```yaml
version: "v1.0.0"
enable_upgrade_check: false
```

### 3. 开发环境

开发环境通常使用 `dev-build` 版本，不会触发升级提示：

```yaml
version: "dev-build"
enable_upgrade_check: false
```

## 版本比较逻辑

- 当客户端版本号 ≠ 系统配置的版本号时，显示升级提示
- `dev-build` 版本永远不会显示升级提示（无论是客户端还是系统版本）
- 只有当 `enable_upgrade_check` 为 true 时才会检查

## 前端显示

当需要升级时：
- 客户端版本号旁边会显示一个黄色警告图标（动画脉冲效果）
- 点击版本号弹出的详情中会显示升级提示信息
- 提示信息会显示需要升级到的版本号

## 示例

### 开发环境
```yaml
version: "dev-build"
enable_upgrade_check: false
```

### 生产环境 - 统一版本管理
```yaml
version: "v1.0.0"
enable_upgrade_check: true
```

所有客户端和服务器都应该是 `v1.0.0`，不是这个版本的会提示升级。

### 版本升级场景
假设你发布了新版本 `v1.2.0`：

1. 更新服务器的配置文件：
```yaml
version: "v1.2.0"
enable_upgrade_check: true
```

2. 重启服务器，服务器版本变为 `v1.2.0`

3. 所有版本为 `v1.0.0` 或其他版本的客户端会看到升级提示

4. 客户端升级到 `v1.2.0` 后，升级提示消失

## 注意事项

1. **配置文件不存在时**：程序会使用默认配置（version: v0.1.0, enable_upgrade_check: true）
2. **版本号格式**：建议使用语义化版本号（如 v1.2.3）
3. **构建时版本**：如果配置文件中 `version` 为空，将使用构建时通过 `-ldflags` 注入的版本号
4. **配置优先级**：配置文件 version > 构建时注入 > 默认值(dev-build)
5. **版本统一**：所有客户端和服务器应该使用相同的版本号，这样才能准确判断是否需要升级

## API 响应

版本信息通过 `/api/v1/user/platform_info` 接口返回，临时存储在 `githubProxyUrl` 字段中（格式：`原值|服务器版本|系统版本|是否启用检查`）。

在后续版本中会迁移到专用字段。
