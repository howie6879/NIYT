## Read the novel in your terminal - NIYT

> 这几天在看`go`，所以这是个练手项目，在终端里看小说

这是个无聊的项目
这真是个无聊的项目

### 1.How to use?

``` shell

git clone https://github.com/howie6879/NIYT.git
cd NIYT
go run main.go

```

首先输入小说名称，调用第三方搜索，利用本地`json`文件解析，终端就会显示小说源：

![demo01](./images/demo01.jpg)

利用`get 0`进入第一个源，获取最新章节：
![demo02](./images/demo02.jpg)

选择章节进行阅读，如阅读最新章节，`get 0`:
![demo03](./images/demo03.jpg)

`show` 会以表格的形式再次展示当前源：
![demo04](./images/demo04.jpg)

**享受阅读吧**