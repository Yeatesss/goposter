# goposter
海报生成器
### 目前支持
* 图片
* 文字
* 横线

### 准备工作
* 1. 拉取代码包进项目 go get github.com/Yeate/goposter@v0.1.1
* 2. 创建海报包的配置文件,配置文件支持yaml 路径例如：/home/admin/abc.yaml 

```
#配置文件
disk: "local" #海报存放方式 local:本地磁盘 oss:阿里云OSS

oss:          #阿里云OSS配置项
  endpoint:
  access_key_id:
  access_key_secret:
  bucket_name:
  
img_tmp_dir: tmp/image/ #临时文件存放地址

```

* 3. 代码入口中进行配置文件初始化

```
cfg := config.NewConfig("/home/admin", "abc")
_ = cfg.InitConfig()
```

### 生成海报代码样例


```
//初始化配置文件
cfg := config.NewConfig("/Users/abc/coding/aaa", "config")
_ = cfg.InitConfig()
//创建海报实例
poster := module.NewPoster()
//加载字体文件
pinfang, _ := os.Open("/Users/aaa/coding/bbb/pingfangsr.ttf")
//设置海报地址
poster.Background = "http://img.aiimg.com/uploads/allimg/180707/1-1PFG64119.jpg"
//设置最终海报输出路径
poster.SavePath = "images/"
poster.SaveName = "test_poster.png"
//海报额外附着图
poster.Images = append(poster.Images, module.Image{Url: "http://n.sinaimg.cn/sinacn16/580/w690h690/20180414/0939-fzcyxmu4864171.jpg", X: 327, Y: 175, Width: 227, Height: 227，CircleClip:false})
//海报文字
text := module.Text{Color: "#080808", Text: "测试海报", X: 334, Y: 653, FontSize: 50}
//设置字体
text.SetFont(pinfang)
poster.Texts = append(poster.Texts, text)
poster.Lines = append(poster.Lines, module.Line{StartX: 0, StartY: 0, EndX: 199, EndY: 50, Width: 2, Color: "#000000"})

err := poster.Draw()
```
![](https://gitee.com/ye3245/oss/raw/master/uPic/test_poster.jpg)

###结构体解析
####海报

| 字段名             | 类型      | 备注                     |
|-----------------|---------|------------------------|
| Width           | float64 | 宽（当Background不为空的时候失效） |
| Height          | float64 | 高（当Background不为空的时候失效） |
| BackgroundColor | string  | 背景色（存在背景图时候会被覆盖）       |
| Background      | string  | 背景图                    |
| Texts           | []Text  | 插入的文字集合                |
| Images          | []Image | 插入的图片集合                |
| SavePath        | string  | 海报保存路径                 |
| SaveName        | string  | 保存的文件名                 |

####横线

| 字段名             | 类型      | 备注                     |
|-----------------|---------|------------------------|
| StartX           | float64 | 起始横坐标 |
| StartY          | float64 | 起始纵坐标 |
| EndX | float64  | 结束横坐标       |
| EndY      | float64  | 结束纵坐标                    |
| Width           | int  | 线的粗细                |
| Color          | string | 线的颜色                |

####文字
| 字段名             | 类型      | 备注                     |
|-----------------|---------|------------------------|
| X           |int | 横坐标 左上原点 |
| Y          |int | 纵坐标 左上原点 |
| Text |string  | 文字       |
| Width      |float64  | 宽度                   |
| FontSize           |int  | 字体大小                |
| Color          |string | 字体颜色               |
| LineHeight        |int  | 行高                 |
| TextAlign        |string  | 文字对齐方式                 |


####图片
| 字段名             | 类型      | 备注                     |
|-----------------|---------|------------------------|
| X           |int | 横坐标 左上原点 |
| Y          |int | 纵坐标 左上原点 |
| Url |string  | 图片网络地址       |
| Width      |float64  | 宽度，不填写原始大小                   |
| Height           |int  | 高度    ，不填写原始大小           |
| CircleClip           |bool  |图片切成圆形 true:切 false:不切           |





