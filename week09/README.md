stickpackage中包含对应3种ICoder处理粘包的具体实现，实现中只是参照其解决思想写的code，
其中边界值等条件可能考虑不周，解决粘包现象的核心就是“合理的分割”tcp建连后传递过来的数据包。