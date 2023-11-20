## 开发

```bash
# 克隆项目
git clone https://gitee.com/OptimisticDevelopers/Ruoyi-Go.git

# 进入项目目录
cd frontend/ruoyi-ui

# 安装依赖
npm install

# 建议不要直接使用 cnpm 安装依赖，会有各种诡异的 bug。可以通过如下操作解决 npm 下载速度慢的问题
npm install --registry=https://registry.npmmirror.com

# 启动服务
npm run dev
```

浏览器访问 http://localhost:80

## 发布

```bash
# 构建测试环境
npm run build:stage

# 构建生产环境
npm run build:prod
```



```xml
>node -v
v18.13.0 
启动报错，设置下面进行运行
>set NODE_OPTIONS=--openssl-legacy-provider

>npm run dev
```

修改了后台访问路径 
https://blog.csdn.net/u012069313/article/details/129530029
