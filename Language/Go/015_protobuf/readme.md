Protocol Buffers（protobuf）是一种语言无关、平台无关的数据序列化格式，它使用.proto文件定义数据结构和消息格式。以下是一些protobuf的规则：

# 语法声明：
在.proto文件的开头，需要声明使用的protobuf语法版本。目前最常用的版本是proto3。
```
syntax = "proto3";
```
# 消息类型定义：
使用message关键字定义消息类型。消息类型可以包含字段和其他嵌套消息类型。
```
message Person {
string name = 1;
int32 age = 2;
}
```
# 字段定义：
字段用于存储消息中的数据。每个字段都有一个唯一的数字标识符和一个类型。可以使用不同的类型，如整数、浮点数、布尔值、字符串等。
```
message Person {
string name = 1;
int32 age = 2;
repeated string hobbies = 3;
}
```

# 字段规则：
字段可以具有不同的规则，如required、optional和repeated。required表示字段必须存在，optional表示字段可以存在也可以不存在，repeated表示字段可以重复出现多次。
```
message Person {
required string name = 1;
optional int32 age = 2;
repeated string hobbies = 3;
}
```

# 枚举类型定义：
使用enum关键字定义枚举类型。枚举类型定义了一组可能的值。
```
enum Gender {
UNKNOWN = 0;
MALE = 1;
FEMALE = 2;
}
```

# 导入其他.proto文件：
可以使用import语句导入其他.proto文件，以
以使用其中定义的消息类型。
```
import "other.proto";

message Person {
string name = 1;
other.OtherMessage other_message = 2;
}
```

# 注释：
可以使用//或/* */来添加注释，以提供关于消息类型、字段或其他元素的说明。
```
// 这是一个人的消息类型
message Person {
string name = 1; // 姓名
int32 age = 2; // 年龄
}
```
