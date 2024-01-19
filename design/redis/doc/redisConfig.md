## 参数说明
### maxmemory-policy

+ volatile-lru：对设置 TTL 过期时间的 key，使用 LRU 算法淘汰并删除。 
+ allkeys-lru：对所有的 key，使用 LRU 算法淘汰删除。 
+ volatile-random：对设置TTL 过期时间的 key，随机地淘汰删除。 
+ allkeys-random：对所有的 key，随机地淘汰删除。 
+ volatile-ttl：对设置TTL 过期时间的 key，淘汰删除即将到达过期时间的 key。 
+ noeviction：不淘汰删除任何 key，在写操作时返回错误信息。


### client
+ timeout: 当客户端连接闲置时间达到该指定值时，将关闭连接，单位为秒（s）。
+ no_loose_disabled-commands: 设置禁用命令，可根据业务需求禁用某些高危命令或高时间复杂度的命令，例如FLUSHALL、FLUSHDB、KEYS、HGETALL、EVAL、EVALSHA、SCRIPT等。
+ password

### persistence
+ appendfsync: AOF（AppendOnly File）持久化功能的fsync频率，仅在appendonly参数开启时生效，默认为everysec，不支持修改。