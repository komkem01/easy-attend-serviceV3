// Package console implements CLI commands for the application
package console

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

// Column represents a SQL table column definition
type Column struct {
	Name     string
	Type     string
	Nullable string
	Comment  string
	IsPK     bool
}

// SQLSyntaxError represents a syntax validation error
type SQLSyntaxError struct {
	Line    int
	Column  int
	Message string
	Context string
}

// FuzzyMatch represents a potential typo correction with confidence score
type FuzzyMatch struct {
	Original   string
	Correction string
	Distance   int
	Confidence float64
}

var (
	// Comprehensive PostgreSQL keywords and data types dictionary
	postgresqlKeywords = map[string]string{
		// DDL Keywords
		"CREATE":   "CREATE",
		"ALTER":    "ALTER",
		"DROP":     "DROP",
		"TRUNCATE": "TRUNCATE",
		"RENAME":   "RENAME",
		"ADD":      "ADD",
		"MODIFY":   "MODIFY",
		"CHANGE":   "CHANGE",

		// Table and Structure
		"TABLE":      "TABLE",
		"INDEX":      "INDEX",
		"VIEW":       "VIEW",
		"DATABASE":   "DATABASE",
		"SCHEMA":     "SCHEMA",
		"SEQUENCE":   "SEQUENCE",
		"TRIGGER":    "TRIGGER",
		"FUNCTION":   "FUNCTION",
		"PROCEDURE":  "PROCEDURE",
		"COLUMN":     "COLUMN",
		"CONSTRAINT": "CONSTRAINT",

		// DML Keywords
		"SELECT":    "SELECT",
		"INSERT":    "INSERT",
		"INTO":      "INTO",
		"UPDATE":    "UPDATE",
		"DELETE":    "DELETE",
		"MERGE":     "MERGE",
		"UPSERT":    "UPSERT",
		"WITH":      "WITH",
		"RECURSIVE": "RECURSIVE",

		// Query Clauses
		"FROM":   "FROM",
		"WHERE":  "WHERE",
		"HAVING": "HAVING",
		"GROUP":  "GROUP",
		"ORDER":  "ORDER",
		"BY":     "BY",
		"LIMIT":  "LIMIT",
		"OFFSET": "OFFSET",
		"FETCH":  "FETCH",
		"FIRST":  "FIRST",
		"NEXT":   "NEXT",
		"ONLY":   "ONLY",

		// Join Types
		"JOIN":    "JOIN",
		"INNER":   "INNER",
		"LEFT":    "LEFT",
		"RIGHT":   "RIGHT",
		"FULL":    "FULL",
		"OUTER":   "OUTER",
		"CROSS":   "CROSS",
		"NATURAL": "NATURAL",
		"ON":      "ON",
		"USING":   "USING",

		// Set Operations
		"UNION":     "UNION",
		"INTERSECT": "INTERSECT",
		"EXCEPT":    "EXCEPT",
		"ALL":       "ALL",
		"DISTINCT":  "DISTINCT",

		// Data Types
		"INTEGER":     "INTEGER",
		"INT":         "INT",
		"INT2":        "INT2",
		"INT4":        "INT4",
		"INT8":        "INT8",
		"BIGINT":      "BIGINT",
		"SMALLINT":    "SMALLINT",
		"TINYINT":     "TINYINT",
		"SERIAL":      "SERIAL",
		"BIGSERIAL":   "BIGSERIAL",
		"SMALLSERIAL": "SMALLSERIAL",

		"VARCHAR":   "VARCHAR",
		"CHAR":      "CHAR",
		"CHARACTER": "CHARACTER",
		"TEXT":      "TEXT",
		"CLOB":      "CLOB",

		"DECIMAL": "DECIMAL",
		"NUMERIC": "NUMERIC",
		"FLOAT":   "FLOAT",
		"FLOAT4":  "FLOAT4",
		"FLOAT8":  "FLOAT8",
		"DOUBLE":  "DOUBLE",
		"REAL":    "REAL",
		"MONEY":   "MONEY",

		"BOOLEAN": "BOOLEAN",
		"BOOL":    "BOOL",
		"BIT":     "BIT",

		"DATE":        "DATE",
		"TIME":        "TIME",
		"TIMESTAMP":   "TIMESTAMP",
		"TIMESTAMPTZ": "TIMESTAMPTZ",
		"DATETIME":    "DATETIME",
		"INTERVAL":    "INTERVAL",

		"BLOB":      "BLOB",
		"BINARY":    "BINARY",
		"VARBINARY": "VARBINARY",
		"BYTEA":     "BYTEA",

		"JSON":  "JSON",
		"JSONB": "JSONB",
		"XML":   "XML",
		"UUID":  "UUID",

		"ARRAY":      "ARRAY",
		"GEOMETRY":   "GEOMETRY",
		"GEOGRAPHY":  "GEOGRAPHY",
		"POINT":      "POINT",
		"POLYGON":    "POLYGON",
		"LINESTRING": "LINESTRING",

		// Constraints and Modifiers
		"PRIMARY":    "PRIMARY",
		"FOREIGN":    "FOREIGN",
		"UNIQUE":     "UNIQUE",
		"CHECK":      "CHECK",
		"DEFAULT":    "DEFAULT",
		"NULL":       "NULL",
		"NOT":        "NOT",
		"KEY":        "KEY",
		"REFERENCES": "REFERENCES",
		"CASCADE":    "CASCADE",
		"RESTRICT":   "RESTRICT",
		"SET":        "SET",
		"ACTION":     "ACTION",

		"AUTO":      "AUTO",
		"INCREMENT": "INCREMENT",
		"IDENTITY":  "IDENTITY",
		"GENERATED": "GENERATED",
		"ALWAYS":    "ALWAYS",
		"STORED":    "STORED",
		"VIRTUAL":   "VIRTUAL",

		"DEFERRABLE": "DEFERRABLE",
		"INITIALLY":  "INITIALLY",
		"DEFERRED":   "DEFERRED",
		"IMMEDIATE":  "IMMEDIATE",

		// Logical Operators
		"AND":     "AND",
		"OR":      "OR",
		"IN":      "IN",
		"EXISTS":  "EXISTS",
		"BETWEEN": "BETWEEN",
		"LIKE":    "LIKE",
		"ILIKE":   "ILIKE",
		"SIMILAR": "SIMILAR",
		"REGEXP":  "REGEXP",
		"IS":      "IS",
		"AS":      "AS",

		// Case Expressions
		"CASE": "CASE",
		"WHEN": "WHEN",
		"THEN": "THEN",
		"ELSE": "ELSE",
		"END":  "END",

		// Aggregate Functions
		"COUNT":      "COUNT",
		"SUM":        "SUM",
		"AVG":        "AVG",
		"MIN":        "MIN",
		"MAX":        "MAX",
		"STDDEV":     "STDDEV",
		"VARIANCE":   "VARIANCE",
		"ARRAY_AGG":  "ARRAY_AGG",
		"STRING_AGG": "STRING_AGG",
		"JSON_AGG":   "JSON_AGG",

		// String Functions
		"SUBSTRING":  "SUBSTRING",
		"LENGTH":     "LENGTH",
		"UPPER":      "UPPER",
		"LOWER":      "LOWER",
		"TRIM":       "TRIM",
		"LTRIM":      "LTRIM",
		"RTRIM":      "RTRIM",
		"CONCAT":     "CONCAT",
		"REPLACE":    "REPLACE",
		"SPLIT_PART": "SPLIT_PART",

		// Null Functions
		"COALESCE": "COALESCE",
		"NULLIF":   "NULLIF",
		"GREATEST": "GREATEST",
		"LEAST":    "LEAST",

		// Transaction Control
		"BEGIN":       "BEGIN",
		"COMMIT":      "COMMIT",
		"ROLLBACK":    "ROLLBACK",
		"SAVEPOINT":   "SAVEPOINT",
		"RELEASE":     "RELEASE",
		"TRANSACTION": "TRANSACTION",
		"WORK":        "WORK",

		// Isolation Levels
		"ISOLATION":    "ISOLATION",
		"LEVEL":        "LEVEL",
		"READ":         "READ",
		"WRITE":        "WRITE",
		"SERIALIZABLE": "SERIALIZABLE",
		"REPEATABLE":   "REPEATABLE",
		"COMMITTED":    "COMMITTED",
		"UNCOMMITTED":  "UNCOMMITTED",

		// Window Functions
		"WINDOW":    "WINDOW",
		"OVER":      "OVER",
		"PARTITION": "PARTITION",
		"ROWS":      "ROWS",
		"RANGE":     "RANGE",
		"PRECEDING": "PRECEDING",
		"FOLLOWING": "FOLLOWING",
		"CURRENT":   "CURRENT",
		"ROW":       "ROW",
		"UNBOUNDED": "UNBOUNDED",

		// Permissions and Security
		"GRANT":      "GRANT",
		"REVOKE":     "REVOKE",
		"PRIVILEGES": "PRIVILEGES",
		"ROLE":       "ROLE",
		"USER":       "USER",
		"PASSWORD":   "PASSWORD",
		"OWNER":      "OWNER",
		"PUBLIC":     "PUBLIC",
		"USAGE":      "USAGE",
		"EXECUTE":    "EXECUTE",
		"CONNECT":    "CONNECT",

		// Miscellaneous
		"COMMENT":    "COMMENT",
		"COLLATE":    "COLLATE",
		"CHARSET":    "CHARSET",
		"ENGINE":     "ENGINE",
		"INHERITS":   "INHERITS",
		"TABLESPACE": "TABLESPACE",
		"VACUUM":     "VACUUM",
		"ANALYZE":    "ANALYZE",
		"EXPLAIN":    "EXPLAIN",
		"COPY":       "COPY",
		"LISTEN":     "LISTEN",
		"NOTIFY":     "NOTIFY",

		// PostgreSQL specific
		"RETURNING":    "RETURNING",
		"CONFLICT":     "CONFLICT",
		"DO":           "DO",
		"NOTHING":      "NOTHING",
		"EXCLUDED":     "EXCLUDED",
		"LATERAL":      "LATERAL",
		"MATERIALIZED": "MATERIALIZED",
		"REFRESH":      "REFRESH",
		"CONCURRENTLY": "CONCURRENTLY",
		"IF":           "IF",
		"EXTENSION":    "EXTENSION",
		"TYPE":         "TYPE",
		"ENUM":         "ENUM",
		"DOMAIN":       "DOMAIN",
		"CAST":         "CAST",
	}

	// Common multi-word SQL phrases that should be treated as units
	multiWordKeywords = map[string]string{
		"NOT NULL":        "NOT NULL",
		"PRIMARY KEY":     "PRIMARY KEY",
		"FOREIGN KEY":     "FOREIGN KEY",
		"UNIQUE KEY":      "UNIQUE KEY",
		"ALTER TABLE":     "ALTER TABLE",
		"CREATE TABLE":    "CREATE TABLE",
		"DROP TABLE":      "DROP TABLE",
		"CREATE INDEX":    "CREATE INDEX",
		"DROP INDEX":      "DROP INDEX",
		"ADD COLUMN":      "ADD COLUMN",
		"DROP COLUMN":     "DROP COLUMN",
		"ORDER BY":        "ORDER BY",
		"GROUP BY":        "GROUP BY",
		"INNER JOIN":      "INNER JOIN",
		"LEFT JOIN":       "LEFT JOIN",
		"RIGHT JOIN":      "RIGHT JOIN",
		"FULL JOIN":       "FULL JOIN",
		"FULL OUTER JOIN": "FULL OUTER JOIN",
		"CROSS JOIN":      "CROSS JOIN",
		"AUTO INCREMENT":  "AUTO INCREMENT",
		"INSERT INTO":     "INSERT INTO",
		"DELETE FROM":     "DELETE FROM",
		"ON CONFLICT":     "ON CONFLICT",
		"DO NOTHING":      "DO NOTHING",
		"DO UPDATE":       "DO UPDATE",
		"IF EXISTS":       "IF EXISTS",
		"IF NOT EXISTS":   "IF NOT EXISTS",
		"SET DEFAULT":     "SET DEFAULT",
		"SET NULL":        "SET NULL",
	}

	// Special short-form corrections that might not match well with fuzzy matching
	shortFormCorrections = map[string]string{
		// Short data type corrections
		"VCH":   "VARCHAR",
		"VARCH": "VARCHAR",
		"VARC":  "VARCHAR",
		"VARH":  "VARCHAR",
		"KE":    "KEY",
		"UID":   "UUID",
		"NUL":   "NULL",
		"ULL":   "NULL",
		"TEX":   "TEXT",
		"TXT":   "TEXT",
		"JSO":   "JSON",
		"BOO":   "BOOLEAN",
		"BOOL":  "BOOL",
		"INT":   "INT",
		"INTG":  "INTEGER",
		"FLT":   "FLOAT",
		"DEC":   "DECIMAL",
		"NUM":   "NUMERIC",
		"TS":    "TIMESTAMP",
		"DT":    "DATETIME",

		// Short keyword corrections
		"AD":   "ADD", // Common typo: ALTER TABLE ... ad COLUMN
		"SEL":  "SELECT",
		"INS":  "INSERT",
		"UPD":  "UPDATE",
		"DEL":  "DELETE",
		"CR":   "CREATE",
		"ALT":  "ALTER",
		"DR":   "DROP",
		"TBL":  "TABLE",
		"IDX":  "INDEX",
		"COL":  "COLUMN",
		"PK":   "PRIMARY",
		"FK":   "FOREIGN",
		"UK":   "UNIQUE",
		"REF":  "REFERENCES",
		"DEF":  "DEFAULT",
		"COM":  "COMMENT",
		"COMM": "COMMENT",
		"JOI":  "JOIN",
		"WHR":  "WHERE",
		"FRM":  "FROM",
		"ORD":  "ORDER",
		"GRP":  "GROUP",
		"HAV":  "HAVING",
		"LIM":  "LIMIT",
		"OFF":  "OFFSET",

		// Context-specific corrections
		"SE":    "SET",    // Common in UPDATE ... se column = value
		"I":     "IS",     // Common in WHERE ... i NOT NULL
		"O":     "ON",     // Common in JOIN ... o condition
		"VALUE": "VALUES", // Common typo: INSERT ... value
		"TRU":   "TRUE",   // Common typo: boolean value
		"FALS":  "FALSE",  // Common typo: boolean value
	}

	// Protected identifiers that should not be corrected
	protectedIdentifiers = []string{
		"USERS", "USER", "PRODUCTS", "ORDERS", "CUSTOMERS", "ITEMS",
		"NAME", "EMAIL", "PASSWORD", "ADDRESS", "PHONE", "PRICE", "QUANTITY",
		"CREATED", "UPDATED", "ACTIVE", "STATUS", "TYPE", "CATEGORY",
		"FIRST", "LAST", "MIDDLE", "AGE", "DATE", "TIME", "YEAR", "MONTH",
		"TOTAL", "AMOUNT", "VALUE", "COST", "RATE", "LEVEL", "RANK",
		"CODE", "TITLE", "DESCRIPTION", "NOTES", "COMMENT", "TAG",
		"YOUNG", "ADULT", "OLD", "TRUE", "FALSE", "YES", "NO",
		"DATA", "INFO", "DETAILS", "RECORD", "ENTRY", "ITEM", "ELEMENT",
		"TEST", "TEMP", "DEMO", "SAMPLE", "EXAMPLE", "MOCK",
		"ID", "KEY", "REF", "LINK", "URL", "PATH", "FILE", "IMAGE",
		"ROLE", "PERMISSION", "ACCESS", "TOKEN", "SESSION", "CACHE",
		"LOG", "EVENT", "MESSAGE", "NOTIFICATION", "ALERT",
		"PRIMARY", "FOREIGN", "UNIQUE", "INDEX", // SQL constraint keywords
		"ROAD", "STREET", "AVENUE", "LANE", "WAY", "DRIVE", // Address terms
		"GROUP", "ORDER", "LIMIT", "OFFSET", "PARTITION", // SQL keywords that might be column names
		// Common names and string values that should not be fuzzy-matched
		"JOHN", "JANE", "MARY", "ROBERT", "MICHAEL", "WILLIAM", "DAVID", "RICHARD", "CHARLES", "JOSEPH",
		"DOE", "SMITH", "JOHNSON", "WILLIAMS", "BROWN", "JONES", "GARCIA", "MILLER", "DAVIS", "RODRIGUEZ",
		"ADMIN", "GUEST", "SYSTEM", "DEFAULT", "NULL", "EMPTY", "TEMP", "TEST",
	}
)

// levenshteinDistance calculates the edit distance between two strings
func levenshteinDistance(s1, s2 string) int {
	s1 = strings.ToUpper(s1)
	s2 = strings.ToUpper(s2)

	if len(s1) == 0 {
		return len(s2)
	}
	if len(s2) == 0 {
		return len(s1)
	}

	matrix := make([][]int, len(s1)+1)
	for i := range matrix {
		matrix[i] = make([]int, len(s2)+1)
	}

	// Initialize first row and column
	for i := 0; i <= len(s1); i++ {
		matrix[i][0] = i
	}
	for j := 0; j <= len(s2); j++ {
		matrix[0][j] = j
	}

	// Fill the matrix
	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			cost := 0
			if s1[i-1] != s2[j-1] {
				cost = 1
			}

			matrix[i][j] = min(
				matrix[i-1][j]+1,      // deletion
				matrix[i][j-1]+1,      // insertion
				matrix[i-1][j-1]+cost, // substitution
			)
		}
	}

	return matrix[len(s1)][len(s2)]
}

// min returns the minimum of three integers
func min(a, b, c int) int {
	if a < b && a < c {
		return a
	}
	if b < c {
		return b
	}
	return c
}

// minMax returns both min and max of two integers
func minMax(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

// calculateConfidence calculates confidence score for a fuzzy match
func calculateConfidence(original, candidate string, distance int) float64 {
	if distance == 0 {
		return 1.0
	}

	maxLen := math.Max(float64(len(original)), float64(len(candidate)))
	if maxLen == 0 {
		return 0.0
	}

	// Base similarity
	baseSimilarity := 1.0 - (float64(distance) / maxLen)

	originalUpper := strings.ToUpper(original)
	candidateUpper := strings.ToUpper(candidate)

	// Length penalty - if lengths differ significantly, reduce confidence
	lenDiff := math.Abs(float64(len(original)) - float64(len(candidate)))
	lengthPenalty := 1.0 - (lenDiff / maxLen * 0.5)

	// First character bonus - SQL keywords usually have consistent first chars
	firstCharBonus := 0.0
	if len(originalUpper) > 0 && len(candidateUpper) > 0 {
		if originalUpper[0] == candidateUpper[0] {
			firstCharBonus = 0.2
		}
	}

	// Common prefix bonus
	prefixLen := 0
	minLen := int(math.Min(float64(len(originalUpper)), float64(len(candidateUpper))))
	for i := 0; i < minLen && originalUpper[i] == candidateUpper[i]; i++ {
		prefixLen++
	}
	prefixBonus := float64(prefixLen) / maxLen * 0.3

	// Penalize if the distance is too high relative to word length
	distancePenalty := 1.0
	if len(original) <= 4 && distance > 1 {
		distancePenalty = 0.3 // Very harsh penalty for short words
	} else if len(original) <= 6 && distance > 2 {
		distancePenalty = 0.5 // Harsh penalty for medium words
	} else if distance > 3 {
		distancePenalty = 0.6 // Penalty for long distance
	}

	confidence := (baseSimilarity * lengthPenalty * distancePenalty) + firstCharBonus + prefixBonus

	// Ensure confidence is between 0 and 1
	if confidence > 1.0 {
		confidence = 1.0
	}
	if confidence < 0.0 {
		confidence = 0.0
	}

	return confidence
}

// findBestMatch finds the best fuzzy match for a word using the PostgreSQL dictionary
func findBestMatch(word string, maxDistance int, minConfidence float64) *FuzzyMatch {
	if len(word) < 3 {
		return nil
	}

	// Skip if the word is already a valid SQL keyword
	upperWord := strings.ToUpper(word)
	if _, exists := postgresqlKeywords[upperWord]; exists {
		return nil // Don't "correct" valid SQL words
	}

	var bestMatch *FuzzyMatch

	// Search through all PostgreSQL keywords
	for _, correctWord := range postgresqlKeywords {
		// Skip if candidate is exactly the same (no typo)
		if strings.EqualFold(word, correctWord) {
			continue
		}

		distance := levenshteinDistance(word, correctWord)

		// Dynamic distance threshold based on word length
		maxAllowedDistance := 1
		if len(word) >= 5 {
			maxAllowedDistance = 2
		}
		if len(word) >= 8 {
			maxAllowedDistance = maxDistance
		}

		if distance <= maxAllowedDistance {
			confidence := calculateConfidence(word, correctWord, distance)

			// Length similarity check
			minLen, maxLen := minMax(len(word), len(correctWord))
			lengthRatio := float64(minLen) / float64(maxLen)
			if lengthRatio < 0.5 {
				continue
			}

			if confidence >= minConfidence {
				if bestMatch == nil || confidence > bestMatch.Confidence ||
					(confidence == bestMatch.Confidence && distance < bestMatch.Distance) {
					bestMatch = &FuzzyMatch{
						Original:   word,
						Correction: correctWord,
						Distance:   distance,
						Confidence: confidence,
					}
				}
			}
		}
	}

	return bestMatch
}

// intelligentAutoCorrect applies intelligent auto-correction using dictionary-based fuzzy matching
func intelligentAutoCorrect(line string) (string, int) {
	fixCount := 0
	fixedLine := line

	// First, apply multi-word fixes (including partial matches like "PRIMARY KE")
	fixedLine, multiWordFixes := applyMultiWordFixes(fixedLine)
	fixCount += multiWordFixes

	// Second, apply short-form corrections (exact matches for short abbreviations)
	fixedLine, shortFormFixes := applyShortFormCorrections(fixedLine)
	fixCount += shortFormFixes

	// Third, apply single-word fuzzy matching for remaining issues
	words := extractSQLWords(fixedLine)

	for _, word := range words {
		if len(word) < 2 || isLikelyUserIdentifier(word) {
			continue
		}

		// Skip if already corrected by short-form fixes
		upperWord := strings.ToUpper(word)
		if _, exists := shortFormCorrections[upperWord]; exists {
			continue
		}

		// Skip if word is already correct
		if _, exists := postgresqlKeywords[upperWord]; exists {
			continue
		}

		// Find best match with conservative parameters
		match := findBestMatch(word, 2, 0.85)

		if match != nil && match.Original != match.Correction {
			// Apply the fix using regex with word boundaries
			pattern := `(?i)\b` + regexp.QuoteMeta(word) + `\b`
			re := regexp.MustCompile(pattern)

			if re.MatchString(fixedLine) {
				newLine := re.ReplaceAllString(fixedLine, match.Correction)
				if newLine != fixedLine {
					fixedLine = newLine
					fixCount++
					fmt.Printf("    Auto-corrected: %s → %s (confidence: %.2f)\n",
						word, match.Correction, match.Confidence)
				}
			}
		}
	}

	return fixedLine, fixCount
}

// applyShortFormCorrections applies exact short-form corrections
func applyShortFormCorrections(line string) (string, int) {
	fixCount := 0
	fixedLine := line

	for shortForm, correct := range shortFormCorrections {
		pattern := `(?i)\b` + regexp.QuoteMeta(shortForm) + `\b`
		re := regexp.MustCompile(pattern)

		if re.MatchString(fixedLine) {
			newLine := re.ReplaceAllString(fixedLine, correct)
			if newLine != fixedLine {
				fixedLine = newLine
				fixCount++
				fmt.Printf("    Short-form fix: %s → %s\n", shortForm, correct)
			}
		}
	}

	return fixedLine, fixCount
}

// applyMultiWordFixes fixes common multi-word SQL phrases and handles partial matches
func applyMultiWordFixes(line string) (string, int) {
	fixCount := 0
	fixedLine := line

	// First, handle partial multi-word fixes (like "PRIMARY KE" -> "PRIMARY KEY")
	partialMultiWordFixes := map[string]string{
		`PRIMARY\s+KE\b`:  "PRIMARY KEY",
		`FOREIGN\s+KE\b`:  "FOREIGN KEY",
		`UNIQUE\s+KE\b`:   "UNIQUE KEY",
		`ORDER\s+B\b`:     "ORDER BY",
		`GROUP\s+B\b`:     "GROUP BY",
		`INNER\s+JOI\b`:   "INNER JOIN",
		`LEFT\s+JOI\b`:    "LEFT JOIN",
		`RIGHT\s+JOI\b`:   "RIGHT JOIN",
		`FULL\s+JOI\b`:    "FULL JOIN",
		`CROSS\s+JOI\b`:   "CROSS JOIN",
		`INSERT\s+INT\b`:  "INSERT INTO",
		`DELETE\s+FRO\b`:  "DELETE FROM",
		`NOT\s+NUL\b`:     "NOT NULL",
		`ALTER\s+TABL\b`:  "ALTER TABLE",
		`CREATE\s+TABL\b`: "CREATE TABLE",
		`DROP\s+TABL\b`:   "DROP TABLE",
		`CREATE\s+INDE\b`: "CREATE INDEX",
		`DROP\s+INDE\b`:   "DROP INDEX",
		`ADD\s+COLUM\b`:   "ADD COLUMN",
		`DROP\s+COLUM\b`:  "DROP COLUMN",
	}

	for pattern, correct := range partialMultiWordFixes {
		re := regexp.MustCompile(`(?i)` + pattern)
		if re.MatchString(fixedLine) {
			newLine := re.ReplaceAllString(fixedLine, correct)
			if newLine != fixedLine {
				fixedLine = newLine
				fixCount++
				fmt.Printf("    Multi-word fix: %s\n", correct)
			}
		}
	}

	// Apply complete multi-word fixes using the dictionary
	for incorrect, correct := range multiWordKeywords {
		// Create case-insensitive pattern for multi-word phrases
		incorrectPattern := strings.ReplaceAll(incorrect, " ", `\s+`)
		pattern := `(?i)\b` + incorrectPattern + `\b`
		re := regexp.MustCompile(pattern)

		if re.MatchString(fixedLine) {
			newLine := re.ReplaceAllString(fixedLine, correct)
			if newLine != fixedLine {
				fixedLine = newLine
				fixCount++
			}
		}
	}

	// Handle common spacing issues (concatenated keywords)
	spacingFixes := map[string]string{
		"NOTNULL":     "NOT NULL",
		"PRIMARYKEY":  "PRIMARY KEY",
		"FOREIGNKEY":  "FOREIGN KEY",
		"ALTERTABLE":  "ALTER TABLE",
		"CREATETABLE": "CREATE TABLE",
		"INSERTINTO":  "INSERT INTO",
		"DELETEFROM":  "DELETE FROM",
		"ORDERBY":     "ORDER BY",
		"GROUPBY":     "GROUP BY",
		"INNERJOIN":   "INNER JOIN",
		"LEFTJOIN":    "LEFT JOIN",
		"RIGHTJOIN":   "RIGHT JOIN",
		"FULLJOIN":    "FULL JOIN",
	}

	for incorrect, correct := range spacingFixes {
		pattern := `(?i)\b` + regexp.QuoteMeta(incorrect) + `\b`
		re := regexp.MustCompile(pattern)
		if re.MatchString(fixedLine) {
			newLine := re.ReplaceAllString(fixedLine, correct)
			if newLine != fixedLine {
				fixedLine = newLine
				fixCount++
			}
		}
	}

	return fixedLine, fixCount
} // isLikelySQL checks if a word looks like it could be a SQL keyword
func isLikelySQL(word string) bool {
	word = strings.ToUpper(word)

	// Skip very short words unless they're common SQL keywords
	if len(word) < 3 {
		shortKeywords := []string{"AS", "BY", "ON", "OR", "IN", "IS"}
		for _, kw := range shortKeywords {
			if word == kw {
				return true
			}
		}
		return false
	}

	// Skip if it looks like a user-defined identifier (contains numbers, underscores in the middle)
	if strings.Contains(word, "_") && len(word) > 3 {
		return false
	}

	if regexp.MustCompile(`\d`).MatchString(word) && !regexp.MustCompile(`^(INT|FLOAT|VARCHAR)\d*$`).MatchString(word) {
		return false
	}

	// Skip if it's in quotes (string literal)
	if strings.HasPrefix(word, "'") || strings.HasPrefix(word, "\"") {
		return false
	}

	return true
}

// extractSQLWords extracts potential SQL keywords from a line
func extractSQLWords(line string) []string {
	// Remove comments
	line = regexp.MustCompile(`--.*$`).ReplaceAllString(line, "")
	line = regexp.MustCompile(`/\*.*?\*/`).ReplaceAllString(line, "")

	// Split by whitespace and common SQL delimiters
	words := regexp.MustCompile(`[^\w]+`).Split(line, -1)

	var sqlWords []string
	for _, word := range words {
		word = strings.TrimSpace(word)
		if len(word) > 0 && isLikelySQL(word) {
			sqlWords = append(sqlWords, word)
		}
	}

	return sqlWords
}

func formatSQLCMD() *cobra.Command {
	var overwrite bool
	var checkOnly bool
	var fixSyntax bool
	var noFix bool

	cmd := &cobra.Command{
		Use:   "format-sql [pattern]",
		Short: "Format SQL files with proper indentation and automatic syntax fixing",
		Long:  "Format SQL CREATE TABLE statements with aligned columns and proper indentation. Supports glob patterns like *.sql. Automatically fixes syntax errors by default.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			pattern := args[0]
			// Default to fixing syntax unless --no-fix is specified
			if !noFix {
				fixSyntax = true
			}
			return formatSQLFiles(pattern, overwrite, checkOnly, fixSyntax)
		},
	}

	cmd.Flags().BoolVarP(&overwrite, "overwrite", "w", true, "Overwrite the original file")
	cmd.Flags().BoolVarP(&checkOnly, "check", "c", false, "Only check syntax without formatting")
	cmd.Flags().BoolVarP(&noFix, "no-fix", "n", false, "Disable automatic syntax fixing")
	cmd.Flags().BoolVarP(&fixSyntax, "fix", "f", false, "Explicitly enable automatic syntax fixing (default behavior)")

	return cmd
}

func formatSQLFiles(pattern string, overwrite bool, checkOnly bool, fixSyntax bool) error {
	// Use filepath.Glob to find matching files
	matches, err := filepath.Glob(pattern)
	if err != nil {
		return fmt.Errorf("failed to match pattern %s: %w", pattern, err)
	}

	if len(matches) == 0 {
		// Try as a single file path
		if _, err := os.Stat(pattern); err == nil {
			matches = []string{pattern}
		} else {
			return fmt.Errorf("no files found matching pattern: %s", pattern)
		}
	}

	if checkOnly {
		fmt.Printf("Checking syntax for %d file(s):\n", len(matches))
	} else if fixSyntax {
		fmt.Printf("Fixing syntax for %d file(s):\n", len(matches))
	} else {
		fmt.Printf("Found %d file(s) to format:\n", len(matches))
	}

	successCount := 0
	errorCount := 0

	for _, filePath := range matches {
		if checkOnly {
			fmt.Printf("Checking: %s\n", filePath)
		} else if fixSyntax {
			fmt.Printf("Fixing: %s\n", filePath)
		} else {
			fmt.Printf("Formatting: %s\n", filePath)
		}

		if err := processSQL(filePath, overwrite, checkOnly, fixSyntax); err != nil {
			fmt.Printf("  ❌ Error: %v\n", err)
			errorCount++
		} else {
			if checkOnly {
				fmt.Printf("  ✅ Syntax OK\n")
			} else if fixSyntax {
				fmt.Printf("  ✅ Fixed and formatted\n")
			} else {
				fmt.Printf("  ✅ Success\n")
			}
			successCount++
		}
	}

	if checkOnly {
		fmt.Printf("\nCompleted: %d valid, %d errors\n", successCount, errorCount)
	} else if fixSyntax {
		fmt.Printf("\nCompleted: %d fixed, %d errors\n", successCount, errorCount)
	} else {
		fmt.Printf("\nCompleted: %d successful, %d errors\n", successCount, errorCount)
	}

	if errorCount > 0 {
		if checkOnly {
			return fmt.Errorf("syntax errors found in %d file(s)", errorCount)
		} else {
			return fmt.Errorf("failed to process %d file(s)", errorCount)
		}
	}

	return nil
}

func processSQL(filePath string, overwrite bool, checkOnly bool, fixSyntax bool) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	// Check syntax before fixes (when fixing is enabled)
	if fixSyntax {
		if syntaxErrors := validateSQLSyntax(lines); len(syntaxErrors) > 0 {
			fmt.Printf("    Found %d syntax issue(s) to fix:\n", len(syntaxErrors))
			for _, syntaxErr := range syntaxErrors {
				fmt.Printf("      Line %d: %s\n", syntaxErr.Line, syntaxErr.Message)
			}
		}
	}

	// Apply automatic fixes if requested
	if fixSyntax {
		fixedLines, fixCount := autoFixSyntax(lines)
		if fixCount > 0 {
			fmt.Printf("    Applied %d syntax fix(es)\n", fixCount)
			lines = fixedLines
		}
	}

	// Check syntax after fixes (or original if no fixes)
	if syntaxErrors := validateSQLSyntax(lines); len(syntaxErrors) > 0 {
		if fixSyntax {
			fmt.Printf("    Remaining issues after fixes:\n")
		}
		for _, syntaxErr := range syntaxErrors {
			fmt.Printf("      Line %d: %s\n", syntaxErr.Line, syntaxErr.Message)
			if syntaxErr.Context != "" {
				fmt.Printf("        Context: %s\n", syntaxErr.Context)
			}
		}
		if !fixSyntax {
			return fmt.Errorf("syntax validation failed with %d error(s)", len(syntaxErrors))
		}
	} else if fixSyntax {
		fmt.Printf("    ✅ All syntax issues resolved\n")
	}

	// If only checking syntax, we're done
	if checkOnly {
		return nil
	}

	// Proceed with formatting using the potentially fixed lines
	return formatSQLFileFromLines(filePath, lines, overwrite)
}

func autoFixSyntax(lines []string) ([]string, int) {
	var fixedLines []string
	fixCount := 0

	// FIRST: Fix bracket issues before any other processing
	fixedLines, bracketFixes := fixBracketIssues(lines)
	fixCount += bracketFixes

	// Process each line with intelligent auto-correction
	var finalLines []string
	for _, line := range fixedLines {
		fixedLine := line

		// Skip lines that likely contain only string values (pure VALUES clauses)
		upperLine := strings.ToUpper(fixedLine)
		if strings.HasPrefix(strings.TrimSpace(upperLine), "VALUES") &&
			(strings.Contains(fixedLine, "'") || strings.Contains(fixedLine, "\"")) {
			finalLines = append(finalLines, fixedLine)
			continue
		}

		// Skip comment lines
		trimmedLine := strings.TrimSpace(line)
		if strings.HasPrefix(trimmedLine, "--") || trimmedLine == "" {
			finalLines = append(finalLines, fixedLine)
			continue
		}

		// Apply intelligent auto-correction
		fixedLine, lineFixes := intelligentAutoCorrect(fixedLine)
		fixCount += lineFixes

		// Fix spacing issues
		fixedLine = fixSpacing(fixedLine)

		// Convert to proper SQL case (uppercase keywords)
		fixedLine = uppercaseSQLKeywords(fixedLine)

		finalLines = append(finalLines, fixedLine)
	}

	// Fix missing semicolon at end of SQL statements (after all other fixes)
	finalLines, missingSemicolonFixes := fixMissingSemicolons(finalLines)
	fixCount += missingSemicolonFixes

	return finalLines, fixCount
}

// isLikelyUserIdentifier checks if a word is likely a user-defined identifier
func isLikelyUserIdentifier(word string) bool {
	word = strings.ToUpper(word)

	// Check against protected identifiers
	for _, identifier := range protectedIdentifiers {
		if word == identifier {
			return true
		}
	}

	// Skip words that contain numbers or underscores (likely identifiers)
	if strings.Contains(word, "_") || regexp.MustCompile(`\d`).MatchString(word) {
		return true
	}

	// Skip single quoted strings
	if strings.HasPrefix(word, "'") && strings.HasSuffix(word, "'") {
		return true
	}

	// Skip if it looks like a function call
	if strings.HasSuffix(word, "()") {
		return true
	}

	// Skip very common English words that might appear in SQL
	englishWords := []string{
		"THE", "AND", "FOR", "ARE", "BUT", "NOT", "YOU", "ALL", "CAN", "HER", "WAS", "ONE", "OUR", "HAD",
		"BY", "WORD", "BUT", "WHAT", "SOME", "WE", "CAN", "OUT", "OTHER", "WERE", "ALL", "THERE", "WHEN", "UP", "USE",
		"YOUR", "HOW", "SAID", "AN", "EACH", "WHICH", "SHE", "DO", "HAS", "IF", "WILL", "UP", "OTHER", "ABOUT", "OUT", "MANY", "THEN", "THEM",
	}

	for _, englishWord := range englishWords {
		if word == englishWord {
			return true
		}
	}

	return false
}

// applyCommonFixes is now replaced by intelligentAutoCorrect
// This function is kept for compatibility but delegates to the new system
func applyCommonFixes(line string) (string, int) {
	return intelligentAutoCorrect(line)
}

// fixSpacing fixes common spacing issues in SQL
func fixSpacing(line string) string {
	// Fix multiple spaces to single space
	re := regexp.MustCompile(`\s+`)
	line = re.ReplaceAllString(line, " ")

	// Fix spacing around parentheses
	line = regexp.MustCompile(`\s*\(\s*`).ReplaceAllString(line, " (")
	line = regexp.MustCompile(`\s*\)\s*`).ReplaceAllString(line, ") ")

	// Fix spacing around commas
	line = regexp.MustCompile(`\s*,\s*`).ReplaceAllString(line, ", ")

	return strings.TrimSpace(line)
}

// uppercaseSQLKeywords converts SQL keywords and data types to uppercase using the PostgreSQL dictionary
func uppercaseSQLKeywords(line string) string {
	// Skip empty lines or comment lines
	trimmed := strings.TrimSpace(line)
	if trimmed == "" || strings.HasPrefix(trimmed, "--") {
		return line
	}

	result := line

	// Process each PostgreSQL keyword for uppercase conversion
	for _, correctWord := range postgresqlKeywords {
		// Create case-insensitive pattern with word boundaries
		pattern := `(?i)\b` + regexp.QuoteMeta(correctWord) + `\b`
		re := regexp.MustCompile(pattern)

		// Find and replace all matches
		result = re.ReplaceAllStringFunc(result, func(match string) string {
			// Check if we're inside single quotes by counting quotes before this position
			matchIndex := strings.Index(result, match)
			if matchIndex < 0 {
				// Try case-insensitive search
				lowerResult := strings.ToLower(result)
				lowerMatch := strings.ToLower(match)
				matchIndex = strings.Index(lowerResult, lowerMatch)
			}

			if matchIndex >= 0 {
				beforeMatch := result[:matchIndex]
				quoteCount := strings.Count(beforeMatch, "'")

				// If odd number of quotes, we're inside a string literal
				if quoteCount%2 == 1 {
					return match // Keep original case
				}
			}

			// Convert to uppercase
			return strings.ToUpper(match)
		})
	}

	return result
}

func fixCommaAndSemicolonIssues(line string) (string, int) {
	fixCount := 0
	fixedLine := line

	// Fix trailing comma before closing parenthesis (common in CREATE TABLE)
	// Example: "PRIMARY KEY (id)," -> "PRIMARY KEY (id)"
	trailingCommaRegex := regexp.MustCompile(`(?i),(\s*\)\s*(?:COMMENT|;|$))`)
	if trailingCommaRegex.MatchString(fixedLine) {
		fixedLine = trailingCommaRegex.ReplaceAllString(fixedLine, "$1")
		fixCount++
	}

	// Fix trailing comma before PRIMARY KEY, FOREIGN KEY, etc.
	// Example: "name VARCHAR NOT NULL," followed by "PRIMARY KEY (id)" -> remove comma
	constraintRegex := regexp.MustCompile(`(?i),(\s*(?:PRIMARY\s+KEY|FOREIGN\s+KEY|UNIQUE|INDEX|KEY|CONSTRAINT)\s*\()`)
	if constraintRegex.MatchString(fixedLine) {
		fixedLine = constraintRegex.ReplaceAllString(fixedLine, "$1")
		fixCount++
	}

	// Fix missing comma between column definitions
	// Example: "id UUID NOT NULL\n    name VARCHAR" -> "id UUID NOT NULL,\n    name VARCHAR"
	missingCommaRegex := regexp.MustCompile(`(?i)(NOT\s+NULL|NULL|DEFAULT\s+\w+)(\s+[a-zA-Z_][a-zA-Z0-9_]*\s+(?:VARCHAR|INT|UUID|TIMESTAMP|BOOLEAN|JSONB|FLOAT|DECIMAL))`)
	if missingCommaRegex.MatchString(fixedLine) {
		fixedLine = missingCommaRegex.ReplaceAllString(fixedLine, "$1,$2")
		fixCount++
	}

	// Fix multiple commas in a row
	// Example: "name VARCHAR,," -> "name VARCHAR,"
	multipleCommaRegex := regexp.MustCompile(`,{2,}`)
	if multipleCommaRegex.MatchString(fixedLine) {
		fixedLine = multipleCommaRegex.ReplaceAllString(fixedLine, ",")
		fixCount++
	}

	// Fix semicolon in wrong places (inside CREATE TABLE)
	// Example: "name VARCHAR; NOT NULL" -> "name VARCHAR NOT NULL"
	wrongSemicolonRegex := regexp.MustCompile(`(?i);(\s*(?:NOT\s+NULL|NULL|DEFAULT|COMMENT))`)
	if wrongSemicolonRegex.MatchString(fixedLine) {
		fixedLine = wrongSemicolonRegex.ReplaceAllString(fixedLine, " $1")
		fixCount++
	}

	// Fix missing semicolon at end of CREATE TABLE statement
	// This is handled in a separate function that processes the entire content

	return fixedLine, fixCount
}

func fixMissingSemicolons(lines []string) ([]string, int) {
	if len(lines) == 0 {
		return lines, 0
	}

	fixCount := 0
	fixedLines := make([]string, len(lines))
	copy(fixedLines, lines)

	// Join all lines to work with the complete SQL statement
	content := strings.Join(lines, "\n")

	// Check if this looks like a SQL statement (contains CREATE TABLE, ALTER TABLE, etc.)
	sqlStatementRegex := regexp.MustCompile(`(?is)(CREATE\s+TABLE|ALTER\s+TABLE|INSERT\s+INTO|UPDATE\s+|DELETE\s+FROM)`)

	if sqlStatementRegex.MatchString(content) {
		// Find the last non-empty line
		lastNonEmptyIndex := -1
		for i := len(fixedLines) - 1; i >= 0; i-- {
			line := strings.TrimSpace(fixedLines[i])
			if line != "" && !strings.HasPrefix(line, "--") { // Skip comments
				lastNonEmptyIndex = i
				break
			}
		}

		if lastNonEmptyIndex >= 0 {
			lastLine := strings.TrimSpace(fixedLines[lastNonEmptyIndex])

			// Check if the last line doesn't end with semicolon
			if !strings.HasSuffix(lastLine, ";") {
				fixedLines[lastNonEmptyIndex] = lastLine + ";"
				fixCount++
			}
		}
	}

	return fixedLines, fixCount
}

func fixBracketIssues(lines []string) ([]string, int) {
	if len(lines) == 0 {
		return lines, 0
	}

	fixCount := 0
	fixedLines := make([]string, len(lines))
	copy(fixedLines, lines)

	// Enhanced bracket fixing with better multi-line support
	for i := 0; i < len(fixedLines); i++ {
		line := fixedLines[i]
		trimmedLine := strings.TrimSpace(line)

		// Check for PRIMARY KEY with missing closing parenthesis followed by COMMENT
		// Pattern: "PRIMARY KEY (column" followed by "COMMENT..." on next line
		primaryKeyPattern := regexp.MustCompile(`(?i)^\s*PRIMARY\s+KEY\s*\(\s*(\w+)\s*$`)
		if primaryKeyPattern.MatchString(trimmedLine) {
			matches := primaryKeyPattern.FindStringSubmatch(trimmedLine)
			if len(matches) > 1 {
				columnName := matches[1]

				// Check if the next line starts with COMMENT - this means we need to add table closing bracket too
				if i+1 < len(fixedLines) {
					nextLine := strings.TrimSpace(fixedLines[i+1])
					if regexp.MustCompile(`(?i)^\s*COMMENT\s+`).MatchString(nextLine) {
						// Need to add both PRIMARY KEY closing ) and table closing )
						fixedLines[i] = fmt.Sprintf("    PRIMARY KEY (%s)", columnName)
						// Insert the table closing bracket before COMMENT
						newLines := make([]string, len(fixedLines)+1)
						copy(newLines[:i+1], fixedLines[:i+1])
						newLines[i+1] = ")"
						copy(newLines[i+2:], fixedLines[i+1:])
						fixedLines = newLines
						fixCount += 2
						fmt.Printf("    Fixed PRIMARY KEY and table closing brackets: (%s) (Line %d)\n", columnName, i+1)
					} else {
						// Only PRIMARY KEY closing parenthesis needed
						fixedLines[i] = fmt.Sprintf("    PRIMARY KEY (%s)", columnName)
						fixCount++
						fmt.Printf("    Fixed missing closing parenthesis in PRIMARY KEY: (%s) (Line %d)\n", columnName, i+1)
					}
				} else {
					// Just fix PRIMARY KEY at end of file
					fixedLines[i] = fmt.Sprintf("    PRIMARY KEY (%s)", columnName)
					fixCount++
					fmt.Printf("    Fixed missing closing parenthesis in PRIMARY KEY: (%s) (Line %d)\n", columnName, i+1)
				}
			}
		}

		// Handle case where PRIMARY KEY is split across lines
		if regexp.MustCompile(`(?i)^\s*PRIMARY\s+KEY\s*\(\s*$`).MatchString(trimmedLine) {
			if i+1 < len(fixedLines) {
				nextLine := strings.TrimSpace(fixedLines[i+1])
				if regexp.MustCompile(`^\w+\s*$`).MatchString(nextLine) {
					fixedLines[i] = fmt.Sprintf("    PRIMARY KEY (%s)", nextLine)
					fixedLines[i+1] = "" // Remove the separate column name line
					fixCount++
					fmt.Printf("    Fixed PRIMARY KEY with missing closing parenthesis: (%s) (Line %d)\n", nextLine, i+1)
				}
			}
		}
	}

	// Second pass: General bracket balancing for any remaining issues
	openBrackets := 0
	closeBrackets := 0

	// Count total brackets across all lines
	for _, line := range fixedLines {
		if strings.TrimSpace(line) != "" {
			openBrackets += strings.Count(line, "(")
			closeBrackets += strings.Count(line, ")")
		}
	}

	// If we still have unbalanced brackets, try to fix them intelligently
	if openBrackets > closeBrackets {
		missingClosing := openBrackets - closeBrackets
		fixed := false

		// Find the COMMENT line to insert closing brackets before it
		for i, line := range fixedLines {
			trimmedLine := strings.TrimSpace(line)
			if regexp.MustCompile(`(?i)^\s*COMMENT\s+`).MatchString(trimmedLine) {
				// Insert missing closing brackets before COMMENT
				for j := 0; j < missingClosing; j++ {
					newLines := make([]string, len(fixedLines)+1)
					copy(newLines[:i+j], fixedLines[:i+j])
					newLines[i+j] = ")"
					copy(newLines[i+j+1:], fixedLines[i+j:])
					fixedLines = newLines
				}
				fixCount += missingClosing
				fmt.Printf("    Added %d missing closing bracket(s) before COMMENT (Line %d)\n", missingClosing, i+1)
				fixed = true
				break
			}
		}

		// If no COMMENT found, try to find the end of CREATE TABLE statement
		if !fixed {
			for i, line := range fixedLines {
				trimmedLine := strings.TrimSpace(line)
				// Look for end of CREATE TABLE (line with just ");")
				if trimmedLine == ");" || (strings.HasSuffix(trimmedLine, ")") && strings.Contains(trimmedLine, ";")) {
					// Insert missing closing brackets before the existing ");"
					for j := 0; j < missingClosing; j++ {
						newLines := make([]string, len(fixedLines)+1)
						copy(newLines[:i+j], fixedLines[:i+j])
						newLines[i+j] = ")"
						copy(newLines[i+j+1:], fixedLines[i+j:])
						fixedLines = newLines
					}
					fixCount += missingClosing
					fmt.Printf("    Added %d missing closing bracket(s) before end of CREATE TABLE (Line %d)\n", missingClosing, i+1)
					fixed = true
					break
				}
			}
		}

		// Last resort: only add at end if we're dealing with a single CREATE TABLE statement
		if !fixed {
			// Check if this looks like a single CREATE TABLE (most lines have column definitions)
			createTableLines := 0
			totalLines := len(fixedLines)
			for _, line := range fixedLines {
				if regexp.MustCompile(`(?i)(CREATE\s+TABLE|PRIMARY\s+KEY|FOREIGN\s+KEY|\w+\s+\w+.*,\s*$)`).MatchString(strings.TrimSpace(line)) {
					createTableLines++
				}
			}

			// Only add brackets at end if majority of lines are CREATE TABLE related
			if createTableLines > totalLines/2 {
				for j := 0; j < missingClosing; j++ {
					fixedLines = append(fixedLines, ")")
				}
				fixCount += missingClosing
				fmt.Printf("    Added %d missing closing bracket(s) at end of CREATE TABLE\n", missingClosing)
			}
		}
	}

	// Remove empty lines that were created during processing
	var cleanedLines []string
	for _, line := range fixedLines {
		if strings.TrimSpace(line) != "" || line == "" {
			cleanedLines = append(cleanedLines, line)
		}
	}

	return cleanedLines, fixCount
}

func formatSQLFileFromLines(filePath string, lines []string, overwrite bool) error {
	formatted := formatSQL(lines)

	if overwrite {
		return writeToFile(filePath, formatted)
	}

	return nil
}

func validateSQLSyntax(lines []string) []SQLSyntaxError {
	var errors []SQLSyntaxError
	content := strings.Join(lines, "\n")

	// Comprehensive syntax validation - check various SQL patterns
	errors = append(errors, checkSQLStatementStructure(lines)...)
	errors = append(errors, checkDataTypes(lines)...)
	errors = append(errors, checkSQLKeywords(lines)...)
	errors = append(errors, checkContextSpecificErrors(lines)...)
	errors = append(errors, checkParenthesesBalance(content)...)
	errors = append(errors, checkCommaUsage(lines)...)
	errors = append(errors, checkQuoteBalance(content)...)

	return errors
}

// checkSQLStatementStructure validates the structure of SQL statements
func checkSQLStatementStructure(lines []string) []SQLSyntaxError {
	var errors []SQLSyntaxError

	for i, line := range lines {
		trimmed := strings.TrimSpace(line)
		upperLine := strings.ToUpper(trimmed)

		// Check CREATE TABLE syntax
		if strings.HasPrefix(upperLine, "CREATE TABLE") {
			parts := strings.Fields(trimmed)
			if len(parts) < 3 {
				errors = append(errors, SQLSyntaxError{
					Line:    i + 1,
					Message: "CREATE TABLE statement missing table name",
					Context: trimmed,
				})
			}
		}

		// Check ALTER TABLE syntax
		if strings.HasPrefix(upperLine, "ALTER TABLE") {
			// Look for common patterns like "ALTER TABLE table_name ad COLUMN"
			if regexp.MustCompile(`(?i)ALTER\s+TABLE\s+\w+\s+ad\s+COLUMN`).MatchString(trimmed) {
				errors = append(errors, SQLSyntaxError{
					Line:    i + 1,
					Message: "Syntax error: 'ad' should be 'ADD' in ALTER TABLE statement",
					Context: trimmed,
				})
			}
		}

		// Check UPDATE syntax
		if strings.HasPrefix(upperLine, "UPDATE") {
			// Look for "UPDATE table se column" instead of "UPDATE table SET column"
			if regexp.MustCompile(`(?i)UPDATE\s+\w+\s+se\s+\w+`).MatchString(trimmed) {
				errors = append(errors, SQLSyntaxError{
					Line:    i + 1,
					Message: "Syntax error: 'se' should be 'SET' in UPDATE statement",
					Context: trimmed,
				})
			}
		}

		// Check INSERT syntax
		if strings.HasPrefix(upperLine, "INSERT") {
			// Look for "INSERT INT" instead of "INSERT INTO"
			if regexp.MustCompile(`(?i)INSERT\s+INT\s+\w+`).MatchString(trimmed) {
				errors = append(errors, SQLSyntaxError{
					Line:    i + 1,
					Message: "Syntax error: 'INSERT INT' should be 'INSERT INTO'",
					Context: trimmed,
				})
			}
			// Look for "value" instead of "VALUES"
			if regexp.MustCompile(`(?i)INSERT\s+INTO\s+\w+.*\s+value\s*\(`).MatchString(trimmed) {
				errors = append(errors, SQLSyntaxError{
					Line:    i + 1,
					Message: "Syntax error: 'value' should be 'VALUES' in INSERT statement",
					Context: trimmed,
				})
			}
		}

		// Check JOIN syntax
		if regexp.MustCompile(`(?i)(INNER|LEFT|RIGHT|FULL)\s+JOIN`).MatchString(trimmed) {
			// Look for "JOIN table o condition" instead of "JOIN table ON condition"
			if regexp.MustCompile(`(?i)JOIN\s+\w+\s+\w+\s+o\s+`).MatchString(trimmed) {
				errors = append(errors, SQLSyntaxError{
					Line:    i + 1,
					Message: "Syntax error: 'o' should be 'ON' in JOIN condition",
					Context: trimmed,
				})
			}
		}

		// Check WHERE clause syntax
		if regexp.MustCompile(`(?i)WHERE\s+\w+\s+i\s+NOT`).MatchString(trimmed) {
			errors = append(errors, SQLSyntaxError{
				Line:    i + 1,
				Message: "Syntax error: 'i' should be 'IS' in WHERE clause",
				Context: trimmed,
			})
		}
	}

	return errors
}

// checkSQLKeywords validates SQL keywords for common typos
func checkSQLKeywords(lines []string) []SQLSyntaxError {
	var errors []SQLSyntaxError

	for i, line := range lines {
		trimmed := strings.TrimSpace(line)

		// Extract words from the line
		words := extractSQLWords(trimmed)

		for _, word := range words {
			// Skip very short words unless they're known problematic ones
			if len(word) < 2 {
				continue
			}

			// Skip if it's likely a user identifier
			if isLikelyUserIdentifier(word) {
				continue
			}

			upperWord := strings.ToUpper(word)

			// Check if it's a known short-form error
			if correction, exists := shortFormCorrections[upperWord]; exists {
				errors = append(errors, SQLSyntaxError{
					Line:    i + 1,
					Message: fmt.Sprintf("Possible typo: '%s' should be '%s'", word, correction),
					Context: trimmed,
				})
				continue
			}

			// Check if word is already correct
			if _, exists := postgresqlKeywords[upperWord]; exists {
				continue
			}

			// Try fuzzy matching for more complex typos
			match := findBestMatch(word, 2, 0.70) // Lower threshold for error detection
			if match != nil {
				errors = append(errors, SQLSyntaxError{
					Line: i + 1,
					Message: fmt.Sprintf("Possible typo: '%s' should be '%s' (confidence: %.0f%%)",
						word, match.Correction, match.Confidence*100),
					Context: trimmed,
				})
			}
		}
	}

	return errors
}

// checkContextSpecificErrors checks for context-specific SQL errors
func checkContextSpecificErrors(lines []string) []SQLSyntaxError {
	var errors []SQLSyntaxError

	for i, line := range lines {
		trimmed := strings.TrimSpace(line)

		// Check for incomplete PRIMARY KEY/FOREIGN KEY
		if regexp.MustCompile(`(?i)PRIMARY\s+KE\b`).MatchString(trimmed) {
			errors = append(errors, SQLSyntaxError{
				Line:    i + 1,
				Message: "Incomplete keyword: 'PRIMARY KE' should be 'PRIMARY KEY'",
				Context: trimmed,
			})
		}

		if regexp.MustCompile(`(?i)FOREIGN\s+KE\b`).MatchString(trimmed) {
			errors = append(errors, SQLSyntaxError{
				Line:    i + 1,
				Message: "Incomplete keyword: 'FOREIGN KE' should be 'FOREIGN KEY'",
				Context: trimmed,
			})
		}

		// Check for boolean value typos
		if regexp.MustCompile(`(?i)=\s*tru\b`).MatchString(trimmed) {
			errors = append(errors, SQLSyntaxError{
				Line:    i + 1,
				Message: "Boolean value typo: 'tru' should be 'TRUE'",
				Context: trimmed,
			})
		}

		// Check for data type abbreviations in wrong context
		if regexp.MustCompile(`(?i)\bvch\s*\(`).MatchString(trimmed) {
			errors = append(errors, SQLSyntaxError{
				Line:    i + 1,
				Message: "Data type abbreviation: 'vch' should be 'VARCHAR'",
				Context: trimmed,
			})
		}
	}

	return errors
}

func checkDataTypes(lines []string) []SQLSyntaxError {
	var errors []SQLSyntaxError

	for i, line := range lines {
		trimmed := strings.TrimSpace(line)

		// Check for multi-word keyword issues
		for incorrect, correct := range multiWordKeywords {
			incorrectPattern := strings.ReplaceAll(incorrect, " ", `\s+`)
			pattern := `(?i)\b` + incorrectPattern + `\b`
			re := regexp.MustCompile(pattern)
			if re.MatchString(trimmed) {
				// Only report if it's not already correct
				if !strings.Contains(strings.ToUpper(trimmed), correct) {
					errors = append(errors, SQLSyntaxError{
						Line:    i + 1,
						Message: fmt.Sprintf("Multi-word keyword: '%s' should be '%s'", incorrect, correct),
						Context: trimmed,
					})
				}
			}
		}
	}

	return errors
}

func checkParenthesesBalance(content string) []SQLSyntaxError {
	var errors []SQLSyntaxError

	openCount := strings.Count(content, "(")
	closeCount := strings.Count(content, ")")

	if openCount != closeCount {
		errors = append(errors, SQLSyntaxError{
			Line:    1,
			Message: fmt.Sprintf("Unbalanced parentheses: %d opening, %d closing", openCount, closeCount),
			Context: "Check parentheses balance",
		})
	}

	return errors
}

func checkCommaUsage(lines []string) []SQLSyntaxError {
	var errors []SQLSyntaxError

	for i, line := range lines {
		trimmed := strings.TrimSpace(line)

		// Check for missing comma at end of column definition
		if i < len(lines)-1 {
			nextTrimmed := strings.TrimSpace(lines[i+1])

			// If this line looks like a column definition and next line too,
			// but current line doesn't end with comma
			if isColumnDefinition(trimmed) && isColumnDefinition(nextTrimmed) {
				if !strings.HasSuffix(trimmed, ",") {
					errors = append(errors, SQLSyntaxError{
						Line:    i + 1,
						Message: "Missing comma at end of column definition",
						Context: trimmed,
					})
				}
			}
		}

		// Check for trailing comma before closing parenthesis
		if strings.HasSuffix(trimmed, ",") {
			if i < len(lines)-1 {
				nextTrimmed := strings.TrimSpace(lines[i+1])
				if strings.HasPrefix(nextTrimmed, ")") {
					errors = append(errors, SQLSyntaxError{
						Line:    i + 1,
						Message: "Trailing comma before closing parenthesis",
						Context: trimmed,
					})
				}
			}
		}
	}

	return errors
}

func checkQuoteBalance(content string) []SQLSyntaxError {
	var errors []SQLSyntaxError

	singleQuotes := strings.Count(content, "'")
	doubleQuotes := strings.Count(content, "\"")

	if singleQuotes%2 != 0 {
		errors = append(errors, SQLSyntaxError{
			Line:    1,
			Message: "Unbalanced single quotes",
			Context: "Check quote pairing",
		})
	}

	if doubleQuotes%2 != 0 {
		errors = append(errors, SQLSyntaxError{
			Line:    1,
			Message: "Unbalanced double quotes",
			Context: "Check quote pairing",
		})
	}

	return errors
}

func isColumnDefinition(line string) bool {
	trimmed := strings.TrimSpace(line)
	if trimmed == "" || trimmed == "(" || strings.HasPrefix(trimmed, ")") {
		return false
	}

	// Skip PRIMARY KEY constraints
	if strings.HasPrefix(strings.ToUpper(trimmed), "PRIMARY KEY") {
		return false
	}

	// Skip CREATE TABLE line
	if strings.HasPrefix(strings.ToUpper(trimmed), "CREATE TABLE") {
		return false
	}

	// Basic check: should have at least identifier and type
	parts := strings.Fields(trimmed)
	return len(parts) >= 2
}

func formatSQL(lines []string) []string {
	var result []string
	content := strings.Join(lines, "\n")

	// Find CREATE TABLE statements using regex
	createTableRegex := regexp.MustCompile(`(?is)CREATE\s+TABLE\s+([^\s(]+)\s*\((.*?)\)\s*(?:COMMENT\s*['"]([^'"]*)['"]\s*)?;`)

	// Find all CREATE TABLE matches with their positions
	matches := createTableRegex.FindAllStringSubmatchIndex(content, -1)

	if len(matches) == 0 {
		// No CREATE TABLE found, return original lines
		return lines
	}

	lastEnd := 0

	for _, matchIndices := range matches {
		// Add content before the CREATE TABLE statement
		beforeContent := content[lastEnd:matchIndices[0]]
		if beforeContent != "" {
			beforeLines := strings.Split(strings.TrimSpace(beforeContent), "\n")
			for _, line := range beforeLines {
				trimmed := strings.TrimSpace(line)
				if trimmed != "" {
					result = append(result, trimmed)
				}
			}
		}

		// Extract the CREATE TABLE match details
		tableName := strings.TrimSpace(content[matchIndices[2]:matchIndices[3]])
		columnsBlock := strings.TrimSpace(content[matchIndices[4]:matchIndices[5]])

		var tableComment string
		if len(matchIndices) > 7 && matchIndices[6] != -1 {
			tableComment = content[matchIndices[6]:matchIndices[7]]
		}

		// Parse individual column definitions
		var columnDefs []string

		// Split by lines and clean up
		columnLines := strings.Split(columnsBlock, "\n")
		for _, line := range columnLines {
			cleaned := strings.TrimSpace(line)
			if cleaned != "" {
				columnDefs = append(columnDefs, cleaned)
			}
		}

		// Format the complete CREATE TABLE statement
		formatted := formatCreateTableStatement("CREATE TABLE "+tableName, columnDefs, tableComment)
		result = append(result, formatted...)

		// Add empty line after CREATE TABLE for readability
		result = append(result, "")

		lastEnd = matchIndices[1]
	}

	// Add any remaining content after the last CREATE TABLE
	afterContent := content[lastEnd:]
	if afterContent != "" {
		afterLines := strings.Split(strings.TrimSpace(afterContent), "\n")
		for _, line := range afterLines {
			trimmed := strings.TrimSpace(line)
			if trimmed != "" {
				result = append(result, trimmed)
			}
		}
	}

	return result
}

func formatCreateTableStatement(tableName string, columnDefs []string, tableComment string) []string {
	var result []string

	// Add CREATE TABLE line
	result = append(result, tableName)
	result = append(result, "(")

	if len(columnDefs) == 0 {
		result = append(result, ")")
		if tableComment != "" {
			result[len(result)-1] += fmt.Sprintf(` COMMENT '%s';`, tableComment)
		} else {
			result[len(result)-1] += ";"
		}
		return result
	}

	// Parse and format column definitions
	var columns []Column
	var primaryKeyCol string
	maxNameLen := 0
	maxTypeLen := 0
	maxNullLen := 0

	for _, colDef := range columnDefs {
		col := parseColumnDefinition(colDef)
		if col.IsPK {
			primaryKeyCol = col.Name
			continue // Skip adding PRIMARY KEY constraint as a column
		}
		columns = append(columns, col)

		if len(col.Name) > maxNameLen {
			maxNameLen = len(col.Name)
		}
		if len(col.Type) > maxTypeLen {
			maxTypeLen = len(col.Type)
		}
		if len(col.Nullable) > maxNullLen {
			maxNullLen = len(col.Nullable)
		}
	}

	// Format each column with proper alignment
	for i, col := range columns {
		line := fmt.Sprintf("    %-*s %-*s %-*s",
			maxNameLen, col.Name,
			maxTypeLen, col.Type,
			maxNullLen, col.Nullable)

		if col.Comment != "" {
			line += fmt.Sprintf(" COMMENT '%s'", col.Comment)
		}

		// Add comma except for last column (unless we have PRIMARY KEY)
		if i < len(columns)-1 || primaryKeyCol != "" {
			line += ","
		}

		result = append(result, line)
	}

	// Add PRIMARY KEY if found
	if primaryKeyCol != "" {
		result = append(result, fmt.Sprintf("    PRIMARY KEY (%s)", primaryKeyCol))
	}

	// Add closing parenthesis and table comment
	closeLine := ")"
	if tableComment != "" {
		closeLine += fmt.Sprintf(` COMMENT '%s'`, tableComment)
	}
	closeLine += ";"
	result = append(result, closeLine)

	return result
}

func parseColumnDefinition(colDef string) Column {
	// Remove trailing comma if present
	colDef = strings.TrimRight(colDef, ",")

	var col Column

	// Check if this is a standalone PRIMARY KEY constraint
	if strings.Contains(strings.ToUpper(colDef), "PRIMARY KEY") && !strings.Contains(strings.ToUpper(colDef), "NOT NULL PRIMARY KEY") {
		re := regexp.MustCompile(`PRIMARY\s+KEY\s*\(\s*(\w+)\s*\)`)
		matches := re.FindStringSubmatch(colDef)
		if len(matches) > 1 {
			col.Name = matches[1]
			col.IsPK = true
			return col
		}
	}

	// Parse regular column definition
	// Handle quoted column names like "group"
	parts := strings.Fields(colDef)
	if len(parts) < 2 {
		return col
	}

	col.Name = parts[0]
	col.Type = parts[1]

	// Check for inline PRIMARY KEY
	upperColDef := strings.ToUpper(colDef)
	if strings.Contains(upperColDef, "PRIMARY KEY") {
		// This column has PRIMARY KEY constraint inline
		// We'll handle it as a regular column for now
	}

	// Find NULL/NOT NULL
	if strings.Contains(upperColDef, "NOT NULL") {
		col.Nullable = "NOT NULL"
	} else if strings.Contains(upperColDef, "NULL") && !strings.Contains(upperColDef, "NOT NULL") {
		col.Nullable = "NULL"
	}

	// Extract comment - handle Thai characters properly
	commentRe := regexp.MustCompile(`(?i)COMMENT\s+['"]([^'"]+)['"]`)
	matches := commentRe.FindStringSubmatch(colDef)
	if len(matches) > 1 {
		col.Comment = matches[1]
	}

	return col
}

func writeToFile(filePath string, lines []string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for _, line := range lines {
		if _, err := writer.WriteString(line + "\n"); err != nil {
			return fmt.Errorf("failed to write to file: %w", err)
		}
	}

	return nil
}
