# IP地理位置信息离线解析

## 使用
1. UI界面可供单独查询

2. 打开界面API开关可以进行api接口查询，请求方式如下：
    Get请求：
    
    ```
    http://localhost:45555/info?ip=103.132.53.10
    ```
    
    成功：
    
    ```
    {
        "ip": "103.132.53.10",
        "region": "印度尼西亚|0|大雅加达|雅加达|0",
        "duration": "0s"
    }
    ```
    
    失败：
    
    ```
    {
        "ip": "123456",
        "region": "",
        "duration": "",
        "error": "IP输入有误！"
    }
    ```
    
    
## 底层
使用 [Ip2region](https://gitee.com/lionsoul/ip2region) 作为底层支持，感谢原作者[@lionsoul](https://gitee.com/lionsoul)，请支持

