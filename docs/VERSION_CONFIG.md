# 版本配置说明

## 配置文件位置

版本配置文件位于 `etc/version.yaml`

## 配置项说明

```yaml
# 当前版本号（可自定义）
# 如果设置了此值，将覆盖构建时注入的版本号
# 如果为空，将使用构建时通过 -ldflags 注入的版本号
version: ""

# 是否启用版本升级检查
# true: 启用升级提示功能
# false: 禁用升级提示
enable_upgrade_check: true

# 建议的最新版本号（用于升级提示）
# 当客户端版本低于此版本时，会在前端显示升级提示
latest_version: "v0.1.0"
```

## 使用场景

### 1. 自定义版本号

如果你想在不重新编译的情况下修改服务器显示的版本号：

```yaml
version: "v1.0.0-custom"
enable_upgrade_check: true
latest_version: "v1.0.0-custom"
```

### 2. 设置升级提示

当你发布新版本后，想提示用户升级：

```yaml
version: ""  # 保持为空，使用构建版本
enable_upgrade_check: true
latest_version: "v1.2.0"  # 设置为最新版本号
```

### 3. 禁用升级提示

如果不想显示升级提示：

```yaml
version: ""
enable_upgrade_check: false
latest_version: "v0.1.0"
```

## 版本比较逻辑

- 客户端版本与 `latest_version` 不同时，会显示升级提示
- `dev-build` 版本不会显示升级提示
- 只有当 `enable_upgrade_check` 为 true 时才会检查

## 前端显示

当需要升级时：
- 客户端版本号旁边会显示一个黄色警告图标（动画脉冲效果）
- 点击版本号弹出的详情中会显示升级提示信息
- 提示信息会显示建议升级到的版本号

## 示例

### 开发环境
```yaml
version: "dev-build"
enable_upgrade_check: false
latest_version: "v0.1.0"
```

### 生产环境
```yaml
version: ""  # 使用 git tag 构建的版本
enable_upgrade_check: true
latest_version: "v1.0.0"
```

### 自托管环境
```yaml
version: "v1.0.0-selfhost"
enable_upgrade_check: false  # 自托管通常不需要升级提示
latest_version: "v1.0.0"
```

## 注意事项

1. **配置文件不存在时**：程序会使用默认配置（启用升级检查，latest_version 为 v0.1.0）
2. **版本号格式**：建议使用语义化版本号（如 v1.2.3）
3. **构建时版本**：通过 `build.sh` 脚本使用 `-ldflags` 注入，基于 git tag
4. **配置优先级**：配置文件 > 构建时注入 > 默认值(dev-build)

## API 响应

版本信息通过 `/api/v1/user/platform_info` 接口返回，临时存储在 `githubProxyUrl` 字段中（格式：`原值|服务器版本|最新版本|是否启用检查`）。

在后续版本中会迁移到专用字段。
