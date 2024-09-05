<p align="center">
    <img src="misc/avatar.png" width="200" height="200" alt="paper-airplane" style="border-radius: 40px 40px 40px 40px;">
</p>

<div align="center">

# paper-airplane

_✨ 基于 QQ 官方 SDK 开发的 QQBot ✨_

</div>

---

十分简陋喵。正在测试喵。

## 食用

1. 将 `misc/config.example.yaml` 改名为 `config.yaml` 并填入相关信息，放入根目录
2. 执行 `make gen` 生成 DAO 代码
3. 执行 `make run` 本地开始调试，若需 release 构建则执行 `Makefile` 中相关构建目标，或自行指定交叉编译环境

## 沙箱环境

默认情况下，paper-airplane 会运行在沙箱环境中，除非将当前主机 IP 加入机器人管理端，不然无法再非沙箱环境使用

若 `export BOT_MODE=release` 则脱离沙箱环境

e.g.

```bash
BOT_MODE=release ./paper-airplane
```

## Features

- [x] SparkLite AI 接入
