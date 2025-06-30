# gpkg/pkg

## design principles

使用`NewXXX`的方式创建对象, 创建对象时使用`WithXXX`的方式设置参数(options模式), 如果想直接创建对象, 直接使用`&struct{...}`创建

gpkg 应该使用 面向对象的方法设计, 还是使用面向接口的方法设计, 暂时不确定


## inspiration

- [spf13/cobra](https://github.com/spf13/cobra)
- [spf13/viper](https://github.com/spf13/viper)
- [spf13/pflag](https://github.com/spf13/pflag)
- [bep/simplecobra](https://github.com/bep/simplecobra)
- [gohugoio/hugo](https://github.com/gohugoio/hugo)
- [uber-go/zap](https://github.com/uber-go/zap)
- [jianghushinian/gokit](https://github.com/jianghushinian/gokit/)
- [marmotedu/iam](https://github.com/marmotedu/iam)
- [onexstack](https://github.com/onexstack/onexstack)
- [gorm.io/gorm](https://github.com/go-gorm/gorm.)