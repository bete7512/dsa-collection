
import (
	"regexp"
    "strings"
)

func isPalindrome(s string) bool {
    	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	    cleaned := re.ReplaceAllString(s, "")
        cleaned = strings.ToLower(cleaned)
        j := len(cleaned) - 1
        for i := 0; i < len(cleaned); i++{
            if cleaned[i] != cleaned[j]{
                return false
            }
            j--
        }
        return true
}