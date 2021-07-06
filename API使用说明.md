# 国密GM/T Go API使用说明

## Go包安装

```bash
go get -u github.com/kingstenzzz/SM-improvement
```

## SM3密码杂凑算法 - SM3 cryptographic hash algorithm
- 遵循的SM3标准号为： GM/T 0004-2012
- g package：`github.com/kingstenzzz/SM-improvement/sm3` 
- `type SM3 struct` 是原生接口hash.Hash的一个实现

### 代码示例

```Go
    data := "test"
    h := sm3.New()
    h.Write([]byte(data))
    sum := h.Sum(nil)
    fmt.Printf("digest value is: %x\n",sum)
```


### 具体功能测试代码参考
```Go
github.com/kingstenzzz/SM-improvement/sm3/sm3_test.go  //sm3算法
```
