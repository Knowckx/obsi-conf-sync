

```plaintext
📁 src
├── 📁 apps // 程序入口
│   └── 📁 tools // 工具入口 快速debug验证
├── 📁 conf  // 配置和全局上下文
│   ├── 📄 conf.py
│   └── 📄 runtime.py
├── 📁 pb 
│   ├── 📁 protos // protoc生成的文件
│   ├── 📄 pb_client/pb_service.py    // 定义服务handle
│   ├── 📄 pb_mapper.py   // pb类型的拆箱和装箱
│   └── 📄 rpcs.py       // 对pb接口的实现
├── 📁 infa // 跨项目依赖
├── 📁 inner // 内部代码
└── 📁 temp // 
```

```plaintext
📁 inner // 内部代码
├── 📁 utils // 小包装
└── 📁 cache // 缓存处理
└── 📁 trade // 交易
    └── 📁 tmod // 交易模型的定义
└── 📁 types // 基础定义
    └── 📁 enums // 枚举值
    └── 📁 instrument // 交易标的
    

```

