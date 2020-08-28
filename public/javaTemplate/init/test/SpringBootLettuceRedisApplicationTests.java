package {{.PackageName}};

import {{.PackageName}}.utils.DateFormatEnum;
import {{.PackageName}}.utils.DateUtil;
import {{.PackageName}}.utils.SnowFlake;
import {{.PackageName}}.utils.SnowFlakeUtil;
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

/**
 * @author {{.Author}}
 * @date {{.NowDate}}
 */
@RunWith(SpringRunner.class)
@SpringBootTest
public class SpringBootLettuceRedisApplicationTests {

    private static final Logger log = LoggerFactory.getLogger(SpringBootLettuceRedisApplicationTests.class);

    @Autowired
    private StringRedisTemplate stringRedisTemplate;

    @Autowired
    private RedisTemplate<String, Serializable> redisCacheTemplate;


    @Test
    public void redisTest() {
        // 测试线程安全
        ExecutorService executorService = Executors.newFixedThreadPool(1000);
        String testIdKey = "test_userId";
        String testName = "test_name";
        //主键生成
        stringRedisTemplate.opsForValue().set(testIdKey, "10000");
        IntStream.range(0, 100).forEach(i ->
                executorService.execute(() -> stringRedisTemplate.opsForValue().increment(testIdKey, 1))
        );
        // 简单key value获取
        String userId = stringRedisTemplate.opsForValue().get(testIdKey);
        log.info("[主键生成userId] - [{}]", userId);
        stringRedisTemplate.opsForValue().set(testName, "fang");
        String name = stringRedisTemplate.opsForValue().get(testName);
        log.info("[字符缓存结果] - [{}]", name);
        //  以下只演示整合，具体Redis命令可以参考官方文档，Spring Data Redis 只是改了个名字而已，Redis支持的命令它都支持
        String userIdKey = "testUser:" + userId;
        redisCacheTemplate.opsForValue().set(userIdKey, "fangfang");
        // 对应 String（字符串）
        String fangName = (String) redisCacheTemplate.opsForValue().get(userIdKey);
        log.info("[对象缓存结果] - [{}]", fangName);
    }

    @Test
    public void SnowFlakeTest() {
        log.info("雪花算法起始时间:{}",DateUtil.formatDateTime(SnowFlake.START_STMP, DateFormatEnum.DATE_YYYY_MM_DD_HH_MM_SS));
        for (int i = 0; i < 10; i++) {
            log.info("雪花算法生成id：{}",SnowFlakeUtil.getNextId());
        }
    }
}
