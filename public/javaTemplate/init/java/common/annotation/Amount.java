package {{.PackageName}}.common.annotation;


import com.github.code.admin.common.validator.AmountBigDecimalValidator;
import com.github.code.admin.common.validator.AmountStringValidator;

import javax.validation.Constraint;
import javax.validation.Payload;
import java.lang.annotation.ElementType;
import java.lang.annotation.Retention;
import java.lang.annotation.RetentionPolicy;
import java.lang.annotation.Target;

@Target({ElementType.FIELD, ElementType.METHOD})
@Retention(RetentionPolicy.RUNTIME)
@Constraint(validatedBy = {AmountBigDecimalValidator.class, AmountStringValidator.class})
/**
 * @author {{.Author}}
 * @date {{.NowDate}}
 */
public @interface Amount {

    String message() default "Amount requires two precision numbers";

    Class<?>[] groups() default {};

    Class<? extends Payload>[] payload() default {};

}
