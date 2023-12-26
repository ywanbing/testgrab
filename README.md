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
# 生成文件的类型 目前只支持 markdown
gen_file_type: markdown

```
运行 `testGrab.exe` 即可

最后生成文件在 `./docs` 目录下

代码比较简单，可以直接拉下来自己修改，或者提issue，我会尽快回复

