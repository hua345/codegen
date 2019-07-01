package com.github.code.admin;

import org.junit.Test;
import org.junit.runner.RunWith;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.data.redis.core.StringRedisTemplate;
import org.springframework.test.context.junit4.SpringRunner;

import java.io.Serializable;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.stream.IntStream;

@RunWith(SpringRunner.class)
@SpringBootTest( classes = Application.class)
public class SpringBootLettuceRedisApplicationTests {

    private static final Logger log = LoggerFactory.getLogger(SpringBootLettuceRedisApplicationTests.class);

    @Autowired
    private StringRedisTemplate stringRedisTemplate;

    @Autowired
    private RedisTemplate<String, Serializable> redisCacheTemplate;


    @Test
    public void get() {
        // 测试线程安全
        ExecutorService executorService = Executors.newFixedThreadPool(1000);
        //主键生成
        stringRedisTemplate.opsForValue().set("userId", "10000");
        IntStream.range(0, 1000).forEach(i ->
                executorService.execute(() -> stringRedisTemplate.opsForValue().increment("userId", 1))
        );
        // 简单key value获取
        String userId = stringRedisTemplate.opsForValue().get("userId");
        log.info("[主键生成userId] - [{}]", userId);
        stringRedisTemplate.opsForValue().set("name", "fang");
        String name = stringRedisTemplate.opsForValue().get("name");
        log.info("[字符缓存结果] - [{}]", name);
        //  以下只演示整合，具体Redis命令可以参考官方文档，Spring Data Redis 只是改了个名字而已，Redis支持的命令它都支持
        String userIdKey = "user:" + userId;
        redisCacheTemplate.opsForValue().set(userIdKey, "fangfang");
        // 对应 String（字符串）
        String fangName = (String) redisCacheTemplate.opsForValue().get(userIdKey);
        log.info("[对象缓存结果] - [{}]", fangName);
    }
}
