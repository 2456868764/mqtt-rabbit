package errcode

var (
	NoRuleError      = Code(4001, "规则集不存在")
	RuleExistedError = Code(4001, "规则集名称已经存在，请换个名称")
)
