# ProtoBuf详解-使用指南

参考官方文档：[Protocol Buffers Developer Guide](https://developers.google.com/protocol-buffers/docs/overview)

## 1. 简介

ProtoBuf 全称 Protocol Buffers，是谷歌[开源](https://github.com/protocolbuffers/protobuf)的一种语言无关、平台无关、可扩展的序列化结构数据的方法，它可用于通信协议、数据存储。



类比 XML、JSON 等：

|        | XML                  | JSON               | ProtoBuf                                         |
| ------ | -------------------- | ------------------ | ------------------------------------------------ |
| 可读性 | 标签形式，可读性较高 | 键值对，可读性较高 | 原始格式为二进制，可读性低(.proto文件可读性较高) |
| 速度   | 慢                   | 中                 | 高速（比JSON快 3~5倍，比XML快20～100倍）         |
| 体积   | 大，内容较冗余       | 中                 | 小（比XML小3～10倍）                             |



Protoctol 的特点：

* **语言无关、平台无关**：支持 Java、C++、Python、Go 等多种语言，支持多个平台
* **高效**：比 XML 更小（3 ~ 10倍）、更快（20 ~ 100倍）、更为简单
* **扩展性、兼容性好**：可以更新数据结构，而不影响和破坏原有的旧程序
* **代码生成机制**：可以通过代码生成代码



## 2. 用法

### 2.1.  .proto 文件

ProtoBuf 通过在 .proto 文件中定义 protocol buffer message 类型，来指定如何对序列化信息进行结构化。

每一个 protocol buffer message 是一个信息的小逻辑记录。



### 2.2. Message 类型

#### 2.2.1 语法（proto3）

官方英文文档：[https://developers.google.com/protocol-buffers/docs/proto3](https://developers.google.com/protocol-buffers/docs/proto3)

```protobuf
// 指定使用 proto3 语法
syntax = "proto3";

message 名称 {
	[字段规则] 类型 名称 = 字段编号;
	// 例如：string name = 1;
}
// 一个 .proto 文件中可以用多个 message 类型
// 注释使用 // 或 /*...*/
```

* 关键字 `message` 定义一个由多个消息字段 （field）组成的实体结构
* **消息字段（field）**包括：字段规则、类型、名称、编号
* **字段规则**：singular、repeated
* **类型**：标量类型、合成类型
* **字段编号**：0 ~ 536870911（除去 19000 到 19999 之间的数字）



#### 2.2.2. 字段规则说明

* **singular**：字段出现 0 次或 1 次，超过 1 次
* **repeated**：字段可以出现任意次（包括 0 次），其中重复值的顺序会被保留

> 在proto3中，`repeated` 数字类型的字段默认使用 `packed` 编码。
>
> 您可以在[协议缓冲区编码中](https://developers.google.com/protocol-buffers/docs/encoding.html#packed)找到更多有关 `packed` 编码的信息。



#### 2.2.3. 保留字段

使用 `reserved` 指定保留的字段，保留一些曾经使用过后来删除或注释掉的编号或名称，避免未来有人修改这个 `.proto` 文件的时候，重复使用这些编号或名称，不然如果以后万一因某些原因加载到相同 `.proto` 的旧版本时可能会导致出现一些数据损坏或隐蔽的bug。确保不会发生这种情况的一种方法是指定保留已删除条目的数值（或名称，名称可能导致JSON序列化问题）。如果将来有任何用户尝试使用这些标识符，则 protocol buffer 编译器会发生警告。

```protobuf
message Foo {
  reserved 2, 15, 9 to 11; // 指定保留字段编号
  reserved "foo", "bar"; // 指定保留字段名称
}
```

> 字段编号跟名称不能出现在同一行，例如错误的例子：reserved 2, 15, "foo";



#### 2.2.4. 默认值

解析 message 时，如果编码的消息不包含特定的 singular 元素，则已解析对象中的相应字段将设置为该字段的默认值。

下列的每个类型对应的默认值：

| 类型             | 默认值                            |
| ---------------- | --------------------------------- |
| 字符串（stirng） | 空字符串                          |
| bytes            | 空 bytes                          |
| 布尔值（bool）   | false                             |
| 数值             | 0                                 |
| 枚举             | 默认是第一个定义的枚举值，必须为0 |

> 对于 message 类型，如果未设置字段。它的确切值取决于使用的语言。有关详细信息，请参见[ generated code guide](https://developers.google.com/protocol-buffers/docs/reference/overview)。
>
>  `repeated` 字段的默认值为空（生成一个相应语言的空列表）

>注意：对于标量类型的 message 字段，一旦被解析，如果值是等于默认值的（例如：布尔类型字段值等于false），无法确认它的值是被显式设置的还是本来就没有设置值。例如你不应该定义 boolean 的默认值 false 作为任何行为的触发方式。也应该注意如果一个标量类型的 message 字段被设置为默认值，这个值将不会被序列化传输。



#### 2.2.5. 例子

定义了一个包含 "person" 相关信息的 message

```protobuf
// 指定使用 proto3 语法，不指定的话，默认会用 proto2 语法，现在基本上都用 proto3
syntax = "proto3";

// 定义一个名称为 SearchRequest 的消息类型
message SearchRequest {
  string query = 1; // 定义类型为 string，名称为 query ，编号为 1 的字段（field）
  int32 page_number = 2;
  int32 result_per_page = 3;
}
```





### 2.3. 使用 .proto 可以生成什么？

| 语言        | 生成文件             | 描述                                                         |
| ----------- | -------------------- | ------------------------------------------------------------ |
| C++         | .h 和 .cc            | 生成的文件包含 `.proto` 文件中描述的每种 message 类型对应的类 |
| Java        | .java                | 每个 message 类型生成一个 .java 文件（类），以及用于创建 message 类实例的特殊 Builder 类 |
| Python      |                      | Python有点不同，Python编译器为.proto文件中的每个message类型生成一个含有静态描述符的模块，该模块与一个元类（metaclass）在运行时（runtime）被用来创建所需的Python数据访问类。 |
| Go          | .pb.go               | 每个 message 类型在 .pb.go 中会声明对应的类型                |
| Ruby        | .rb                  | 生成一个包含所有 mesaage 类型的 Ruby 模块文件                |
| Objective-C | pbobjc.h 和 pbobjc.m | 文件中描述的每种 message 类型都有一个类。                    |
| C#          | .cs                  | 文件中描述的每种 message 类型都有一个类。                    |
| Dart        | .pb.dart             | 文件中描述的每种 message 类型都有一个类。                    |



### 2.4. 枚举（Enum）

#### 2.4.1. 使用

当你定义一个 message 类型时，可能希望一个字段是从一个预定义的列表里面取值，此时就可以使用枚举。

例如：想在 `SearchRequest` 中添加一个 `corpus` 字段，这个字段是值可能是 `UNIVERSAL`, `WEB`, `IMAGES`, `LOCAL`, `NEWS`, `PRODUCTS` 或 `VIDEO`。可以在 message 中定义一个 `enum` 以及对应的每个可能值的常量在实现这个操作。

```protobuf
message SearchRequest {
  string query = 1;
  int32 page_number = 2;
  int32 result_per_page = 3;
  enum Corpus {
    UNIVERSAL = 0;
    WEB = 1;
    IMAGES = 2;
    LOCAL = 3;
    NEWS = 4;
    PRODUCTS = 5;
    VIDEO = 6;
  }
  Corpus corpus = 4;
}
```

> 每个 enum 的第一个常量必须映射为 0，并且这个映射为 0 的常量必须存在：
>
> * 必须有一个零值，这个我们可以把这个 0 作为默认值
> * 这个零值必须作为第一个常量元素，为了兼容proto2，枚举类的第一个值总是默认值



当需要分配两个相同的值给不同的常量时，可以使用别名。

使用别名的时候，必须将 `allow_alias` 选项设置为 `true`，否则编译器遇到别名的时候会产生一个错误信息。

```protobuf
// option allow_alias = true 设置后，允许 STARTED、RUNNING 映射同样的值
message MyMessage {
  enum EnumAllowingAlias {
    option allow_alias = true;
    UNKNOWN = 0;
    STARTED = 1;
    RUNNING = 1;
  }
}
enum EnumNotAllowingAlias {
  UNKNOWN = 0;
  STARTED = 1;
  // RUNNING = 1;  // Uncommenting this line will cause a compile error inside Google and a warning message outside.
}
```

枚举器常量必须在 32 位整型值的范围内。因为`enum` 值传输使用 [ varint 编码](https://developers.google.com/protocol-buffers/docs/encoding)，使用负值效率不高，因此不推荐使用。如上例所示，可以在 一个 message 定义的内部或外部定义枚举——这些枚举可以在 .proto 文件中的任何 message 定义里重用。当然也可以在一个 message 中声明一个枚举类型，而在另一个不同的 message 中使用它——采用 *`MessageType.EnumType`* 的语法格式。

当使用 protocol buffer 编译器编译一个使用了 `enum` 的 .proto 文件时，生成的代码中将有一个对应的enum（对Java或C++来说），或者一个特殊的EnumDescriptor类（对 Python来说），后者被用来在运行时生成的类中创建一系列的整型值符号常量（symbolic constants）。

> **注意**：生成的代码可能受到特定于枚举器数量的语言的限制（一种语言的下限为几千）。请查看您计划使用的语言的限制。

在反序列化的过程中，无法识别的枚举值会被保存在消息中，如何表示该值取决于语言。在支持具有超出指定范围的值的开放式枚举类型的语言（例如C ++和Go）中，未知的枚举值仅存储为其基础整数表示形式。在具有封闭枚举类型的语言（例如Java）中，使用枚举中的一个类型来表示未识别的值，并且可以使用特殊的访问器访问基础整数。无论哪种情况，如果 message 已序列化，则无法识别的值仍将与 message 一起序列化。

更多关于如何在你的应用程序的中使用枚举的信息，请查看所选择语言的 [generated code guide](http://code.google.com/intl/zh-CN/apis/protocolbuffers/docs/reference/overview.html。)



#### 2.4.2. 保留字段

使用 `reserved` 指定保留的字段，保留一些曾经使用过后来删除或注释掉的编号或名称，避免未来有人修改这个 `.proto` 文件的时候，重复使用这些编号或名称，不然如果以后万一因某些原因加载到相同 `.proto` 的旧版本时可能会导致出现一些数据损坏或隐蔽的bug。确保不会发生这种情况的一种方法是指定保留已删除条目的数值（或名称，名称可能导致JSON序列化问题）。如果将来有任何用户尝试使用这些标识符，则 protocol buffer 编译器会发生警告。

```
enum Foo {
  reserved 2, 15, 9 to 11, 40 to max;
  reserved "FOO", "BAR";
}
```

> 可以使用 max 关键字指定保留的数值范围达到最大可能值。
>
> 不能在同一个 reserverd 语句中混合使用字段名和数值。



### 2.5. 使用 message 作为另一个 message 的字段类型

#### 2.5.1. 用法

字段类型可以使用另一个 message 类型。

例如：在 `SearchResponse`  message 中包含 `Result` 

```protobuf
// 定义一个 SearchResponse message 类型
// 将 Result 作为 SearchResponse 的字段类型
message SearchResponse {
  repeated Result results = 1; // 指定字段类型为 Result
}
// 在同一 .proto 文件中，定义 Result message 类型
message Result {
  string url = 1;
  string title = 2;
  repeated string snippets = 3;
}
```

#### 2.5.2. 引入

上例是在同一文件中使用另一个 message 作为字段类型，同样也可以通过引入，使用其他 .proto 文件中定义的 message。

```
import "myproject/other_protos.proto";
```

默认只能使用直接引入的 .proto 文件中的定义。但有时可能需要将 .proto 文件移动到新的位置，此时可以在旧的位置放置一个虚拟的 .proto 文件，使用 import public 语句转发所有引入的 .proto 文件，而不是直接移动 .proto 文件并修改每个引用的路径。

引入的 .proto 文件如果包含 import public ，则会传递转发这个 import public 引入的依赖。

例如：

new.proto

```protobuf
// 新的位置的 .proto 文件
// 所有原来 .proto 的内容都转移到此文件
```

old.proto

```protobuf
// 这个是原来被引用的就文件
import public "new.proto"; // 引入新位置的文件
import "other.proto";
```

client.proto *其他引用了 old.proto 的文件*

```protobuf
import "old.proto"; // 引入旧位置的 .proto 文件
// 可以使用 new.proto 和 old.proto 中定义的内容，但不能使用 other.proto 中定义的内容
```



### 2.6. 





## 3. 标量类型列表

[https://developers.google.com/protocol-buffers/docs/proto3#scalar](https://developers.google.com/protocol-buffers/docs/proto3#scalar)



















​	