> 最开始Lua解释器的状态是完全隐藏在API后面的，散落在各种全局变量里。由于某些宿主环境(比如Web服务器)需要同时使用多个Lua解释器实例，Lua 3.1引入了lua State结构体，对解释器状态进行了封装，从而使得用户可以在多个解释器实例之间切换。

> Lua State是Lua API非常核心的概念，全部的API函数都是围绕Lua State进行操作

> 而Lua State内部封装的最为基础的一个状态就是虚拟栈(后面我们称其为Lua栈)。Lua栈是宿主语言 (对于官方Lua来说是C语言，对于本书来说是Go语言)和Lua语言进行**沟通的桥梁**。

> lua 代码里，变量是不携带类型信息的，变量的值携带类型信息