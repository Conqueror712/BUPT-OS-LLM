# BUPT-OS-LLM
> BUPT 2023 autumn operating system course big assignment - LLM local deployment project
>

## 项目描述

这个项目是基于百川API，使用Go语言和Gin框架封装了一个自定义的LLM工具，并提供了一个基于Bootstrap的可交互的Web应用程序。

通过这个工具，用户可以输入文本或问题，并获得LLM生成的语言模型输出。用户可以与LLM进行对话、提问、获得文本摘要等。

> 为什么选择这些？

## 项目结构

```
.
├── client.py		    # Web Server Python Version 
├── client.go		    # Web Server Golang Version
├── templates
│   └── index.html		# 前端界面的HTML模板文件，包含用户输入和LLM输出的显示
│   └── script.js		# 处理一些响应的文件
└── static
    └── style.css		# 前端界面的CSS样式文件，用于美化界面
```

## 安装和配置

1. 克隆项目到本地：

```
git clone git@github.com:Conqueror712/BUPT-OS-LLM.git
```

2. 进入项目目录：

```
cd BUPT-OS-LLM
```

3. 配置LLM API密钥：

在`没想好`文件中，将以下代码行中的`YOUR_API_KEY`替换为你的LLM API密钥：

> 获取密钥可以前往：[URL](https://platform.baichuan-ai.com/console/apikey)  点击**创建API Key**即可获取
> 

```
const apiKey = "YOUR_API_KEY"
```

## 运行应用程序

1. 在命令行中执行以下命令启动应用程序：

```
没想好
```

2. 在浏览器中访问http://localhost:13240，即可打开应用程序的前端界面。

## 使用示例

1. 在输入框中输入你的问题或文本。
2. 点击**"发送"**按钮，应用程序将发送请求到LLM API并等待响应。
3. LLM生成的文本将显示在界面的输出框中。
4. 可以重复上述步骤，与LLM进行对话、提问等。

## 贡献

如果你发现了bug，或者有任何改进意见或建议，请提交issue。欢迎参与项目的开发和贡献！

## 许可证

这个项目基于 MIT license 开源。
