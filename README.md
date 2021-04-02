# go 工具集
未完待续...

## number包
```go
import "github.com/wangxudong123/assist/number"

```
<!-- Todo -->


## str包
```go
import "github.com/wangxudong123/assist/str"

```
- 拼接字符串 `func Join(arg ...interface{}) (string, error)`

    ```go
    str.Join(p_int,p_int64,p_string)
    ```
#
## time包
```go
import "github.com/wangxudong123/assist/time"
```
#### 时间转换
- 创建当前时间对象`func NewUnixNow() *theTime`

    ```go
    NewUnixNow()
    ```
- 创建时间戳对象`NewUnix(unix int64) *theTime`

    ```go
    NewUnix(1617264318)
    ``` 
  
- 创建格式化时间对象`NewFormat(t string) (*theTime, error)`
    ```go
    NewFormat("2021-04-01 12:00:00")
    ```

- 创建iso8601时间对象`NewISO8601(iso8601 string) (*theTime, error)`
    ```go
    NewIso8601("2019-09-11T01:54:23+00:00")
    ```
- 格式化时间 `func (t *theTime) ToFormat() string`

- 转时间戳 `func (t *theTime) ToUnix() int64`

- 转iso8601时间表达式 `func (t *theTime) ToIso8601() string`

- 通过时间模板格式化时间 `func (t *theTime) ToFormatForTpl(tpl string) string`
