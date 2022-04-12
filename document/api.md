# fighting接口文档

# api相关介绍

host:http://39.106.3.178

port:678

## 方法描述：发布加油动态

URL地址：/record

请求方法：POST

请求体

| 字段 | 类型 | 说明 |
| --- | --- | --- |
| receiver | string | 接收者 |
| content | string | 内容 |

请求体示例：

```json
{
  "receiver": "lxp",
  "content": "fighting"
}
```

响应体

| 字段 | 类型 | 说明 |
| --- | --- | --- |
| data | json | 相关数据 |

响应体示例

```json
{
  "data": null
}
```

## 方法描述：评论回复

URL地址：/comment

请求方法：POST

请求体

| 字段 | 类型 | 说明 |
| --- | --- | --- |
| record_id | uint | 记录id |
| content | string | 内容 |

请求体示例

```json
{
  "record_id": 1,
  "content": "fighting"
}
```

响应体

| 字段 | 类型 | 说明 |
| --- | --- | --- |
| data | json | 相关数据 |

响应体示例

```json
{
  "data": null
}
```

## 方法描述：获取记录总数

URL地址：/records/number

请求方法：GET

无请求体

响应体

| 字段 | 类型 | 说明 |
| --- | --- | --- |
| data | json | 相关数据 |

```json
{
  "data": null
}
```

## 方法描述：获取记录

URL地址：/records

请求方法：GET

请求体：/records?page=1

响应体

| 字段 | 类型 | 说明 |
| --- | --- | --- |
| data | json | 相关数据 |

```json
{
  "data": null
}
```