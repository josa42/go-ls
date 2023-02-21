package lsp

// Params for the CodeActionRequest
type CodeActionParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"` // The document in which the command was invoked.
	Range        Range                  `json:"range"`        // The range for which the command was invoked.
	Context      CodeActionContext      `json:"context"`      // Context carrying additional information.
}

// The kind of a code action.
//
// Kinds are a hierarchical list of identifiers separated by `.`, e.g. `"refactor.extract.function"`.
//
// The set of kinds is open and client needs to announce the kinds it supports to the server during
// initialization.
type CodeActionKind string

// A set of predefined code action kinds
const (

	// Empty kind.
	CAKEmpty CodeActionKind = ""

	// Base kind for quickfix actions: 'quickfix'
	CAKQuickFix CodeActionKind = "quickfix"

	// Base kind for refactoring actions: 'refactor'
	CAKRefactor CodeActionKind = "refactor"

	// Base kind for refactoring extraction actions: 'refactor.extract'
	//
	// Example extract actions:
	//
	// - Extract method
	// - Extract function
	// - Extract variable
	// - Extract interface from class
	// - ...
	CAKRefactorExtract CodeActionKind = "refactor.extract"

	// Base kind for refactoring inline actions: 'refactor.inline'
	//
	// Example inline actions:
	//
	// - Inline function
	// - Inline variable
	// - Inline constant
	// - ...
	CAKRefactorInline CodeActionKind = "refactor.inline"

	// Base kind for refactoring rewrite actions: 'refactor.rewrite'
	//
	// Example rewrite actions:
	//
	// - Convert JavaScript function to class
	// - Add or remove parameter
	// - Encapsulate field
	// - Make method static
	// - Move method to base class
	// - ...
	CAKRefactorRewrite CodeActionKind = "refactor.rewrite"

	// Base kind for source actions: `source`
	//
	// Source code actions apply to the entire file.
	CAKSource CodeActionKind = "source"

	// Base kind for an organize imports source action: `source.organizeImports`
	CAKSourceOrganizeImports CodeActionKind = "source.organizeImports"
)

// Contains additional diagnostic information about the context in which
// a code action is run.
type CodeActionContext struct {
	// An array of diagnostics.
	Diagnostics []Diagnostic `json:"diagnostics"`

	// Requested kind of actions to return.
	//
	// Actions not of this kind are filtered out by the client before being shown. So servers
	// can omit computing them.
	Only *[]CodeActionKind `json:"only,omitempty"`
}

// A code action represents a change that can be performed in code, e.g. to fix a problem or
// to refactor code.
//
// A CodeAction must set either `edit` and/or a `command`. If both are supplied, the `edit` is applied first, then the `command` is executed.
type CodeAction struct {

	// A short, human-readable, title for this code action.
	Title string `json:"title"`

	// The kind of the code action.
	//
	// Used to filter code actions.
	Kind *CodeActionKind `json:"kind,omitempty"`

	// The diagnostics that this code action resolves.
	Diagnostics *[]Diagnostic `json:"diagnostics,omitempty"`

	// The workspace edit this code action performs.
	Edit *WorkspaceEdit `json:"edit,omitempty"`

	// A command this code action executes. If a code action
	// provides an edit and a command, first the edit is
	// executed and then the command.
	Command *Command `json:"command,omitempty"`
}
