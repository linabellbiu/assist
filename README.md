# go 工具集
未完待续...


## number包

```go
import "github.com/wangxudong123/assist/number"
```
#### 数字类型转换

- string转float64`func StringToFloat64(value string) (f float64)`

- int64转float64`func Int64ToFloat64(value int64) (f float64)`

- string转float64`func StringToInt64(value string) (i int64)`

- string转int`func StringToInt(value string) (i int)`

- float64转int64`func Float64ToInt64(value float64) (i int64)`

- float64转string`func Float64ToString(value float64) (s string)`

- int64转string`func Int64ToString(value int64) (s string)`

- int64转float64`func IntToFloat64(value int) (f float64)`

#### 处理小数位 
- 保留小数位`func RoundFloat64(value float64, exp int) (num float64)`
  > exp 保留的位数 


#### 浮点型数字比较
- 左 = 右`func (c Cmp) Equal(l, r float64) bool`
  ```go
  res := Cpm.Equal(3, 3)
  t.Log(res) //true
  ```
- 左 > 右`func (c Cmp) Greater(l, r float64) bool`

- 左 > 右`func (c Cmp) Smaller(l, r float64) bool`

- 左 >= 右`func (c Cmp) GreaterOrEqual(l, r float64) bool`

- 左 <= 右`func (c Cmp) SmallerOrEqual(l, r float64) bool`

#### 排序
- 初始化数组`func NewNumbers(n ...interface{}) number`
  ```go
  //数组可以包含float64,float32,int,int64,int32
  NewNumbers(3, 44, 3453, float64(4.999999999), float64(2.0000001))
  ```
- 找出最大值`func (n number) Max() float64`
  
  ```go
  max := NewNumbers(3, 44, 3453, float64(4.999999999), float64(2.0000001)).Max()
  t.Log(max) //3453
  ```
- 找出最小值`func (n number) Min() float64`

  ```go
  min := NewNumbers(3, 44, 3453, float64(4.999999999), float64(2.0000001)).Min()
  t.Log(min)//2.0000001
  ```

 - 从大到小排列 `func (n number) OrderDesc() []interface{}`
   > 返回数组的元素类型是原始的类型，没有做过转换
  ```go
  desc := NewNumbers(3, 44, 3453, float64(4.999999999), float64(2.0000001)).OrderDesc()
  t.Log(desc) //[3453 44 4.999999999 3 2.0000001]
  ```

 - 从小到大排列`func (n number) OrderAsc() []interface{}`
  > 返回数组的元素类型是原始的类型，没有做过转换
  ```go
  asc := NewNumbers(3, 44, 3453, float64(4.999999999), float64(2.0000001)).OrderAsc()
  t.Log(asc)//[2.0000001 3 4.999999999 44 3453]
  ```
#
## str包

```go
import "github.com/wangxudong123/assist/str"
```
#### 字符串处理
- 拼接字符串`func Join(arg ...interface{}) (string, error)`
  
  ```go
  Join(p_int,p_int64,p_string)
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
