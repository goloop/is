package is

import (
	"regexp"
	"strings"
)

var (
	slugPattern = `[\p{L}\p{N}]{1,}([\p{L}\p{N}.-]{1,}[\p{L}\p{N}]{1,})?`
	localRegex  = regexp.MustCompile(`^` + slugPattern + `$`)
	domainRegex = regexp.MustCompile(`^` + slugPattern + `\.[\p{L}\p{N}]{2,}$`)
)

// Email returns true if email has valid format.
func Email(email string) bool {
	// Check the overall length.
	if len(email) > 254 {
		return false
	}

	// Split the email into local and domain parts.
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}
	local, domain := parts[0], parts[1]

	// Check the length of the local part.
	if len(local) > 64 || len(local) < 1 {
		return false
	}

	// Check if both parts are valid.
	if !localRegex.MatchString(local) || !domainRegex.MatchString(domain) {
		return false
	}

	return true
}

/*
package is

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/goloop/g"
)

var (
	chunkRegex = regexp.MustCompile(`^[\p{L} 0-9.\-\\@,?]+$`)
	hostRegex  = regexp.MustCompile(
		`^(?i:[\p{L}0-9]([\p{L}0-9\-]{0,61}[\p{L}0-9])?\.)*` +
			`[\p{L}0-9]([\p{L}0-9\-]{0,61}[\p{L}0-9])?\.[\p{L}]{2,8}$`)
	emailStrictRegex = regexp.MustCompile(
		`^[\p{L}0-9]{1,}([\p{L}0-9.-]+[\p{L}0-9]{1,}|[\p{L}0-9]+)?` +
			`@[\p{L}0-9.\-]+\.[\p{L}]{2,}$`)
	emailExpandedRegex = regexp.MustCompile(
		`(?m)^(((((((((\s? +)?(\(((\s? +)?(([!-'*-[\]-~]*)|` +
			`(\\([ -~]|\s))))*(\s? +)?\)))(\s? +)?)|(\s? +))?` +
			`([\p{L}0-9!#-'*+\/=?^_\x60{|}~-])+((((\s? +)?(\(((\s? +)?` +
			`(([!-'*-[\]-~]*)|(\\([ -~]|\s))))*(\s? +)?\)))(\s? +)?)|` +
			`(\s? +))?)|(((((\s? +)?(\(((\s? +)?(([!-'*-[\]-~]*)|` +
			`(\\([ -~]|\s))))*(\s? +)?\)))(\s? +)?)|(\s? +))?"((\s? +)?` +
			`(([!#-[\]-~])|(\\([ -~]|\s))))*(\s? +)?"))?)?(((((\s? +)?` +
			`(\(((\s? +)?(([!-'*-[\]-~]*)|(\\([ -~]|\s))))*(\s? +)?\)))` +
			`(\s? +)?)|(\s? +))?<(((((((\s? +)?(\(((\s? +)?(([!-'*-[\]-~]*)|` +
			`(\\([ -~]|\s))))*(\s? +)?\)))(\s? +)?)|(\s? +))?` +
			`(([\p{L}0-9!#-'*+\/=?^_\x60{|}~-])+` +
			`(\.([\p{L}0-9!#-'*+\/=?^_\x60{|}~-])+)*)` +
			`((((\s? +)?(\(((\s? +)?(([!-'*-[\]-~]*)|(\\([ -~]|\s))))*` +
			`(\s? +)?\)))(\s? +)?)|(\s? +))?)|(((((\s? +)?(\(((\s? +)?` +
			`(([!-'*-[\]-~]*)|(\\([ -~]|\s))))*(\s? +)?\)))(\s? +)?)|` +
			`(\s? +))?"((\s? +)?(([!#-[\]-~])|(\\([ -~]|\s))))*(\s? +)?"))` +
			`@((((((\s? +)?(\(((\s? +)?(([!-'*-[\]-~]*)|(\\([ -~]|\s))))*` +
			`(\s? +)?\)))(\s? +)?)|(\s? +))?(([\p{L}0-9!#-'*+\/=?^_` +
			`\x60{|}~-])+(\.([\p{L}0-9!#-'*+\/=?^_\x60{|}~-])+)*)((((\s? +)` +
			`?(\(((\s? +)?(([!-'*-[\]-~]*)|(\\([ -~]|\s))))*(\s? +)?\)))` +
			`(\s? +)?)|(\s? +))?)|(((((\s? +)?(\(((\s? +)?(([!-'*-[\]-~]*)|` +
			`(\\([ -~]|\s))))*(\s? +)?\)))(\s? +)?)|(\s? +))?\[((\s? +)?` +
			`([!-Z^-~]))*(\s? +)?\]((((\s? +)?(\(((\s? +)?(([!-'*-[\]-~]*)|` +
			`(\\([ -~]|\s))))*(\s? +)?\)))(\s? +)?)|(\s? +))?)))>((((\s? +)?` +
			`(\(((\s? +)?(([!-'*-[\]-~]*)|(\\([ -~]|\s))))*(\s? +)?\)))` +
			`(\s? +)?)|(\s? +))?))|(((((((\s? +)?(\(((\s? +)?` +
			`(([!-'*-[\]-~]*)|(\\([ -~]|\s))))*(\s? +)?\)))(\s? +)?)|` +
			`(\s? +))?(([\p{L}0-9!#-'*+\/=?^_\x60{|}~-])+` +
			`(\.([\p{L}0-9!#-'*+\/=?^_\x60{|}~-])+)*)((((\s? +)?(\(((\s? +)` +
			`?(([!-'*-[\]-~]*)|(\\([ -~]|\s))))*(\s? +)?\)))(\s? +)?)|` +
			`(\s? +))?)|(((((\s? +)?(\(((\s? +)?(([!-'*-[\]-~]*)|(\\([ -~]|` +
			`\s))))*(\s? +)?\)))(\s? +)?)|(\s? +))?"((\s? +)?(([!#-[\]-~])|` +
			`(\\([ -~]|\s))))*(\s? +)?"))@((((((\s? +)?(\(((\s? +)?` +
			`(([!-'*-[\]-~]*)|(\\([ -~]|\s))))*(\s? +)?\)))(\s? +)?)|` +
			`(\s? +))?(([\p{L}0-9!#-'*+\/=?^_\x60{|}~-])+` +
			`(\.([\p{L}0-9!#-'*+\/=?^_\x60{|}~-])+)*)((((\s? +)?` +
			`(\(((\s? +)?(([!-'*-[\]-~]*)|(\\([ -~]|\s))))*(\s? +)?\)))` +
			`(\s? +)?)|(\s? +))?)|(((((\s? +)?(\(((\s? +)?(([!-'*-[\]-~]*)|` +
			`(\\([ -~]|\s))))*(\s? +)?\)))(\s? +)?)|(\s? +))?\[((\s? +)?` +
			`([!-Z^-~]))*(\s? +)?\]((((\s? +)?(\(((\s? +)?(([!-'*-[\]-~]*)|` +
			`(\\([ -~]|\s))))*(\s? +)?\)))(\s? +)?)|(\s? +))?))))$`)
)

// isHostValid checks if the host is valid.
func isHostValid(host string) bool {
	return hostRegex.MatchString(host)
}

// getQuotesChunks returns all the quoted strings in the given string.
func getQuotesChunks(s string) ([]string, error) {
	re := regexp.MustCompile(`"(.*?)"`)
	matches := re.FindAllStringSubmatch(s, -1)

	if len(matches) == 0 {
		return []string{}, nil
	}

	results := make([]string, 0, len(matches))
	for _, match := range matches {
		results = append(results, match[1])
	}

	return results, nil
}

// The emailExpanded checks if the string is an email.
func emailExpanded(email string) bool {
	// At least 1 char before @, 3 chars at least after.
	i := strings.LastIndexByte(email, '@')
	if i <= 0 || i > len(email)-4 {
		return false
	}

	// Check for a valid local and domain part.
	local, host := email[:i], email[i+1:]
	if len(local) > 64 || len(host) < 4 {
		return false
	}

	// Check for valid local and domain part.
	if strings.HasPrefix(local, " ") || strings.HasSuffix(local, " ") {
		return false
	}

	if strings.HasPrefix(host, " ") || strings.HasSuffix(host, " ") {
		return false
	}

	if !isHostValid(host) {
		return false
	}

	chunks, err := getQuotesChunks(local)
	if err != nil {
		return false
	}

	// Check if there's at least one non-space character inside a quoted string.
	for _, chunk := range chunks {
		if len(chunk) == 0 || strings.TrimSpace(chunk) == "" {
			return false
		}

		if !chunkRegex.MatchString(chunk) {
			return false
		}

		// Replace quoted chunks with a valid placeholder.
		email = strings.ReplaceAll(email, fmt.Sprintf("\"%s\"", chunk), "go")
	}

	// Check if unquoted parts of the local part contain consecutive,
	// leading or trailing dots.
	unquotedParts := strings.Split(email[:strings.IndexByte(email, '@')], "\"")
	for _, part := range unquotedParts {
		if strings.HasPrefix(part, ".") ||
			strings.HasSuffix(part, ".") || strings.Contains(part, "..") {
			return false
		}
	}

	return emailExpandedRegex.MatchString(email)
}

// The emailStrict checks if the string is an email.
func emailStrict(email string) bool {
	// At least 1 char before @, 3 chars at least after.
	i := strings.LastIndexByte(email, '@')
	if i <= 0 || i > len(email)-4 {
		return false
	}

	// Check for a valid local and domain part.
	local, host := email[:i], email[i+1:]
	if len(local) > 64 || len(host) < 4 {
		return false
	}

	// Check for valid local and domain part.
	if strings.HasPrefix(local, " ") || strings.HasSuffix(local, " ") {
		return false
	}

	if strings.HasPrefix(host, " ") || strings.HasSuffix(host, " ") {
		return false
	}

	if !isHostValid(host) {
		return false
	}

	return emailStrictRegex.MatchString(email)
}

func Email(email string, extended ...bool) bool {
	if l := len(email); l < 3 || l > 254 {
		return false
	}

	// Convert to lowercase.
	email = strings.ToLower(email)

	if g.All(extended...) {
		return emailExpanded(email)
	}

	return emailStrict(email)
}
*/
