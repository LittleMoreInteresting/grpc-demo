
> 前面章节中gRPC的服务都没有提供证书
支持，因此客户端在链接服务器中通过 grpc.WithInsecure() 选项跳过了对服务器证书的验证。没
有启用证书的gRPC服务在和客户端进行的是明文通讯，信息面临被任何第三方监听的风险。为了保障
gRPC通信不被第三方监听篡改或伪造，我们可以对服务器启动TLS加密特性。

> 通过一个安全可靠的根证书分别对服务器和客户端的证书进行
签名。这样客户端或服务器在收到对方的证书后可以通过根证书进行验证证书的有效性。


## Errors

- 2022/09/06 10:03:53 rpc error: code = Unavailable desc = connection error: desc = "transport: authentication handshake failed: x509: certificate relies on legacy C
  ommon Name field, use SANs instead"

> 因为 go 1.15 版本开始废弃 CommonName，因此推荐使用 SAN 证书。 如果想兼容之前的方式，需要设置环境变量 GODEBUG 为 x509ignoreCN=0。
> 推荐重新生成版本 https://blog.csdn.net/weixin_40280629/article/details/113563351?spm=1001.2101.3001.6661.1&utm_medium=distribute.pc_relevant_t0.none-task-blog-2%7Edefault%7ECTRLIST%7ERate-1-113563351-blog-109230584.pc_relevant_multi_platform_whitelistv4eslandingctr&depth_1-utm_source=distribute.pc_relevant_t0.none-task-blog-2%7Edefault%7ECTRLIST%7ERate-1-113563351-blog-109230584.pc_relevant_multi_platform_whitelistv4eslandingctr&utm_relevant_index=1

