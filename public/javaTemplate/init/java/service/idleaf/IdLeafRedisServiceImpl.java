package {{.PackageName}}.service.idleaf;

import {{.PackageName}}.model.LeafAlloc;
import {{.PackageName}}.utils.SnowFlakeUtil;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.redis.core.StringRedisTemplate;
import org.springframework.data.redis.support.atomic.RedisAtomicLong;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;
import org.springframework.util.StringUtils;

import java.util.Map;
import java.util.Set;
import java.util.concurrent.*;

/**
 * @author {{.Author}}
 * @date {{.NowDate}}
 */
@Service
public class IdLeafRedisServiceImpl implements IdLeafServiceFactory {
    private static final Logger logger = LoggerFactory.getLogger(IdLeafRedisServiceImpl.class);

    private Map<String, SegmentBuffer> leafMap = new ConcurrentHashMap<String, SegmentBuffer>();

    private ExecutorService service = new ThreadPoolExecutor(3, 10, 60L, TimeUnit.SECONDS, new SynchronousQueue<Runnable>(), new IdLeafThreadFactory());

    private final static String ID_LEAF_PREFIX = "idLeaf:";

    private final static Integer DEFAULT_STEP = 1000;
    @Autowired
    private StringRedisTemplate stringRedisTemplate;

    public static class IdLeafThreadFactory implements ThreadFactory {

        private static int threadInitNumber = 0;

        private static synchronized int nextThreadNum() {
            return threadInitNumber++;
        }

        @Override
        public Thread newThread(Runnable r) {
            return new Thread(r, "Thread-IdLeaf-" + nextThreadNum());
        }
    }

    public void initAllLeafFromDb() {
        Set<String> leafKeys = stringRedisTemplate.keys(ID_LEAF_PREFIX);
        if (StringUtils.isEmpty(leafKeys)) {
            return;
        }
        leafKeys.stream().forEach(item -> {
            initLeafFromDb(item);
        });
    }

    @Override
    public Long getIdByBizTag(final String bizTag) {
        if (!leafMap.containsKey(bizTag)) {
            synchronized (leafMap) {
                if (!leafMap.containsKey(bizTag)) {
                    initLeafFromDb(bizTag);
                }
            }
        }
        return getIdFromSegmentBuffer(leafMap.get(bizTag));
    }

    /**
     * 初始化IdLeaf
     */
    public void initLeafFromDb(String bizTag) {
        SegmentBuffer buffer = new SegmentBuffer();
        Segment segment = buffer.getCurrent();
        LeafAlloc leafAlloc = updateMaxIdAndGetLeafAlloc(bizTag);
        buffer.setStep(leafAlloc.getStep());
        buffer.setKey(bizTag);
        long currentId = leafAlloc.getMaxId() - buffer.getStep();
        segment.getCurrentId().set(currentId);
        segment.setMax(leafAlloc.getMaxId());
        segment.setStep(buffer.getStep());
        leafMap.put(bizTag, buffer);
    }

    /**
     * 更新idLeaf
     */
    public void updateLeafFromDb(String bizTag, Segment segment) {
        SegmentBuffer buffer = segment.getBuffer();
        LeafAlloc leafAlloc = updateMaxIdAndGetLeafAlloc(bizTag);
        buffer.setStep(leafAlloc.getStep());
        long currentId = leafAlloc.getMaxId() - buffer.getStep();
        segment.getCurrentId().set(currentId);
        segment.setMax(leafAlloc.getMaxId());
        segment.setStep(buffer.getStep());
    }

    @Transactional(rollbackFor = Exception.class)
    public LeafAlloc updateMaxIdAndGetLeafAlloc(String bizTag) {
//        RedisAtomicLong entityIdCounter = new RedisAtomicLong(bizTag, stringRedisTemplate.getConnectionFactory());
//        Long increment = entityIdCounter.getAndIncrement();
        Long increment = stringRedisTemplate.opsForValue().increment(bizTag);
        LeafAlloc leafAlloc = new LeafAlloc();
        leafAlloc.setBizTag(bizTag);
        leafAlloc.setMaxId(increment * DEFAULT_STEP);
        leafAlloc.setStep(DEFAULT_STEP);
        return leafAlloc;
    }

    public Long getIdFromSegmentBuffer(final SegmentBuffer buffer) {
        while (true) {
            buffer.rLock().lock();
            try {
                final Segment segment = buffer.getCurrent();
                // 加载另外一个id段
                if (!buffer.isNextReady() && (segment.getAvailableIdRange() < 0.8 * segment.getStep()) && buffer.getThreadRunning().compareAndSet(false, true)) {
                    service.execute(new Runnable() {
                        @Override
                        public void run() {
                            Segment next = buffer.getSegments()[buffer.nextPos()];
                            boolean updateOk = false;
                            try {
                                updateLeafFromDb(buffer.getKey(), next);
                                updateOk = true;
                                logger.info("update segment {} from db {}", buffer.getKey(), next);
                            } catch (Exception e) {
                                logger.warn(buffer.getKey() + " updateSegmentFromDb exception", e);
                            } finally {
                                if (updateOk) {
                                    buffer.wLock().lock();
                                    buffer.setNextReady(true);
                                    buffer.getThreadRunning().set(false);
                                    buffer.wLock().unlock();
                                } else {
                                    buffer.getThreadRunning().set(false);
                                }
                            }
                        }
                    });
                }
                long value = segment.getCurrentId().getAndIncrement();
                if (value < segment.getMax()) {
                    return value;
                }
            } finally {
                buffer.rLock().unlock();
            }
            waitAndSleep(buffer);
            buffer.wLock().lock();
            try {
                final Segment segment = buffer.getCurrent();
                long value = segment.getCurrentId().getAndIncrement();
                if (value < segment.getMax()) {
                    return value;
                }
                if (buffer.isNextReady()) {
                    buffer.switchPos();
                    buffer.setNextReady(false);
                } else {
                    logger.error("leaf生成id异常:{}，使用雪花算法生成!", buffer);
                    return SnowFlakeUtil.getNextId();
                }
            } finally {
                buffer.wLock().unlock();
            }
        }
    }

    private void waitAndSleep(SegmentBuffer buffer) {
        while (buffer.getThreadRunning().get()) {
            try {
                TimeUnit.MILLISECONDS.sleep(6);
            } catch (InterruptedException e) {
                logger.warn("Thread {} Interrupted", Thread.currentThread().getName());
                break;
            }
        }
    }
}
