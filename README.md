# SORM

## 简介✨✨✨✨✨✨

​		sorm是一款基于go数据库操作包的orm框架，此项目并非完全的orm框架项目。而更像是一个开发日记，记录每次框架更新迭代的过程，版本号以v开发，例如 v0,v1,v2。

使用日记方式进行框架开发，旨在 😁

1. 加强对数据库系统了解，go语言与数据库交互     
2. 向大家记录开发框架全过程，供大家使用
3. 为大家提供pr机会，本人目前的技术水准应该还有待提高，欢迎大家指出🤠
4. 体验处理issue，与pr
5. 在实际的框架开发中，增强自己对go语言的理解

## 当前支持功能

1. 数据库的连接
2. SQL语句的执行
3. 结构体解析为schema
4. 根据结构体创建数据表
5. 表的删除与是否存在检查

6. CRUD语句的生成与实现
7. 关键词之间的链式调用
8. 钩子的实现
9. 事务实现

## 版本功能更新记录🚧🚧

### v1: 2022.12.25 🥽

> 1. ORM中logger开发，实现logger分级以及彩色打印
> 1. session会话实现，sql语句的生成与执行，单条记录查询，多条记录查询
> 1. 数据库连接

### v2: 2022.12.26 🐱‍🏍

> 1. 不同类型数据库的适配
> 1. mysql数据类型与go数据类型映射
> 1. 数据库表模式的定义
> 1. 数据库表相关操作，创建，删除，是否存在

### v3: 2022.12.29 🍿

> 1.不同关键词字句的生成
>
> 2.CURD对应SQL语句的生成与执行
>
> 3.关键词之间实现链式调用

### v4: 2023.01.01🙆‍♂️

> 1.钩子的实现
>
> 2.事务调用

## 功能实例

### 1.根据结构体创建数据表

```go
package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"sorm"
)

type User struct {
	Name string `sorm:"PRIMARY KEY"`
	Age  int
}

func main() {
	dsn := "root:root@tcp(127.0.0.0:3306)/sorm"
	engine, _ := sorm.NewEngine("mysql", dsn)
	defer engine.Close()

	s := engine.NewSession()
	err := s.Model(&User{}).CreateTable()
	if err != nil {
		panic(err)
	}
	fmt.Println(s.RefTable().Name, "创建成功")
	fmt.Printf("字段: %s", s.RefTable().FieldNames)
}
```

![image-20221228165543750](https://xingqiu-tuchuang-1256524210.cos.ap-shanghai.myqcloud.com/1770/image-20221228165543750.png)

![image-20221228165613917](https://xingqiu-tuchuang-1256524210.cos.ap-shanghai.myqcloud.com/1770/image-20221228165613917.png)

### 2.向表中插入数据INSERT

```go
package main

import (
   "fmt"
   _ "github.com/go-sql-driver/mysql"
   "sorm"
)

type User struct {
   Name string `sorm:"PRIMARY KEY"`
   Age  int
}

func main() {
	dsn := "root:root@tcp(127.0.0.0:3306)/sorm"
   engine, _ := sorm.NewEngine("mysql", dsn)
   defer engine.Close()

   s := engine.NewSession().Model(&User{})
   user1 := User{
      Name: "张三",
      Age:  18,
   }
   user2 := User{
      Name: "李四",
      Age:  19,
   }
   user3 := User{
      Name: "王五",
      Age:  18,
   }
   result, err := s.Insert(user1, user2, user3)
   if err != nil {
      panic(err)
   }
   fmt.Println("插入数据条数:", result)
}
```

![image-20221228170404636](https://xingqiu-tuchuang-1256524210.cos.ap-shanghai.myqcloud.com/1770/image-20221228170404636.png)

![image-20221228170450116](https://xingqiu-tuchuang-1256524210.cos.ap-shanghai.myqcloud.com/1770/image-20221228170450116.png)

### 3.条件搜索FIND

```go
package main

import (
   "fmt"
   _ "github.com/go-sql-driver/mysql"
   "sorm"
)

type User struct {
   Name string `sorm:"PRIMARY KEY"`
   Age  int
}

func main() {
	dsn := "root:root@tcp(127.0.0.0:3306)/sorm"
   engine, _ := sorm.NewEngine("mysql", dsn)
   defer engine.Close()
   var users []User
   session := engine.NewSession().Model(&User{})
   fmt.Println("-----where限制")
   session.Where("Age = ?", "18").Find(&users)
   for _, user := range users {
      fmt.Printf("%v\n", user)
   }
   users = []User{}
   fmt.Println("-----order限制")
   session.Where("Age = ?", "18").Order("Name desc").Find(&users)
   for _, user := range users {
      fmt.Printf("%v\n", user)
   }
   users = []User{}

   fmt.Println("-----limit限制")
   session.Where("Age = ?", "18").Order("Name desc").Limit(1).Find(&users)
   for _, user := range users {
      fmt.Printf("%v\n", user)
   }
   users = []User{}

   fmt.Println("-----offset限制")
   session.Where("Age = ?", "18").Order("Name desc").Limit(1).Offset(1).Find(&users)
   for _, user := range users {
      fmt.Printf("%v\n", user)
   }

}
```

![image-20221228171622769](https://xingqiu-tuchuang-1256524210.cos.ap-shanghai.myqcloud.com/1770/image-20221228171622769.png)

### 4.INSERT语句

```go
func TestInsertSQLBuild(t *testing.T) {
   user1 := User{
      Name: "张三",
      Age:  18,
   }
   user2 := User{
      Name: "李四",
      Age:  19,
   }
   session := Engine.NewSession().Model(&User{})
   insert, err := session.Insert(user1, user2)
   if err != nil {
      t.Error(err)
   }
   fmt.Println(insert)
}
```

![image-20221228141529106](https://xingqiu-tuchuang-1256524210.cos.ap-shanghai.myqcloud.com/1770/image-20221228141529106.png)

![image-20221228142334524](https://xingqiu-tuchuang-1256524210.cos.ap-shanghai.myqcloud.com/1770/image-20221228142334524.png)

### 5.FIND语句

```go
func TestFindSQLBuild(t *testing.T) {
   session := Engine.NewSession().Model(&User{})
   var users []User
   err := session.Where("Age =?", 18).Order("Age desc").Limit(2).Offset(0).Find(&users)
   if err != nil {
      t.Error(err)
   }
   for _, user := range users {
      fmt.Printf("%v\n", user)
   }
}
```

### 6.UPDATE语句

```go
func TestUpdateSQLBuild(t *testing.T) {
   session := Engine.NewSession().Model(&User{})
   result, err := session.Where("Age =?", 20).Update("Name", "ccccc")
   if err != nil {
      t.Error(err)
   }
   fmt.Println(result)
}
```

![image-20221229112233342](https://xingqiu-tuchuang-1256524210.cos.ap-shanghai.myqcloud.com/1770/image-20221229112233342.png)

![image-20221229112256212](https://xingqiu-tuchuang-1256524210.cos.ap-shanghai.myqcloud.com/1770/image-20221229112256212.png)

### 7.DELETE语句

```go
func TestDeleteSQLBuild(t *testing.T) {
   session := Engine.NewSession().Model(&User{})
   result, err := session.Where("Age = ?", 20).Delete()
   if err != nil {
      t.Error(err)
   }
   fmt.Println(result)
}
```

![image-20221229112803620](https://xingqiu-tuchuang-1256524210.cos.ap-shanghai.myqcloud.com/1770/image-20221229112803620.png)

### 8.钩子的使用

```go
package test

import (
   "fmt"
   "sorm/log"
   "sorm/session"
   "testing"
)

type Teacher struct {
   Name string `sorm:"PRIMARY KEY"`
   Age  int
}

func (teacher *Teacher) BeforeInsert(s *session.Session) error {
   log.Info("BEFORE INSERT")
   return nil
}

func (teacher *Teacher) AfterInsert(s *session.Session) error {
   log.Info("AFTER INSERT")
   return nil
}

func TestInsertHooks(t *testing.T) {
   s := Engine.NewSession().Model(&Teacher{})
   err := s.DropTable()
   if err != nil {
      t.Error(err)
   }
   err = s.CreateTable()
   if err != nil {
      t.Error(err)
   }
   result, err := s.Insert(&Teacher{
      Name: "张三",
      Age:  10,
   }, &Teacher{
      Name: "李四",
      Age:  20,
   })
   if err != nil {
      t.Error(err)
   }
   fmt.Println(result)
}
```

​		Teacher结构体绑定的钩子很简单，就是打印两条语句，如果想干点其它操作的话可以通过传入的session来完成	

![image-20230101102425898](https://xingqiu-tuchuang-1256524210.cos.ap-shanghai.myqcloud.com/1770/image-20230101102425898.png)

### 9.事务使用

### 

```go
import (
   "fmt"
   _ "github.com/go-sql-driver/mysql"
   "sorm"
   "sorm/session"
)

type User struct {
   Name string `sorm:"PRIMARY KEY"`
   Age  int
}

func main() {
   dsn := "root:root@tcp(127.0.0.1:3306)/sorm"
   engine, _ := sorm.NewEngine("mysql", dsn)
   defer engine.Close()
   user := User{
      Name: "ccc",
      Age:  1,
   }
   s := engine.NewSession().Model(&User{})
   fmt.Println("事务第一次调用")
   // 1.第一次插入user
   _, err := s.Transaction(func(s *session.Session) (result interface{}, err error) {
      _, err = s.Insert(&user)
      return
   })
   fmt.Println("事务第二次调用")
   // 2.第二次插入user
   _, err = s.Transaction(func(s *session.Session) (result interface{}, err error) {
      _, err = s.Insert(&user)
      return
   })
   if err != nil {
      fmt.Println(err)
   }
}
```

![image-20230101100900172](https://xingqiu-tuchuang-1256524210.cos.ap-shanghai.myqcloud.com/1770/image-20230101100900172.png)

## 开发日记

### <a href="https://www.yuque.com/c_pluto/rsz2ys/gnydi1edhi6k7af1?singleDoc# 《Go_SORM开发日记(一)—SQL生成》">🚗Go_SORM开发日记(一)—SQL生成</a>

### <a href="https://www.yuque.com/c_pluto/rsz2ys/qx33m4iuyzeyl5no?singleDoc# 《Go_SORM开发日记(二)—不同数据库之间差异屏蔽》">🚓Go_SORM开发日记(二)—结构体解析成为表</a>

### <a href="https://www.yuque.com/c_pluto/rsz2ys/sp28lohra8yxhr45?singleDoc# 《3.Go_SORM开发日记(三)—不同关键词字句的生成》">🚕Go_SORM开发日记(三)—不同关键词字句的生成</a>

### <a href="https://www.yuque.com/c_pluto/rsz2ys/yvvs2lvgqk3h9n2h?singleDoc# 《4.Go_SROM开发日记(四)—钩子与事务》">🚌Go_SORM开发日记(四)—钩子和事务实现</a>
