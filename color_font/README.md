# 输出颜色字体
## 命令行
```
fmt.Sprintf("%c[%d;%d;%dm%s%c[0m", 0x1B, conf, bg, text, msg, 0x1B)
```
## 效果
![](image/1.png)