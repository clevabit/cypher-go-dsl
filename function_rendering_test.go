package cypher

import (
	"testing"
)

func TestCountInReturnClause(t *testing.T) {
	userNode := ANode("User").NamedByString("u")
	statement, err := Match(userNode).Returning(Count(userNode)).Build()
	if err != nil {
		t.Errorf("error when rendering statement: %s", err)
	}
	query, _ := NewRenderer().Render(statement)
	if query != "MATCH (u:`User`) RETURN Count(u)" {
		t.Errorf("query is not MatchPhrase:\n %s", query)
	}
}
