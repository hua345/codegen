package {{.PackageName}}.model;

import lombok.Getter;
import lombok.Setter;

import java.util.Date;

/**
 * @author {{.Author}}
 * @date {{.NowDate}}
 */
@Getter
@Setter
public class LeafAlloc {

    private String bizTag;

    private Long maxId;

    private Integer step;

    private String description;

    private Date updateTime;

}