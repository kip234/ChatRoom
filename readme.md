# 产品说明书？

## 客户端使用

### 启动



| 参数 | 值              | 描述     |
| ---- | --------------- | -------- |
| -h   | localhost:8080  | 主机地址 |
| -n   | string          | 用户名   |
| -p   | 123             | 用户密码 |
| -u   | 1               | 用户ID   |
| -r   | login、register | 请求     |

示例：

- 注册

```shell
client -h=localhost:8080 -r=register -n=kip -p=123
```

- 登录

```shell
client -h=localhost:8080 -r=login -u=1 -p=123
```

### 后续使用

#### 发消息

> 直接在终端输入后回车就行

#### room操作

| 指令     | 参数   | 描述           |
| -------- | ------ | -------------- |
| /lstroom | 无     | 房间列表       |
| /inroom  | 房间名 | 加入房间       |
| /outroom | 房间名 | 退出房间       |
| /nwroom  | 房间名 | 新建房间并加入 |

#### 强行解释

> 成功登陆后用户处于大厅(lobby)中，默认情况下任何用户都可以接收到大厅内的消息
>
> 当执行/inroom后就可以接收这个房间内的消息，当前发送的消息也将被视为该房间内的消息
>
> 当执行/outroom后会退出该房间不再接收消息，之后发送的消息视为在大厅发送的消息
>
> 可以使用/inroom进入多个房间并接收消息，直到/outroom后才会停止接收

#### 其他指令

> \#q、#exit 退出

## 服务端使用

> MySQL Redis…

## 其他东西

> 关于资源释放：
>
> >  ChatRoom/server/Handlers/funcs.go:68
> >
> > ChatRoom/Models/connpool.go:34