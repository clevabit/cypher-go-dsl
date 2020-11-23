package cypher_go_dsl


func NewNode(primaryLabel string, additionalLabel ...string) Node  {
	var labels = make([]string, 0)
	labels = append(labels, primaryLabel)
	labels = append(labels, additionalLabel...)
	return Node{
		labels: labels,
	}
}
func MatchCondition(element ...PatternElement ) OngoingReadingWithoutWhere {
	return NewDefaultBuilder().Match(element...)
}