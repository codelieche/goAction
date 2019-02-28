
## 构建命令

```bash
cd entry
GOOS=linux GOARCH=amd64 go build ./main.go&& echo "`date +%"F %T"`: 构建成功" || echo "`date +%"F %T"`: 构建失败！！！"
```
