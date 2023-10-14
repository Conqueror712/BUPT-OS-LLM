# BUPT-OS-LLM
BUPT 2023 autumn operating system course big assignment - LLM local deployment project

# 项目名称

## 项目描述

这个项目是基于百川API，使用Go语言和Gin框架封装了一个自定义的LLM工具，并提供了一个基于Bootstrap的可交互的Web应用程序。

通过这个工具，用户可以输入文本或问题，并获得LLM生成的语言模型输出。用户可以与LLM进行对话、提问、获得文本摘要等。

> 为什么选择这些？

- 选择百川的理由：
- 选择Go语言的理由：
- 选择Bootstrap的理由：

## 项目结构

```
.
├── go.mod				# 应用程序的入口文件，包含主要的启动逻辑
├── go.sum
├── main.go
├── handlers
│   └── api.go			# 实现了与LLM API的交互逻辑，包括发送请求和解析响应
├── models
│   └── request.go		# 定义了请求模型，用于将用户输入转换为LLM API所需的格式
├── templates
│   └── index.html		# 前端界面的HTML模板文件，包含用户输入和LLM输出的显示
└── static
    └── style.css		# 前端界面的CSS样式文件，用于美化界面
```

## 安装和配置

1. 克隆项目到本地：

```
git clone xxx
```

2. 进入项目目录：

```
cd xxx
```

3. 安装依赖：

```
go mod download
```

4. 配置LLM API密钥：

在handlers/api.go文件中，将以下代码行中的YOUR_API_KEY替换为你的LLM API密钥：

> 获取密钥可以前往：URL

```
const apiKey = "YOUR_API_KEY"
```

## 运行应用程序

1. 在命令行中执行以下命令启动应用程序：

```
go run main.go
```

2. 在浏览器中访问http://localhost:13240，即可打开应用程序的前端界面。

## 使用示例

1. 在输入框中输入你的问题或文本。
2. 点击"提交"按钮，应用程序将发送请求到LLM API并等待响应。
3. LLM生成的文本将显示在界面的输出框中。
4. 可以重复上述步骤，与LLM进行对话、提问等。

## 贡献

如果你发现了bug，或者有任何改进意见或建议，请提交issue。欢迎参与项目的开发和贡献！

## 许可证

这个项目基于 MIT license 开源。
