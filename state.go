package main

import (
	"fmt"
	"strings"
)

func buildState() {
	usernames = splitSliceAtDelimiter(fullUsernames, ":", 0)
	usernameIDs = splitSliceAtDelimiter(fullUsernames, ":", 1)
	usernameBodies = buildGraphQLSingle()
	activeSession = createActiveSession()
	usernameChangeRequests = buildUsernameChangeRequests()
	usernameMatches = make([][]byte, len(usernames))
	for i, username := range usernames {
		usernameMatches[i] = []byte(fmt.Sprintf(`"username":"%s"`, username))
	}
}

func createActiveSession() []string {
	for _, session := range activeSessions {
		info := strings.Split(session, ":")
		return append(info[:4], strings.Join(info[4:], ":"))
	}
	panic("Please restart A/C with at least 1 active Instagram session")
}

func buildGraphQLSingle() []string {
	bodies := make([]string, 0, len(usernames))
	for _, username := range usernames {
		bodies = append(bodies, fmt.Sprintf(`https://i.instagram.com/graphql_www?doc_id=6881983411865519&variables={"username":"%s"}`, username))
	}
	return bodies
}

func buildGraphQLClientRedirect(batch []string) string {
	ids := strings.Join(batch, `\",\"credential_type\":\"none\",\"token\":\"\"},{\"uid\":\"`)
	return fmt.Sprintf(`https://i.instagram.com/graphql_www?doc_id=7765850536785467&variables={"input":{"params":"{\"account_list\":[{\"uid\":\"%s\",\"credential_type\":\"none\",\"token\":\"\"}]}"}}`, ids)
}

func buildGraphQLOneTap(batch []string) string {
	return fmt.Sprintf(`https://i.instagram.com/graphql/query?__a=1&doc_id=24047751574823150&variables={"one_tap":true,"one_tap_user_ids":["%s"]}`, strings.Join(batch, `","`))
}

func buildGraphQLSuggestions(batch []string) string {
	if len(batch) == 0 {
		return ""
	}
	return fmt.Sprintf(`doc_id=25391252800555418&__a=1&variables={"input":{"contactpoint":{"sensitive_string_value":"%s@example.com"}}}`, batch[0])
}
