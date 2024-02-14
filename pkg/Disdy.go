package pkg

import "errors"

const (
	H1 = iota
	H2
	H3
	H4
	H5
	H6
)

var (
	ErrHTMLTypeInvalid = errors.New("HTML Type Invalid")
)

func IsValidTypeHTML(strType string) bool {
	if strType == "INVALID" {
		return false
	} else {
		return true
	}
}
func ConvertTypeHTMLToString(TypeHTML int) string {
	j, err := ConvertTypeHTML(TypeHTML)
	if err != nil {
		return "INVALID"
	} else {
		return *j
	}
}
func ConvertTypeHTML(TypeHTML int) (*string, error) {
	ty := "INVALID"
	switch TypeHTML {
	case H1:
		ty = "h1"
		break
	case H2:
		ty = "h2"
		break
	case H3:
		ty = "h3"
		break
	case H4:
		ty = "h4"
		break
	case H5:
		ty = "h5"
		break
	case H6:
		ty = "h6"
		break
	default:
		ty = "INVALID"
		break
	}
	if ty == "INVALID" {
		return nil, ErrHTMLTypeInvalid
	}
	return &ty, nil
}

var Document = NewDocumentDOM()
