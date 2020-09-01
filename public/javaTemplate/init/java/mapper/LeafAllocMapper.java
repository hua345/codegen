package {{.PackageName}}.mapper;

import {{.PackageName}}.model.LeafAlloc;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Param;
import org.apache.ibatis.annotations.Update;

import java.util.List;

/**
 * @author {{.Author}}
 * @date {{.NowDate}}
 */
@Mapper
public interface LeafAllocMapper {
    /**
     * 查询所有数据
     */
    List<LeafAlloc> selectAllLeafAlloc();

    @Update("UPDATE leaf_alloc SET max_id = max_id + step WHERE biz_tag = #{tag}")
    void updateMaxId(@Param("tag") String tag);

    int deleteByPrimaryKey(String bizTag);

    int insert(LeafAlloc record);

    LeafAlloc selectByPrimaryKey(String bizTag);

    int updateByPrimaryKey(LeafAlloc record);
}