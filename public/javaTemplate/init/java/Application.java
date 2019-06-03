package {{.PackageName}};

import org.mybatis.spring.annotation.MapperScan;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;{{ if .SupportSwagger }}
import springfox.documentation.swagger2.annotations.EnableSwagger2;{{ end }}

/**
 * @author {{.Author}}
 * @date {{.NowDate}}
 */{{ if .SupportSwagger }}
@EnableSwagger2{{ end }}
@SpringBootApplication
@MapperScan("{{.PackageName}}.mapper")
public class Application {

    public static void main(String[] args) {
        SpringApplication.run(Application.class, args);
    }
}
