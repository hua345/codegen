package {{.PackageName}}.service.idleaf;

import java.util.concurrent.atomic.AtomicLong;

/**
 * @author {{.Author}}
 * @date {{.NowDate}}
 */
public class Segment {
    private AtomicLong currentId = new AtomicLong(0);

    private volatile long max;

    private volatile int step;

    private SegmentBuffer buffer;

    public Segment(SegmentBuffer buffer) {
        this.buffer = buffer;
    }

    public AtomicLong getCurrentId() {
        return currentId;
    }

    public void setCurrentId(AtomicLong currentId) {
        this.currentId = currentId;
    }

    public long getMax() {
        return max;
    }

    public void setMax(long max) {
        this.max = max;
    }

    public int getStep() {
        return step;
    }

    public void setStep(int step) {
        this.step = step;
    }

    public SegmentBuffer getBuffer() {
        return buffer;
    }

    /**
     * 可用Id范围
     */
    public long getAvailableIdRange() {
        return this.getMax() - getCurrentId().get();
    }

    @Override
    public String toString() {
        StringBuilder sb = new StringBuilder("Segment(");
        sb.append("value:");
        sb.append(currentId);
        sb.append(",max:");
        sb.append(max);
        sb.append(",step:");
        sb.append(step);
        sb.append(")");
        return sb.toString();
    }
}
