package markdown

import (
	"github.com/charmbracelet/glamour/ansi"

	"github.com/go-fn/fn/internal/util"
)

func CustomDarkStyleConfig() ansi.StyleConfig {
	return ansi.StyleConfig{
		Document: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				BlockPrefix: "",
				BlockSuffix: "",
				Prefix:      "",
			},
			Margin:      util.Ptr(uint(0)),
			IndentToken: util.Ptr(""),
			Indent:      util.Ptr(uint(0)),
		},
		Paragraph: ansi.StyleBlock{
			Margin:      util.Ptr(uint(0)),
			IndentToken: util.Ptr(""),
			Indent:      util.Ptr(uint(0)),
			StylePrimitive: ansi.StylePrimitive{
				BlockPrefix: "",
				BlockSuffix: "",
				Prefix:      "",
			},
		},
		BlockQuote: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{},
			Indent:         util.Ptr(uint(1)),
			IndentToken:    util.Ptr("│ "),
		},
		List: ansi.StyleList{
			LevelIndent: 2,
		},
		Heading: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				BlockSuffix: "",
				Color:       util.Ptr("#666CA6"),
				Bold:        util.Ptr(true),
			},
		},
		H1: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				Prefix: "# ",
				Color:  util.Ptr("#666CA6"),
				Bold:   util.Ptr(true),
			},
		},
		H2: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				Prefix: "## ",
			},
		},
		H3: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				Prefix: "### ",
			},
		},
		H4: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				Prefix: "#### ",
			},
		},
		H5: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				Prefix: "##### ",
			},
		},
		H6: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				Prefix: "###### ",
				Color:  util.Ptr("35"),
				Bold:   util.Ptr(false),
			},
		},
		Strikethrough: ansi.StylePrimitive{
			CrossedOut: util.Ptr(true),
		},
		Emph: ansi.StylePrimitive{
			Italic: util.Ptr(true),
		},
		Strong: ansi.StylePrimitive{
			Bold: util.Ptr(true),
		},
		HorizontalRule: ansi.StylePrimitive{
			Color:  util.Ptr("240"),
			Format: "\n--------\n",
		},
		Item: ansi.StylePrimitive{
			BlockPrefix: "• ",
		},
		Enumeration: ansi.StylePrimitive{
			BlockPrefix: ". ",
		},
		Task: ansi.StyleTask{
			StylePrimitive: ansi.StylePrimitive{},
			Ticked:         "[✓] ",
			Unticked:       "[ ] ",
		},
		Link: ansi.StylePrimitive{
			Color:       util.Ptr("#666CA6"),
			Underline:   util.Ptr(false),
			BlockPrefix: "",
			BlockSuffix: "",
			Format:      "",
		},
		LinkText: ansi.StylePrimitive{
			Color: util.Ptr("#666CA6"),
			Bold:  util.Ptr(true),
		},
		Image: ansi.StylePrimitive{
			Underline: util.Ptr(false),
			Color:     util.Ptr("#666CA6"),
			Bold:      util.Ptr(false),
			Format:    "",
		},
		ImageText: ansi.StylePrimitive{
			Color:  util.Ptr("#666CA6"),
			Bold:   util.Ptr(true),
			Format: "",
		},
		Code: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				Color:       util.Ptr("#E2E1ED"),
				Overlined:   util.Ptr(true),
				Prefix:      "`",
				Suffix:      "`",
				BlockPrefix: "",
				BlockSuffix: "",
			},
		},
		CodeBlock: ansi.StyleCodeBlock{
			StyleBlock: ansi.StyleBlock{
				StylePrimitive: ansi.StylePrimitive{
					Color: util.Ptr("244"),
				},
				Margin: util.Ptr(uint(0)),
			},
			Chroma: &ansi.Chroma{
				Text: ansi.StylePrimitive{
					Color: util.Ptr("#C4C4C4"),
				},
				Error: ansi.StylePrimitive{
					Color:           util.Ptr("#F1F1F1"),
					BackgroundColor: util.Ptr("#F05B5B"),
				},
				Comment: ansi.StylePrimitive{
					Color: util.Ptr("#676767"),
				},
				CommentPreproc: ansi.StylePrimitive{
					Color: util.Ptr("#FF875F"),
				},
				Keyword: ansi.StylePrimitive{
					Color: util.Ptr("#00AAFF"),
				},
				KeywordReserved: ansi.StylePrimitive{
					Color: util.Ptr("#FF5FD2"),
				},
				KeywordNamespace: ansi.StylePrimitive{
					Color: util.Ptr("#FF5F87"),
				},
				KeywordType: ansi.StylePrimitive{
					Color: util.Ptr("#6E6ED8"),
				},
				Operator: ansi.StylePrimitive{
					Color: util.Ptr("#EF8080"),
				},
				Punctuation: ansi.StylePrimitive{
					Color: util.Ptr("#E8E8A8"),
				},
				Name: ansi.StylePrimitive{
					Color: util.Ptr("#C4C4C4"),
				},
				NameBuiltin: ansi.StylePrimitive{
					Color: util.Ptr("#FF8EC7"),
				},
				NameTag: ansi.StylePrimitive{
					Color: util.Ptr("#B083EA"),
				},
				NameAttribute: ansi.StylePrimitive{
					Color: util.Ptr("#7A7AE6"),
				},
				NameClass: ansi.StylePrimitive{
					Color:     util.Ptr("#F1F1F1"),
					Underline: util.Ptr(true),
					Bold:      util.Ptr(true),
				},
				NameDecorator: ansi.StylePrimitive{
					Color: util.Ptr("#FFFF87"),
				},
				NameFunction: ansi.StylePrimitive{
					Color: util.Ptr("#00D787"),
				},
				LiteralNumber: ansi.StylePrimitive{
					Color: util.Ptr("#6EEFC0"),
				},
				LiteralString: ansi.StylePrimitive{
					Color: util.Ptr("#C69669"),
				},
				LiteralStringEscape: ansi.StylePrimitive{
					Color: util.Ptr("#AFFFD7"),
				},
				GenericDeleted: ansi.StylePrimitive{
					Color: util.Ptr("#FD5B5B"),
				},
				GenericEmph: ansi.StylePrimitive{
					Italic: util.Ptr(true),
				},
				GenericInserted: ansi.StylePrimitive{
					Color: util.Ptr("#00D787"),
				},
				GenericStrong: ansi.StylePrimitive{
					Bold: util.Ptr(true),
				},
				GenericSubheading: ansi.StylePrimitive{
					Color: util.Ptr("#777777"),
				},
				Background: ansi.StylePrimitive{
					BackgroundColor: util.Ptr("#373737"),
				},
			},
		},
		Table: ansi.StyleTable{
			StyleBlock: ansi.StyleBlock{
				StylePrimitive: ansi.StylePrimitive{
					Format:  " ",
					Conceal: util.Ptr(true),
				},
			},
			CenterSeparator: util.Ptr(""),
			ColumnSeparator: util.Ptr(""),
			RowSeparator:    util.Ptr(""),
		},
		DefinitionDescription: ansi.StylePrimitive{
			BlockPrefix: " ",
		},
	}
}
