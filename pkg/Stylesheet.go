package pkg

import (
	"reflect"
	"strings"
)

type CSStylesheet struct {
	AlignContent             string `default:".None"`
	AlignItems               string `default:".None"`
	AlignSelf                string `default:".None"`
	All                      string `default:".None"`
	Animation                string `default:".None"`
	AnimationDelay           string `default:".None"`
	AnimationDirection       string `default:".None"`
	AnimationDuration        string `default:".None"`
	AnimationFillMode        string `default:".None"`
	AnimationIteration       string `default:".None"`
	AnimationName            string `default:".None"`
	AnimationPlayState       string `default:".None"`
	AnimationTiming          string `default:".None"`
	BackdropFilter           string `default:".None"`
	BackfaceVisibility       string `default:".None"`
	Background               string `default:".None"`
	BackgroundAttachment     string `default:".None"`
	BackgroundBlendMode      string `default:".None"`
	BackgroundClip           string `default:".None"`
	BackgroundColor          string `default:".None"`
	BackgroundImage          string `default:".None"`
	BackgroundOrigin         string `default:".None"`
	BackgroundPosition       string `default:".None"`
	BackgroundRepeat         string `default:".None"`
	BackgroundSize           string `default:".None"`
	BlockSize                string `default:".None"`
	Border                   string `default:".None"`
	BorderBlock              string `default:".None"`
	BorderBlockColor         string `default:".None"`
	BorderBlockEnd           string `default:".None"`
	BorderBlockEndColor      string `default:".None"`
	BorderBlockEndStyle      string `default:".None"`
	BorderBlockEndWidth      string `default:".None"`
	BorderBlockStart         string `default:".None"`
	BorderBlockStartColor    string `default:".None"`
	BorderBlockStartStyle    string `default:".None"`
	BorderBlockStartWidth    string `default:".None"`
	BorderBlockStyle         string `default:".None"`
	BorderBlockWidth         string `default:".None"`
	BorderBottom             string `default:".None"`
	BorderBottomColor        string `default:".None"`
	BorderBottomLeftRadius   string `default:".None"`
	BorderBottomRightRadius  string `default:".None"`
	BorderBottomStyle        string `default:".None"`
	BorderBottomWidth        string `default:".None"`
	BorderCollapse           string `default:".None"`
	BorderColor              string `default:".None"`
	BorderImage              string `default:".None"`
	BorderImageOutset        string `default:".None"`
	BorderImageRepeat        string `default:".None"`
	BorderImageSlice         string `default:".None"`
	BorderImageSource        string `default:".None"`
	BorderImageWidth         string `default:".None"`
	BorderInline             string `default:".None"`
	BorderInlineColor        string `default:".None"`
	BorderInlineEnd          string `default:".None"`
	BorderInlineEndColor     string `default:".None"`
	BorderInlineEndStyle     string `default:".None"`
	BorderInlineEndWidth     string `default:".None"`
	BorderInlineStart        string `default:".None"`
	BorderInlineStartColor   string `default:".None"`
	BorderInlineStartStyle   string `default:".None"`
	BorderInlineStartWidth   string `default:".None"`
	BorderInlineStyle        string `default:".None"`
	BorderInlineWidth        string `default:".None"`
	BorderLeft               string `default:".None"`
	BorderLeftColor          string `default:".None"`
	BorderLeftStyle          string `default:".None"`
	BorderLeftWidth          string `default:".None"`
	BorderRadius             string `default:".None"`
	BorderRight              string `default:".None"`
	BorderRightColor         string `default:".None"`
	BorderRightStyle         string `default:".None"`
	BorderRightWidth         string `default:".None"`
	BorderSpacing            string `default:".None"`
	BorderStyle              string `default:".None"`
	BorderTop                string `default:".None"`
	BorderTopColor           string `default:".None"`
	BorderTopLeftRadius      string `default:".None"`
	BorderTopRightRadius     string `default:".None"`
	BorderTopStyle           string `default:".None"`
	BorderTopWidth           string `default:".None"`
	BorderWidth              string `default:".None"`
	Bottom                   string `default:".None"`
	BoxDecorationBreak       string `default:".None"`
	BoxShadow                string `default:".None"`
	BoxSizing                string `default:".None"`
	BreakAfter               string `default:".None"`
	BreakBefore              string `default:".None"`
	BreakInside              string `default:".None"`
	CaptionSide              string `default:".None"`
	CaretColor               string `default:".None"`
	Clear                    string `default:".None"`
	Clip                     string `default:".None"`
	ClipPath                 string `default:".None"`
	Color                    string `default:".None"`
	ColumnCount              string `default:".None"`
	ColumnFill               string `default:".None"`
	ColumnGap                string `default:".None"`
	ColumnRule               string `default:".None"`
	ColumnRuleColor          string `default:".None"`
	ColumnRuleStyle          string `default:".None"`
	ColumnRuleWidth          string `default:".None"`
	ColumnSpan               string `default:".None"`
	ColumnWidth              string `default:".None"`
	Columns                  string `default:".None"`
	Contain                  string `default:".None"`
	Content                  string `default:".None"`
	CounterIncrement         string `default:".None"`
	CounterReset             string `default:".None"`
	Cursor                   string `default:".None"`
	Direction                string `default:".None"`
	Display                  string `default:".None"`
	EmptyCells               string `default:".None"`
	Filter                   string `default:".None"`
	Flex                     string `default:".None"`
	FlexBasis                string `default:".None"`
	FlexDirection            string `default:".None"`
	FlexFlow                 string `default:".None"`
	FlexGrow                 string `default:".None"`
	FlexShrink               string `default:".None"`
	FlexWrap                 string `default:".None"`
	Float                    string `default:".None"`
	Font                     string `default:".None"`
	FontFamily               string `default:".None"`
	FontFeatureSettings      string `default:".None"`
	FontKerning              string `default:".None"`
	FontLanguageOverride     string `default:".None"`
	FontSize                 string `default:".None"`
	FontSizeAdjust           string `default:".None"`
	FontStretch              string `default:".None"`
	FontStyle                string `default:".None"`
	FontSynthesis            string `default:".None"`
	FontVariant              string `default:".None"`
	FontVariantCaps          string `default:".None"`
	FontVariantEastAsian     string `default:".None"`
	FontVariantLigatures     string `default:".None"`
	FontVariantNumeric       string `default:".None"`
	FontVariantPosition      string `default:".None"`
	FontWeight               string `default:".None"`
	Gap                      string `default:".None"`
	Grid                     string `default:".None"`
	GridArea                 string `default:".None"`
	GridAutoColumns          string `default:".None"`
	GridAutoFlow             string `default:".None"`
	GridAutoRows             string `default:".None"`
	GridColumn               string `default:".None"`
	GridColumnEnd            string `default:".None"`
	GridColumnGap            string `default:".None"`
	GridColumnStart          string `default:".None"`
	GridGap                  string `default:".None"`
	GridRow                  string `default:".None"`
	GridRowEnd               string `default:".None"`
	GridRowGap               string `default:".None"`
	GridRowStart             string `default:".None"`
	GridTemplate             string `default:".None"`
	GridTemplateAreas        string `default:".None"`
	GridTemplateColumns      string `default:".None"`
	GridTemplateRows         string `default:".None"`
	HangingPunctuation       string `default:".None"`
	Height                   string `default:".None"`
	Hyphens                  string `default:".None"`
	ImageOrientation         string `default:".None"`
	ImageRendering           string `default:".None"`
	InlineSize               string `default:".None"`
	Inset                    string `default:".None"`
	InsetBlock               string `default:".None"`
	InsetBlockEnd            string `default:".None"`
	InsetBlockStart          string `default:".None"`
	InsetInline              string `default:".None"`
	InsetInlineEnd           string `default:".None"`
	InsetInlineStart         string `default:".None"`
	Isolation                string `default:".None"`
	JustifyContent           string `default:".None"`
	Left                     string `default:".None"`
	LetterSpacing            string `default:".None"`
	LineBreak                string `default:".None"`
	LineHeight               string `default:".None"`
	ListStyle                string `default:".None"`
	ListStyleImage           string `default:".None"`
	ListStylePosition        string `default:".None"`
	ListStyleType            string `default:".None"`
	Margin                   string `default:".None"`
	MarginBlock              string `default:".None"`
	MarginBlockEnd           string `default:".None"`
	MarginBlockStart         string `default:".None"`
	MarginBottom             string `default:".None"`
	MarginInline             string `default:".None"`
	MarginInlineEnd          string `default:".None"`
	MarginInlineStart        string `default:".None"`
	MarginLeft               string `default:".None"`
	MarginRight              string `default:".None"`
	MarginTop                string `default:".None"`
	Mask                     string `default:".None"`
	MaskBorder               string `default:".None"`
	MaskBorderMode           string `default:".None"`
	MaskBorderOutset         string `default:".None"`
	MaskBorderRepeat         string `default:".None"`
	MaskBorderSlice          string `default:".None"`
	MaskBorderSource         string `default:".None"`
	MaskBorderWidth          string `default:".None"`
	MaskClip                 string `default:".None"`
	MaskComposite            string `default:".None"`
	MaskImage                string `default:".None"`
	MaskMode                 string `default:".None"`
	MaskOrigin               string `default:".None"`
	MaskPosition             string `default:".None"`
	MaskRepeat               string `default:".None"`
	MaskSize                 string `default:".None"`
	MaskType                 string `default:".None"`
	MaxBlockSize             string `default:".None"`
	MaxHeight                string `default:".None"`
	MaxInlineSize            string `default:".None"`
	MaxWidth                 string `default:".None"`
	MinBlockSize             string `default:".None"`
	MinHeight                string `default:".None"`
	MinInlineSize            string `default:".None"`
	MinWidth                 string `default:".None"`
	MixBlendMode             string `default:".None"`
	ObjectFit                string `default:".None"`
	ObjectPosition           string `default:".None"`
	Offset                   string `default:".None"`
	OffsetAnchor             string `default:".None"`
	OffsetBlock              string `default:".None"`
	OffsetBlockEnd           string `default:".None"`
	OffsetBlockStart         string `default:".None"`
	OffsetInline             string `default:".None"`
	OffsetInlineEnd          string `default:".None"`
	OffsetInlineStart        string `default:".None"`
	OffsetDistance           string `default:".None"`
	OffsetPath               string `default:".None"`
	OffsetRotate             string `default:".None"`
	Opacity                  string `default:".None"`
	Order                    string `default:".None"`
	Orphans                  string `default:".None"`
	Outline                  string `default:".None"`
	OutlineColor             string `default:".None"`
	OutlineOffset            string `default:".None"`
	OutlineStyle             string `default:".None"`
	OutlineWidth             string `default:".None"`
	Overflow                 string `default:".None"`
	OverflowAnchor           string `default:".None"`
	OverflowBlock            string `default:".None"`
	OverflowInline           string `default:".None"`
	OverflowWrap             string `default:".None"`
	OverflowX                string `default:".None"`
	OverflowY                string `default:".None"`
	OverscrollBehavior       string `default:".None"`
	OverscrollBehaviorBlock  string `default:".None"`
	OverscrollBehaviorInline string `default:".None"`
	OverscrollBehaviorX      string `default:".None"`
	OverscrollBehaviorY      string `default:".None"`
	Padding                  string `default:".None"`
	PaddingBlock             string `default:".None"`
	PaddingBlockEnd          string `default:".None"`
	PaddingBlockStart        string `default:".None"`
	PaddingBottom            string `default:".None"`
	PaddingInline            string `default:".None"`
	PaddingInlineEnd         string `default:".None"`
	PaddingInlineStart       string `default:".None"`
	PaddingLeft              string `default:".None"`
	PaddingRight             string `default:".None"`
	PaddingTop               string `default:".None"`
	PageBreakAfter           string `default:".None"`
	PageBreakBefore          string `default:".None"`
	PageBreakInside          string `default:".None"`
	PaintOrder               string `default:".None"`
	Perspective              string `default:".None"`
	PerspectiveOrigin        string `default:".None"`
	PlaceContent             string `default:".None"`
	PlaceItems               string `default:".None"`
	PlaceSelf                string `default:".None"`
	PointerEvents            string `default:".None"`
	Position                 string `default:".None"`
	Quotes                   string `default:".None"`
	Resize                   string `default:".None"`
	Right                    string `default:".None"`
	Rotate                   string `default:".None"`
	RowGap                   string `default:".None"`
	RubyAlign                string `default:".None"`
	RubyMerge                string `default:".None"`
	RubyPosition             string `default:".None"`
	Scale                    string `default:".None"`
	ScrollBehavior           string `default:".None"`
	ScrollMargin             string `default:".None"`
	ScrollMarginBlock        string `default:".None"`
	ScrollMarginBlockEnd     string `default:".None"`
	ScrollMarginBlockStart   string `default:".None"`
	ScrollMarginBottom       string `default:".None"`
	ScrollMarginInline       string `default:".None"`
	ScrollMarginInlineEnd    string `default:".None"`
	ScrollMarginInlineStart  string `default:".None"`
	ScrollMarginLeft         string `default:".None"`
	ScrollMarginRight        string `default:".None"`
	ScrollMarginTop          string `default:".None"`
	ScrollPadding            string `default:".None"`
	ScrollPaddingBlock       string `default:".None"`
	ScrollPaddingBlockEnd    string `default:".None"`
	ScrollPaddingBlockStart  string `default:".None"`
	ScrollPaddingBottom      string `default:".None"`
	ScrollPaddingInline      string `default:".None"`
	ScrollPaddingInlineEnd   string `default:".None"`
	ScrollPaddingInlineStart string `default:".None"`
	ScrollPaddingLeft        string `default:".None"`
	ScrollPaddingRight       string `default:".None"`
	ScrollPaddingTop         string `default:".None"`
	ScrollSnapAlign          string `default:".None"`
	ScrollSnapStop           string `default:".None"`
	ScrollSnapType           string `default:".None"`
	ShapeImageThreshold      string `default:".None"`
	ShapeMargin              string `default:".None"`
	ShapeOutside             string `default:".None"`
	TabSize                  string `default:".None"`
	TableLayout              string `default:".None"`
	TextAlign                string `default:".None"`
	TextAlignLast            string `default:".None"`
	TextCombineUpright       string `default:".None"`
	TextDecoration           string `default:".None"`
	TextDecorationColor      string `default:".None"`
	TextDecorationLine       string `default:".None"`
	TextDecorationSkip       string `default:".None"`
	TextDecorationSkipInk    string `default:".None"`
	TextDecorationStyle      string `default:".None"`
	TextDecorationThickness  string `default:".None"`
	TextEmphasis             string `default:".None"`
	TextEmphasisColor        string `default:".None"`
	TextEmphasisPosition     string `default:".None"`
	TextEmphasisStyle        string `default:".None"`
	TextIndent               string `default:".None"`
	TextJustify              string `default:".None"`
	TextOrientation          string `default:".None"`
	TextOverflow             string `default:".None"`
	TextRendering            string `default:".None"`
	TextShadow               string `default:".None"`
	TextSizeAdjust           string `default:".None"`
	TextTransform            string `default:".None"`
	TextUnderlinePosition    string `default:".None"`
	Top                      string `default:".None"`
	TouchAction              string `default:".None"`
	Transform                string `default:".None"`
	TransformBox             string `default:".None"`
	TransformOrigin          string `default:".None"`
	TransformStyle           string `default:".None"`
	Transition               string `default:".None"`
	TransitionDelay          string `default:".None"`
	TransitionDuration       string `default:".None"`
	TransitionProperty       string `default:".None"`
	TransitionTiming         string `default:".None"`
	Translate                string `default:".None"`
	UnicodeBidi              string `default:".None"`
	UserSelect               string `default:".None"`
	VerticalAlign            string `default:".None"`
	Visibility               string `default:".None"`
	WhiteSpace               string `default:".None"`
	Width                    string `default:".None"`
	WillChange               string `default:".None"`
	WordBreak                string `default:".None"`
	WordSpacing              string `default:".None"`
	WordWrap                 string `default:".None"`
	WritingMode              string `default:".None"`
	ZIndex                   string `default:".None"`
}

func ConvertSheetToSlice(sheet CSStylesheet) map[string]string {
	cssType := reflect.TypeOf(sheet)
	cssValue := reflect.ValueOf(sheet)

	// Inicializar el slice de mapas
	var cssMapSlice map[string]string = make(map[string]string)

	// Recorrer todos los campos de la estructura
	for i := 0; i < cssType.NumField(); i++ {
		field := cssType.Field(i)
		value := cssValue.Field(i).String()

		// Ignorar las propiedades con el valor por defecto
		if value != ".None" && value != "" && strings.TrimSpace(value) != "" {

			// Convertir el nombre del campo a un estilo CSS válido
			cssPropertyName := toCSSPropertyName(field.Name)

			cssMapSlice[cssPropertyName] = value
		}
	}
	return cssMapSlice
}
func toCSSPropertyName(a string) string {
	//A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X, Y, Z
	a = strings.ReplaceAll(a, "A", "-a")
	a = strings.ReplaceAll(a, "B", "-b")
	a = strings.ReplaceAll(a, "C", "-c")
	a = strings.ReplaceAll(a, "D", "-d")
	a = strings.ReplaceAll(a, "E", "-e")
	a = strings.ReplaceAll(a, "F", "-f")
	a = strings.ReplaceAll(a, "G", "-g")
	a = strings.ReplaceAll(a, "H", "-h")
	a = strings.ReplaceAll(a, "I", "-i")
	a = strings.ReplaceAll(a, "J", "-j")
	a = strings.ReplaceAll(a, "K", "-k")
	a = strings.ReplaceAll(a, "L", "-l")
	a = strings.ReplaceAll(a, "M", "-m")
	a = strings.ReplaceAll(a, "N", "-n")
	a = strings.ReplaceAll(a, "O", "-o")
	a = strings.ReplaceAll(a, "P", "-p")
	a = strings.ReplaceAll(a, "Q", "-q")
	a = strings.ReplaceAll(a, "R", "-r")
	a = strings.ReplaceAll(a, "S", "-s")
	a = strings.ReplaceAll(a, "T", "-t")
	a = strings.ReplaceAll(a, "U", "-u")
	a = strings.ReplaceAll(a, "V", "-v")
	a = strings.ReplaceAll(a, "W", "-w")
	a = strings.ReplaceAll(a, "X", "-x")
	a = strings.ReplaceAll(a, "Y", "-y")
	a = strings.ReplaceAll(a, "Z", "-z")
	a = strings.TrimPrefix(a, "-")
	return a

}
