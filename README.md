# testGrab

## 使用方法
修改 `config_env.yaml` 到 `config.yaml`

填写自己需要的配置
```yaml
#用户名
name: {准考证号}
#密码
pwd: {身份证号码}
#批次ID
batch_id: {登录首页网址上面的batch_id}
#每科循环次数
loop_num: 10
#是否简单模式（简单模式下：选择题只保留正确的选项，默认不是简单模式）
simple: false
# 生成文件的类型 目前支持 markdown html pdf
gen_file_type: markdown

```

需要注意：生成pdf需要`wkhtmltopdf`，我已经打包在tools目录下。

执行前的应该在 `testgrab.exe` 目录下有 `./tools/wkhtmltopdf.exe`，如果没有请自行下载。

运行 `testGrab.exe` 即可

最后生成文件在 `./docs` 目录下

代码比较简单，可以直接拉下来自己修改，或者提issue，我会尽快回复

## 特性
1. 可以选择只保留选择题的答案，或者保留所有选项
2. 提取多次考试的答案，并自动去重，生成一个文件
3. 支持带有图片的题目和答案
4. 生成的`md`文件可以很方便转为pdf
5. 现在支持直接生成pdf和html格式
6. 生成pdf需要`wkhtmltopdf`，我已经打包在tools目录下，可以自行下载

支持的课程（本人考试科目），稳定输出
- 高等数学（工专）
- 工程经济
- 计算机网络技术（由于题目和md格式有冲突，自己修改一下文件）
- 数据库及其应用
- 动态网络编程基础

其他课程未测试，欢迎提issue
- 英语不支持

