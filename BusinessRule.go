package rule

type BusinessRule struct {
	// Context (data) for the rule
	Context map[string]interface{}
	// Callback for condition
	conditionCallback func(ctx map[string]interface{}) bool
}

// Shortcut to initialize the rule with context
func NewBusinessRule(context map[string]interface{}) *BusinessRule {
	o := &BusinessRule{}
	o.Context = context
	return o
}

// Sets the context for the rule
func (br *BusinessRule) SetContext(context map[string]interface{}) *BusinessRule {
	br.Context = context
	return br
}

// Sets the condition for the rule
func (br *BusinessRule) SetCondition(conditionCallback func(ctx map[string]interface{}) bool) *BusinessRule {
	br.conditionCallback = conditionCallback
	return br
}

// Returns true if the rule evaluates to false
func (br *BusinessRule) Fails() bool {
	return !br.Evaluate()
}

// Returns true if the rule evaluates to true
func (br *BusinessRule) Passes() bool {
	return br.Evaluate()
}

// Evaluates the condition and returns the result
func (br *BusinessRule) Evaluate() bool {
	return br.conditionCallback(br.Context)
}
