# pictureProxy

goframe project

go run main.go 

直接运行 

注：大陆地区需开启外网http代理 

localhost:8199/picture/normal 随机涩图 

localhost:8199/picture/r18 随机r18涩图 

1C0.5G环境无压力，内存大的机器可以修改[app/api/hello.go的37行和38行](https://github.com/yangge2333/r18PictureProxy/blob/main/app/api/hello.go#L37)的通道缓存，开的越大，图片缓存的越多

随机涩图api文档：https://api.lolicon.app/#/setu

-------

## release包的食用方法

准备带有systemctl和make的Linux环境

解压release包

在项目路径下执行 make install

更新执行 make update
