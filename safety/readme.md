

## Errors

- 2022/09/06 10:03:53 rpc error: code = Unavailable desc = connection error: desc = "transport: authentication handshake failed: x509: certificate relies on legacy C
  ommon Name field, use SANs instead"

> 因为 go 1.15 版本开始废弃 CommonName，因此推荐使用 SAN 证书。 如果想兼容之前的方式，需要设置环境变量 GODEBUG 为 x509ignoreCN=0。
> 推荐重新生成版本 https://blog.csdn.net/weixin_40280629/article/details/113563351?spm=1001.2101.3001.6661.1&utm_medium=distribute.pc_relevant_t0.none-task-blog-2%7Edefault%7ECTRLIST%7ERate-1-113563351-blog-109230584.pc_relevant_multi_platform_whitelistv4eslandingctr&depth_1-utm_source=distribute.pc_relevant_t0.none-task-blog-2%7Edefault%7ECTRLIST%7ERate-1-113563351-blog-109230584.pc_relevant_multi_platform_whitelistv4eslandingctr&utm_relevant_index=1

