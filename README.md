# log

log interface for go module

https://github.com/go-logr/logr

- 1: 支持项目级别替换日志

- 2: 支持模块级别替换日志

```text
三方框架含有自有日志模块，通过SetLogger修改公式自有模块的日志实现，整合在一起。
```

### 优点：

- 基于接口实现
- 模块级别输出控制

### 缺点：

- 基于 log4go 输出的文件和行号不准

### 快速使用

- 参考 Test&&Example

### Todo

- [x] 日志添加模块名字
- [] 适配其他日志库
  - [x] Kratos
  - [x] log4go
