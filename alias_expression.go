package cypher

import (
	"errors"
)

type AliasedExpression struct {
	ExpressionContainer
	delegate Expression
	alias    string
	key      string
	notNil   bool
	err      error
}

func AliasedExpressionCreate(delegate Expression, alias string) AliasedExpression {
	if delegate == nil || !delegate.isNotNil() {
		return AliasedExpressionError(errors.New("expression to alias can't be nil"))
	}
	if delegate.GetError() != nil {
		return AliasedExpressionError(delegate.GetError())
	}
	if alias == "" {
		return AliasedExpressionError(errors.New("the alias may not be empty"))
	}
	a := AliasedExpression{
		delegate: delegate,
		alias:    alias,
		notNil:   true,
	}
	a.key = getAddress(&a)
	a.ExpressionContainer = ExpressionWrap(a)
	return a
}

func AliasedExpressionError(err error) AliasedExpression {
	return AliasedExpression{
		err: err,
	}
}

func (aliased AliasedExpression) GetError() error {
	return aliased.err
}

func (aliased AliasedExpression) isNotNil() bool {
	return aliased.notNil
}

func (aliased AliasedExpression) GetExpressionType() ExpressionType {
	return EXPRESSION
}

func (aliased AliasedExpression) Aliased(newAlias string) AliasedExpression {
	if newAlias == "" {
		return AliasedExpressionError(errors.New("the alias may not be empty"))
	}
	return AliasedExpressionCreate(aliased.delegate, newAlias)
}

func (aliased AliasedExpression) accept(visitor *CypherRenderer) {
	visitor.enter(aliased)
	NameOrExpression(aliased.delegate).accept(visitor)
	visitor.leave(aliased)
}

func (aliased AliasedExpression) getKey() string {
	return aliased.key
}

func (aliased AliasedExpression) enter(renderer *CypherRenderer) {
	if _, visited := renderer.visitableToAliased[aliased.delegate.getKey()]; visited {
		renderer.append(EscapeIfNecessary(aliased.alias))
	}
}

func (aliased AliasedExpression) leave(renderer *CypherRenderer) {
	if _, visited := renderer.visitableToAliased[aliased.delegate.getKey()]; !visited {
		renderer.append(" AS ")
		renderer.append(EscapeIfNecessary(aliased.alias))
	}
}

func (aliased AliasedExpression) GetAlias() string {
	return aliased.alias
}

func (aliased AliasedExpression) AsName() SymbolicName {
	return SymbolicNameCreate(aliased.alias)
}
