# 公式リンク
- [Java Platform Standard Edition 8 Documentation](https://docs.oracle.com/javase/8/docs/)
- [The Java® Virtual Machine Specification](https://docs.oracle.com/javase/specs/jvms/se8/html/index.html)
- [JDK 21 Documentation - Home](https://docs.oracle.com/en/java/javase/21/)
- [Java SE Specifications](https://docs.oracle.com/javase/specs/index.html)
- [Java Virtual Machine Technology Overview](https://docs.oracle.com/en/java/javase/21/vm/java-virtual-machine-technology-overview.html)


## 第2章　classファイルとデコード

### classファイルフォーマット
```c
ClassFile {
    u4       magic;
    u2       minor_version;
    u2       major_version;
    u2       constant_pool_count;
    up_info  constant_pool[constant_pool_count-1];
    u2       access_flags;
    ,,,
    ,,
}
```

classファイルのフォーマットに従って、実装を進める。
上の資料を読んで、該当箇所を確認。
class file formatの章があり、これをデコードして実行するイメージ。