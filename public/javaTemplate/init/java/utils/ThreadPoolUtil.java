package {{.PackageName}}.utils;

import lombok.extern.slf4j.Slf4j;

import java.util.concurrent.*;
import java.util.concurrent.atomic.AtomicLong;

/**
 * @author {{.Author}}
 * @date {{.NowDate}}
 */
@Slf4j
public class ThreadPoolUtil {
    public static class MyDiscardPolicy implements RejectedExecutionHandler {
        public MyDiscardPolicy() {
        }

        @Override
        public void rejectedExecution(Runnable r, ThreadPoolExecutor e) {
            log.error("WorkServer rejectedExecution 线程池队列已满");
        }
    }

    public static class MyNameThreadFactory implements ThreadFactory {

        private final String poolName;
        private AtomicLong count = new AtomicLong(1);

        public MyNameThreadFactory(String poolName) {
            this.poolName = poolName;
        }

        @Override
        public Thread newThread(Runnable r) {
            Thread t = new Thread(r, poolName + "my-thread-" + count.getAndIncrement());
            //设置为非守护线程
            if (t.isDaemon()) {
                t.setDaemon(false);
            }
            return t;
        }
    }

    private volatile static ExecutorService threadPool;

    public static ExecutorService getInstance() {
        if (null == threadPool) {
            synchronized (ThreadPoolUtil.class) {
                if (null == threadPool) {
                    ThreadFactory namedThreadFactory = new MyNameThreadFactory("my-threadPool-");
                    threadPool = new ThreadPoolExecutor(5, 10,
                            0, TimeUnit.SECONDS,
                            new ArrayBlockingQueue<>(512),
                            namedThreadFactory,
                            new MyDiscardPolicy());
                }
            }
        }
        return threadPool;
    }
}
