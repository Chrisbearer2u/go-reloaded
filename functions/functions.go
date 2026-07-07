package functions

import (
	"strconv"
	"strings"
	"unicode"
)

//
// ===============================
// Main Pipeline
// ===============================
//

// ProcessText runs the full text-processing pipeline.
func ProcessText(text string) string {

	// Step 1: Split the input text into tokens
	words := tokenize(text)

	// Step 2: Convert numbers followed by (hex) or (bin)
	words = convertNumbers(words)

	// Step 3: Apply case transformation tags
	words = applyCaseTags(words)

	// Step 4: Attach punctuation to the previous word
	words = attachPunctuation(words)

	// Step 5: Merge quoted text
	words = mergeQuotes(words)

	// Step 6: Fix "a" → "an"
	words = correctArticles(words)

	// Join tokens into the final string
	return strings.Join(words, " ")
}

//
// ===============================
// Tokenization
// ===============================
//

// tokenize splits the text into words, punctuation, quotes and tags.
func tokenize(text string) []string {

	var tokens []string
	var buf strings.Builder

	// Helper function to flush buffered word
	flush := func() {
		if buf.Len() > 0 {
			tokens = append(tokens, buf.String())
			buf.Reset()
		}
	}

	for i := 0; i < len(text); i++ {

		ch := text[i]

		switch {

		// Detect tag start: (up), (cap,2), (hex)
		case ch == '(':
			flush()

			j := i
			for j < len(text) && text[j] != ')' {
				j++
			}

			if j < len(text) {
				tokens = append(tokens, text[i:j+1])
				i = j
			}

		// Handle spaces
		case ch == ' ':
			flush()

		// Handle quotes
		case ch == '\'':
			flush()
			tokens = append(tokens, "'")

		// Handle punctuation groups (..., !!, ?!)
		case strings.ContainsRune(".,!?:;", rune(ch)):
			flush()

			j := i
			for j < len(text) && strings.ContainsRune(".,!?:;", rune(text[j])) {
				j++
			}

			tokens = append(tokens, text[i:j])
			i = j - 1

		// Default: build normal word
		default:
			buf.WriteByte(ch)
		}
	}

	// Flush last buffered word
	flush()

	return tokens
}

//
// ===============================
// Number Conversion
// ===============================
//

// convertNumbers converts numbers followed by (hex) or (bin).
func convertNumbers(words []string) []string {

	var result []string

	for i := 0; i < len(words); i++ {

		// Check next token for conversion tag
		if i+1 < len(words) {

			base := numberBase(words[i+1])

			// If valid numeric base exists
			if base != 0 {

				// Attempt conversion
				if n, err := strconv.ParseInt(words[i], base, 64); err == nil {

					// Append converted decimal number
					result = append(result, strconv.FormatInt(n, 10))

					i++ // Skip tag
					continue
				}
			}
		}

		result = append(result, words[i])
	}

	return result
}

// numberBase identifies conversion base.
func numberBase(tag string) int {

	switch tag {

	case "(hex)":
		return 16

	case "(bin)":
		return 2
	}

	return 0
}

//
// ===============================
// Case Transformations
// ===============================
//

// applyCaseTags processes (up), (low), (cap) transformations.
func applyCaseTags(words []string) []string {

	var result []string

	for _, w := range words {

		tag, count, ok := parseCaseTag(w)

		// If not a case tag
		if !ok {
			result = append(result, w)
			continue
		}

		// Determine start index for transformation
		start := len(result) - count
		if start < 0 {
			start = 0
		}

		// Apply transformation
		for i := start; i < len(result); i++ {
			result[i] = transformCase(tag, result[i])
		}
	}

	return result
}

// parseCaseTag extracts tag name and count.
func parseCaseTag(token string) (string, int, bool) {

	if !strings.HasPrefix(token, "(") || !strings.HasSuffix(token, ")") {
		return "", 0, false
	}

	content := token[1 : len(token)-1]

	if !(strings.HasPrefix(content, "up") ||
		strings.HasPrefix(content, "low") ||
		strings.HasPrefix(content, "cap")) {
		return "", 0, false
	}

	parts := strings.Split(content, ",")

	tag := parts[0]
	count := 1

	if len(parts) == 2 {
		if n, err := strconv.Atoi(strings.TrimSpace(parts[1])); err == nil {
			count = n
		}
	}

	return tag, count, true
}

// transformCase applies the correct transformation.
func transformCase(tag, word string) string {

	switch tag {

	case "up":
		return strings.ToUpper(word)

	case "low":
		return strings.ToLower(word)

	case "cap":
		return capitalize(word)
	}

	return word
}

// capitalize makes the first letter uppercase.
func capitalize(s string) string {

	r := []rune(strings.ToLower(s))

	if len(r) > 0 {
		r[0] = unicode.ToUpper(r[0])
	}

	return string(r)
}

//
// ===============================
// Punctuation
// ===============================
//

// attachPunctuation attaches punctuation to previous word.
func attachPunctuation(words []string) []string {

	var result []string

	for _, w := range words {

		if isPunctuation(w) && len(result) > 0 {

			result[len(result)-1] += w

		} else {

			result = append(result, w)
		}
	}

	return result
}

// isPunctuation checks if token is punctuation.
func isPunctuation(s string) bool {

	if s == "" {
		return false
	}

	for _, r := range s {

		if !strings.ContainsRune(".,!?:;", r) {
			return false
		}
	}

	return true
}

//
// ===============================
// Quotes
// ===============================
//

// mergeQuotes joins tokens between apostrophes.
func mergeQuotes(words []string) []string {

	var result []string

	for i := 0; i < len(words); i++ {

		if words[i] != "'" {
			result = append(result, words[i])
			continue
		}

		j := i + 1
		var content []string

		// Collect words until closing quote
		for j < len(words) && words[j] != "'" {
			content = append(content, words[j])
			j++
		}

		// If closing quote exists
		if j < len(words) {

			result = append(result, "'"+strings.Join(content, " ")+"'")
			i = j
			continue
		}

		result = append(result, "'")
	}

	return result
}

//
// ===============================
// Article Correction
// ===============================
//

// correctArticles converts "a" to "an" when needed.
func correctArticles(words []string) []string {

	for i := 0; i < len(words)-1; i++ {

		if !strings.EqualFold(words[i], "a") {
			continue
		}

		if len(words[i+1]) == 0 {
			continue
		}

		first := strings.ToLower(string(words[i+1][0]))

		if strings.ContainsAny(first, "aeiouh") {

			if words[i] == "A" {
				words[i] = "An"
			} else {
				words[i] = "an"
			}
		}
	}

	return words
}
