package {{.PackageName}}.common.validator;

import {{.PackageName}}.common.annotation.Amount;

import javax.validation.ConstraintValidator;
import javax.validation.ConstraintValidatorContext;
import java.math.BigDecimal;
import java.util.regex.Pattern;

/**
 * @author {{.Author}}
 * @date {{.NowDate}}
 */
public class AmountBigDecimalValidator implements ConstraintValidator<Amount, BigDecimal> {

    /**
     * 表示金额的正则表达式
     */
    private String moneyReg = "^\\d+(\\.\\d{1,2})?$";
    private Pattern moneyPattern = Pattern.compile(moneyReg);

    @Override
    public void initialize(Amount amount) {

    }

    @Override
    public boolean isValid(BigDecimal value, ConstraintValidatorContext arg1) {
        if (null == value) {
            return true;
        }
        return moneyPattern.matcher(value.toString()).matches();
    }
}
