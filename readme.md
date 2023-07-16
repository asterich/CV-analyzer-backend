ubuntu下部署：

```bash
sudo apt-get install poppler-utils
sudo apt-get install libreoffice libreoffice-l10n-zh-cn libreoffice-help-zh-cn
```

```bash
python3 -m pip install -r requirements.txt
```

```bash
go run src/main.go
```

*关于config*

具体配置文件在config/config.ini中，其中：

```ini
[server] # 关于服务器部署
AppMode = debug # debug模式和release模式
HttpPort = 8090 # 服务器端口
JwtKey = 324234234 # jwt密钥
MaxLoginTime = 240 # 登录过期时间
[database] # 数据库配置
Db = sqlite # 数据库类型
DbName = softbei # 数据库名称
DbPath = ./data/softbei_test.db # 数据库路径
[upload] # 文件上传配置
UploadPath = ./static/ # 文件路径
[redis] # redis配置
RedisAddr = 127.0.0.1:6379 # redis地址
RedisPassword = # redis密码
RedisDB = 0 # redis DB
```
