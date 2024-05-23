# compressURL

# 简介
一个使用 [goframe](https://goframe.org/display/gf) 创建的短链生成服务。
# 详细介绍

[猛击查看掘金文章](https://juejin.cn/post/7345310754470887458)

![image.png](https://p1-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/e6d9b9e7cb8f420e80cd5d48839aa4ea~tplv-k3u1fbpfcp-jj-mark:0:0:0:0:q75.image#?w=2174&h=828&s=193679&e=png&b=ffffff)

# 数据库统计索引优化
```sql
-- 创建索引以优化查询性能
CREATE INDEX idx_short_url_visits_on_created_at ON short_url_visits(created_at);
CREATE INDEX idx_short_url_visits_on_short_url ON short_url_visits(short_url);
```