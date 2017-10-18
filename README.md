# selpg-projec

项目的具体要求在该网址中：https://www.ibm.com/developerworks/cn/linux/shell/clutil/index.html
该项目主要分为两个阶段:
1.获取命令行参数
2.分析命令行参数，实现功能


1.获取命令行参数
  使用flag包来获取参数，会使得参数形式为“-s=1”，与原题中的“-s1”有一点出入，但影响不大，就这样了。。。
  将参数存储在一个结构体args中，如下：
  
  type args struct {
  sPage int
  ePage int
  pageLen int
  pageType bool
  inputFile string
  dest string
}

2.分析参数实现功能
  由于文件的格式有两种（定长、分页符），所以使用type_1和type_2两个函数分别处理两种情况，两个函数其实结构相似，主要是在循环读页时，一个以定长判断一页然后循环读下一页，另一个直到读到分页符之前都不终止这一轮循环
  
具体内容见代码和测试截图（因为没有打印机，所以某些功能无法测试）
