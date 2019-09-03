package {{.PackageName}}.common.validator;

import {{.PackageName}}.common.annotation.Amount;

import javax.validation.ConstraintValidator;
import javax.validation.ConstraintValidatorContext;
import java.util.regex.Pattern;

/**
 * @author {{.Author}}
 * @date {{.NowDate}}
 */
public class AmountStringValidator implements ConstraintValidator<Amount, String> {

    /**
     * 表示金额的正则表达式
     */
    private String moneyReg = "^\\d+(\\.\\d{1,2})?$";
    private Pattern moneyPattern = Pattern.compile(moneyReg);

    @Override
    public void initialize(Amount amount) {

    }

    @Override
    public boolean isValid(String value, ConstraintValidatorContext arg1) {
        if (null == value) {
            return true;
        }
        return moneyPattern.matcher(value).matches();
    }
}
