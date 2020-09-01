package {{.PackageName}};

import {{.PackageName}}.service.idleaf.IdLeafMysqlServiceImpl;
import {{.PackageName}}.service.idleaf.IdLeafRedisServiceImpl;
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
import org.springframework.test.context.ActiveProfiles;
import org.springframework.test.context.junit4.SpringRunner;

import java.util.concurrent.*;

/**
 * @author {{.Author}}
 * @date {{.NowDate}}
 */
@RunWith(SpringRunner.class)
@SpringBootTest
@ActiveProfiles("dev")
public class IdGenerateTests {

    private static final Logger log = LoggerFactory.getLogger(IdGenerateTests.class);

    @Autowired
    private IdLeafMysqlServiceImpl leaf;

    @Autowired
    private IdLeafRedisServiceImpl redisLeaf;

    @Test
    public void SnowFlakeTest() {
        log.info("雪花算法起始时间:{}", DateUtil.formatDateTime(SnowFlake.START_STMP, DateFormatEnum.DATE_YYYY_MM_DD_HH_MM_SS));
        Long startMS = System.currentTimeMillis();
        for (int i = 0; i < 1000000; i++) {
            SnowFlakeUtil.getNextId();
        }
        Long endMS = System.currentTimeMillis();
        log.info("雪花算法生成100万id耗时：{}ms", endMS - startMS);
    }

    @Test
    public void testMysqlLeaf() throws Exception {
        ExecutorService executorService = new ThreadPoolExecutor(2, 2,
                0, TimeUnit.SECONDS,
                new ArrayBlockingQueue<>(512), // 使用有界队列，避免OOM
                new ThreadPoolExecutor.DiscardPolicy());
        CountDownLatch latch = new CountDownLatch(2);
        Long startMS = System.currentTimeMillis();
        executorService.submit(new Runnable() {
            @Override
            public void run() {
                for (int i = 0; i < 100000; i++) {
                    leaf.getIdByBizTag("leaf-segment-test");
                }
                latch.countDown();
            }
        });
        executorService.submit(new Runnable() {
            @Override
            public void run() {
                for (int i = 0; i < 100000; i++) {
                    redisLeaf.getIdByBizTag("leaf-segment-test");
                }
                latch.countDown();
            }
        });
        latch.await();
        Long endMS = System.currentTimeMillis();
        log.info("mysql leaf算法生成20万id耗时：{}ms", endMS - startMS);
        executorService.shutdown();
    }

    @Test
    public void testRedisLeaf() throws Exception {
        ExecutorService executorService = new ThreadPoolExecutor(2, 2,
                0, TimeUnit.SECONDS,
                new ArrayBlockingQueue<>(512), // 使用有界队列，避免OOM
                new ThreadPoolExecutor.DiscardPolicy());
        CountDownLatch latch = new CountDownLatch(2);
        Long startMS = System.currentTimeMillis();
        executorService.submit(new Runnable() {
            @Override
            public void run() {
                for (int i = 0; i < 100000; i++) {
                    redisLeaf.getIdByBizTag("leaf-segment-test");
                }
                latch.countDown();
            }
        });
        executorService.submit(new Runnable() {
            @Override
            public void run() {
                for (int i = 0; i < 100000; i++) {
                    redisLeaf.getIdByBizTag("leaf-segment-test");
                }
                latch.countDown();
            }
        });
        latch.await();
        Long endMS = System.currentTimeMillis();
        log.info("redis leaf算法生成20万id耗时：{}ms", endMS - startMS);
        executorService.shutdown();
    }
}
