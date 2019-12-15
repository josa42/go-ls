package lsp

import "github.com/sourcegraph/go-lsp"

type DocumentSymbolParams struct {
	TextDocument lsp.TextDocumentIdentifier `json:"textDocument"`
}

type SymbolKind int

const (
	SKFile          SymbolKind = 1
	SKModule        SymbolKind = 2
	SKNamespace     SymbolKind = 3
	SKPackage       SymbolKind = 4
	SKClass         SymbolKind = 5
	SKMethod        SymbolKind = 6
	SKProperty      SymbolKind = 7
	SKField         SymbolKind = 8
	SKConstructor   SymbolKind = 9
	SKEnum          SymbolKind = 10
	SKInterface     SymbolKind = 11
	SKFunction      SymbolKind = 12
	SKVariable      SymbolKind = 13
	SKConstant      SymbolKind = 14
	SKString        SymbolKind = 15
	SKNumber        SymbolKind = 16
	SKBoolean       SymbolKind = 17
	SKArray         SymbolKind = 18
	SKObject        SymbolKind = 19
	SKKey           SymbolKind = 20
	SKNull          SymbolKind = 21
	SKEnumMember    SymbolKind = 22
	SKStruct        SymbolKind = 23
	SKEvent         SymbolKind = 24
	SKOperator      SymbolKind = 25
	SKTypeParameter SymbolKind = 26
)

/**
 * Represents programming constructs like variables, classes, interfaces etc. that appear in a document. Document symbols can be
 * hierarchical and they have two ranges: one that encloses its definition and one that points to its most interesting range,
 * e.g. the range of an identifier.
 */
type DocumentSymbol struct {

	/**
	 * The name of this symbol. Will be displayed in the user interface and therefore must not be
	 * an empty string or a string only consisting of white spaces.
	 */
	Name string `json:"name"`

	/**
	 * More detail for this symbol, e.g the signature of a function.
	 */
	Detail string `json:"detail,omitempty"`

	/**
	 * The kind of this symbol.
	 */
	Kind SymbolKind `json:"kind"`

	/**
	 * Indicates if this symbol is deprecated.
	 */
	Deprecated bool `json:"deprecated,omitempty"`

	/**
	 * The range enclosing this symbol not including leading/trailing whitespace but everything else
	 * like comments. This information is typically used to determine if the clients cursor is
	 * inside the symbol to reveal in the symbol in the UI.
	 */
	Range lsp.Range `json:"range"`

	/**
	 * The range that should be selected and revealed when this symbol is being picked, e.g the name of a function.
	 * Must be contained by the `range`.
	 */
	SelectionRange lsp.Range `json:"selectionRange"`

	/**
	 * Children of this symbol, e.g. properties of a class.
	 */
	Children []DocumentSymbol `json:"children,omitempty"`
}

/**
//  * Represents information about programming constructs like variables, classes,
//  * interfaces etc.
//  */
// type SymbolInformation struct {
// 	/**
// 	 * The name of this symbol.
// 	 */
// 	Name string `json:"name"`
//
// 	/**
// 	 * The kind of this symbol.
// 	 */
// 	Kind int `json:"kind"`
//
// 	/**
// 	 * Indicates if this symbol is deprecated.
// 	 */
// 	Deprecated bool `json:"deprecated,omitempty"`
//
// 	/**
// 	 * The location of this symbol. The location's range is used by a tool
// 	 * to reveal the location in the editor. If the symbol is selected in the
// 	 * tool the range's start information is used to position the cursor. So
// 	 * the range usually spans more then the actual symbol's name and does
// 	 * normally include things like visibility modifiers.
// 	 *
// 	 * The range doesn't have to denote a node range in the sense of a abstract
// 	 * syntax tree. It can therefore not be used to re-construct a hierarchy of
// 	 * the symbols.
// 	 */
// 	Location lsp.Location `json:"location"`
//
// 	/**
// 	 * The name of the symbol containing this symbol. This information is for
// 	 * user interface purposes (e.g. to render a qualifier in the user interface
// 	 * if necessary). It can't be used to re-infer a hierarchy for the document
// 	 * symbols.
// 	 */
// 	ContainerName string `json:"containerName,omitempty"`
// }
