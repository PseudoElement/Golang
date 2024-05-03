package chats_queries

import (
	"fmt"
	"strings"
)

func (cq *ChatsQueries) getMembersArrayToDB(members ...string) string {
	joined := strings.Join(members, ", ")
	result := fmt.Sprintf("ARRAY[%v]", joined)
	return result
}
